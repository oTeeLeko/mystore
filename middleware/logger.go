package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/oTeeLeko/mystore/util"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		util.LogActivity(ctx)
	}
}
