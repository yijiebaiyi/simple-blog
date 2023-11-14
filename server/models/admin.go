package models

import (
	_func "gin_admin/pkg/func"
	"gin_admin/pkg/setting"
)

const (
	ADMIN_STATE_ALLOW = 10
	ADMIN_STATE_FORBIDDEN = 20
)

type Admin struct {
	Model
	AdminName     string `json:"admin_name"`
	AdminPassword string `json:"admin_password"`
	AdminState    int    `json:"admin_state"`
	LastLoginTime int    `json:"last_login_time"`
	LastLoginIp   int    `json:"last_login_ip"`
}

func GetAdminByName(name string) (admin Admin) {
	db.Where("admin_name = ?", name).First(&admin)
	return
}

func GetAdminById(id int) (admin Admin) {
	db.Where("id = ?", id).First(&admin)
	return
}

func GetAdmins(pageOffset int, pageSize int, maps interface{}) (admins []Admin) {
	db.Where(maps).Select("id, admin_name, admin_state, last_login_time, last_login_ip").Offset(pageOffset).Limit(pageSize).Find(&admins)
	return
}

func ExistAdminById(id int) bool {
	var admin Admin
	db.Where("id = ?", id).First(&admin)
	if admin.ID < 1 {
		return false
	}
	return true
}

func ExistAdminByName(adminName string) bool {
	var admin Admin
	db.Where("admin_name = ?", adminName).First(&admin)
	if admin.ID < 1 {
		return false
	}
	return true
}

func GetPasswordOfMd5(password string) string {
	return _func.Md5Encode(password + setting.Md5Salt)
}

func AddAdmin(admin Admin) bool {
	ormDb := db.Create(&admin)
	return isSuccess(ormDb)
}

func EditAdmin(id int, data interface{}) bool {
	ormDb := db.Model(&Admin{}).Where("id = ?", id).Updates(data)
	return isSuccess(ormDb)
}
