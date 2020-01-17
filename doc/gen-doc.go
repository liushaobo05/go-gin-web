package main

import (
	"fmt"
	"go-gin-web/pkg/parse"
	"os"
	"text/template"
)

var data = `
    name: sigin
    method: POST
    version: v1
    path: /signIn
    auth: false
    summary: "用户登陆"
    description: "用户登陆接口"
    params:
        body:
            type: object
            items:
                username:
                    type: string
                    required: true
                    desc: 用户名
                password:
                    type: string
                    required: true
                    desc: 密码
`

type API struct {
	Name        string `yaml:"name"`
	Method      string `yaml:"method"`
	Version     string `yaml:"version"`
	Path        string `yaml:"path"`
	Auth        bool   `yaml:"auth"`
	Summary     string `yaml:"summary"`
	Description string `yaml:"description"`
	Params      Parmas `yaml:"params"`
}

// 请求参数
type Parmas struct {
	Query  Query  `yaml:"query"`
	Header Header `yaml:"header"`
	Body   Body   `yaml:"body"`
}

type Header struct {
	Items []Items `yaml:"items"`
}

type Query struct {
	Items []Items `yaml:"items"`
}

type Items struct {
	Name     string `yaml:"name"`
	Value    string `yaml:"value"`
	Desc     string `yaml:"desc"`
	Default  string `yaml:"default"`
	Required bool   `yaml:"required"`
}

type Body struct {
	Param
}

// Param 表示参数类型
type Param struct {
	Type     string           `yaml:"type"`
	Example  string           `yaml:"example"`
	Required bool             `yaml:"required"`
	Items    map[string]Param `yaml:"items"`
	Desc     string           `yaml:"desc"`
}

func NewApi(path string) (API, error) {
	var api API

	// byteStr, err := ioutil.ReadFile(path)
	// if err != nil {
	// 	return api, err
	// }

	// if err := yaml.Unmarshal(byteStr, &api); err != nil {
	// 	return api, err
	// }

	data, err := parse.LoadFile(path)
	if err != nil {
		return api, err
	}
	// api := new(API)
	err = data.Get("auth").Get("sigin").GetStruct(&api)

	fmt.Println("dddddddd", api)

	return api, nil
}

func main() {
	//t, err := template.ParseFiles("./swagger.tpl", "./meta.tpl")
	//if err != nil {
	//	fmt.Println("模版加载失败", err)
	//	panic(err)
	//}

	data, _ := NewApi("demo.yaml")

	fmt.Printf("ddddd========%+v", data)

	//if err = t.ExecuteTemplate(os.Stdout, "index", data); err != nil {
	//	panic(err)
	//}
}
