package index

import (
	"gin_admin/models"
	"gin_admin/pkg/e"
	"gin_admin/pkg/setting"
	"gin_admin/pkg/util"

	"github.com/gin-gonic/gin"
)

func GetTags(ctx *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	data["lists"] = models.GetTags(util.GetPageOffset(ctx), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	JsonReturn(ctx, e.SUCCESS, "", data)
}
