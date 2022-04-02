package routers

import (
	"github.com/gin-gonic/gin"
	"mingdeng/controller/api"
	"mingdeng/middleware"
	"mingdeng/setting"
)

func SetupRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	//r.Static("/static", "static")
	//r.LoadHTMLGlob("templates/*")
	//r.GET("/", api.IndexHandler)
	r.POST("/auth", api.AuthHandler)

	v1Group := r.Group("v1", middleware.JWTAuthMiddleware())
	{
		// 用户
		// 添加
		v1Group.POST("/admin", api.CreateAdmin)
		// 查看所有的用户
		v1Group.GET("/admin", api.GetAdminList)
		// 修改某一个用户
		v1Group.PUT("/admin/:id", api.UpdateAAdmin)
		// 删除某一个用户
		v1Group.DELETE("/admin/:id", api.DeleteAAdmin)

		/*公司*/
		//初始化公司表
		v1Group.POST("/company/init", api.ComInit)
		//获取公司列表
		v1Group.GET("/company", api.GetCompanyList)
		//创建公司
		v1Group.POST("/company", api.CreateACompany)
		//修改公司1
		v1Group.PUT("/company/:id", api.UpdateCompanyById1)
		//修改公司2
		v1Group.PATCH("/company", api.UpdateCompanyById2)
		//删除公司
		v1Group.DELETE("/company/:id", api.DeleteACompany)
	}
	return r
}
