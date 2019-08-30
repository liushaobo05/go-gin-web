package validate

import (
	"errors"
	"fmt"
)

type Validator struct {
	data     map[string]string   //要校验的数据
	rules    map[string]*vRule   //规则列表，key为字段名
	messages map[string][]string //错误输出
}

type vRule struct {
	required bool
	vr       ValidateRuler
}

//校验规则接口，支持自定义规则
type ValidateRuler interface {
	Check(data string) error
}

//内置规则结构，实现ValidateRuler接口
type normalRule struct {
	key      string
	dataType string
	rules    []string
}

//创建校验器对象
func NewValidator(data interface{}) *Validator {
	v := &Validator{data: data}
	v.rule = make(map[string]*vRule)
	v.messages = make(map[string][]string)
	return v
}

//添加内置的校验规则
func (v *Validator) AddRule(key string, required bool, dataType string, rules ...string) {
	nr := &normalRule{key, dataType}
	nr.rules = append(nr.rules, rules)
	v.rule[key] = &vRule{nr, required} //默认required = true
}

//执行检查
func (v *Validator) Check() (errs map[string]error) {
	errs = make(map[string]error)
	for k, v := range v.rule {
		data, exists := v.data[k]
		if !exists { //无值
			if v.required { //如果必填，报错
				errs[k] = errors.New("data error: required field miss")
			}
		} else { //有值判断规则
			if err := v.vr.Check(data); err != nil { //调用ValidateRuler接口的Check方法来检查
				errs[k] = err
			}
		}
	}
	return errs
}

func (v *normalRule) Check(data string) (Err error) {
	if v.params == "" {
		Err = errors.New("rule error: params wrong of rule")
		return
	}
	switch v.rule {
	case "string":
		//字符串，根椐params判断长度的最大值和最小值
	case "number":
		//判断是否整数数字
		//判断最大值和最小值是否在params指定的范围
	case "list":
		//判断值是否在params指定的列表
	case "regular":
		//是否符合正则表达式
	default:
		Err = errors.New(fmt.Sprintf("rule error: not support of rule=%s", v.rule))
	}
	return
}
