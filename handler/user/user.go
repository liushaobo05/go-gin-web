package user

import (
	"encoding/json"
	"fmt"
	"go-gin-web/dao/userDao"
	"go-gin-web/model"
	"go-gin-web/pkg/common"
	"go-gin-web/pkg/config"
	"go-gin-web/pkg/errMsg"
	"go-gin-web/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	serverCfg = config.ServerCfg
)

// 用户登陆
func SignIn(c *gin.Context) {
	// 初始化
	var (
		reqData = make(map[string]interface{}, 0)
		// resObj  siginRes
		resData gin.H
		req     = common.Req{C: c}
		res     = common.Res{C: c}
	)

	// 数据解析
	if err := req.ParseParams(reqData); err != nil {
		resData = gin.H{
			"username": reqData["username"].(string),
		}

		fmt.Println("参数解析", err)

		res.SendJSON(http.StatusInternalServerError, errMsg.ERROR, resData)
		return
	}

	// 获取请求数据及数据校验
	// errs := req.Parse(reqData)
	// switch errData := errs.(type) {
	// case map[string]string:
	// 	if len(errData) != 0 {
	// 		res.SendJSON(http.StatusBadRequest, errMsg.INVALID_PARAMS, errData)
	// 		return
	// 	}
	// default:
	// 	res.SendJSON(http.StatusInternalServerError, errMsg.ERROR, errData)
	// 	return
	// }

	// 获取用户信息
	var params = map[string]interface{}{
		"username": reqData["username"].(string),
	}

	user, err := userDao.GetUser(params)
	if err != nil {
		resData = gin.H{
			"username": reqData["username"],
		}

		res.SendJSON(http.StatusUnauthorized, errMsg.UNAUTHORIZED, resData)
		return
	}

	// checkpassword
	if err := util.Compare(user.PassHash, reqData["password"].(string)); err != nil {
		resData = gin.H{
			"id":       user.Id,
			"username": user.Username,
		}

		res.SendJSON(http.StatusUnauthorized, errMsg.UNAUTHORIZED, resData)
		return
	}

	// 判断是否被禁
	if user.Status == 1 {
		resData = gin.H{
			"id":       user.Id,
			"username": user.Username,
		}

		res.SendJSON(http.StatusForbidden, errMsg.FORBIDDEN, resData)
		return
	}

	// 生成token
	tokenString, err := util.GenerateToken(user.Id, user.Username, []byte(serverCfg.JwtSecret))
	if err != nil {
		resData = gin.H{
			"id":       user.Id,
			"username": user.Username,
		}

		fmt.Println("gen token", err)

		res.SendJSON(http.StatusUnauthorized, errMsg.UNAUTHORIZED, resData)
		return
	}

	userBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	// 优化TODO (封装到service中)
	loginUserKey := util.GetCacheKey("userLogin", user.Id)

	RedisConn := model.RedisPool.Get()
	defer RedisConn.Close()

	if _, redisErr := RedisConn.Do("SET", loginUserKey, userBytes, "EX", serverCfg.TokenMaxAge); redisErr != nil {
		fmt.Println("redis set failed: ", redisErr.Error())
	}

	// 请求返回
	resData = gin.H{
		"id":        user.Id,
		"username":  user.Username,
		"authToken": tokenString,
	}

	res.SendJSON(http.StatusOK, errMsg.SUCCESS, resData)
}

// // 用户注册
func SignUp(c *gin.Context) {
	var (
		reqData = make(map[string]interface{}, 0)
		// resObj  siginRes
		resData gin.H
		req     = common.Req{C: c}
		res     = common.Res{C: c}
	)

	// 数据解析
	if err := req.ParseParams(reqData); err != nil {
		resData = gin.H{
			"username": reqData["username"].(string),
		}

		fmt.Println("参数解析", err)

		res.SendJSON(http.StatusInternalServerError, errMsg.ERROR, resData)
		return
	}

	var user model.User
	if err := model.DB.Where("username = ? OR email = ? OR phone= ?", reqData["username"].(string), reqData["email"].(string), reqData["phone"].(string)).Find(&user).Error; err == nil {
		resData = gin.H{
			"username": reqData["username"].(string),
			"email":    reqData["email"].(string),
			"phone":    reqData["phone"].(string),
		}

		res.SendJSON(http.StatusInternalServerError, errMsg.ERROR, resData)
		return
	}

	user.Id = util.GenShortUuid()
	user.Username = reqData["username"].(string)
	user.PassHash = util.Encrypt(reqData["password"].(string))
	user.Name = reqData["name"].(string)
	user.Phone = reqData["phone"].(string)
	user.Email = reqData["email"].(string)

	if err := model.DB.Create(&user).Error; err != nil {
		resData = gin.H{
			"username": reqData["username"].(string),
			"email":    reqData["email"].(string),
			"phone":    reqData["phone"].(string),
		}

		res.SendJSON(http.StatusInternalServerError, errMsg.ERROR, resData)
		return
	}

	// 请求返回
	resData = gin.H{
		"id":       user.Id,
		"username": user.Username,
		"emial":    user.Email,
		"phone":    user.Phone,
	}

	res.SendJSON(http.StatusOK, errMsg.SUCCESS, resData)
}

// // 退出登陆
func Signout(c *gin.Context) {
	var (
		// resObj  siginRes
		resData gin.H
		res     = common.Res{C: c}
	)

	// 优化点
	userId, exists := c.Get("id")
	if exists {
		RedisConn := model.RedisPool.Get()
		defer RedisConn.Close()

		// 优化TODO (封装到service中)
		loginUserKey := util.GetCacheKey("userLogin", userId.(string))

		if _, err := RedisConn.Do("DEL", loginUserKey); err != nil {
			fmt.Println("redis delelte failed:", err)
		}
	}

	// 请求返回
	resData = gin.H{
		"id": userId,
	}

	res.SendJSON(http.StatusOK, errMsg.SUCCESS, resData)
}
