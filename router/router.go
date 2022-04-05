package router

import (
	"github.com/gin-gonic/gin"
	"okex/api"
	"okex/api/account"
	"okex/api/market"
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
		r.POST("/trade/order_info", trade.OrderInfoHttp)         // 订单详情
		r.POST("/trade/orders_pending", trade.OrderPendingHttp)  // 未成交订单
		r.POST("/trade/orders_history", trade.OrderHistoryHttp)  // 历史订单记录
		r.POST("/trade/order", trade.OrderHttp)                 // 下单
		r.POST("/trade/cancel_order", trade.CancelOrderHttp)    // 撤单

		// 行情
		r.POST("/market/tickers", market.TickersHttp)           // 所有产品行情

		r.GET("/test", api.TestHttp)
	}

	return g
}
