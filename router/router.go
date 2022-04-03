package router

import (
	"github.com/gin-gonic/gin"
	"okex/api"
	"okex/api/account"
	"okex/api/trade"
)

// Load 加载路由
func Load(g *gin.Engine) *gin.Engine {
	g.Use(api.Recovery()) // 引用中间件 不会panic防止500错误

	r := g.Group("/okex/")   // v5版本
	{
		// 账户相关
		r.GET("/account/balance", account.BalanceHttp)          // 账户余额
		r.GET("/account/positions", account.PositionsHttp)      // 持仓信息

		// 交易
		r.POST("/trade/order", trade.OrderHttp)                 // 下单


		r.GET("/test", api.TestHttp)
	}

	return g
}
