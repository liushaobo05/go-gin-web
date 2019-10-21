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
	Summary     string `yaml:"summary,omitempty"`
	Description string `yaml:"description,omitempty"`
	Params      Parmas `yaml:"parmas,omitempty"`
}

// 请求参数
type Parmas struct {
	Query  Query  `yaml:"query,omitempty"`
	Header Header `yaml:"header,omitempty"`
	Body   Body   `yaml:"body,omitempty"`
}

type Header struct {
	Items []Items `yaml:"items,omitempty"`
}

type Query struct {
	Items []Items `yaml:"items,omitempty"`
}

type Items struct {
	Name     string `yaml:"name,omitempty"`
	Value    string `yaml:"value,omitempty"`
	Desc     string `yaml:"desc,omitempty"`
	Default  string `yaml:"default,omitempty"`
	Required bool   `yaml:"required,omitempty"`
}

type Body struct {
	Param
}

// Param 表示参数类型
type Param struct {
	Type     string            `yaml:"type,omitempty"`
	Example  string            `yaml:"example,omitempty"`
	Required bool              `yaml:"required,omitempty"`
	Items    map[string]*Param `yaml:"items,omitempty"`
	Desc     string            `yaml:"desc,omitempty"`
}

func NewApi(path string) (API, error) {
	var api API

	routerCfg, err := parse.LoadFile(path)
	if err != nil {
		return api, err
	}

	if err = routerCfg.GetPath("auth.sigin").GetStruct(&api); err != nil {
		return api, err
	}

	return api, nil
}

func main() {
	t, err := template.ParseFiles("./swagger.yaml", "./meta.tpl")
	if err != nil {
		fmt.Println("模版加载失败", err)
		panic(err)
	}

	data, _ := NewApi("./router.yaml")

	fmt.Printf("ddddd========%+v", data)

	if err = t.ExecuteTemplate(os.Stdout, "index", data); err != nil {
		panic(err)
	}
}
