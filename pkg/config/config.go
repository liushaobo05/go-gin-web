package config

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type cfg struct {
	Name string
}

// 全局配置
var Cfg Config

// server配置
var ServerCfg ServerConfig

// mysql配置
var MyCfg DBConfig

// redis
var RedisCfg RedisConfig

// cors
var CorsCfg CorsConfig

var RateCfg RateConfig

func Init(path string) error {
	c := cfg{
		Name: path,
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	c.parse()

	// 监控配置文件变化并热加载程序
	c.watchConfig()

	return nil
}

func (c *cfg) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("MESSAGE")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	return nil
}

// 监控配置文件变化并热加载程序
func (c *cfg) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}

// 解析配置
func (c *cfg) parse() {
	viper.Unmarshal(&Cfg)
	viper.UnmarshalKey("server", &ServerCfg)    // 将配置解析到server变量
	viper.UnmarshalKey("mysql", &MyCfg)         // 将配置解析到mysql变量
	viper.UnmarshalKey("redis", &RedisCfg)      // 将配置解析到redis变量
	viper.UnmarshalKey("cors", &CorsCfg)        // 将配置解析到cors变量
	viper.UnmarshalKey("rateLimiter", &RateCfg) // 将配置解析到rateLimiter变量
}
