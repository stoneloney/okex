package watch

import (
	"fmt"
	"okex/model"
	"okex/service"
	"time"
)

var (
	MarketSvr *service.MarketSvr
)

// 实时获取usdt价格
func WatchUSDT() {
	// 创建定时器
	ticker := time.NewTicker(time.Second * 2)
	go func() {
		for {
			<-ticker.C
			UsdtDo()
		}
	}()
}

func UsdtDo() {
	params := model.TickerReq {
		InstId: "ORS-USDT",
	}

	marketSvr := GetMarketSvr()
	res, err := marketSvr.Ticker(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
}

func GetMarketSvr() *service.MarketSvr {
	if MarketSvr != nil {
		return MarketSvr
	}
	return new(service.MarketSvr)
}

