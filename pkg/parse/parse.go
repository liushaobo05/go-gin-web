package parse

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"path"
	"reflect"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

type Parser struct {
	data interface{}
}

func LoadFile(filePath string) (*Parser, error) {
	byteStr, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	p := &Parser{}

	ext := path.Ext(filePath)

	if ext == ".json" {
		p.data, err = NewJson(byteStr)

		if err != nil {
			return p, err
		}
	} else if ext == ".yaml" {
		p.data, err = NewYaml(byteStr)

		if err != nil {
			return p, err
		}
		// todo fixbug
	} else if ext == ".toml" {
		p.data, err = NewYaml(byteStr)

		if err != nil {
			return p, err
		}
	} else {
		return p, errors.New("file is not support")
	}

	return p, nil
}

func NewFromReader(r io.Reader) (interface{}, error) {
	p := new(Parser)
	dec := json.NewDecoder(r)
	err := dec.Decode(&p.data)
	return p, err
}

func NewJson(body []byte) (interface{}, error) {
	p := new(Parser)
	err := p.UnmarshalJSON(body)
	if err != nil {
		return nil, err
	}
	return p.data, nil
}

func NewToml(body []byte) (interface{}, error) {
	p := new(Parser)
	if err := toml.Unmarshal(body, &p.data); err != nil {
		return nil, err
	}

	return p.data, nil
}

func NewYaml(body []byte) (interface{}, error) {
	p := new(Parser)
	if err := yaml.Unmarshal(body, &p.data); err != nil {
		return nil, err
	}
	return p.data, nil
}

func (p *Parser) MarshalJSON() ([]byte, error) {
	return json.Marshal(&p.data)
}

func (p *Parser) MarshalYaml() ([]byte, error) {
	return yaml.Marshal(&p.data)
}

func (p *Parser) UnmarshalJSON(bt []byte) error {
	return json.Unmarshal(bt, &p.data)
}

func (p *Parser) UnmarshalYaml(bt []byte) error {
	return yaml.Unmarshal(bt, &p.data)
}

// pretty
func (p *Parser) EncodePretty() ([]byte, error) {
	return json.MarshalIndent(&p.data, "", "  ")
}

func (p *Parser) Get(key string) *Parser {
	m, err := p.Map()
	if err == nil {
		if val, ok := m[key]; ok {
			return &Parser{val}
		}
	}
	return &Parser{nil}
}

// getpath
func (p *Parser) GetPath(path string) *Parser {
	jin := p
	r := strings.Split(path, ".")
	for _, v := range r {
		jin = jin.Get(v)
	}
	return jin
}

// get index
func (p *Parser) GetIndex(index int) *Parser {
	a, err := p.array()
	if err == nil {
		if len(a) > index {
			return &Parser{a[index]}
		}
	}
	return &Parser{nil}
}

// Int coerces into an int
func (p *Parser) GetInt() (int, error) {
	switch p.data.(type) {
	case float32, float64:
		return int(reflect.ValueOf(p.data).Float()), nil
	case int, int8, int16, int32, int64:
		return int(reflect.ValueOf(p.data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(p.data).Uint()), nil
	}
	return 0, errors.New("invalid value type")
}

// Int64 coerces into an int64
func (p *Parser) GetInt64() (int64, error) {
	switch p.data.(type) {
	case float32, float64:
		return int64(reflect.ValueOf(p.data).Float()), nil
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(p.data).Int(), nil
	case uint, uint8, uint16, uint32, uint64:
		return int64(reflect.ValueOf(p.data).Uint()), nil
	}
	return 0, errors.New("invalid value type")
}

// Uint64 coerces into an uint64
func (p *Parser) GetUint64() (uint64, error) {
	switch p.data.(type) {
	case float32, float64:
		return uint64(reflect.ValueOf(p.data).Float()), nil
	case int, int8, int16, int32, int64:
		return uint64(reflect.ValueOf(p.data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return reflect.ValueOf(p.data).Uint(), nil
	}
	return 0, errors.New("invalid value type")
}

func (p *Parser) GetFloat64() (float64, error) {
	switch p.data.(type) {
	case float32, float64:
		return reflect.ValueOf(p.data).Float(), nil
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(p.data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(p.data).Uint()), nil
	}
	return 0, errors.New("invalid value type")
}

// Bool type asserts to `bool`
func (p *Parser) GetBool() (bool, error) {
	if s, ok := (p.data).(bool); ok {
		return s, nil
	}
	return false, errors.New("type assertion to bool failed")
}

// String type asserts to `string`
func (p *Parser) GetString() (string, error) {
	if s, ok := (p.data).(string); ok {
		return s, nil
	}
	return "", errors.New("type assertion to string failed")
}

// Bytes type asserts to `[]byte`
func (p *Parser) GetBytes() ([]byte, error) {
	if s, ok := (p.data).(string); ok {
		return []byte(s), nil
	}
	return nil, errors.New("type assertion to []byte failed")
}

// Array type asserts to an `array`
func (p *Parser) GetArray() ([]interface{}, error) {
	return p.array()
}

func (p *Parser) array() ([]interface{}, error) {
	if a, ok := (p.data).([]interface{}); ok {
		return a, nil
	}
	return nil, errors.New("type assertion to []interface{} failed")
}

// Map type asserts to `map`
func (p *Parser) GetMap() (map[interface{}]interface{}, error) {
	return p.Map()
}

func (p *Parser) Map() (map[interface{}]interface{}, error) {
	if m, ok := (p.data).(map[interface{}]interface{}); ok {
		return m, nil
	}

	return nil, errors.New("type assertion to map[string]interface{} failed")
}

func (p *Parser) GetValue() interface{} {
	return p.Interface()
}

func (p *Parser) GetStruct(data interface{}) error {
	bytes, err := p.MarshalYaml()
	if err != nil {
		return err
	}

	fmt.Println("解析", string(bytes))

	err = yaml.Unmarshal(bytes, data)
	if err != nil {
		return err
	}

	return nil
}

// Interface returns the underlying data
func (p *Parser) Interface() interface{} {
	return p.data
}

// empty Parser
func New() *Parser {
	return &Parser{
		data: make(map[string]interface{}),
	}
}

func (p *Parser) CheckGet(key string) (*Parser, bool) {
	m, err := p.Map()
	if err == nil {
		if val, ok := m[key]; ok {
			return &Parser{val}, true
		}
	}
	return nil, false
}

// set
func (p *Parser) Set(key string, val interface{}) {
	m, err := p.Map()
	if err != nil {
		return
	}
	m[key] = val
}

// Del modifies `Json` map by deleting `key` if it is present.
func (p *Parser) Del(key string) {
	m, err := p.Map()
	if err != nil {
		return
	}
	delete(m, key)
}
