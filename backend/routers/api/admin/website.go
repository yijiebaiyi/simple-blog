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

func GetWebsites(ctx *gin.Context) {
	ads := models.GetWebsites(util.GetPageOffset(ctx), setting.PageSize)
	JsonReturn(ctx, e.SUCCESS, "", ads)
}

func EditWebsite(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()
	websiteName, _ := ctx.GetPostForm("name")
	websiteDescription, _ := ctx.GetPostForm("description")

	var valid validation.Validation
	valid.Required(id, "id").Message("请选择数据")
	valid.Required(websiteName, "name").Message("请输入网站名称")

	valid.MaxSize(websiteName, 20, "name").Message("网站名称不能超过20个字符")
	valid.MaxSize(websiteDescription, 500, "description").Message("网站描述不能超过500个字符")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			JsonReturn(ctx, e.INVALID_PARAMS, err.Message, nil)
			return
		}
	}

	// 校验数据是否存在
	if !models.ExistWebsiteById(id) {
		JsonReturn(ctx, e.ERROR_NOT_EXIST_AD, "", nil)
		return
	}

	// 数据修改
	var website models.Website
	website.WebsiteName = websiteName
	website.WebsiteDescription = websiteDescription

	if models.EditWebsite(id, website) {
		JsonReturn(ctx, e.SUCCESS, "", nil)
		return
	}
	JsonReturn(ctx, e.ERROR, "", nil)
}
