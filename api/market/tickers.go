package market

import (
	"github.com/gin-gonic/gin"
	"okex/api"
	"okex/model"
	"okex/service"
)

// ======= 所有产品行情 ======

func TickersHttp(c *gin.Context) {
	api.DoHttpProcess(new(TickersApi), c)
}

type TickersApi struct {
	api.Base

	apiParams model.TickersReq
}

func (a *TickersApi) ProcessHttp() {
	if err := a.Ctx.ShouldBind(&a.apiParams); err != nil {
		a.Response(1001, nil, "", "")
		return
	}
	res, err := new(service.MarketSvr).Tickers(a.apiParams)
	if err != nil {
		a.Response(3001, "", "查询错误", err.Error())
		return
	}

	a.Response(0, res, "success", "")
}
