package api

import (
	"github.com/gin-gonic/gin"
	"okex/api/strategy"
)

func TestHttp(c *gin.Context) {
	DoHttpProcess(new(TestApi), c)
}

type TestApi struct {
	Base
}

func (a *TestApi) ProcessHttp() {
	new(strategy.StrategyOne).Init().Run()
}