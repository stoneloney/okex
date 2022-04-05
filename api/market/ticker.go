package market

import (
	"github.com/gin-gonic/gin"
	"okex/api"
	"okex/model"
	"okex/service"
)

// ======= 单个产品行情 ======

func TickerHttp(c *gin.Context) {
	api.DoHttpProcess(new(TickerApi), c)
}

type TickerApi struct {
	api.Base

	apiParams model.TickerReq
}

func (a *TickerApi) ProcessHttp() {
	if err := a.Ctx.ShouldBind(&a.apiParams); err != nil {
		a.Response(1001, nil, "", "")
		return
	}
	res, err := new(service.MarketSvr).Ticker(a.apiParams)
	if err != nil {
		a.Response(3001, "", "查询错误", err.Error())
		return
	}

	a.Response(0, res, "success", "")
}
