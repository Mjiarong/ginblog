package v1

import (
	"fmt"
	"ginblog/model"
	"ginblog/utils/errmsg"
	"ginblog/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

//查询用户是否存在
func UserExist(c *gin.Context)  {

}

//添加用户
func AddUser(c *gin.Context)  {
	var data model.User
	_=c.ShouldBindJSON(&data)
	msg,code := validator.Validate(&data) //数据验证
	if code!=errmsg.SUCCESS{
		c.JSON(http.StatusOK,gin.H{
			"status":code,
			"message":msg,
		})
		return
	}

	if code=model.CheckUserByName(data.Username);code==errmsg.SUCCESS{
		model.CreateUser(&data)
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})

}

// GetUserInfo 查询单个用户
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var maps = make(map[string]interface{})
	data, code := model.GetUser(id)
	maps["username"] = data.Username
	maps["role"] = data.Role
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    maps,
			"total":   1,
			"message": errmsg.GetErrMsg(code),
		},
	)

}

//搜索用户
func SearchUser(c *gin.Context)  {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")

	data, total := model.SearchUser(username, pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)

}

//查询用户列表
func GetUsers(c *gin.Context)  {
	pageSize,_:=strconv.Atoi(c.Query("pagesize"))
	pageNum,_:=strconv.Atoi(c.Query("pagenum"))
	data, total :=model.GetUsers(pageSize,pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"total":total,
		"message":errmsg.GetErrMsg(code),
	})
}

//编辑用户
func EditUser(c *gin.Context)  {
	var data model.User
	id,_:=strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	if code=model.CheckUserByID(id);code==errmsg.SUCCESS{
		code=model.EditUser(id,&data)
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}

//删除用户
func DeleteUser(c *gin.Context)  {
	id,_:=strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}