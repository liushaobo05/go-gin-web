package main

import (
	"fmt"
	"go-gin-web/pkg/util"
	"reflect"
)

type Validate struct {
	data interface{}
	rule map[string]interface{}
	msg string
}

// 创建对象
func NewValidate(data interface{}, rule map[string]interface{}) *Validate {
	return &Validate{
		data: data,
		rule: rule,
	}
}

func (v *Validate) Check() string {
	myContext := &Context{v}
	validate := &ValidateExec{
		F:   make(FilterFunc, 0),
		Ctx: myContext,
	}

	for key, _ := range v.rule {
		f := bakedInValidators[key]
		validate.register(f)
	}

	validate.verify()

    return validate.Ctx.msg
}

type Context struct {
	*Validate
}

type CheckFunc func(*Context) bool

var bakedInValidators = map[string]CheckFunc{
	"IsRequired": IsRequired,
	"len":      Len,
	"min":      Min,
	"max":      Max,
}

type FilterFunc []CheckFunc

type ValidateExec struct {
	F   FilterFunc
	Ctx *Context
}

func (e *ValidateExec) register(f CheckFunc) {
	e.F = append(e.F, f)
}

func (e *ValidateExec) verify() {
	for _, f := range e.F {
		res := f(e.Ctx)
		if res == false {
			return
		}
	}
}

func IsRequired(ctx *Context) bool {
    if ctx.rule["IsRequired"].(bool) {
    	if util.IsEmpty(ctx.data) {
			ctx.msg = "不能为空"
			return false
		}
	}
	return true
}

func Len(ctx *Context) bool {
	flag := true
	switch v := ctx.data.(type) {
	case string:
		flag = (len(v) == ctx.rule["len"].(int))
	default:
		{
			refValue := reflect.ValueOf(v)
			switch refValue.Kind() {
			case reflect.Slice, reflect.Array, reflect.Map:
				{
					flag = (refValue.Len() == ctx.rule["len"].(int))
				}
			}
		}
	}

	if !flag {
       ctx.msg = "长度校验失败"
       return false
	}

	return true
}

func Max(ctx *Context) bool {
	ctx.msg = "最大值check失败"
	return true
}

func Min(ctx *Context) bool {
	ctx.msg = "最小值check失败"
	return false
}

func main() {
	rule := map[string]interface{}{
		"IsRequired": true,
		"max": 5,
		"min": 3,
	}

	v := NewValidate(12, rule)
	errMsg := v.Check()
	fmt.Println("==========", errMsg)
}
