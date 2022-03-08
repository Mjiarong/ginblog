package v1

import (
	"fmt"
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//查询分类名是否存在

//添加分类
func AddCategory(c *gin.Context)  {
	var data model.Category
	_=c.ShouldBindJSON(&data)
	if code=model.CheckCategory(data.Name);code==errmsg.SUCCESS{
		model.CreateCategory(&data)
		c.JSON(http.StatusOK,gin.H{
			"status":code,
			"data":data,
			"message":errmsg.GetErrMsg(code),
		})
	}else if code==errmsg.ERROR_CATENAME_USED{
		c.JSON(http.StatusOK,gin.H{
			"status":code,
			"data":data,
			"message":errmsg.GetErrMsg(code),
		})
	}
}


//查询分类列表
func GetCategory(c *gin.Context)  {
	pageSize,_:=strconv.Atoi(c.Query("pagesize"))
	pageNum,_:=strconv.Atoi(c.Query("pagenum"))
	data:=model.GetCategory(pageSize,pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}

//编辑分类
func EditCategory(c *gin.Context)  {
	var data model.Category
	id,_:=strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	if code=model.CheckCategory(data.Name);code==errmsg.SUCCESS{
		model.EditCategory(id,&data)
	}else if code==errmsg.ERROR_CATENAME_USED{
		c.Abort()
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}

//删除分类
func DeleteCategory(c *gin.Context)  {
	id,_:=strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	code = model.DeleteCategory(id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}
