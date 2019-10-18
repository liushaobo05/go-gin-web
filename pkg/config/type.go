package config

// server
type ServerConfig struct {
	Env         string `yaml:"env"`
	SiteName    string `yaml:"siteName"`
	APIPrefix   string `yaml:"apiPrefix"`
	Port        string `yaml:"port"`
	Domain      string `yaml:"domain"`
	JwtSecret   string `yaml:"jwtSecret"`
	TokenMaxAge int    `yaml:"tokenMaxAge"`
	CronTask    bool   `yaml:"cronTask"`
}

// db
type DBConfig struct {
	Dialect      string `yaml:"dialect"`
	Database     string `yaml:"database"`
	User         string `yaml:"auth"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Charset      string `yaml:"charset"`
	URL          string `yaml:"url"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
}

// redis
type RedisConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Password  string `yaml:"password"`
	URL       string `yaml:"url"`
	MaxIdle   int    `yaml:"maxIdle"`
	MaxActive int    `yaml:"maxActive"`
}

// mongo
type MongoConfig struct {
	URL      string `yaml:"url"`
	Database string `yaml:"database"`
}

type Config struct {
	ServerCfg ServerConfig `yaml:"server"`
	DBCfg     DBConfig     `yaml:"mysql"`
	RedisCfg  RedisConfig  `yaml:"redis"`
	// mongo mongoCfg
	CorsCfg CorsConfig `yaml:"cors"`
	RateCfg RateConfig `yaml:"rateLimiter"`
}

// CORS 跨域请求配置参数
type CorsConfig struct {
	Enable           bool     `yaml:"enable"`
	AllowOrigins     []string `yaml:"allowOrigins"`
	AllowMethods     []string `yaml:"allowMethods"`
	AllowHeaders     []string `yaml:"allowHeaders"`
	AllowCredentials bool     `yaml:"allowCredentials"`
	MaxAge           int      `yaml:"maxAge"`
}

type RateConfig struct {
	Enable bool `yaml:"enable"`
	Count  int
}
