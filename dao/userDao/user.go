package userDao

import (
	"go-gin-web/model"
)

// 获取用户
func GetUser(params map[string]interface{}) (model.User, error) {
	var user model.User

	if err := model.DB.Where("username=?", params["username"]).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// redis
