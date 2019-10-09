package router

import (
	"go-gin-web/handler/auth"
	"go-gin-web/handler/demo"
	"go-gin-web/pkg/config"

	"github.com/gin-gonic/gin"
)

func RouterV1(router *gin.Engine) {
	var (
		serverCfg = config.ServerCfg
	)

	apiPrefix := serverCfg.APIPrefix

	api := router.Group(apiPrefix)
	{
		// api.POST("/secretKey/createSecretKey", middleware.SigninRequired, auth.CreateSecretKey)
		// api.POST("/secretKey/UpdateSecretKey", middleware.SigninRequired, auth.UpdateSecretKey)
		// api.POST("/secretKey/ForbiddenSecretKey", middleware.SigninRequired, auth.ForbiddenSecretKey)
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

// 挂载路由
func RouterMount() {
	Router("demo.test", demo.Ping)
	Router("auth.sigin", auth.SignIn)
	Router("auth.sigup", auth.SignUp)
	Router("auth.sigout", auth.SignOut)
	Router("auth.createSecretKey", auth.CreateSecretKey)
	Router("auth.updateSecretKey", auth.UpdateSecretKey)
	Router("auth.forbiddenSecretKey", auth.ForbiddenSecretKey)
}
