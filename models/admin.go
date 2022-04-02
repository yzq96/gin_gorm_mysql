package models

import (
	"mingdeng/database"
)

type Admin struct {
	ID         uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Username   string `json:"username" gorm:"column:username" form:"username"`
	Password   string `json:"password" gorm:"column:password" form:"password"`
	Name       string `json:"name" gorm:"column:name" form:"name"`
	Mobile     string `json:"mobile" gorm:"column:mobile" form:"mobile"`
	CreateTime string `json:"create_time"`
}

/*Todo 增删改查*/

//CreateAAdmin 创建一个admin
func CreateAAdmin(admin Admin) (err error) {
	err = database.DB.Create(&admin).Error
	return
}

// 获取列表信息
func GetAllAdmin() (adminList []*Admin, err error) {
	adminList = make([]*Admin, 1, 15)
	err = database.DB.Find(&adminList).Error
	if err != nil {
		return nil, err
	}
	return
}

//GetAAdmin
func GetAAdmin(id string) (admin *Admin, err error) {
	admin = new(Admin)
	err = database.DB.Where("id = ?", id).First(admin).Error
	if err != nil {
		return nil, err
	}
	return
}

//GetAAdminByWhere
func GetAAdminByWhere(field, value string) (admin *Admin, err error) {
	admin = new(Admin)
	err = database.DB.Where(field+" = ?", value).First(admin).Error
	if err != nil {
		return nil, err
	}
	return
}

//UpdateAAdmin
func UpdateAAdmin(admin *Admin) (err error) {
	err = database.DB.Save(&admin).Error
	return
}

//DeleteAAdmin
func DeleteAAdmin(id string) (err error) {
	err = database.DB.Where("id = ?", id).Delete(&Admin{}).Error
	return
}
