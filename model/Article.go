package model

import (
	"fmt"
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

//搜索文章
func SearchArt(title string, pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var  total int64
	err = db.Select("article.id,title, img, article.created_at, article.updated_at, `desc`, Category.name").Order("id Asc").Joins("Category").Where("title LIKE ?",
		title+"%",
	).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
	//单独计数
	db.Model(&articleList).Where("title LIKE ?",
		title+"%",
	).Count(&total)
	fmt.Println(total)
	if err != nil {
		return nil, errmsg.ERROR,0
	}
	return articleList, errmsg.SUCCESS,total
}

//查询分类下所有文章
func GetCateArt(cid int,pageSize int,pageNum int) ([]Article,int, int64){
	var cateArt []Article
	var total int64
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


	err=db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid=?",cid).Find(&cateArt).Error
	db.Model(&cateArt).Where("cid =?", cid).Count(&total)
	if err!=nil{
		return  cateArt,errmsg.ERROR_ARTICLE_OF_CATE_NOT_EXIST,0
	}

	return cateArt,errmsg.SUCCESS,total
}


//查询文章列表
func GetArticle(pageSize int,pageNum int) ([]Article,int, int64) {
	var art []Article
	var total int64
	if pageSize==0{pageSize=-1}//Cancel limit condition with -1
	//if pageNum==0{pageNum=-1}//Cancel offset condition with -1
	err=db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&art).Error
	db.Model(&art).Count(&total)
	if err!=nil&&err!=gorm.ErrRecordNotFound{
		return  nil,errmsg.ERROR,0
	}
	return art,errmsg.SUCCESS,total
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

