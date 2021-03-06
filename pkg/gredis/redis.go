package gredis

import (
	"blog/pkg/setting"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin/json"
	"time"
)

var RedisConn *redis.Pool

func Setup() error {
	RedisConn = &redis.Pool{
		MaxIdle:     setting.RedisSetting.MaxIdle,
		MaxActive:   setting.RedisSetting.MaxActive,
		IdleTimeout: setting.RedisSetting.IdleTimeout,
		Dial: func() (conn redis.Conn, e error) {
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
					err = c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}

func Set(key string, data interface{}, time int) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return false, err
	}
	reply, err := redis.Bool(conn.Do("SET", key, value))
	conn.Do("EXPIRE", key, time)
	return reply, err
}

func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return reply
}

func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return reply, err
}

func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bool(conn.Do("DEL", key))
	return reply, err
}

func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err := Delete(key)
		if err != nil {
			return err
		}
	}
	return nil
}
