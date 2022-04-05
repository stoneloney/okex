package trade

import (
	"github.com/gin-gonic/gin"
	"okex/api"
	"okex/model"
	"okex/service"
)

// ======= 订单详情 ======

func OrderInfoHttp(c *gin.Context) {
	api.DoHttpProcess(new(OrderInfoApi), c)
}

type OrderInfoApi struct {
	api.Base

	apiParams model.TradeOrderInfoReq
}

/**
  订单详情
  @ instId  产品ID，如 BTC-USD-190927  （必填)
  @ ordId   订单ID  ordId和clOrdId必须传一个，若传两个，以ordId为主  (选填)
  @ clOrdId  用户自定义ID (选填)
*/

func (a *OrderInfoApi) ProcessHttp() {
	if err := a.Ctx.ShouldBind(&a.apiParams); err != nil {
		a.Response(1001, nil, "", "")
		return
	}
	res, err := new(service.TradeSvr).OrderInfo(a.apiParams)
	if err != nil {
		a.Response(3001, "", "查询错误", err.Error())
		return
	}

	a.Response(0, res, "success", "")
}
