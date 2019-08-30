package middleware

import (
	"errors"
	"go-gin-web/pkg/util"
)

func getUser(c *gin.Context) (model.User, error) {
	var user model.User
	tokenString, cookieErr := c.Cookie("token")

	if cookieErr != nil {
		return user, errors.New("未登录")
	}

	if claims, err := util.ParseToken(tokenString, []byte(jwtSecret)); err != nil {
		userID := int(claims["id"].(float64))
		var err error
		user, err = model.UserFromRedis(userID)
		if err != nil {
			return user, errors.New("未登录")
		}
		return user, nil
	}

	return user, errors.New("未登录")
}

// SigninRequired 必须是登录用户
func SigninRequired(c *gin.Context) {
	// 初始化
	var (
		reqData = make(map[string]interface{}, 0)
		// resObj  siginRes
		resData gin.H
		req     = common.Req{C: c}
		res     = common.Res{C: c}
	)

	var user model.User
	var err error

	if user, err = getUser(c); err != nil {
		SendErrJSON("未登录", model.ErrorCode.LoginTimeout, c)
		return
	}
	c.Set("userId", userId)
	c.Next()
}
