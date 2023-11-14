package models

type Website struct {
	Model
	WebsiteName        string `json:"website_name"`
	WebsiteDescription string `json:"website_description"`
}

func GetWebsites(pageOffset int, pageSize int) (websites []Website) {
	db.Offset(pageOffset).Limit(pageSize).Find(&websites)
	return
}

func ExistWebsiteById(id int) bool {
	var website Website
	db.Where("id = ?", id).First(&website)
	if website.ID < 1 {
		return false
	}
	return true
}

func EditWebsite(id int, data interface{}) bool {
	ormDb := db.Model(&Website{}).Where("id = ?", id).Updates(data)
	return isSuccess(ormDb)
}
