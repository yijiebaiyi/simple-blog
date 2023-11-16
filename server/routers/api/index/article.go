package index

import (
	"gin_admin/middleware/httpresponse"
	"gin_admin/models"
	"gin_admin/pkg/setting"
	"gin_admin/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func ArticlesList(ctx *gin.Context) {
	maps := make(map[string]interface{})
	articles := models.ArticlesList(util.GetPageOffset(ctx), setting.PageSize, maps)
	httpresponse.SuccessResponse(ctx, articles)
}

func ArticlesDetail(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()
	article := models.ArticlesDetail(id)
	httpresponse.SuccessResponse(ctx, article)
}
