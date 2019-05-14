package admin_service

import (
	"gin-modules/pkg/redis"
	"gin-modules/pkg/setting"
	"github.com/go-redis/redis"
)

var redisCache pkg_redis.RedisCache

func init() {
	redisCache = pkg_redis.RedisCache{Config: redis.Options{Addr: pkg_setting.Setting.RedisCache.Host, Password: pkg_setting.Setting.RedisCache.Password, DB: pkg_setting.Setting.RedisCache.DB,}}
}
