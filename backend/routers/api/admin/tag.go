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

func GetTags(ctx *gin.Context) {
	if !IsLogin(ctx) {
		JsonReturn(ctx, e.ERROR_NOT_LOGIN, "", nil)
		return
	}

	tagName := ctx.Query("tag_name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if tagName != "" {
		maps["tag_name"] = tagName
	}

	code := e.SUCCESS
	data["lists"] = models.GetTags(util.GetPageOffset(ctx), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	JsonReturn(ctx, code, "", data)
}

func AddTag(ctx *gin.Context) {
	if !IsLogin(ctx) {
		JsonReturn(ctx, e.ERROR_NOT_LOGIN, "", nil)
		return
	}

	tagName, _ := ctx.GetPostForm("tag_name")

	valid := validation.Validation{}
	valid.Required(tagName, "tag_name").Message("名称不能为空")
	valid.MaxSize(tagName, 100, "tag_name").Message("名称最长为100字符")

	code := e.SUCCESS
	msg := ""
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			code = e.INVALID_PARAMS
			msg = err.Message

			JsonReturn(ctx, code, msg, nil)
			return
		}
	}

	if !models.ExistTagByName(tagName) {
		if !models.AddTag(tagName) {
			code = e.ERROR
		}
	} else {
		code = e.ERROR_EXIST_TAG
	}

	JsonReturn(ctx, code, msg, nil)
}

//修改文章标签
func EditTag(ctx *gin.Context) {
	if !IsLogin(ctx) {
		JsonReturn(ctx, e.ERROR_NOT_LOGIN, "", nil)
		return
	}

	id := com.StrTo(ctx.Param("id")).MustInt()
	tagName, _ := ctx.GetPostForm("tag_name")

	valid := validation.Validation{}
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(tagName, "tag_name").Message("名称不能为空")
	valid.MaxSize(tagName, 100, "tag_name").Message("名称最长为20字符")

	code := e.SUCCESS
	msg := ""
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			code = e.INVALID_PARAMS
			msg = err.Message

			JsonReturn(ctx, code, msg, nil)
			return
		}
	}

	if models.ExistTagByID(id) {
		if models.ExistTagByName(tagName) {
			code = e.ERROR_EXIST_TAG
		} else {
			models.EditTag(id, tagName)
		}
	} else {
		code = e.ERROR_NOT_EXIST_TAG
	}

	JsonReturn(ctx, code, msg, nil)
}

//删除文章标签
func DeleteTag(ctx *gin.Context) {
	if !IsLogin(ctx) {
		JsonReturn(ctx, e.ERROR_NOT_LOGIN, "", nil)
		return
	}

	id := com.StrTo(ctx.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("请选择数据")

	code := e.SUCCESS
	msg := ""
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			code = e.INVALID_PARAMS
			msg = err.Message

			JsonReturn(ctx, code, msg, nil)
			return
		}
	}

	if models.ExistTagByID(id) {
		models.DeleteTag(id)
	} else {
		code = e.ERROR_NOT_EXIST_TAG
	}

	JsonReturn(ctx, code, msg, nil)
}
