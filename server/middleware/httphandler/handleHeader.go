package httphandler

import (
	"gin_admin/middleware/httpresponse"

	"github.com/gin-gonic/gin"
)

func HandelHeader() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 处理预检请求
		if ctx.Request.Method == "OPTIONS" {
			httpresponse.SuccessResponse(ctx)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
