package userDao

import (
	"go-gin-web/model"
)

// 获取用户信息
func GetUser(params map[string]interface{}) (model.User, error) {
	var user model.User

	if err := model.DB.Where(params).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func GetRoles(userId string) (string, error) {
	var (
		role string
	)

	row := model.DB.Raw("SELECT role.name from tb_main_role role left join tb_main_user_role user on role.id = user.roleId where user.userId = ?", userId).Row()

	row.Scan(&role)

	return role, nil
}

// 添加用户
func CreateUser(user model.User) error {
	err := model.DB.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

// 创建secret
func CreateSecret(secret model.Secret) error {
	err := model.DB.Create(&secret).Error
	if err != nil {
		return err
	}

	return nil
}

// 更新secret
func UpdateSecret(params map[string]interface{}, data interface{}) error {
	var secret model.Secret
	err := model.DB.Model(&secret).Where("id = ? AND userId = ?", params["id"].(string), params["userId"].(string)).Updates(data).Error
	if err != nil {
		return err
	}

	return nil
}

// 获取secret
func GetSecret(params map[string]interface{}) (model.Secret, error) {
	var secret model.Secret

	if err := model.DB.Where("id= ? AND userId= ?", params["id"], params["userId"]).First(&secret).Error; err != nil {
		return secret, err
	}

	return secret, nil
}

// 列表 todo 分页实践
