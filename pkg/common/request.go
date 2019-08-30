package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

const (
	defaultMemory = 32 << 20 // 32 MB
)

type Req struct {
	C *gin.Context
}

// 解析路由参数
func (req *Req) Params(name string) (string, bool) {
	return req.C.Params.Get(name)
}

// 解析query参数
func (req *Req) Query(name string) (string, bool) {
	return req.C.GetQuery("id")
}

// 解析body
func (req *Req) Parse(data interface{}) interface{} {
	errMsg := make(map[string]string)

	rules := govalidator.MapData{
		"username": []string{"required", "between:3,20"},
		"password": []string{"required", "between:3,10"},
	}

	opts := govalidator.Options{
		Request:         req.C.Request,
		Rules:           rules,
		RequiredDefault: true,
	}

	if bingData, ok := data.(map[string]interface{}); ok {
		if t, ok := req.C.Request.Header["Content-Type"]; ok {
			if t[0] == "application/json" || t[0] == "text/plain" {
				opts.Data = &bingData
				validator := govalidator.New(opts)
				errsMap := validator.ValidateJSON()

				for errkey, errText := range errsMap {
					fmt.Println(errText)
					errMsg[errkey] = strings.Join(errText, ":")
				}
			} else {
				validator := govalidator.New(opts)
				errsMap := validator.Validate()

				if err := req.BindMap(bingData); err != nil {
					return errors.New("serverErr")
				}

				for errkey, errText := range errsMap {
					fmt.Println(errText)
					errMsg[errkey] = strings.Join(errText, ":")
				}
			}
		}
	} else {
		return errors.New("serverErr")
	}

	return errMsg
}

func (req *Req) ParseParams(data map[string]interface{}) error {
	if err := req.BindMap(data); err != nil {
		return errors.New("serverErr")
	}

	return nil
}

//BindMap decodes body to map
func (req *Req) BindMap(obj map[string]interface{}) error {
	//Provides JSON and Form content types
	if req.C.Request.Method == "GET" {
		return nil
	}
	contentType := req.C.ContentType()
	switch contentType {
	case gin.MIMEJSON:
		return req.jsonBind(obj)
	case gin.MIMEPOSTForm, gin.MIMEMultipartPOSTForm:
		return req.formBind(obj)
	default:
		return errors.New("Unsupported content type")
	}
}

//jsonBind binds json into map
func (req *Req) jsonBind(obj map[string]interface{}) (err error) {
	body, err := ioutil.ReadAll(req.C.Request.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, &obj)
}

//jsonBind binds form into json
func (req *Req) formBind(obj map[string]interface{}) (err error) {
	if err = req.C.Request.ParseForm(); err != nil {
		return err
	}
	req.C.Request.ParseMultipartForm(defaultMemory)
	for k, v := range req.C.Request.Form {
		obj[k] = v[0] //take first
	}
	return nil
}
