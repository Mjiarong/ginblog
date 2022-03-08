package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role int `gorm:"type:int,DEFAULT:2" json:"role" validate:"required,gte=2" label:"用户角色编号"`
}

//查询用户是否存在
func CheckUserByName(username string)  int {
	var info User
	db.Select("id").Where("username=?",username).First(&info)
	if info.ID >0{
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}


func CheckUserByID(id int)  int {
	var info User
	err:=db.Select("username").Where("id=?",id).First(&info).Error
	if err!=nil&&err==gorm.ErrRecordNotFound{
		return errmsg.ERROR_USERNAME_NOT_EXIST
	}
	return errmsg.SUCCESS
}

func CreateUser(data *User)  int {
	err:=db.Create(&data).Error
	if err!=nil{
		return  errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询用户列表
func GetUsers(pageSize int,pageNum int) []User {
	var users []User
	if pageSize==0{pageSize=-1}//Cancel limit condition with -1
	//if pageNum==0{pageNum=-1}//Cancel offset condition with -1
	err=db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&users).Error
	if err!=nil&&err!=gorm.ErrRecordNotFound{
		return  nil
	}
	return users
}

//编辑用户
func EditUser(id int,data *User) int {
/*	var tmap = make(map[string]interface{})
	tmap["username"] = data.Username
	tmap["role"] = data.Role*/
	//err=db.Model(&User{}).Where("id=?",id).Update(tmap).Error
	err=db.Model(&User{}).Select("Username","Role").Where("id=?",id).Updates(data).Error
	if err!=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除用户
func DeleteUser(id int)int {
	err=db.Where("id=?",id).Delete(&User{}).Error
	if err!=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//密码加密
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {//hook
	u.Password=EncPassword(u.Password)
	return
}

func EncPassword(password string)string{
	const PwHashByte=10
	salt:=make([]byte,8)
	salt = []byte{0xc8, 0x28, 0xf2, 0x58, 0xa7, 0x6a, 0xad, 0x7b}

	HashPw,err:=scrypt.Key([]byte(password),salt,1<<14,8,1,10)
	if err!=nil{
		log.Fatal(err)
	}
	fpw:=base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

//登录验证
func CheckLogin(username string,password string) int {
	var info User
	db.Where("username=?",username).First(&info)
	if info.ID==0{
		return errmsg.ERROR_USERNAME_NOT_EXIST
	}
	if EncPassword(password) != info.Password{
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if info.Role != 1{
		return errmsg.ERROR_USER_NO_PERMISSION
	}
	return errmsg.SUCCESS
}