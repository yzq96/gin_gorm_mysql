package api

import (
	"fmt"
	"mingdeng/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//func IndexHandler(c *gin.Context) {
//	c.HTML(http.StatusOK, "index.html", nil)
//}
func AdminInit(c *gin.Context) {
	models.InitAdmin()
	return
}

func CreateAdmin(c *gin.Context) {
	//1.1 创建一个用于接收参数的todo结构体
	var admin models.Admin
	//1.2 接收整合前端传来的参数 gin自带的好方法 c.BindJSON()
	//err := c.ShouldBind(&admin)
	err := c.ShouldBindJSON(&admin)
	if err != nil {
		fmt.Printf("c.BindJSON failed,err:%v\n", err)
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	//2. 存入数据库
	err = models.CreateAAdmin(admin)
	if err != nil {
		fmt.Printf("c.BindJSON failed,err:%v\n", err)
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, admin)
	}
}
func GetAdminList(c *gin.Context) {
	adminList, err := models.GetAllAdmin()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, adminList)
	}
}
func UpdateAAdmin(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
	}
	admin, err := models.GetAAdmin(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	//???????????????
	err = c.BindJSON(&admin)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	err = models.UpdateAAdmin(admin)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "更新失败"})
	} else {
		c.JSON(http.StatusOK, admin)
	}
}
func DeleteAAdmin(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
	}
	err := models.DeleteAAdmin(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"id": "delete"})
	}
}
