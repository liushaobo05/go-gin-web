package router

import (
	"go-gin-web/pkg/parse"
)

// 富文本
type Richtext string

// Type 表示数据类型
type Type uint8

// 表示支持的各种数据类型
const (
	Number Type = iota
	Bool
	String
	Array
	Object
)

var (
	typeStringMap = map[Type]string{
		Number: "number",
		Bool:   "bool",
		String: "string",
		Array:  "array",
		Object: "object",
	}

	stringTypeMap = map[string]Type{
		"number": Number,
		"bool":   Bool,
		"string": String,
		"array":  Array,
		"object": Object,
	}
)

// api
type API struct {
	Name        string `yaml:name`
	Version     string `yaml:version`
	Method      string `yaml:method`
	Path        string `yaml:path`
	Auth        bool   `yaml:auth`
	Summary     string `yaml:summary`
	Description string `yaml:description`

	// Request []*Request
	// Responses []*Responses
	// Tags []string
}

// 请求参数
type Request struct {
	Param
	Headers []*Header
}

// Header 报头信息
type Header struct {
	Name        string
	Description string
	Value       string
}

// Path 路径信息
type Path struct {
	Path    string
	Params  []*Param
	Queries []*Param
}

// Param 表示参数类型
type Param struct {
	Name        string
	Type        Type
	Default     string
	Required    bool
	Summary     string
	Description string
}

func NewApi(path string) (*API, error) {
	data, err := parse.LoadFile(path)
	if err != nil {
		return nil, err
	}
	api := new(API)
	err = data.Get("auth").Get("demo").GetStruct(api)
	if err != nil {
		return api, err
	}

	return api, nil
}
