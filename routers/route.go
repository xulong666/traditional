package router

import (
	"traditional/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"traditional/middleware"
)

var r *gin.Engine

func SetupRouter() {
	r = gin.Default()
	r .Use(cors.Default())

	user := r.Group("/user")
	{
		user.POST("/register", models.Register)
		user.POST("/login", models.Login)
		user.GET("/info", middleware.AuthMiddleware, models.Getuserinfo)
		user.GET("/myarticle", middleware.AuthMiddleware, models.MyArticle) //获取自己的文章列表
	}

	article := r.Group("/article")
	{
		article.GET("/index", models.GetAllArticle)     //首页获取所有文章列表
		article.POST("/upload", middleware.AuthMiddleware, models.ArticleUpload) //上传文章
		article.GET("/getarticke/:user", models.GetarticleOfUser) //获取某用户的文章列表
	//	article.GET("/:articleid", models.Getarticle)  //获取某篇文章内容

	}

}

func Run() {
	r.Run(":8000")
}
