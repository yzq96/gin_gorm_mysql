package models

import (
	"fmt"
	"mingdeng/database"

	"gorm.io/gorm"
)

type Company struct {
	Name string `json:"name" gorm:"column:name;type:string;size:50;unique" form:"name"`
	gorm.Model
}

func (Company) TableName() string {
	return "company"
}

func InitCompany() (err error) {
	err = database.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Company{})
	fmt.Println(err)
	return
}

func CreateACompany(company Company) (err error) {
	err = database.DB.Create(&company).Error
	return
}

func GetAllCompany() (companyList []*Company, err error) {
	companyList = make([]*Company, 1, 100)
	err = database.DB.Find(&companyList).Error
	if err != nil {
		return nil, err
	}
	return
}

func UpdateACompany1(id int64, company *Company) (err error) {
	err = database.DB.Where("id", id).Updates(&company).Error
	return
}
func UpdateACompany2(company *Company) (err error) {
	err = database.DB.Save(&company).Error
	return
}

func DeleteACompanyById(id string) (err error) {
	err = database.DB.Debug().Where("id", id).Delete(&Company{}).Error
	return
}
