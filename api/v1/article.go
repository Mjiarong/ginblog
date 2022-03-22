package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加文章
func AddArticle(c *gin.Context)  {
	var data model.Article
	_=c.ShouldBindJSON(&data)
	code := model.CreateArticle(&data)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}

//搜索文章
func SearchArt(c *gin.Context)  {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	title := c.Query("title")

	data, code ,total := model.SearchArt(title, pageSize, pageNum)
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)

}

//查询单个分类下的文章
func GetCateArt(c *gin.Context)  {
	pageSize,_:=strconv.Atoi(c.Query("pagesize"))
	pageNum,_:=strconv.Atoi(c.Query("pagenum"))
	id,_:=strconv.Atoi(c.Param("id"))
	data,code, total:=model.GetCateArt(id,pageSize,pageNum)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"total":total,
		"message":errmsg.GetErrMsg(code),
	})
}

//查询单个文章
func GetArtInfo(c *gin.Context)  {
	id,_:=strconv.Atoi(c.Param("id"))
	data,code:=model.GetArtInfo(id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}

//查询文章列表
func GetArticle(c *gin.Context)  {
	pageSize,_:=strconv.Atoi(c.Query("pagesize"))
	pageNum,_:=strconv.Atoi(c.Query("pagenum"))
	data,code, total:=model.GetArticle(pageSize,pageNum)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"total":total,
		"message":errmsg.GetErrMsg(code),
	})
}

//编辑文章
func EditArticle(c *gin.Context)  {
	var data model.Article
	id,_:=strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code:=model.EditArticle(id,&data)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}

//删除文章
func DeleteArticle(c *gin.Context)  {
	id,_:=strconv.Atoi(c.Param("id"))
	code = model.DeleteArticle(id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}
