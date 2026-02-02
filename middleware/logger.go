package middleware

import (
	"bytes"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/oTeeLeko/mystore/util"
)

// func Logger() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		ctx.Next()
// 		util.LogActivity(ctx)
// 	}
// }

// AccessLogger สำหรับ log การเข้าถึง
func AccessLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// เก็บ request body ไว้สำหรับ error logging
		var bodyBytes []byte
		if ctx.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(ctx.Request.Body)
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			ctx.Set("requestBody", string(bodyBytes))
		}

		ctx.Next()
		// Log ทุกครั้งไม่ว่า LogLevel จะเป็นอะไร
		util.LogActivity(ctx)
	}
}
