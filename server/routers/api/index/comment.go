package index

import (
	"gin_admin/models"
	"gin_admin/pkg/e"
	"gin_admin/pkg/setting"
	"gin_admin/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func CommentsListByArticleId(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("article_id")).MustInt()
	comments := models.CommentsListByArticleId(util.GetPageOffset(ctx), setting.PageSize, id)
	ctx.Next()
	ctx.Header("Access-Control-Allow-Origin", "*")
	JsonReturn(ctx, e.SUCCESS, "", comments)
}

func CommentsCreate(ctx *gin.Context) {

	requestData := make(map[string]interface{})
	ctx.BindJSON(&requestData)

	comment := models.Comment{}
	comment.Content = requestData["Content"].(string)
	comment.Name = requestData["Name"].(string)
	// comment.ArticleId = requestData["articleId"]

	models.CommentsCreate(comment)
	ctx.Next()
	ctx.Header("Access-Control-Allow-Origin", "*")

	JsonReturn(ctx, e.SUCCESS, "", nil)
}
