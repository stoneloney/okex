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
		r.GET("/trade/order_info", trade.OrderInfoHttp)         // 订单详情
		r.GET("/trade/orders_pending", trade.OrderPendingHttp)  // 未成交订单
		r.GET("/trade/orders_history", trade.OrderHistoryHttp)  // 历史订单记录
		r.POST("/trade/order", trade.OrderHttp)                 // 下单
		r.POST("/trade/cancel_order", trade.CancelOrderHttp)    // 撤单



		r.GET("/test", api.TestHttp)
	}

	return g
}
