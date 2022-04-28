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
  

##  Docker部署

### 一、如何安装docker
以ubuntu 18.04 LTS为例
https://blog.csdn.net/liangcsdn111/article/details/115405223

### 二、拉取镜像和创建镜像和容器编排
### Mysql服务器的镜像
```shell
#首先确定mysql是否能被搜素到，这步可以跳过，也可以在dockerhub.com中搜索
$ docker search mysql

#拉取镜像
docker pull mysql  #这里默认是拉取的最新版本，如果需要特定版本可以在镜像后面添加tag，具体版本信息可以在dockerhub.com查询

#特定版本拉取,比如要拉取8.0.22(版本号一定要是官方放出的版本号，否则是查找不到的)
docker pull mysql:8.0.22

#这时可以查看下拉取的镜像
docker images

#运行镜像
docker run -d -p 3306:3306 -v /my/own/datadir:/var/lib/mysql --name ginblog-mysql -e MYSQL_ROOT_PASSWORD=admin123  mysql

# -d 表示后台运行，并返回容器id
# -p 3006:3306 表示端口映射，具体为 -p 主机端口：容器端口
# --name 给容器取个名字
# -e MYSQL_ROOT_PASSWORD=password 给mysql root管理员设置密码
# -v /my/own/datadir:/var/lib/mysql 添加数据卷/my/own/datadir是主机的数据库路径 /var/lib/mysql是容器中的数据库路径，这一步非常重要

#进入容器配置
docker exec -it ginblog-mysql bash

root@ed9345077e02:/# mysql -u root -p
Enter password:
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 8
Server version: 8.0.22 MySQL Community Server - GPL
Copyright (c) 2000, 2020, Oracle and/or its affiliates. All rights reserved.
Oracle is a registered trademark of Oracle Corporation and/or its affiliates.
Other names may be trademarks of their respective owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql>

# 之后就和一般情况下mysql的操作一样了。
```



### 制作ginblog项目镜像

- 首相要拉取我们的ginblog项目

```shell
# 新建一个项目文件夹，在你认为任何适合的地方都可以

$ cd /
$ mkdir app

# 我们这里利用git来远程同步

$ git clone 项目地址
```

- 编写Dockerfile

```dockerfile
FROM golang:latest
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

WORKDIR $GOPATH/src/ginblog
COPY . $GOPATH/src/ginblog

RUN go build .

EXPOSE 3000

ENTRYPOINT ["./ginblog"]
```

- 配置ginblog的config

```ini
# config/config.ini

# DbHost = ginblog-mysql 是为了后面容器互通做准备，对应的是mysql容器的name

Db = mysql
DbHost = ginblog-mysql 
DbPort = 3306
DbUser = ginblog
DbPassWord = admin123
DbName = ginblog
```

### 生成镜像

最后一步，就是生成我们的ginblog docker image了，这部很简单，运行下列命令

```shell
$ docker build -t ginblog .
$ docker run -d -p 3000:3000 --name ginblog ginblog

#这样访问服务器IP:3000 就可以访问网站了
```
