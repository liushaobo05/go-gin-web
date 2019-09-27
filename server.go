package main

import (
	"go-gin-web/model"
	"go-gin-web/pkg/config"
	"go-gin-web/router"

	"github.com/gin-gonic/gin"
)

func init() {
	// 配置

	// init config
	if err := config.Init("./conf/config.yaml"); err != nil {
		panic(err)
	}

	// init db
	model.Init()
}

func main() {
	var (
		serverCfg = config.ServerCfg
	)

	if serverCfg.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)

		gin.DisableConsoleColor()

		// 日志分割
		// logFile, err := os.OpenFile(config.ServerConfig.LogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		// if err != nil {
		// 	fmt.Printf(err.Error())
		// 	os.Exit(-1)
		// }
		// gin.DefaultWriter = io.MultiWriter(logFile)
	}

	// create router
	app := gin.New()

	// Global middleware
	app.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one
	app.Use(gin.Recovery())

	// 挂载路由
	router.RouterMount()

	// 加载
	router.RouterLoad(app)

	// 定时任务开启
	// todo

	// 启动服务
	app.Run(":" + serverCfg.Port)
}
