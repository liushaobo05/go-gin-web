package model

import (
	"fmt"
	"os"
	"time"

	"go-gin-web/pkg/config"

	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库连接
var DB *gorm.DB

// RedisPool Redis连接池
var RedisPool *redis.Pool

// MongoDB 数据库连接
// var MongoDB *mgo.Database

func initDB() {
	var (
		dbCfg     = config.MyCfg
		serverCfg = config.ServerCfg
	)

	if dbCfg.URL == "" {
		dbCfg.URL = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			dbCfg.User,
			dbCfg.Password,
			dbCfg.Host,
			dbCfg.Port,
			dbCfg.Database,
			dbCfg.Charset)
	}

	db, err := gorm.Open(dbCfg.Dialect, dbCfg.URL)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	if serverCfg.Env == "dev" {
		db.LogMode(true)
	}
	db.DB().SetMaxIdleConns(dbCfg.MaxIdleConns)
	db.DB().SetMaxOpenConns(dbCfg.MaxOpenConns)
	DB = db
}

func initRedis() {
	var (
		redisCfg = config.RedisCfg
	)

	RedisPool = &redis.Pool{
		MaxIdle:     redisCfg.MaxIdle,
		MaxActive:   redisCfg.MaxActive,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisCfg.URL, redis.DialPassword(redisCfg.Password))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}

/*
 * mgo文档 http://labix.org/mgo
 */
// func initMongo() {
// 	if config.MongoConfig.URL == "" {
// 		return
// 	}
// 	session, err := mgo.Dial(config.MongoConfig.URL)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		os.Exit(-1)
// 	}
// 	// Optional. Switch the session to a monotonic behavior.
// 	session.SetMode(mgo.Monotonic, true)
// 	MongoDB = session.DB(config.MongoConfig.Database)
// }

func Init() {
	initDB()
	initRedis()
	// initMongo()
}
