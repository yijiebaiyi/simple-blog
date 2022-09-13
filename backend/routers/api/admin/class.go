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

func GetClasses(ctx *gin.Context) {
	className := ctx.Query("name")
	classState := com.StrTo(ctx.Query("state")).MustInt()
	maps := make(map[string]interface{})

	if className != "" {
		maps["class_name"] = className
	}
	if classState != 0 {
		maps["class_state"] = classState
	}

	classes := models.GetClasses(util.GetPageOffset(ctx), setting.PageSize, maps)
	JsonReturn(ctx, e.SUCCESS, "", classes)
}

func EditClass(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()
	className, _ := ctx.GetPostForm("name")
	classStateTmp, _ := ctx.GetPostForm("state")
	classState := com.StrTo(classStateTmp).MustInt()
	if classState == 0 {
		classState = models.CLASS_STATE_ON
	}

	var valid validation.Validation
	valid.Required(id, "id").Message("请选择数据")
	valid.Required(className, "name").Message("分栏名称必填")
	valid.Required(classState, "state").Message("分栏状态必填")
	valid.MaxSize(className, 20, "name").Message("分栏最大长度不能超过20")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			JsonReturn(ctx, e.INVALID_PARAMS, err.Message, nil)
			return
		}
	}

	class := models.GetClassById(id)
	if class.ID < 1 {
		JsonReturn(ctx, e.ERROR_NOT_EXIST_CLASS, "", nil)
		return
	}

	if class.ClassName != className && models.ExistClassByName(className) {
		JsonReturn(ctx, e.ERROR_EXIST_CLASS, "", nil)
		return
	}

	class.ClassName = className
	class.ClassState = classState
	if models.EditClass(id, class) {
		JsonReturn(ctx, e.SUCCESS, "", nil)
	} else {
		JsonReturn(ctx, e.ERROR, "", nil)
	}
}

func AddClass(ctx *gin.Context) {
	className, _ := ctx.GetPostForm("name")
	classStateTmp, _ := ctx.GetPostForm("state")
	classState := com.StrTo(classStateTmp).MustInt()
	if classState == 0 {
		classState = models.CLASS_STATE_ON
	}

	var valid validation.Validation
	valid.Required(className, "name").Message("分栏名称必填")
	valid.Required(classState, "state").Message("分栏状态必填")
	valid.MaxSize(className, 20, "name").Message("分栏最大长度不能超过20")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			JsonReturn(ctx, e.INVALID_PARAMS, err.Message, nil)
			return
		}
	}

	if models.ExistClassByName(className) {
		JsonReturn(ctx, e.ERROR_EXIST_CLASS, "", nil)
		return
	}

	var class models.Class
	class.ClassName = className
	class.ClassState = classState
	if models.AddClass(class) {
		JsonReturn(ctx, e.SUCCESS, "", nil)
	} else {
		JsonReturn(ctx, e.ERROR, "", nil)
	}
}

func DeleteClass(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()

	var valid validation.Validation
	valid.Required(id, "id").Message("请选择数据")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			JsonReturn(ctx, e.INVALID_PARAMS, err.Message, nil)
			return
		}
	}

	if !models.ExistClassById(id) {
		JsonReturn(ctx, e.ERROR_NOT_EXIST_CLASS, "", nil)
		return
	}

	if models.DeleteClass(id) {
		JsonReturn(ctx, e.SUCCESS, "", nil)
	} else {
		JsonReturn(ctx, e.ERROR, "", nil)
	}
}
