package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-gin-web/pkg/util"
	"go-gin-web/dao/cache"
	"go-gin-web/pkg/config"
	"go-gin-web/pkg/common"
	"go-gin-web/model"
	"net/http"
	"github.com/gin-gonic/gin"
	"go-gin-web/pkg/errMsg"
)

var (
	serverCfg = config.ServerCfg
)

func getUser(c *gin.Context) (string, error) {
	tokenString := c.Request.Header.Get("xAuthToken")
	if tokenString == "" {
		return "", errors.New("未登录")
	}

	userId, err := util.ParseToken(tokenString, []byte(serverCfg.JwtSecret))
	if err != nil {
		fmt.Println("=====", err)
		return "", err
	}

	// todo 优化
	RedisConn := model.RedisPool.Get()
	defer RedisConn.Close()

	// 优化TODO (封装到service中)
	loginUserKey := util.GetCacheKey("userLogin", userId)

	userBytes, err := cache.Get(loginUserKey)
	if err != nil {
		return "", errors.New("未登录")
	}

	var user model.User
	bytesErr := json.Unmarshal(userBytes, &user)
	if bytesErr != nil {
		return "", errors.New("未登录")
	}

	return userId, nil
}

// SigninRequired 必须是登录用户
func SigninRequired(c *gin.Context) {
	// 初始化
	var (
		res     = common.Res{C: c}
	)

	userId, err := getUser(c)
	if err != nil {
		res.SendJSON(http.StatusUnauthorized, errMsg.UNAUTHORIZED, err)

		c.Abort()
		return
	}

	c.Set("userId", userId)
	c.Next()
}
