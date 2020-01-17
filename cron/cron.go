package cron

import (
	"go-gin-web/model"
	"go-gin-web/pkg/config"

	"github.com/robfig/cron"
)

var cronMap = map[string]func(){}

func init() {
	var serverCfg = config.ServerCfg

	if serverCfg.Env != model.DevelopmentMode {
		// 添加任务
		// cronMap["0 0 3 * * *"] = yesterdayCron
	} else {
		// go func() {
		// 	time.Sleep(1 * time.Second)
		// 	yesterdayCron()
		// }()
	}
}

// New 构造cron
func New() *cron.Cron {
	c := cron.New()
	for spec, cmd := range cronMap {
		c.AddFunc(spec, cmd)
	}
	return c
}

// // 自动刷新数据任务计划
// func CronStart() {

// 	// 如果配置文件没有开启就不跑
// 	// logs.Info("任务计划参数->", beego.AppConfig.String("crontab"))
// 	// if beego.AppConfig.String("crontab") != "true" {
// 	// 	return
// 	// }

// 	logs.Info("启动任务计划")
// 	cron := cron.New()
// 	// 应用缓存,每2分钟
// 	cron.AddFunc("1 */2 * * * ?", func() {
// 		app.CacheAppData()
// 	}, "CacheAppData")
// 	// 容器
// 	cron.AddFunc("6 * * * * ?", func() {
// 		app.MakeContainerData("")
// 	}, "MakeContainerData")
// 	// node状态写入到缓存
// 	cron.AddFunc("20 */3 * * * ?", func() {
// 		hosts.CronCache()
// 	}, "NodeStatusCache")
// 	// 服务数据写入到缓存
// 	cron.AddFunc("40 * * * * ?", func() {
// 		app.CronServiceCache()
// 	}, "CronServiceCache")
// 	// 仓库镜像写入缓存
// 	cron.AddFunc("1 */3 * * * ?", func() {
// 		registry.UpdateGroupImageInfo()
// 	}, "UpdateGroupImageInfo")
// 	// 集群数据写入缓存
// 	cron.AddFunc("1 */5 * * * ?", func() {
// 		cluster.CacheClusterData()
// 	}, "CacheClusterData")

// 	// 集群数据写入缓存
// 	cron.AddFunc("1 */1 * * * ?", func() {
// 		cluster.CacheClusterHealthData()
// 	}, "CacheClusterHealthData")

// 	// 监控自动扩容
// 	cron.AddFunc("*/30 * * * * ?", func() {
// 		monitor.CronAutoScale()
// 	}, "CronAutoScale")
// 	// 清除无效的job
// 	cron.AddFunc("1 */1 * * * ?", func() {
// 		ci.ClearJob()
// 	}, "ClearJob")
// 	cron.Start()
// }
