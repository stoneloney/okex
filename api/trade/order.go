package trade

import (
	"github.com/gin-gonic/gin"
	"okex/api"
	"okex/service"
)

// ======= 下单 ======

func OrderHttp(c *gin.Context) {
	api.DoHttpProcess(new(OrderApi), c)
}

type OrderApi struct {
	api.Base

	apiParams struct {
		InstId     string `json:"instId" form:"instId"`
		TdMode     string `json:"tdMode" form:"tdMode"`
		Ccy        string `json:"ccy" form:"ccy"`
		ClOrdId    string `json:"clOrdId" form:"clOrdId"`
		Tag        string `json:"tag" form:"tag"`
		Side       string `json:"side" form:"side"`
		PosSide    string `json:"posSide" form:"posSide"`
		OrdType    string `json:"ordType" form:"ordType"`
		Sz         string `json:"sz" form:"sz"`
		Px         string `json:"px" form:"px"`
		ReduceOnly bool   `json:"reduceOnly" form:"reduceOnly"`
		TgtCcy     bool   `json:"tgtCcy" form:"tgtCcy"`
	}
}

func (a *OrderApi) ProcessHttp() {
	if err := a.Ctx.ShouldBind(a.apiParams); err != nil {
		a.Response(1001, nil, "", "")
		return
	}
	res, err := new(service.TradeSvr).Order(a.apiParams)
	if err != nil {
		a.Response(3001, "", "查询错误", err.Error())
		return
	}

	a.Response(0, res, "success", "")
}
