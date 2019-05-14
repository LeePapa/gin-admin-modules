package pkg_redis

import (
	"github.com/go-redis/redis"
	"errors"
	"time"
)

type RedisCache struct {
	Client *redis.Client
	Config redis.Options
}

func (r *RedisCache) Connect() error {
	var err error
	if r.Client != nil {
		_, err = r.Client.Ping().Result()
		if err == nil {
			return nil
		}
	}
	r.Client = redis.NewClient(&r.Config)
	_, err = r.Client.Ping().Result()
	if err != nil {
		return errors.New("连接 redis 服务器失败")
	}
	return nil
}

func (r *RedisCache) Get(key string, defaultValue ...string) (string, error) {
	err := r.Connect()
	if err != nil {
		return "", err
	}
	value, err := r.Client.Get(key).Result()
	if err != nil {
		return "", errors.New("获取失败，key不存在")
	}
	return value, nil
}

func (r *RedisCache) Set(key string, value interface{}, expire time.Duration) error {
	err := r.Connect()
	if err != nil {
		return err
	}
	if r.Client.Set(key, value, expire).Err() != nil {
		return errors.New("缓存设置失败")
	}
	return nil
}

func (r *RedisCache) HSet(key string, hash string, value interface{}) error {
	err := r.Connect()
	if err != nil {
		return err
	}
	if r.Client.HSet(key, hash, value).Err() != nil {
		return errors.New("缓存设置失败")
	}
	return nil
}

func (r *RedisCache) HGet(key string, hash string, defaultValue ...string) (string, error) {
	err := r.Connect()
	if err != nil {
		return "", err
	}
	value, err := r.Client.HGet(key, hash).Result()
	if err != nil {
		return "", errors.New("获取失败，key不存在")
	}
	return value, nil
}

func (r *RedisCache) ExpireTime(key string, expire time.Duration) error {
	err := r.Connect()
	if err != nil {
		return err
	}
	if r.Client.Expire(key, expire).Err() != nil {
		return errors.New("设置key有效时间失败")
	}
	return nil
}

func (r *RedisCache) Delete(key string) error {
	err := r.Connect()
	if err != nil {
		return err
	}
	if r.Client.Del(key).Err() != nil {
		return errors.New("缓存设置失败")
	}
	return nil
}