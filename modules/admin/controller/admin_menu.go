package admin_controller

import (
	"github.com/gin-gonic/gin"
	"gin-modules/modules/admin/service"
	"gin-modules/modules/admin/model"
)

type rbacRoleRule struct {
	RoleName string `form:"role_name" binding:"required"`
	Path     string `form:"path" binding:"required"`
	Method   string `form:"method" binding:"required"`
}

type menuRule struct {
	Id     int    `form:"id"`
	Pid    int    `form:"pid"`
	Level  int    `form:"level" binding:"required"`
	Name   string `form:"name" binding:"required"`
	Icon   string `form:"icon"`
	Key    string `form:"key" binding:"required"`
	Path   string `form:"path" binding:"required"`
	Status int    `form:"status"`
}

//获取我的菜单列表
func GetMyMenuList(ct *gin.Context) {
	menu, err := admin_service.GetMenuList()
	//tokenInfo, _ := ct.Get("claims")
	//menu, err := admin_service.GetMyRoleMenu(tokenInfo.(*middleware.CustomClaims).RoleId)
	if err != nil {
		print(err.Error())
	}
	OutputJson(ct, 200, "success", menu)
}

//获取所有的菜单列表
func GetMenuList(ct *gin.Context) {
	menu, err := admin_service.GetMenuList()
	if err != nil {
		print(err.Error())
	}
	OutputJson(ct, 200, "success", menu)
}

//获取所有的菜单列表树，主要是给
func GetMenuTree(ct *gin.Context) {
	menu, err := admin_service.GetMenuTree()
	if err != nil {
		print(err.Error())
	}
	OutputJson(ct, 200, "success", menu)
}

//获取菜单详情
func GetMenuDetail(ct *gin.Context) {
	id := ct.Query("id")
	menuInfo, has := admin_model.GetMenuInfo("*", "id = ?", id)
	if !has {
		OutputJson(ct, -1, "无法获取到菜单信息", struct{}{})
		return
	}
	_, has = admin_model.GetMenuByOrder("id,`sort`", "`sort` ASC", "pid = ? AND `sort` > ? AND is_del = ?", menuInfo.Pid, menuInfo.Sort, 0)
	if !has {
		menuInfo.Sort = 1
	}
	_, has = admin_model.GetMenuByOrder("id,`sort`", "`sort` DESC", "pid = ? AND `sort` < ? AND is_del = ?", menuInfo.Pid, menuInfo.Sort, 0)
	if !has {
		if menuInfo.Sort == 1 {
			menuInfo.Sort = 0
		} else {
			menuInfo.Sort = -1
		}
	}

	OutputJson(ct, 200, "success", menuInfo)
}

//校验菜单的key
func CheckMenuKey(ct *gin.Context) {
	key := ct.PostForm("key")
	print(key)
	_, has := admin_model.GetMenuInfo("id", "`key` = ? AND is_del = ?", key, 0)
	if has {
		OutputJson(ct, -1, "菜单唯一标识已经存在", struct{}{})
		return
	}
	OutputJson(ct, 200, "success", struct{}{})
}

//菜单保存操作
func MenuSave(ct *gin.Context) {
	var rule menuRule
	if ct.ShouldBind(&rule) != nil {
		OutputJson(ct, -1, "参数异常", struct{}{})
		return
	}
	menuInfo, has := admin_model.GetMenuInfo("id", "`key` = ? AND is_del = ?", rule.Key, 0)
	if rule.Id != 0 {
		if has && menuInfo.ID != rule.Id {
			OutputJson(ct, -1, "菜单唯一标识已经存在", struct{}{})
			return
		}
		if !admin_model.UpdateMenu(map[string]interface{}{
			"pid":    rule.Pid,
			"level":  rule.Level,
			"name":   rule.Name,
			"key":    rule.Key,
			"icon":   rule.Icon,
			"status": rule.Status,
			"path":   rule.Path,
		}, "id = ?", rule.Id) {
			OutputJson(ct, -1, "菜单更新失败", struct{}{})
			return
		}
		admin_service.ClearMenuCache()
		go admin_service.GetMenuList()
		OutputJson(ct, 200, "菜单更新成功", struct{}{})
		return
	}
	if has {
		OutputJson(ct, -1, "菜单唯一标识已经存在", struct{}{})
		return
	}
	id := admin_model.AddMenuBySort(admin_model.AdminMenu{
		Pid:    rule.Pid,
		Level:  rule.Level,
		Name:   rule.Name,
		Key:    rule.Key,
		Icon:   rule.Icon,
		Status: rule.Status,
		Path:   rule.Path,
	})
	if id == 0 {
		OutputJson(ct, -1, "菜单添加失败", struct{}{})
		return
	}
	admin_service.ClearMenuCache()
	go admin_service.GetMenuList()
	OutputJson(ct, 200, "菜单添加成功", struct{}{})
}

