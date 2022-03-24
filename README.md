# ginblog
基于gin框架的博客

介绍
gin+vue 全栈制作一个博客。

这是一个分享全栈制作过程的项目，旨在为有兴趣接触 golang web 开发的朋友分享一些制作经验。

你可以前往 B 站(https://space.bilibili.com/402177130) 观看全栈的制作过程，你也可以留言分享你的观点，非常乐意与你交流。

运行&&部署
克隆项目

git clone git@gitee.com:wejectchan/ginblog.git
or
git clone https://github.com/wejectchen/Ginblog.git
转到下面文件夹下

cd yourPath/ginbolg

安装依赖

go mod tidy
初始化项目配置config.ini
./config/config.ini

[server]
AppMode = debug # debug 开发模式，release 生产模式
HttpPort = :3000 # 项目端口
JwtKey = 89js82js72 #JWT密钥，随机字符串即可

[database]
Db = mysql #数据库类型，不能变更为其他形式
DbHost = 127.0.0.1 # 数据库地址
DbPort = 3306 # 数据库端口
DbUser = ginblog # 数据库用户名
DbPassWord = admin123 # 数据库用户密码
DbName = ginblog # 数据库名

[qiniu]
# 七牛储存信息
AccessKey = # AK
SecretKey = # SK
Bucket = 
QiniuSever =
在database中将sql文件导入数据库

推荐navicat或者其他sql管理工具导入

启动项目

 go run main.go
此时，项目启动，你可以访问页面

首页
http://localhost:3000
后台管理页面
http://localhost:3000/admin

默认管理员:admin  密码:123456
