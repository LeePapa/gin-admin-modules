package main

import (
	"github.com/gin-gonic/gin"
	"gin-modules/modules/admin"
	setting "gin-modules/pkg/setting"
	"os"
	"io"
	"gin-modules/pkg/redis"
)

func init() {
	pkg_redis.Setup()
	admin.InitDB(setting.Setting.Database)
}

func main() {
	gin.SetMode(setting.Setting.Server.Mode)
	if setting.Setting.Log.Enable == 1 {
		gin.DisableConsoleColor()
		f, _ := os.Create(setting.Setting.Log.Path)
		gin.DefaultWriter = io.MultiWriter(f)
	}
	r := gin.Default()
	r = admin.InitRouter(r)
	r.Run(":" + setting.Setting.Server.Port)
}
