package router

import (
	"github.com/gin-gonic/gin"
	"okex/api"
)

// Load 加载路由
func Load(g *gin.Engine) *gin.Engine {
	g.Use(api.Recovery()) // 引用中间件 不会panic防止500错误

	r := g.Group("/okex/")
	{
		// 资金账户信息 v3
		r.GET("/account/v3/wallet", api.AccountWalletHttp)
		// 查看账户余额 v5
		r.GET("/account/v5/balance", api.AccountBalanceHttp)
		r.GET("/test", api.TestHttp)
	}

	return g
}
