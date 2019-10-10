package config

// server
type ServerConfig struct {
	Env         string `json:"env"`
	SiteName    string `json:"siteName"`
	APIPrefix   string `json:"apiPrefix"`
	Port        string `json:"port"`
	Domain      string `json:"domain"`
	JwtSecret   string `json:"jwtSecret"`
	TokenMaxAge int    `json:"tokenMaxAge"`
}

// db
type DBConfig struct {
	Dialect      string `json:"dialect"`
	Database     string `json:"database"`
	User         string `json:"auth"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Charset      string `json:"charset"`
	URL          string `json:"url"`
	MaxIdleConns int    `json:"maxIdleConns"`
	MaxOpenConns int    `json:"maxOpenConns"`
}

// redis
type RedisConfig struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	Password  string `json:"password"`
	URL       string `json:"url"`
	MaxIdle   int    `json:"maxIdle"`
	MaxActive int    `json:"maxActive"`
}

// mongo
type MongoConfig struct {
	URL      string `json:"url"`
	Database string `json:"database"`
}

type Config struct {
	ServerCfg ServerConfig `json:"server"`
	DBCfg     DBConfig     `json:"mysql"`
	RedisCfg  RedisConfig  `json:"redis"`
	// mongo mongoCfg
}
