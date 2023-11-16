package index

import (
	"gin_admin/middleware/httpresponse"
	"gin_admin/models"
	"gin_admin/pkg/setting"
	"gin_admin/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

type CommentsCreateRequest struct {
	Content   string `json:"content"`
	Url       string `json:"url"`
	Email     string `json:"email"`
	Tel       string `json:"tel"`
	Name      string `json:"name" binding:"required"`
	ArticleId int    `json:"articleId"`
}

func CommentsListByArticleId(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("article_id")).MustInt()
	comments := models.CommentsListByArticleId(util.GetPageOffset(ctx), setting.PageSize, id)
	httpresponse.SuccessResponse(ctx, comments)
}

func CommentsCreate(ctx *gin.Context) {
	var err error
	var requestData CommentsCreateRequest
	err = ctx.BindJSON(&requestData)
	if err != nil {
		httpresponse.ErrorResponse(ctx, httpresponse.CODE_BAD_REQUEST, err)
		return
	}

	comment := models.Comment{}
	comment.Content = requestData.Content
	comment.Name = requestData.Name
	comment.Tel = requestData.Tel
	comment.Email = requestData.Email
	comment.Url = requestData.Url
	comment.ArticleId = requestData.ArticleId
	comment.State = models.CommentStateEnum_Enable

	models.CommentsCreate(comment)
	httpresponse.SuccessResponse(ctx)
}
