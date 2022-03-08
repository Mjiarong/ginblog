package server

import (
	"context"
	"fmt"
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var ImgUrl = utils.QiniuServer


func UpLoadFile(file multipart.File,fileSize int64)(string ,int){
	putPolicy := storage.PutPolicy{//上传凭证
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	// 文件上传（表单方式）
	// 最简单的就是上传本地文件，直接指定文件的完整路径即可上传。

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 可选配置
	putExtra := storage.PutExtra{}

	//数据流上传（表单方式）
	//io.Reader对象的上传也是采用Put方法或者PutWithoutKey方法
	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "",errmsg.ERROR
	}
	url:=ImgUrl+ret.Key
	fmt.Println(ret.Key,ret.Hash)
	return url,errmsg.SUCCESS

}

