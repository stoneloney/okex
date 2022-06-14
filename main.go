package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"okex/api/strategy"
	"okex/router"
)

func main() {
	g := gin.Default()
	router.Load(g)

	// 运行策略1
	go func() {
		err := new(strategy.StrategyOne).Init().Run()
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	_ = http.ListenAndServe(":14000", g)
}