package database

import (
	"fmt"
	"mingdeng/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	/*Create*/
	// 迁移 在数据库中创建Users表
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
