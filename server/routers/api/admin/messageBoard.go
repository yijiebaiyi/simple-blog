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

func GetMessageBoards(ctx *gin.Context) {
	messageBoardUser := ctx.Query("user")
	messageBoardContacts := ctx.Query("contacts")

	maps := make(map[string]interface{})
	if messageBoardUser != "" {
		maps["message_board_user"] = messageBoardUser
	}
	if messageBoardContacts != "" {
		maps["message_board_contacts"] = messageBoardContacts
	}

	messageBoards := models.GetMessageBoards(util.GetPageOffset(ctx), setting.PageSize, maps)
	JsonReturn(ctx, e.SUCCESS, "", messageBoards)
}


func DeleteMessageBoard(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("请选择数据")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			JsonReturn(ctx, e.INVALID_PARAMS, err.Message, nil)
			return
		}
	}

	models.DeleteMessageBoard(id)
	JsonReturn(ctx, e.SUCCESS, "", nil)

}
