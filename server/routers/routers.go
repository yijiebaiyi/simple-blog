package routers

import (
	"encoding/gob"
	"gin_admin/models"
	"gin_admin/pkg/setting"
	"gin_admin/routers/api/admin"
	"gin_admin/routers/api/index"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {

	Router = gin.New()

	Router.Use(gin.Logger())

	Router.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	sessionInit()

	//==登录接口
	Router.POST("/api/login", admin.Login)
	Router.GET("/api/login-out", admin.LoginOut)

	apiIndex := Router.Group("/api/index")
	{
		// 标签列表
		apiIndex.GET("/tags", index.GetTags)

		// 文章列表
		apiIndex.GET("/articles", index.ArticlesList)
		// 文章详情
		apiIndex.GET("/articles/:id", index.ArticlesDetail)
		// 文章评论列表
		apiIndex.GET("/comments/:article_id", index.CommentsListByArticleId)
		// 文章评论发表
		apiIndex.POST("/comments", index.CommentsCreate)
	}

	apiAdmin := Router.Group("/api/admin")
	{
		// 管理员：增删改查
		apiAdmin.GET("/admins", admin.GetAdmins)
		apiAdmin.POST("/admins", admin.AddAdmin)
		apiAdmin.PUT("/admins/:id", admin.EditAdmin)
		apiAdmin.GET("/test", admin.Test)

		// 网站信息：增删改查
		apiAdmin.GET("/websites", admin.GetWebsites)
		apiAdmin.PUT("/websites/:id", admin.EditWebsite)

		// 广告：增删改查
		apiAdmin.GET("/ads", admin.GetAds)
		apiAdmin.POST("/ads", admin.AddAd)
		apiAdmin.PUT("/ads/:id", admin.EditAd)
		apiAdmin.DELETE("/ads/:id", admin.DeleteAd)

		// 标签：增删改查
		apiAdmin.GET("/tags", admin.GetTags)
		apiAdmin.POST("/tags", admin.AddTag)
		apiAdmin.PUT("/tags/:id", admin.EditTag)
		apiAdmin.DELETE("/tags/:id", admin.DeleteTag)

		// 文章：增删改查
		apiAdmin.GET("/articles", admin.GetArticles)

		// 分栏：增删改查
		apiAdmin.GET("/classes", admin.GetClasses)
		apiAdmin.POST("/classes", admin.AddClass)
		apiAdmin.PUT("/classes/:id", admin.EditClass)
		apiAdmin.DELETE("/classes/:id", admin.DeleteClass)

		// 留言板：增删改查
		apiAdmin.GET("/message-boards", admin.GetMessageBoards)
		apiAdmin.DELETE("/message-boards/:id", admin.DeleteMessageBoard)
	}
}

func sessionInit() {
	store := sessions.NewCookieStore([]byte(setting.CookieSecret))
	Router.Use(sessions.Sessions(setting.SessionName, store))

	// 注册session结构体存储
	gob.Register(models.Admin{})
}
