package market

import (
	"github.com/gin-gonic/gin"
	"okex/api"
	"okex/model"
	"okex/service"
)

// ======= 产品K线数据 ======

func CandlesHttp(c *gin.Context) {
	api.DoHttpProcess(new(CandlesApi), c)
}

type CandlesApi struct {
	api.Base

	apiParams model.CandlesReq
}

func (a *CandlesApi) ProcessHttp() {
	if err := a.Ctx.ShouldBind(&a.apiParams); err != nil {
		a.Response(1001, nil, "", "")
		return
	}
	res, err := new(service.MarketSvr).Candles(a.apiParams)
	if err != nil {
		a.Response(3001, "", "查询错误", err.Error())
		return
	}

	a.Response(0, res, "success", "")
}
