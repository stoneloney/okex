package strategy

import (
	"fmt"
	"strconv"
	"time"
)

type StrategyOne struct {
	Strategy
}

func (s *StrategyOne) Init() *StrategyOne {
	s.SetCurrency("BTC-USDT")    // 查询BTC
	s.SetPrice(30764.2)             // 设置起始价格
	s.SetPercentageIncrease(1)   // 设置涨幅百分比
	s.SetPercentageDrop(1)       // 设置跌幅百分比


	return s
}
func (s *StrategyOne) Run() {
	// 创建定时器
	ticker := time.NewTicker(time.Second * 60)
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
	fmt.Println(res)

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
		contrastPercentage := (lastPriceFloat - s.price) / s.price * 100   // 转为100

		if contrastPercentage > 0 && contrastPercentage >= s.percentageIncrease {   // 价格增加,触发减仓策略
			fmt.Println(fmt.Sprintf("减仓, contrastPercentage:%v, percentageIncrease:%v, currentPrice:%v, setPrice:%v",
				contrastPercentage,
				s.percentageIncrease,
				lastPriceFloat,
				s.price))
		} else if contrastPercentage < 0 && contrastPercentage >= s.percentageDrop {      // 价格减少,触发补仓策略
			fmt.Println(fmt.Sprintf("补仓, contrastPercentage:%v, percentageIncrease:%v, currentPrice:%v, setPrice:%v",
				contrastPercentage,
				s.percentageIncrease,
				lastPriceFloat,
				s.price))
		} else {    // 保持监控状态
			fmt.Println(fmt.Sprintf("监控, contrastPercentage:%v, currentPrice:%v, setPrice:%v",
				contrastPercentage,
				lastPriceFloat,
				s.price))
		}
	}

}

