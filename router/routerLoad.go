package router

import (
	"fmt"
	"go-gin-web/middleware"
	"go-gin-web/pkg/config"
	"go-gin-web/pkg/parse"

	"github.com/gin-gonic/gin"
)

type Vrouter struct {
	Method  string
	Path    string
	Handles []gin.HandlerFunc
}

var _ROUTER = []*Vrouter{}

func Router(pathRoot string, handle gin.HandlerFunc) {
	routerCfg, err := parse.LoadFile("./router/router.yaml")
	if err != nil {
		fmt.Println("路由配置文件加载错误", err)
	}

	api := new(API)
	err = routerCfg.GetPath(pathRoot).GetStruct(api)
	if err != nil {
		fmt.Println("error", err)
	}

	vrouter := &Vrouter{
		Method: api.Method,
		Path:   api.Path,
	}

	handles := []gin.HandlerFunc{}
	if api.Auth {
		handles = append(handles, middleware.SigninRequired)
	}

	handles = append(handles, handle)

	vrouter.Handles = handles
	_ROUTER = append(_ROUTER, vrouter)
}

func RouterLoad(router *gin.Engine) {
	var (
		serverCfg = config.ServerCfg
	)

	apiPrefix := serverCfg.APIPrefix
	api := router.Group(apiPrefix)

	for _, v := range _ROUTER {
		api.Handle(v.Method, v.Path, v.Handles...)
	}
}
