package main

import (
	"github.com/gin-gonic/gin"
	"gin-modules/modules/admin"
	setting "gin-modules/pkg/setting"
	"os"
	"io"
	"gin-modules/pkg/redis"
	"gin-modules/pkg/db"
)

func init() {
	//公用redis客户端
	pkg_redis.Setup()
	//公用db客户端
	pkg_db.Initialize(setting.Setting.Database)
	//初始化admin模块的db客户端
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
	//初始化admin模块的路由
	r = admin.InitRouter(r)
	//启动服务
	r.Run(":" + setting.Setting.Server.Port)
}
