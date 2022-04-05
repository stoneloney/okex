package trade

import (
	"github.com/gin-gonic/gin"
	"okex/api"
	"okex/service"
)

// ======= 订单详情 ======

func OrderInfoHttp(c *gin.Context) {
	api.DoHttpProcess(new(OrderInfoApi), c)
}

type OrderInfoApi struct {
	api.Base

	apiParams struct {
		InstId   string  `json:"instId" form:"instId"`
		OrdId    string  `json:"ordId" form:"ordId"`
		ClOrdId  string  `json:"clOrdId" form:"clOrdId"`
	}
}

/**
  订单详情
  @ instId  产品ID，如 BTC-USD-190927  （必填)
  @ ordId   订单ID  ordId和clOrdId必须传一个，若传两个，以ordId为主  (选填)
  @ clOrdId  用户自定义ID (选填)
*/

func (a *OrderInfoApi) ProcessHttp() {
	if err := a.Ctx.ShouldBind(a.apiParams); err != nil {
		a.Response(1001, nil, "", "")
		return
	}
	res, err := new(service.TradeSvr).OrderInfo(a.apiParams.InstId, a.apiParams.OrdId, a.apiParams.ClOrdId)
	if err != nil {
		a.Response(3001, "", "查询错误", err.Error())
		return
	}

	a.Response(0, res, "success", "")
}
