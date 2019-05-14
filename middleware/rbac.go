package middleware

import (
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
	"fmt"
	"gin-modules/pkg/setting"
)

type Rbac struct {
}

func (r Rbac) Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims").(*CustomClaims)
		e := initCasbin()
		res, err := e.EnforceSafe(claims.RoleId, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			c.JSON(200, gin.H{"code": 403, "msg": "权限不足" + err.Error(), "data": struct{}{}})
			c.Abort()
			return
		}
		if res {
			c.Next()
		} else {
			c.JSON(200, gin.H{"code": 403, "msg": "很抱歉您没有此权限", "data": struct{}{}})
			c.Abort()
			return
		}
	}
}

//新增权限
func (r Rbac) Add(name string, path string, method string) bool {
	e := initCasbin()
	return e.AddPolicy(name, path, method)
}

//新增权限
func (r Rbac) AddGroup(name string, path string, method string) bool {
	e := initCasbin()
	return e.AddGroupingPolicy(name, path, method)
}

//持久化到数据库
func initCasbin() *casbin.Enforcer {
	a := gormadapter.NewAdapter("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", pkg_setting.Setting.Database.User, pkg_setting.Setting.Database.Password, pkg_setting.Setting.Database.Host, pkg_setting.Setting.Database.Db), true)
	e := casbin.NewEnforcer("/Users/Mr.Zhou/Project/golang/Libraries/src/gin-modules/modules/admin/conf/rbac_model.conf", a)
	//dir, _ := os.Getwd()
	//e := casbin.NewEnforcer(dir+"conf/auth_model.conf", a)
	e.LoadPolicy()
	return e
}
