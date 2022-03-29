package api

import (
	"github.com/gin-gonic/gin"
)

func TestHttp(c *gin.Context) {
	DoHttpProcess(new(TestApi), c)
}

type TestApi struct {
	Base
}

func (a *TestApi) ProcessHttp() {
	a.Response(0, "aaaa", "success")
}