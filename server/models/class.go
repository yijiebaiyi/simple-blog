package models

const (
	CLASS_STATE_ON  = 10
	CLASS_STATE_OFF = 20
)

type Class struct {
	Model
	ClassName  string `json:"class_name"`
	ClassState int    `json:"class_state"`
}

func GetClasses(pageOffset int, pageSize int, maps interface{}) (classes []Class) {
	db.Where(maps).Offset(pageOffset).Limit(pageSize).Find(&classes)
	return
}

func GetClassById(id int) (class Class) {
	db.Where("id = ?", id).First(&class)
	return
}

func ExistClassById(id int) bool {
	var class Class
	db.Where("id = ?", id).First(&class)
	if class.ID < 1 {
		return false
	}
	return true
}

func ExistClassByName(className string) bool {
	var class Class
	db.Where("class_name = ?", className).First(&class)
	if class.ID < 1 {
		return false
	}
	return true
}

func AddClass(class Class) bool {
	ormDb := db.Create(&class)
	return isSuccess(ormDb)
}

func EditClass(id int, data interface{}) bool {
	ormDb := db.Model(&Class{}).Where("id = ?", id).Updates(data)
	return isSuccess(ormDb)
}

func DeleteClass(id int) bool {
	ormDb := db.Where("id = ?", id).Delete(&Class{})
	return isSuccess(ormDb)
}
