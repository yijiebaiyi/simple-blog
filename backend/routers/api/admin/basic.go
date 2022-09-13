package admin

import (
	"gin_admin/pkg/e"
	"gin_admin/pkg/util"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
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

func IsLogin(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	loginAdmin := session.Get(util.SESSION_KEY_LOGIN_ADMIN)
	if loginAdmin == nil {
		return false
	}
	return true
}

// CheckLogin 验证登录
func CheckLogin(ctx *gin.Context) {
	session := sessions.Default(ctx)
	loginAdmin := session.Get(util.SESSION_KEY_LOGIN_ADMIN)
	if loginAdmin == nil {
		// 跳转到登录页
		ctx.Redirect(302, "/login")
	}
}
