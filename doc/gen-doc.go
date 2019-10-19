package main

import (
	"fmt"
	"os"
	"text/template"
)

type router struct {
	Path        string
	Method      string
	Tag         string
	Summary     string
	Description string
	ContentType string
	Type        string
	Comment     string
	Required    bool
	DataType    string
}

func main() {
	t, err := template.ParseFiles("./swagger.yaml", "./meta.tpl")
	if err != nil {
		fmt.Println("模版加载失败", err)
		panic(err)
	}

	data1 := router{
		Path:        "/test",
		Method:      "post",
		Tag:         "auth",
		Summary:     "这是一个测试",
		Description: "swagger doc 渲染",
		ContentType: "application/json",
		Type:        "query",
		Comment:     "请求数据",
		Required:    true,
		DataType:    "array",
	}

	data2 := router{
		Path:        "/test2",
		Method:      "post",
		Tag:         "auth",
		Summary:     "这是一个测试",
		Description: "swagger doc 渲染",
		ContentType: "application/json",
		Type:        "query",
		Comment:     "请求数据",
		Required:    true,
		DataType:    "array",
	}

	var dic = []router{data1, data2}

	if err = t.ExecuteTemplate(os.Stdout, "index", dic); err != nil {
		panic(err)
	}
}
