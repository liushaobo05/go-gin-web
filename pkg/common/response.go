package common

import (
	"go-gin-web/pkg/errMsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Res struct {
	C *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// 返回字符
func (res *Res) SendString(httpCode int, data string) {
	res.C.String(httpCode, data)

	return
}

// 返回json格式数据
func (res *Res) SendJSON(httpCode, errCode int, data interface{}) {
	res.C.JSON(httpCode, Response{
		Code:    errCode,
		Message: errMsg.GetMsg(errCode),
		Data:    data,
	})

	return
}

// 返回html页面
func (res *Res) SendHtml(httpCode int, tpl string, data interface{}) {
	res.C.HTML(http.StatusOK, tpl, data)

	return
}

// 重定向
func (res *Res) Redirect(nextURL string) {
	// 支持内部外部重定向
	res.C.Redirect(http.StatusMovedPermanently, nextURL)

	return
}
