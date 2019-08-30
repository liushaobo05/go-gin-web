package demo

import (
	"go-gin-web/pkg/common"
	"go-gin-web/pkg/errMsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	// 初始化变量
	var (
		res = common.Res{C: c}
	)

	// 数据返回
	data := gin.H{
		"message": "pong",
	}

	res.SendJSON(http.StatusOK, errMsg.SUCCESS, data)
}
