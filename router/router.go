package router

import (
	"go-gin-web/handler/auth"
	"go-gin-web/handler/demo"
	"go-gin-web/middleware"
	"go-gin-web/pkg/config"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	var (
		serverCfg = config.ServerCfg
	)

	apiPrefix := serverCfg.APIPrefix

	api := router.Group(apiPrefix)
	{
		// 示例
		api.GET("/ping", middleware.SigninRequired, demo.Ping)

		// 认证
		api.POST("/signIn", auth.SignIn)
		api.POST("/signUp", auth.SignUp)
		api.POST("/signOut", middleware.SigninRequired, auth.SignOut)

		// AK认证
		api.POST("/secretKey/create", middleware.SigninRequired, auth.CreateSecretKey)
		api.POST("/secretKey/update", middleware.SigninRequired, auth.UpdateSecretKey)
		api.POST("/secretKey/forbidden", middleware.SigninRequired, auth.Forbidden)
		// api.GET("/secretKey/get", middleware.SigninRequired, auth.GetSecret)

		// webhook
		// api.POST("/webhook", middleware.SigninRequired, hook.HandleMsg)

		// 用户
		//api.POST("/user/create", auth.createUser)
		//api.POST("/user/update", auth.updateUser)
		//api.POST("/user/forbidden", user.forbidden)

		// 文章
		//api.POST("/article", auth.SignIn)
	}
}
