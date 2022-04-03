package api

import (
	"github.com/gin-gonic/gin"
	"okex/service"
)

// ======= 查看余额 v5 ====
func AccountBalanceHttp(c *gin.Context) {
	DoHttpProcess(new(AccountBalanceApi), c)
}

type AccountBalanceApi struct {
	Base
}

func (a *AccountBalanceApi) ProcessHttp() {
	ccy := a.ctx.DefaultQuery("ccy", "")
	new(service.AccountSvr).GetBalance(ccy)
}
