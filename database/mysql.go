package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mingdeng/setting"
)

var (
	DB *gorm.DB
)

func InitMySQL(cfg *setting.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Db)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("gorm.Open failed,err:%v\n", err)
		return
	}
	//fmt.Print(DB)
	//return
	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Printf("DB.DB() failed,err:%v\n", err)
		return
	}
	return sqlDB.Ping()
}
func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Printf("DB Close func,DB.DB() failed,err:%v\n", err)
		return
	}
	_ = sqlDB.Close()
}
