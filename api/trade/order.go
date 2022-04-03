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
}

func (a *OrderApi) ProcessHttp() {
	params := map[string]interface{} {
		"aaa": "1111",
	}
	res, err := new(service.TradeSvr).Order(params)
	if err != nil {
		a.Response(3001, "", "查询错误", err.Error())
		return
	}

	a.Response(0, res, "success", "")
}
