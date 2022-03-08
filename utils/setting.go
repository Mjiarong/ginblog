package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMod string
	HttpPort string
	JwtKey string

	DB string
	DbHost string
	DbPort string
	DbUser string
	DbPassWord string
	DbName string

	AccessKey  string
	SecretKey string
	Bucket  string
	QiniuServer  string

	LogFilePath string
)

func init() {
	file,err:=ini.Load("config/config.ini")
	if err!=nil{
		fmt.Println("配置文件读取错误:",err)
	}
	LoadServer(file)
	LoadData(file)
	LoadQiniu(file)
	LoadLogger(file)
}

func LoadServer(file *ini.File){
	AppMod=file.Section("server").Key("AppMode").MustString("debug")
	HttpPort=file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey=file.Section("server").Key("JwtKey").MustString("das1da5sdas2d1a")
}

func LoadData(file *ini.File){
	DB=file.Section("database").Key("DB").MustString("mysql")
	DbHost=file.Section("database").Key("DbHost").MustString("localhost")
	DbPort=file.Section("database").Key("DbPort").MustString("3306")
	DbUser=file.Section("database").Key("DbUser").MustString("root")
	DbPassWord=file.Section("database").Key("DbPassWord").MustString("123456")
	DbName=file.Section("database").Key("DbName").MustString("ginblog")
}

func LoadQiniu(file *ini.File){
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuServer= file.Section("qiniu").Key("QiniuServer").String()
}

func LoadLogger(file *ini.File){
	LogFilePath = file.Section("logger").Key("LogFilePath").MustString("log/log.log")
}