//菜单排序更新
func MenuSortSet(ct *gin.Context) {
	id := ct.PostForm("id")
	operateType := ct.PostForm("type")
	menuInfo, has := admin_model.GetMenuInfo("`sort`,pid", "id = ?", id)
	if !has {
		OutputJson(ct, -1, "获取菜单信息失败", struct{}{})
		return
	}
	if operateType == "up" {
		upMenuInfo, has := admin_model.GetMenuByOrder("id,`sort`", "`sort` ASC", "pid = ? AND `sort` > ? AND is_del = ?", menuInfo.Pid, menuInfo.Sort, 0)
		if !has {
			OutputJson(ct, -1, "菜单已到上限位置，无法移动", struct{}{})
			return
		}
		if admin_model.UpdateMenu(map[string]interface{}{"sort": menuInfo.Sort}, "id = ?", upMenuInfo.ID) &&
			admin_model.UpdateMenu(map[string]interface{}{"sort": upMenuInfo.Sort}, "id = ?", id) {
			admin_service.ClearMenuCache()
			go admin_service.GetMenuList()
			OutputJson(ct, 200, "菜单位置更新成功", struct{}{})
			return
		}
		OutputJson(ct, -1, "菜单位置更新失败", struct{}{})
		return
	} else {
		downMenuInfo, has := admin_model.GetMenuByOrder("id,`sort`", "`sort` DESC", "pid = ? AND `sort` < ? AND is_del = ?", menuInfo.Pid, menuInfo.Sort, 0)
		if !has {
			OutputJson(ct, -1, "菜单已到下限位置，无法移动", struct{}{})
			return
		}
		if admin_model.UpdateMenu(map[string]interface{}{"sort": menuInfo.Sort}, "id = ?", downMenuInfo.ID) &&
			admin_model.UpdateMenu(map[string]interface{}{"sort": downMenuInfo.Sort}, "id = ?", id) {
			admin_service.ClearMenuCache()
			go admin_service.GetMenuList()
			OutputJson(ct, 200, "菜单位置更新成功", struct{}{})
			return
		}
		OutputJson(ct, -1, "菜单位置更新失败", struct{}{})
		return
	}
}

//菜单删除操作
func MenuDelete(ct *gin.Context) {
	id := ct.PostForm("id")
	if id == "" {
		OutputJson(ct, -1, "参数异常", struct{}{})
		return
	}
	if !admin_model.UpdateMenu(map[string]interface{}{"is_del": 1}, "id = ?", id) {
		OutputJson(ct, -1, "菜单删除失败", struct{}{})
	}
	admin_service.ClearMenuCache()
	go admin_service.GetMenuList()
	OutputJson(ct, 200, "菜单删除成功", struct{}{})
}

//菜单显示隐藏操作
func MenuShow(ct *gin.Context) {
	id := ct.PostForm("id")
	status := ct.DefaultPostForm("status", "1")
	if id == "" {
		OutputJson(ct, -1, "参数异常", struct{}{})
		return
	}
	if !admin_model.UpdateMenu(map[string]interface{}{"status": status}, "id = ?", id) {
		OutputJson(ct, -1, "菜单操作失败", struct{}{})
	}
	OutputJson(ct, 200, "菜单操作成功", struct{}{})
}
