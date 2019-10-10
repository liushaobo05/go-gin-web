package main

import (
	"errors"
	"fmt"
	"go-gin-web/middleware"
	"go-gin-web/model"
	"go-gin-web/pkg/config"
	"go-gin-web/pkg/httputil"
	"go-gin-web/router"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	// init config
	if err := config.Init("./conf/config.yaml"); err != nil {
		log.Panic("load config failed", err)
	}

	// init db
	model.Init()
}

func main() {
	var (
		serverCfg = config.ServerCfg
		domain    = serverCfg.Domain
	)

	if serverCfg.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}

	// create router
	app := gin.New()

	// Global middleware
	// Recovery middleware recovers from any panics and writes a 500 if there was one
	app.Use(gin.Recovery())

	app.Use(middleware.RequestId())

	// 控制台日志中间件
	app.Use(middleware.LoggerToConsole())

	// 日志文件中间件
	app.Use(middleware.LoggerToConsole())

	// 请求日志落库

	// 挂载路由
	router.RouterMount()

	// 加载
	router.Health(app)

	router.RouterLoad(app)

	// 自动化测试
	go func() {
		TestUnit(app)
	}()

	// 健康检查
	go func() {
		if err := pingServer(domain); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	// 定时任务开启
	// todo

	// 启动服务
	app.Run(":" + serverCfg.Port)
}

// 检查服务健康
func pingServer(domain string) error {
	for i := 0; i < 4; i++ {
		resp, err := http.Get(domain + "/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("server not running")
}

func TestUnit(app *gin.Engine) error {
	fmt.Println("执行自动化测试")
	statusCode, _ := httputil.Get("/api/v1/ping", app)
	if statusCode == 200 {
		return nil
	}

	return errors.New("unit test fail")
}
