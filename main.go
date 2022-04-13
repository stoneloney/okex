package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"okex/api/watch"
	"okex/router"
)

func main() {
	g := gin.Default()
	router.Load(g)

	// 开一个协程来监控价格变化
	go watch.WatchUSDT()

	_ = http.ListenAndServe(":14000", g)
}