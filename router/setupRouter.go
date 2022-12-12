package router

import (
	"go_forum/controller"
	"go_forum/middleware"

	_ "go_forum/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(r *gin.Engine) {
	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	apiV1 := r.Group("/api/v1")

	user := apiV1.Group("/user") // /api/v1
	{
		user.POST("/register", controller.RegisterHandler)
		user.POST("/login", controller.LoginHandler)
	}

	{
		//社区
		apiV1.GET("/community", controller.CommunityHandler)
		apiV1.GET("/community/:id", controller.CommunityDetailHandler) //  /community/:id  /1  uri参数

		// 帖子
		apiV1.POST("/note", controller.CreatePostHandler)
		apiV1.GET("/note/:id", controller.GetPostDetailHandler) //帖子id
		apiV1.GET("/notelist", controller.GetPostListHandler)
		apiV1.GET("/post2", controller.GetPostListHandler2)

	}

	apiV1.Use(middleware.JWTMiddleware())
	apiV1.GET("/vote", controller.PostVoteController)

	//设置为发布gin.SetMode(gin.ReleaseMode),默认debug模式,终端信息输出 debug test release
}
