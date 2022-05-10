package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"okex/api/strategy"
	"okex/router"
)

func main() {
	g := gin.Default()
	router.Load(g)

	// 运行策略1
	go new(strategy.StrategyOne).Init().Run()

	_ = http.ListenAndServe(":14000", g)
}