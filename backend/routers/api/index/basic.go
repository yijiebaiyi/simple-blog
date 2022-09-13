package index

import (
	"gin_admin/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonReturn(ctx *gin.Context, code int, msg string, data interface{}) {
	if data == nil {
		data = make(map[string]interface{})
	}

	if msg == "" {
		msg = e.GetMsg(code)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
