package router

import (
	"go-gin-web/pkg/parse"
)

// 富文本
type Richtext string

// Type 表示数据类型
type Type uint8

type ParamType uint

// 表示支持的各种数据类型
const (
	Number Type = iota
	Bool
	String
	Array
	Object
)

const (
	Bodu ParamType = iota
	Query
	Path
	Header
)

// api
type API struct {
	Name        string   `yaml:name`
	Version     string   `yaml:version`
	ID          string   `yaml:"id,omitempty"`
	Method      string   `yaml:method`
	Path        string   `yaml:path`
	Auth        bool     `yaml:auth`
	Summary     string   `yaml:summary,omitempty`
	Description string   `yaml:description,omitempty`
	Params      *Parmas  `yaml:"parmas,omitempty"`
	Responses   *Request `yaml:"response,omitempty"`
}

// 请求参数
type Parmas struct {
	Query  *Query
	Header *Header
	Body   *Body
}

type Header struct {
	item []*item
}

type Query struct {
	item []*item
}

type item struct {
	Key      string
	Value    string
	Desc     string
	Default  string
	Required bool
}

type Body struct {
	item
}

// Param 表示参数类型
type Param struct {
	Name     string   `yaml:"name,attr"`
	Type     string   `yaml:"type,attr,omitempty"`
	Value    string   `yaml:"value,attr,omitempty"`
	Default  string   `yaml:"default,attr,omitempty"`
	Required bool     `yaml:"required,attr,omitempty"`
	Items    []*Param `yaml:"param,omitempty"`
	Desc     string   `yaml:"desc,omitempty"`
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
