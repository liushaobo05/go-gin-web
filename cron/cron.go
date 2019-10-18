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
