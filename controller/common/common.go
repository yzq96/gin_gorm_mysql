package common

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonSuccess(c *gin.Context, msg string, data interface{}, total, num int) {
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"msg":    msg,
		"data":   data,
		"total":  total,
		"num":    num,
	})
}

func JsonError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{"status": 0, "msg": msg})
}

func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

