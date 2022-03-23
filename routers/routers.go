package routers

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouters()  {
	gin.SetMode(utils.AppMod)
	r:=gin.New()
	r.Use(middleware.Logger())
	r.Use(middleware.Cors())
	r.Use(gin.Recovery())

	r.LoadHTMLGlob("static/admin/index.html")
	r.Static("admin/static","static/admin/static")
	r.StaticFile("admin/favicon.ico","static/admin/static/favicon.ico")
	r.GET("admin", func(c *gin.Context) {
		c.HTML(200,"index.html",nil)
	})

	Auth:=r.Group("api/v1")//需要鉴权的操作
	Auth.Use(middleware.JwtToken())
	{
		//user模块的路由接口

		Auth.PUT("users/:id",v1.EditUser)
		Auth.DELETE("users/:id",v1.DeleteUser)
		//category模块的路由接口
		Auth.POST("category/add",v1.AddCategory)
		Auth.PUT("category/:id",v1.EditCategory)
		Auth.DELETE("category/:id",v1.DeleteCategory)
		//article模块的路由接口
		Auth.POST("article/add",v1.AddArticle)
		Auth.PUT("article/:id",v1.EditArticle)
		Auth.DELETE("article/:id",v1.DeleteArticle)
		//上传文件
		Auth.POST("upload",v1.UpLoad)
		// 更新个人设置
		Auth.PUT("profile/:id", v1.UpdateProfile)
	}

	routerV1:=r.Group("api/v1")
	{
		routerV1.GET("user/:id",v1.GetUserInfo)
		routerV1.GET("users",v1.GetUsers)
		routerV1.GET("users/search",v1.SearchUser)
		routerV1.POST("user/add",v1.AddUser)//新增注册用户
		routerV1.GET("category",v1.GetCategory)
		routerV1.GET("category/:id",v1.GetCateInfo)
		routerV1.GET("article/search",v1.SearchArt)
		routerV1.GET("article",v1.GetArticle)////查询文章列表
		routerV1.GET("article/:id",v1.GetArtInfo)//获取单个文章
		routerV1.GET("article/list/:id",v1.GetCateArt)//查询单个分类下的文章
		routerV1.POST("login",v1.Login)
		// 获取个人设置信息
		routerV1.GET("profile/:id", v1.GetProfile)
	}

	r.Run(utils.HttpPort)

}