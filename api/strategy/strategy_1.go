package strategy

import (
	"fmt"
	"math"
	"okex/helper"
	"strconv"
	"time"
)

type StrategyOne struct {
	Strategy

	totalAmount float64 // 投入的总金额
	number      float64 // 拥有的币数
	percentage  float64 // 加减仓的百分比 (补仓按价格的百分比，减仓按币数的百分比)
}

func (s *StrategyOne) Init() *StrategyOne {
	s.SetCurrency("BTC-USDT")     // 查询BTC
	s.SetPrice(32338)           // 设置起始价格
	s.SetPercentageIncrease(0.02) // 设置涨幅百分比
	s.SetPercentageDrop(0.02)     // 设置跌幅百分比

	s.percentage = 0.1 // 每次补减百分比 (10%)
	s.totalAmount = 1000000 // 总金额

	return s
}
func (s *StrategyOne) Run() {
	// 创建定时器
	ticker := time.NewTicker(time.Second * 20)
	go func() {
		for {
			<-ticker.C
			s.Do()
		}
	}()
}

func (s *StrategyOne) Do() {
	res, err := s.GetTickerInfo()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if res.Code != "0" {
		fmt.Println(fmt.Sprintf("code is %s, errMsg:%s", res.Code, res.Msg))
		return
	}

	// 最新交易价格
	if len(res.Data) > 0 {
		detail := res.Data[0]
		lastPrice := detail.Last

		lastPriceFloat, _ := strconv.ParseFloat(lastPrice, 32)
		// 执行价格对比
		contrastPercentage := (lastPriceFloat - s.price) / s.price

		if contrastPercentage > 0 && contrastPercentage >= s.percentageIncrease { // 价格增加,触发减仓策略
			fmt.Println(fmt.Sprintf("减仓, contrastPercentage:%v, percentageIncrease:%v, currentPrice:%v, setPrice:%v",
				contrastPercentage,
				s.percentageIncrease,
				lastPriceFloat,
				s.price))

			// 计算卖出数
			if s.number <= 0 {
				fmt.Println(fmt.Sprintf("币数为0, 不能执行减仓操作, number:%v", s.number))
				return
			}

			// 卖出数量
			sellNumber := math.Ceil(s.number * s.percentage)
			if s.number < sellNumber {
				sellNumber = s.number
			}

			// 计算卖出总价
			sellAmount := sellNumber * lastPriceFloat

			// 总币数减少
			s.number -= sellNumber
			// 总金额添加
			s.totalAmount += sellAmount

		} else if contrastPercentage < 0 && contrastPercentage <= s.percentageDrop { // 价格减少,触发补仓策略
			fmt.Println(fmt.Sprintf("补仓, contrastPercentage:%v, percentageIncrease:%v, currentPrice:%v, setPrice:%v",
				contrastPercentage,
				s.percentageIncrease,
				lastPriceFloat,
				s.price))

			if s.totalAmount <= 0 {
				fmt.Println(fmt.Sprintf("金额0, 不能执行补仓操作, amout:%v", s.totalAmount))
				return
			}

			// 计算本次购买使用的金额
			buyAmount := s.totalAmount * s.percentage
			if s.totalAmount < buyAmount {
				buyAmount = s.totalAmount
			}

			// 计算可购买的数量
			buyNumber := math.Floor(buyAmount / lastPriceFloat)

			// 总币数添加
			s.number += buyNumber
			// 总金额减少
			s.totalAmount -= buyAmount

		} else { // 保持监控状态
			fmt.Println(fmt.Sprintf("监控, contrastPercentage:%v, currentPrice:%v, setPrice:%v, number:%v, amount:%v",
				contrastPercentage,
				lastPriceFloat,
				s.price,
				s.number,
				helper.Float64ToString(s.totalAmount)))
		}
	}

}
