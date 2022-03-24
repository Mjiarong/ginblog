package v1

import (
	"ginblog/middleware"
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login后台登录
func Login(c *gin.Context)  {
	var data model.User
	c.ShouldBindJSON(&data)
	var token string
	code:=model.CheckLogin(data.Username,data.Password)
	if code == errmsg.SUCCESS{
		token,code = middleware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
		"token":token,
	})
}

// LoginFront 前台登录
func LoginFront(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)

	data, code:=model.CheckLoginFront(data.Username, data.Password)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data.Username,
		"id":      data.ID,
		"message": errmsg.GetErrMsg(code),
	})
}
