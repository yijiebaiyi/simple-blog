package hander

import "github.com/gin-gonic/gin"

func HandelHeader() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
	}
}
