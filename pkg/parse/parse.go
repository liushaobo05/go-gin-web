package parse

import (
	"encoding/json"
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

type Cfg struct {
	data interface{}
}

func Load(path string) (*Cfg, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return LoadContent(data, util.Ext(path))
}

func LoadContent(data []byte, dataType ...string) (*Cfg, error) {
	t := "json"
	if len(dataType) > 0 {
		t = dataType[0]
	}

	switch t {
	case "yml", "yaml", ".yml", ".yaml":
		return NewYaml(data)
	case "toml", ".toml":
		return NewToml(data)
	}
	return NewJson(data)
}

func New(t string, body []byte) (*Cfg, error) {
	if t == "json" {
		return NewJson(body)
	} else if t == "yaml" {
		return NewYaml(body)
	} else if t == "toml" {
		return NewToml(body)
	}
	return nil, errors.New("invalid type t, t is json or yaml")
}

func NewJson(body []byte) (*Cfg, error) {
	j := new(Cfg)
	if err := json.Unmarshal(body, &j.data); err != nil {
		return nil, err
	}
	return j, nil
}

func NewToml(body []byte) (*Cfg, error) {
	j := new(Cfg)
	if err := toml.Unmarshal(body, &j.data); err != nil {
		return nil, err
	}
	return j, nil
}

func NewYaml(body []byte) (*Cfg, error) {
	j := new(Cfg)
	if err := yaml.Unmarshal(body, &j.data); err != nil {
		return nil, err
	}
	return j, nil
}
