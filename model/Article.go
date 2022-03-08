package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title string `gorm:"type:varchar(100);not null" json:"title"`
	Cid int  `gorm:"type:int;not null" json:"cid"`//文章分类id
	Desc string  `gorm:"type:varchar(200)" json:"desc"`//文章描述
	Content string `gorm:"type:longtext" json:"content"` //文章主题内容
	Img string  `gorm:"type:varchar(100)" json:"img"`//文章图片
}

//新增文章
func CreateArticle(data *Article)  int {
	err:=db.Create(&data).Error
	if err!=nil{
		return  errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询单个文章
func GetArtInfo(id int) (Article,int) {
	var art Article
	err = db.Preload("Category").Where("id=?",id).First(&art).Error
	if err!=nil{
		return  art,errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return art,errmsg.SUCCESS
}

//查询分类下所有文章
func GetCateArt(cid int,pageSize int,pageNum int) ([]Article,int){
	var cateArt []Article
	//var category []Category
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

/*	err=db.Select("name").Where("id=?",cid).First(&category).Error
	if err!=nil{
		return  cateArt,errmsg.ERROR_CATENAME_NOT_EXIST
	}*/


	err=db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid=?",cid).Take(&cateArt).Error
	if err!=nil{
		return  cateArt,errmsg.ERROR_ARTICLE_OF_CATE_NOT_EXIST
	}
	return cateArt,errmsg.SUCCESS
}


//查询文章列表
func GetArticle(pageSize int,pageNum int) ([]Article,int) {
	var art []Article
	if pageSize==0{pageSize=-1}//Cancel limit condition with -1
	//if pageNum==0{pageNum=-1}//Cancel offset condition with -1
	err=db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&art).Error
	if err!=nil&&err!=gorm.ErrRecordNotFound{
		return  nil,errmsg.ERROR
	}
	return art,errmsg.SUCCESS
}

//编辑文章
func EditArticle(id int,data *Article) int {
	err=db.Model(&Article{}).Select("Title","Cid","Desc","Content","Img").Where("id=?",id).Updates(data).Error
	if err!=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除文章
func DeleteArticle(id int)int {
	err=db.Where("id=?",id).Delete(&Article{}).Error
	if err!=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

