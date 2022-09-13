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

func GetAds(ctx *gin.Context) {
	adName := ctx.Query("name")
	adPosition := com.StrTo(ctx.Query("position")).MustInt()

	maps := make(map[string]interface{})
	if adName != "" {
		maps["ad_name"] = adName
	}
	if adPosition != 0 {
		maps["ad_position"] = adPosition
	}

	ads := models.GetAds(util.GetPageOffset(ctx), setting.PageSize, maps)
	JsonReturn(ctx, e.SUCCESS, "", ads)
}

func EditAd(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()
	adName, _ := ctx.GetPostForm("name")
	adPicture, _ := ctx.GetPostForm("picture")
	adRedirect, _ := ctx.GetPostForm("redirect")
	adPosition, _ := ctx.GetPostForm("position")
	adStartAtTmp, _ := ctx.GetPostForm("start_at")
	adEndAtTmp, _ := ctx.GetPostForm("end_at")
	adStartAt := com.StrTo(adStartAtTmp).MustInt()
	adEndAt := com.StrTo(adEndAtTmp).MustInt()

	var valid validation.Validation
	valid.Required(id, "id").Message("请选择数据")
	valid.Required(adName, "name").Message("请输入广告名称")
	valid.Required(adStartAt, "start_at").Message("请选择开始时间")
	valid.Required(adEndAt, "end_at").Message("请选择结束时间")
	valid.Required(adPosition, "position").Message("请选择广告位置")

	valid.MaxSize(adName, 20, "name").Message("广告名称不能超过20个字符")
	valid.MaxSize(adPicture, 2000, "picture").Message("图片路径不能超过2000个字符")
	valid.MaxSize(adRedirect, 2000, "redirect").Message("跳转地址不能超过2000个字符")
	valid.MaxSize(adPosition, 2000, "20").Message("广告位值不能超过20个字符")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			JsonReturn(ctx, e.INVALID_PARAMS, err.Message, nil)
			return
		}
	}

	// 校验数据是否存在
	if !models.ExistAdById(id) {
		JsonReturn(ctx, e.ERROR_NOT_EXIST_AD, "", nil)
		return
	}

	// 数据修改
	var ad models.Ad
	ad.AdName = adName
	ad.AdPicture = adPicture
	ad.AdRedirect = adRedirect
	ad.AdPosition = adPosition
	ad.AdStartAt = adStartAt
	ad.AdEndAt = adEndAt

	if models.EditAd(id, ad) {
		JsonReturn(ctx, e.SUCCESS, "", nil)
		return
	}
	JsonReturn(ctx, e.ERROR, "", nil)
}

func AddAd(ctx *gin.Context) {
	adName, _ := ctx.GetPostForm("name")
	adPicture, _ := ctx.GetPostForm("picture")
	adRedirect, _ := ctx.GetPostForm("redirect")
	adPosition, _ := ctx.GetPostForm("position")
	adStartAtTmp, _ := ctx.GetPostForm("start_at")
	adEndAtTmp, _ := ctx.GetPostForm("end_at")
	adStartAt := com.StrTo(adStartAtTmp).MustInt()
	adEndAt := com.StrTo(adEndAtTmp).MustInt()

	var valid validation.Validation
	valid.Required(adName, "name").Message("请输入广告名称")
	valid.Required(adStartAt, "start_at").Message("请选择开始时间")
	valid.Required(adEndAt, "end_at").Message("请选择结束时间")
	valid.Required(adPosition, "position").Message("请选择广告位置")

	valid.MaxSize(adName, 20, "name").Message("广告名称不能超过20个字符")
	valid.MaxSize(adPicture, 2000, "picture").Message("图片路径不能超过2000个字符")
	valid.MaxSize(adRedirect, 2000, "redirect").Message("跳转地址不能超过2000个字符")
	valid.MaxSize(adPosition, 2000, "20").Message("广告位值不能超过20个字符")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			JsonReturn(ctx, e.INVALID_PARAMS, err.Message, nil)
			return
		}
	}

	// 数据添加
	var ad models.Ad
	ad.AdName = adName
	ad.AdPicture = adPicture
	ad.AdRedirect = adRedirect
	ad.AdPosition = adPosition
	ad.AdStartAt = adStartAt
	ad.AdEndAt = adEndAt

	if models.AddAd(ad) {
		JsonReturn(ctx, e.SUCCESS, "", nil)
		return
	}
	JsonReturn(ctx, e.ERROR, "", nil)
}

func DeleteAd(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("请选择数据")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			JsonReturn(ctx, e.INVALID_PARAMS, err.Message, nil)
			return
		}
	}

	if models.ExistTagByID(id) {
		if models.DeleteAd(id) {
			JsonReturn(ctx, e.SUCCESS, "", nil)
		} else {
			JsonReturn(ctx, e.ERROR, "", nil)
		}
	} else {
		JsonReturn(ctx, e.ERROR_NOT_EXIST_AD, "", nil)
	}

}
