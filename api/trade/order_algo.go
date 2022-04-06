package trade

import (
	"github.com/gin-gonic/gin"
	"okex/api"
	"okex/model"
	"okex/service"
)

// ======= 下单 ======

func OrderAlgoHttp(c *gin.Context) {
	api.DoHttpProcess(new(OrderAlgoApi), c)
}

type OrderAlgoApi struct {
	api.Base

	apiParams model.TradeOrderAlgoReq
}

func (a *OrderAlgoApi) ProcessHttp() {
	if err := a.Ctx.ShouldBind(&a.apiParams); err != nil {
		a.Response(1001, nil, "", "")
		return
	}
	res, err := new(service.TradeSvr).OrderAlgo(a.apiParams)
	if err != nil {
		a.Response(3001, "", "查询错误", err.Error())
		return
	}

	a.Response(0, res, "success", "")
}
