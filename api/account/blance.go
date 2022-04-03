package account

import (
	"github.com/gin-gonic/gin"
	"okex/api"
	"okex/service"
)

// ======= 查看余额 ======

func BalanceHttp(c *gin.Context) {
	api.DoHttpProcess(new(BalanceApi), c)
}

type BalanceApi struct {
	api.Base
}

func (a *BalanceApi) ProcessHttp() {
	ccy := a.Ctx.DefaultQuery("ccy", "")
	res, err := new(service.AccountSvr).GetBalance([]string{ccy})
	if err != nil {
		a.Response(3001, "", "查询错误", err.Error())
		return
	}

	a.Response(0, res, "success", "")
}
