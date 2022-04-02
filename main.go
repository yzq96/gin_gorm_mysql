package main

import (
	"fmt"
	"mingdeng/database"
	"mingdeng/routers"
	"mingdeng/setting"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./main.go conf/config.ini")
		return
	}
	//加载配置文件
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed,err:%v\n", err)
		return
	}

	//创建数据库
	//sql：CREATE DATABASE bubble;
	//连接数据库
	err := database.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed,err:%v\n", err)
		return
	}
	defer database.Close() // defer 在程序最后关闭数据库连接
	//模型绑定
	//err = database.DB.AutoMigrate(&models.Todo{})
	//if err != nil {
	//	fmt.Printf("AutoMigrate model failed,err:%v\n", err)
	//	return
	//}
	//TODO 注册路由
	r := routers.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
