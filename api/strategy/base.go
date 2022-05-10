package strategy

import (
	"fmt"
	"okex/model"
	"okex/service"
	"time"
)

type Strategy struct {
	price              float32 // 设置价格
	percentageIncrease int // 价格涨幅百分比
	percentageDrop     int // 价格跌幅百分比
	currency           string   // 币种
	MarketSvr *service.MarketSvr   // 实例
}

/**
 * 获取市场对象
 */

func (s *Strategy) GetMarketSvr() *service.MarketSvr {
	if s.MarketSvr != nil {
		return s.MarketSvr
	}
	return new(service.MarketSvr)
}

/**
 * 获取市场行情
 */

func (s *Strategy) GetTickerInfo() (model.TickerRsp, error) {
	params := model.TickerReq {
		//InstId: "ORS-USDT",
		InstId: s.currency,
	}

	marketSvr := s.GetMarketSvr()
	res, err := marketSvr.Ticker(params)
	return res, err
}

/**
 * 设置价格
 */

func (s *Strategy) SetPrice(price float32) {
	s.price = price
}

/**
 * 设置币种
 */

func (s *Strategy) SetCurrency(currency string) {
	s.currency = currency
}

/**
 * 设置涨幅百分比
 */

func (s *Strategy) SetPercentageIncrease(percent int) {
	s.percentageIncrease = percent
}

/**
 * 设置跌幅百分比
 */

func (s *Strategy) SetPercentageDrop(percent int) {
	s.percentageDrop = percent
}

/**
 * 运行
 */

func (s *Strategy) Run() {
	// 创建定时器
	ticker := time.NewTicker(time.Second * 2)
	go func() {
		for {
			<-ticker.C
			s.Do()
		}
	}()
}

func (s *Strategy) Do() {
	fmt.Println("do")
}


