package router

import (
	"go-gin-web/handler/demo"
	"go-gin-web/handler/user"
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
		api.GET("/ping", demo.Ping)

		// 用户
		api.POST("/signIn", user.SignIn)
		api.POST("/signUp", user.SignUp)
		api.POST("/signOut", user.Signout)

		// 文章
		api.POST("/article", user.SignIn)
		api.POST("/signUp", user.SignUp)
		api.POST("/signOut", user.Signout)
	}
}
