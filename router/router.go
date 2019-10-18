package router

import (
	"go-gin-web/handler/auth"
	"go-gin-web/handler/demo"

	"github.com/gin-gonic/gin"
)

func Health(router *gin.Engine) {
	router.GET("/health", demo.Ping)
}

// 挂载路由
func RouterMount() {
	Router("demo.test", demo.Ping)
	Router("auth.sigin", auth.SignIn)
	Router("auth.sigup", auth.SignUp)
	Router("auth.sigout", auth.SignOut)
	Router("auth.userInfo", auth.GetUserInfo)
	Router("auth.createSecretKey", auth.CreateSecretKey)
	Router("auth.updateSecretKey", auth.UpdateSecretKey)
	Router("auth.forbiddenSecretKey", auth.ForbiddenSecretKey)
}
