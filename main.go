package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"okex/router"
)

func main() {
	g := gin.Default()
	router.Load(g)

	_ = http.ListenAndServe(":14000", g)
}