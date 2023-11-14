package admin

import (
	"gin_admin/models"
	"gin_admin/pkg/e"
	"gin_admin/pkg/setting"
	"gin_admin/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetArticles(ctx *gin.Context) {
	if !IsLogin(ctx) {
		JsonReturn(ctx, e.ERROR_NOT_LOGIN, "", nil)
		return
	}

	maps := make(map[string]interface{})
	state := ctx.Query("state")
	title := ctx.Query("title")
	if state != "" {
		maps["title_state"] = com.StrTo(state).MustInt()
	}
	if title != "" {
		maps["title_title"] = title
	}

	articles := models.ArticlesList(util.GetPageOffset(ctx), setting.PageSize, maps)
	JsonReturn(ctx, e.SUCCESS, "", articles)
}
