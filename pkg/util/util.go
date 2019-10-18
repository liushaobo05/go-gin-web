package util

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"reflect"
	"time"

	"github.com/goinggo/mapstructure"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

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

// 生成随机字符串
func GenRandStr(mark string, n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	if mark == "" {
		return string(b)
	}

	return fmt.Sprintf("%s-%s", mark, string(b))
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
	return filepath.Ext(path)[1:]
}

// cache key
func GetCacheKey(topic, mark string) string {
	return fmt.Sprintf("%s@%s", topic, mark)
}

// ParseToStr 将map中的键值对输出成querystring形式
func ParseToStr(mp map[string]string) string {
	values := ""
	for key, val := range mp {
		values += "&" + key + "=" + val
	}
	temp := values[1:]
	values = "?" + temp
	return values
}
