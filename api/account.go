package api

import (
	"github.com/gin-gonic/gin"
	"okex/service"
)

// ====  查看资金 v3  ====
func AccountWalletHttp(c *gin.Context) {
	DoHttpProcess(new(AccountWalletApi), c)
}

type AccountWalletApi struct {
	Base
}

func (a *AccountWalletApi) ProcessHttp() {
	new(service.AccountSvr).GetWallet()
}

// ======= 查看余额 v5 ====
func AccountBalanceHttp(c *gin.Context) {
	DoHttpProcess(new(AccountBalanceApi), c)
}

type AccountBalanceApi struct {
	Base
}

func (a *AccountBalanceApi) ProcessHttp() {
	new(service.AccountSvr).GetBalance()
}
