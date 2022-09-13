package admin

import (
	"gin_admin/models"
	"gin_admin/pkg/e"
	"gin_admin/pkg/setting"
	"gin_admin/pkg/util"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func Test(ctx *gin.Context) {
	JsonReturn(ctx, e.SUCCESS, "", make(map[string]interface{}))
}

func GetAdmins(ctx *gin.Context) {
	adminName := ctx.Query("name")
	adminState := com.StrTo(ctx.Query("state")).MustInt()

	maps := make(map[string]interface{})
	if adminName != "" {
		maps["admin_name"] = adminName
	}
	if adminState != 0 {
		maps["admin_state"] = adminState
	}

	admins := models.GetAdmins(util.GetPageOffset(ctx), setting.PageSize, maps)
	JsonReturn(ctx, e.SUCCESS, "", admins)
}

func EditAdmin(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()
	adminName, _ := ctx.GetPostForm("name")
	adminStateTmp, _ := ctx.GetPostForm("state")
	adminState := com.StrTo(adminStateTmp).MustInt()
	adminPassword, _ := ctx.GetPostForm("password")

	var valid validation.Validation
	valid.Required(id, "id").Message("请选择数据")
	valid.Required(adminName, "name").Message("请输入用户名")
	valid.Required(adminPassword, "password").Message("请输入密码")
	valid.Required(adminState, "state").Message("请选择状态")

	valid.MaxSize(adminName, 20, "name").Message("用户名不能超过20个字符")
	valid.MinSize(adminName, 2, "name").Message("用户名不能小于2个字符")
	valid.MaxSize(adminPassword, 18, "password").Message("密码不能超过18个字符")
	valid.MinSize(adminPassword, 3, "password").Message("密码不能小于3个字符")

	valid.AlphaDash(adminPassword, "password").Message("密码非法")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			JsonReturn(ctx, e.INVALID_PARAMS, err.Message, nil)
			return
		}
	}

	// 超管不允许修改
	if id == 1 {
		JsonReturn(ctx, e.ERROR_CAN_NOT_UPDATE, "", nil)
		return
	}

	// 校验数据是否存在
	admin := models.GetAdminById(id)
	if admin.ID < 1 {
		JsonReturn(ctx, e.ERROR_NOT_EXIST_ADMIN, "", nil)
		return
	}

	// 校验用户名是否存在
	if admin.AdminName != adminName {
		if models.ExistAdminByName(adminName) {
			JsonReturn(ctx, e.ERROR_EXIST_ADMIN, "", nil)
			return
		}
	}

	// 数据修改
	admin.AdminName = adminName
	admin.AdminState = adminState
	admin.AdminPassword = models.GetPasswordOfMd5(adminPassword)

	if models.EditAdmin(id, admin) {
		JsonReturn(ctx, e.SUCCESS, "", nil)
		return
	}
	JsonReturn(ctx, e.ERROR, "", nil)
}

func AddAdmin(ctx *gin.Context) {
	adminName, _ := ctx.GetPostForm("name")
	adminStateTmp, _ := ctx.GetPostForm("state")
	adminState := com.StrTo(adminStateTmp).MustInt()
	adminPassword, _ := ctx.GetPostForm("password")

	var valid validation.Validation
	valid.Required(adminName, "name").Message("请输入用户名")
	valid.Required(adminPassword, "password").Message("请输入密码")
	valid.Required(adminState, "state").Message("请选择状态")

	valid.MaxSize(adminName, 20, "name").Message("用户名不能超过20个字符")
	valid.MinSize(adminName, 2, "name").Message("用户名不能小于2个字符")
	valid.MaxSize(adminPassword, 18, "password").Message("密码不能超过18个字符")
	valid.MinSize(adminPassword, 3, "password").Message("密码不能小于3个字符")

	valid.AlphaDash(adminPassword, "password").Message("密码非法")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			JsonReturn(ctx, e.INVALID_PARAMS, err.Message, nil)
			return
		}
	}

	// 校验用户名是否存在
	if models.ExistAdminByName(adminName) {
		JsonReturn(ctx, e.ERROR_EXIST_ADMIN, "", nil)
		return
	}

	// 数据添加
	var admin models.Admin
	admin.AdminName = adminName
	admin.AdminState = adminState
	admin.AdminPassword = models.GetPasswordOfMd5(adminPassword)

	if models.AddAdmin(admin) {
		JsonReturn(ctx, e.SUCCESS, "", nil)
		return
	}
	JsonReturn(ctx, e.ERROR, "", nil)
}
