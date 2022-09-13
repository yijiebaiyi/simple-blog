package admin

import (
	"fmt"
	"gin_admin/models"
	"gin_admin/pkg/e"
	"gin_admin/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/thinkeridea/go-extend/exnet"
	"time"
)

func Login(ctx *gin.Context) {
	adminName, _ := ctx.GetPostForm("admin_name")
	adminPassword, _ := ctx.GetPostForm("admin_password")

	// 验证
	valid := validation.Validation{}
	valid.Required(adminName, "admin_name").Message("请输入用户名")
	valid.Required(adminPassword, "admin_password").Message("请输入密码")
	valid.MaxSize(adminName, 20, "admin_name").Message("用户名不能超过20位")
	valid.MaxSize(adminPassword, 18, "admin_name").Message("密码不能超过18位")
	valid.MinSize(adminPassword, 3, "admin_name").Message("密码不能小于3位")

	var msg string
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			msg = err.Message
			break
		}
		JsonReturn(ctx, e.INVALID_PARAMS, msg, nil)
		return
	}

	// 验证用户名和密码
	admin := models.GetAdminByName(adminName)
	if admin.AdminName == "" {
		JsonReturn(ctx, e.ERROR_NOT_EXIST_ADMIN, "", nil)
		return
	}
	fmt.Println(models.GetPasswordOfMd5(adminPassword))
	if admin.AdminPassword != models.GetPasswordOfMd5(adminPassword) {
		JsonReturn(ctx, e.ERROR_PASSWORD, "", nil)
		return
	}

	// 用户是否被禁用
	if admin.AdminState == models.ADMIN_STATE_FORBIDDEN {
		JsonReturn(ctx, e.ERROR_FORBIDDEN_ADMIN, "", nil)
		return
	}

	// 修改最后登录ip和最后登录时间
	clientIp := exnet.ClientIP(ctx.Request)
	clientIpLong, _ := exnet.IPString2Long(clientIp)
	admin.LastLoginTime = int(time.Now().Unix())
	admin.LastLoginIp = int(clientIpLong)
	models.EditAdmin(admin.ID, admin)

	// 保存到session
	session := sessions.Default(ctx)
	session.Set(util.SESSION_KEY_LOGIN_ADMIN, admin)
	_ = session.Save()
	JsonReturn(ctx, e.SUCCESS, "", nil)
}

func LoginOut(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Delete(util.SESSION_KEY_LOGIN_ADMIN)
	_ = session.Save()
	JsonReturn(ctx, e.SUCCESS, "", nil)
}
