package hook

import "net/http"

func HandleMsg(c *gin.Context) {
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
