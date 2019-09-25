package cache

import (
	"fmt"
	"go-gin-web/model"
	"github.com/garyburd/redigo/redis"
)

func Set(key string, data []byte, expiration int) error {
	RedisConn := model.RedisPool.Get()
	defer RedisConn.Close()

	if _, redisErr := RedisConn.Do("SET", key, data, "EX", expiration); redisErr != nil {
		fmt.Println("redis set failed: ", redisErr.Error())
		return  redisErr
	}

	return nil
}

func Get(key string) ([]byte, error) {
	RedisConn := model.RedisPool.Get()
	defer RedisConn.Close()

	dataBytes, redisErr := redis.Bytes(RedisConn.Do("GET", key))
	if redisErr != nil {
		fmt.Println("redis set failed: ", redisErr.Error())
		return  nil, redisErr
	}

	return dataBytes, nil
}

func Del(key string) error {
	RedisConn := model.RedisPool.Get()
	defer RedisConn.Close()

	if _, redisErr := RedisConn.Do("DEL", key); redisErr != nil {
		fmt.Println("redis set failed: ", redisErr.Error())
		return  redisErr
	}

	return nil
}