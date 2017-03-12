package main

import (
	"github.com/SuperGod/blog/controller"
	"github.com/SuperGod/blog/model"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	userAPI controller.UserController
	blogAPI controller.BlogController
)

func main() {
	model.Init("root:123456@tcp(172.17.0.1:3306)/blog_super?charset=utf8")
	web := gin.New()
	web.Use(gin.Recovery(), gin.Logger())
	store := sessions.NewCookieStore([]byte("secret"))
	web.Use(sessions.Sessions("mysession", store))
	userRouter := web.Group("/user")
	userRouter.POST("/login", userAPI.PostLogin)
	userRouter.GET("/login", userAPI.GetLogin)
	userRouter.GET("/info/:id", userAPI.Info)
	postRouter := web.Group("/blog")
	postRouter.GET("/content/:blogName", blogAPI.GetBlog)
	postRouter.GET("/list", blogAPI.GetBlogList)
	web.Run("127.0.0.1:8880")
}
