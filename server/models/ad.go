package models

type Ad struct {
	Model
	AdName     string `json:"ad_name"`
	AdPicture  string `json:"ad_picture"`
	AdRedirect string `json:"ad_redirect"`
	AdPosition string `json:"ad_position"`
	AdStartAt  int    `json:"ad_start_at"`
	AdEndAt    int    `json:"ad_end_at"`
}

func GetAds(pageOffset int, pageSize int, maps interface{}) (ads []Ad) {
	db.Where(maps).Offset(pageOffset).Limit(pageSize).Find(&ads)
	return
}

func ExistAdById(id int) bool {
	var ad Ad
	db.Where("id = ?", id).First(&ad)
	if ad.ID < 1 {
		return false
	}
	return true
}

func ExistAdByName(adName string) bool {
	var ad Ad
	db.Where("ad_name = ?", adName).First(&ad)
	if ad.ID < 1 {
		return false
	}
	return true
}

func AddAd(ad Ad) bool {
	ormDb := db.Create(&ad)
	return isSuccess(ormDb)
}

func EditAd(id int, data interface{}) bool {
	ormDb := db.Model(&Ad{}).Where("id = ?", id).Updates(data)
	return isSuccess(ormDb)
}

func DeleteAd(id int) bool {
	ormDb := db.Where("id = ?", id).Delete(&Ad{})
	return isSuccess(ormDb)
}
