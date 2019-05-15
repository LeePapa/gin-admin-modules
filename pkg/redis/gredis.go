package pkg_redis

import (
	"time"
	"github.com/gomodule/redigo/redis"
	"gin-modules/pkg/setting"
	"errors"
)

var RedisConn *redis.Pool

func Setup() error {
	RedisConn = &redis.Pool{
		MaxIdle:     pkg_setting.Setting.Redis.MaxIdle,
		MaxActive:   pkg_setting.Setting.Redis.MaxActive,
		IdleTimeout: pkg_setting.Setting.Redis.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", pkg_setting.Setting.Redis.Host)
			if err != nil {
				return nil, err
			}
			if pkg_setting.Setting.Redis.Password != "" {
				if _, err := c.Do("AUTH", pkg_setting.Setting.Redis.Password); err != nil {
					c.Close()
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

func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()
	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return exists
}

func Set(key string, value interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()
	var err error
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	if time > 0 {
		_, err = conn.Do("EXPIRE", key, time)
		if err != nil {
			return err
		}
	}
	return nil
}

func Get(key string) (string, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	reply, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}
	return reply, nil
}

func HSet(key string, hash string, value interface{}) error {
	conn := RedisConn.Get()
	defer conn.Close()
	_, err := conn.Do("HSET", key, hash, value)
	if err != nil {
		return err
	}
	return nil
}

func HGetAll(key string) (map[string]string, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	hashMap, err := redis.StringMap(conn.Do("HGETALL", key))
	if err != nil {
		return nil, err
	}
	return hashMap, nil
}

func HGet(key string, hash string) (string, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	hashString, err := redis.String(conn.Do("HGET", key, hash))
	if err != nil {
		return "", err
	}
	return hashString, nil
}

func ExpireTime(key string, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()
	_, err := conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}
	return nil
}

func Delete(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()
	b, err := redis.Bool(conn.Do("DEL", key))
	if err != nil {
		return false
	}
	return b
}

func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()
	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}
	for _, key := range keys {
		if !Delete(key) {
			return errors.New("删除失败")
		}
	}
	return nil
}
