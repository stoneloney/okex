package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

// ======== APi 基类 ==========

type Base struct {
	Ctx *gin.Context
}

func (a *Base) SetCtx(ctx *gin.Context)  {
	a.Ctx = ctx
}

func(a *Base) Response(code int, data interface{}, msg string, errMsg string) {
	a.Ctx.JSON(http.StatusOK, gin.H{"code": code, "data": data, "msg":msg, "errMsg":errMsg})
	return
}


// ========== 请求基类 ==========

type RequestHandler interface {
	SetCtx(ctx *gin.Context)
	ProcessHttp()
}

func DoHttpProcess(handler RequestHandler, ctx *gin.Context) {
	defer Recovery()

	handler.SetCtx(ctx)
	handler.ProcessHttp()
}

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		// panic的错误输出到log日志
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(string(debug.Stack()))
				c.JSON(http.StatusOK, gin.H{"code": 500, "msg": "500 error", "data": ""})
			}
		}()
		c.Next()
	}
}