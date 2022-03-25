# ginblog
基于gin框架的博客

## 介绍

gin+vue 全栈制作一个博客。

这是一个参照教程制作的全栈项目，仅作为学习交流。
制作过程全程参考[B 站(https://space.bilibili.com/402177130)](https://space.bilibili.com/402177130) ，感谢原作者的无私奉献。



## 目录结构
### 后端源码
```shell
├─  .gitignore
│  go.mod // 项目依赖
│  go.sum
│  LICENSE
│  main.go //主程序
│  README.md
│  tree.txt
│          
├─api         
├─config // 项目配置入口   
├─database  // 数据库备份文件（初始化）
├─log  // 项目日志
├─middleware  // 中间件
├─model // 数据模型层
├─routes
│      router.go // 路由入口    
├─static // 打包静态文件
│  ├─admin  // 后台管理页面      
│  └─front  // 前端展示页面           
├─upload   
└─utils // 项目公用工具库
   │  setting.go 
   ├─errmsg   
   └─validator         
```
### 后端源码
储存于仓库:https://github.com/Mjiarong/ginWeb
```shell
ginWeb // 前端开发源码
  ├─admin  // 后台管理页面      
  ├─front  // 前端展示页面  
  ├─README.md  
  └─LICENSE
```



## 运行&&部署

1. 克隆项目

	```shell
	git clone git@github.com:Mjiarong/ginblog.git

2. 转到下面文件夹下

	cd yourPath/ginbolg


3. 安装依赖

```
go mod tidy
```

4. 初始化项目配置config.ini

```ini
./config/config.ini

[server]
AppMode = debug # debug 开发模式，release 生产模式
HttpPort = :3000 # 项目端口
JwtKey = das1da5sdas2d1a #JWT密钥，随机字符串即可

[database]
Db = mysql #数据库类型，不能变更为其他形式
DbHost = localhost # 数据库地址
DbPort = 3306 # 数据库端口
DbUser = root # 数据库用户名
DbPassWord = 123456 # 数据库用户密码
DbName = ginblog # 数据库名

[qiniu]
# 七牛储存信息
AccessKey = # your AK
SecretKey = # your SK
Bucket = 
QiniuSever =

[logger]
# 日志储存目录
LogFilePath = log/log.log
```

5. 在database中将sql文件导入数据库

	推荐navicat或者其他sql管理工具导入

6. 启动项目

```shell
 go run main.go
```



此时，项目启动，你可以访问页面

```shell
首页
http://localhost:3000
后台管理页面
http://localhost:3000/admin

默认管理员:root  密码:123456
```

## 实现功能

1.  简单的用户管理权限设置
2.  用户密码加密存储
3.  文章的分类和增删查改
4.  列表分页
5.  图片上传七牛云
6.  JWT 认证
7.  自定义日志和日志分页功能
8.  跨域 cors 设置


## 技术栈

- golang
  - Gin web framework
  - gorm
  - jwt-go
  - scrypt
  - logrus
  - gin-contrib/cors
  - go-playground/validator/v10
  - go-ini
- JavaScript
  - vue
  - vue cli
  - vue router
  - ant design vue
  - vuetify
  - axios
  - tinymce
- MySQL version:8.0.26

## 项目预览

- 前端展示页面
  ![](https://github.com/Mjiarong/ginblog/blob/main/upload/front1.jpg)

- 前端展示页面
  ![](https://github.com/Mjiarong/ginblog/blob/main/upload/front2.png)

- 后台登录页面

  ![](https://github.com/Mjiarong/ginblog/blob/main/upload/admin1.png)

- 后台管理页面

  ![](https://github.com/Mjiarong/ginblog/blob/main/upload/admin2.png)
