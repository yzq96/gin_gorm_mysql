package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mingdeng/controller/common"
	"mingdeng/models"
	"net/http"
	"strconv"
)

func ComInit(c *gin.Context) {
	models.InitCompany()
	return
}

func CreateACompany(c *gin.Context) {
	var company models.Company
	err := c.Bind(&company)
	//c.Bind()
	if err != nil {
		fmt.Println(err)
		common.JsonError(c, "参数获取失败")
		return
	}
	err = models.CreateACompany(company)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "创建失败", "error": err})
		common.JsonError(c, "创建失败")
		return
	}
	common.JsonSuccess(c, "创建成功", []string{}, 0, 0)
}

func GetCompanyList(c *gin.Context) {
	list, err := models.GetAllCompany()
	if err != nil {
		common.JsonError(c, "获取列表失败")
		return
	}
	common.JsonSuccess(c, "success", list, 0, 0)
}

func UpdateCompanyById1(c *gin.Context) {
	//id := c.PostForm("id")//form-data
	id, _ := c.Params.Get("id") //query
	idInt64, err := strconv.ParseInt(id, 10, 64)
	var company models.Company
	err = c.ShouldBind(&company)
	if err != nil {
		common.JsonError(c, "参数错误")
		return
	}
	fmt.Printf("UpdateCompanyById1 func,id:%v\n", id)
	err = models.UpdateACompany1(idInt64, &company)
	if err != nil {
		fmt.Printf("UpdateACompany err :%v\n", err)
		common.JsonError(c, "修改失败")
		return
	}
	common.JsonSuccess(c, "suc", company, 0, 0)
}

func UpdateCompanyById2(c *gin.Context) {
	var company models.Company
	err := c.ShouldBind(&company)
	if err != nil {
		fmt.Printf("ShouldBind err:%v\n", err)
		common.JsonError(c, "参数错误")
		return
	}
	fmt.Printf("UpdateCompanyById2 companyInfo:%v\n", company)
	err = models.UpdateACompany2(&company)
	if err != nil {
		fmt.Printf("UpdateACompany err :%v\n", err)
		common.JsonError(c, "修改失败")
		return
	}
	common.JsonSuccess(c, "suc", company, 0, 0)
}
func DeleteACompany(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		common.JsonError(c, "无效id")
		return
	}
	err := models.DeleteACompanyById(id)
	if err != nil {
		fmt.Printf("DeleteACompany err :%v\n", err)
		common.JsonError(c, "删除失败")
		return
	}
	common.JsonSuccess(c, "删除成功", &models.Company{}, 0, 0)
}
