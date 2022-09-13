package models

type Tag struct {
	Model
	Name string `json:"name"`
}

func GetTags(pageOffset int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageOffset).Limit(pageSize).Find(&tags)
	return tags
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func GetTagById(id int) (tag Tag) {
	db.Select("id, tag_name").Where("id = ?", id).First(&tag)
	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("tag_name = ?", name).First(&tag)

	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(name string) bool {
	ormDb := db.Create(&Tag{
		Name: name,
	})
	return isSuccess(ormDb)
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

func EditTag(id int, tagName string) bool {
	ormDb := db.Model(&Tag{}).Where("id = ?", id).Update("Name", tagName)
	return isSuccess(ormDb)
}
