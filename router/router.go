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
		r.GET("/account/wallet", api.AccountWalletHttp)
		r.GET("/test", api.TestHttp)
	}

	return g
}
