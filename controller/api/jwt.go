package api

import (
	"errors"
	"fmt"
	"mingdeng/controller/common"
	"mingdeng/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
	//jwt.StandardClaims
}
type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("夏天夏天悄悄过去")

// GenToken 生成JWT
func GenToken(username string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		username, // 签发人
		//jwt.StandardClaims{
		//	ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
		//	Issuer:    "my-project",                               // 签发人
		//},
		jwt.RegisteredClaims{
			Issuer:    "my-project",
			Subject:   "sub",
			Audience:  []string{"a"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        "id",
		},
	}
	fmt.Printf("MyClaims:%v\n", c)
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
func AuthHandler(c *gin.Context) {
	// 用户发送用户名和密码过来
	var user UserInfo
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}
	admin, err := models.GetAAdminByWhere("username", user.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "用户信息错误",
		})
		return
	}
	// 校验用户名和密码是否正确
	fmt.Println(user.Password + "2008")
	//password md5 加密
	password := common.Md5V(common.Md5V(user.Password) + "2008")
	fmt.Println(password)
	if user.Username == admin.Username && password == admin.Password {
		// 生成Token
		tokenString, _ := GenToken(user.Username)
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"token": tokenString},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg":  "鉴权失败",
	})
	return
}
