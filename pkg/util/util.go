package util

import (
	"fmt"
	"path/filepath"
	"reflect"

	"github.com/goinggo/mapstructure"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// 生成uuid
func GenUuid(mark string) string {
	u2 := uuid.NewV4()

	uuidStr := fmt.Sprintf("%s-%s", mark, u2.String())
	return uuidStr
}

// uuid
func GenShortUuid() string {
	return uuid.NewV4().String()
}

// bcrypt
func Encrypt(source string) string {
	hashedBytes, _ := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes)
}

// Compare
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// map to struct
func MapToStruct(obj interface{}, mapData map[interface{}]interface{}) error {
	if err := mapstructure.Decode(mapData, obj); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// struct to map
func StructToMap(obj interface{}, mapData map[string]interface{}) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	for i := 0; i < t.NumField(); i++ {
		mapData[t.Field(i).Name] = v.Field(i).Interface()
	}
}

// 文件扩展名
func Ext(path string) string {
	return filepath.Ext(path)
}

// cache key
func GetCacheKey(topic, mark string) string {
	return fmt.Sprintf("%s@%d", topic, mark)
}
