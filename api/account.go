package api

import (
	"github.com/gin-gonic/gin"
	"okex/service"
)

func AccountWalletHttp(c *gin.Context) {
	DoHttpProcess(new(AccountWalletApi), c)
}

type AccountWalletApi struct {
	Base
}

func (a *AccountWalletApi) ProcessHttp() {
	new(service.AccountSvr).GetWallet()
}