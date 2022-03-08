package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID uint `gorm:"primary;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

//查询分类是否存在
func CheckCategory(name string)  int {
	var info Category
	db.Select("id").Where("name=?",name).First(&info)
	if info.ID >0{
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

//新增分类
func CreateCategory(data *Category)  int {
	err:=db.Create(&data).Error
	if err!=nil{
		return  errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询分类列表
func GetCategory(pageSize int,pageNum int) []Category {
	var cate []Category
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	err=db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&cate).Error
	if err!=nil&&err!=gorm.ErrRecordNotFound{
		return  nil
	}
	return cate
}

//编辑分类
func EditCategory(id int,data *Category) int {
	err=db.Model(&Category{}).Select("Name").Where("id=?",id).Updates(data).Error
	if err!=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除分类
func DeleteCategory(id int)int {
	err=db.Where("id=?",id).Delete(&Category{}).Error
	if err!=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

