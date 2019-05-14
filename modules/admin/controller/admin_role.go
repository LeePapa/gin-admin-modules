package admin_controller

import (
	"github.com/gin-gonic/gin"
	"gin-modules/middleware"
)

func GetRoleList(ct *gin.Context) {

}

func AddRole(ct *gin.Context) {
	var rule rbacRoleRule
	if ct.ShouldBind(&rule) != nil {
		OutputJson(ct, -1, "参数异常", struct{}{})
		return
	}
	if !new(middleware.Rbac).Add(rule.RoleName, rule.Path, rule.Method) {
		OutputJson(ct, -1, "role权限添加失败了", struct{}{})
		return
	}
	OutputJson(ct, 200, "权限添加成功", struct{}{})
}

func RoleSave(ct *gin.Context) {

}

func RoleDelete(ct *gin.Context) {

}

func RoleSetLock(ct *gin.Context) {

}
