package admin

import (
	"github.com/gin-gonic/gin"
	"gin-modules/modules/admin/controller"
	"gin-modules/middleware"
	"gin-modules/modules/admin/model"
	"gin-modules/pkg/setting"
)

func InitDB(config pkg_setting.DatabaseIni) {
	admin_model.Initialize(config)
}

func InitRouter(r *gin.Engine) *gin.Engine {
	prx := "/api"
	r.POST(prx+"/auth/login", admin_controller.Login)
	r.POST(prx+"/auth/register", admin_controller.Register)
	r.POST(prx+"/public/sendSmsCode", admin_controller.SendSmsCode)
	r.POST(prx+"/role/add", admin_controller.AddRole)
	r.POST(prx+"/role/info", admin_controller.GetMyMenuList)

	//auth := r.Group("/", middleware.JWTAuth(), new(middleware.Rbac).Check())
	auth := r.Group("/", middleware.JWTAuth())
	{
		//菜单相关
		auth.GET(prx+"/admin/menu/getAll", admin_controller.GetMenuList)
		auth.GET(prx+"/admin/menu/getTree", admin_controller.GetMenuTree)
		auth.GET(prx+"/admin/menu/getMy", admin_controller.GetMyMenuList)
		auth.GET(prx+"/admin/menu/detail", admin_controller.GetMenuDetail)
		auth.POST(prx+"/admin/menu/checkKey", admin_controller.CheckMenuKey)
		auth.POST(prx+"/admin/menu/save", admin_controller.MenuSave)
		auth.POST(prx+"/admin/menu/delete", admin_controller.MenuDelete)
		auth.POST(prx+"/admin/menu/setShow", admin_controller.MenuShow)
		auth.POST(prx+"/admin/menu/setSort", admin_controller.MenuSortSet)

		//权限相关
		auth.POST(prx+"/admin/role/save", admin_controller.MenuSortSet)
		auth.POST(prx+"/admin/role/delete", admin_controller.MenuSortSet)
		auth.POST(prx+"/admin/role/setLock", admin_controller.MenuSortSet)
		auth.POST(prx+"/admin/role/getList", admin_controller.MenuSortSet)

		//用户相关
		auth.POST(prx+"/admin/user/save", admin_controller.SaveAdmin)
		auth.POST(prx+"/admin/user/setLock", admin_controller.SetAdminStatus)
		auth.POST(prx+"/admin/user/setRole", admin_controller.SetAdminRole)
		auth.POST(prx+"/admin/user/delete", admin_controller.DeleteAdmin)
		auth.GET(prx+"/admin/user/getList", admin_controller.GetAdminList)
		auth.GET(prx+"/admin/user/getLoginRecordList", admin_controller.GetAdminLoginRecord)

		//用户登录记录相关
		auth.POST(prx+"/admin/loginRecord/getList", admin_controller.MenuSortSet)

		//验证登录权限
		auth.POST(prx+"/auth/check", admin_controller.Auth)

		//demo调试接口
		auth.GET(prx+"/user/info", admin_controller.GetUserInfo)
		auth.GET(prx+"/role", admin_controller.GetUserRole)
		auth.GET(prx+"/service", admin_controller.GetService)
		auth.GET(prx+"/list/search/projects", admin_controller.GetProject)
		auth.GET(prx+"/workplace/activity", admin_controller.GetActivity)
		auth.GET(prx+"/workplace/teams", admin_controller.GetTeams)
		auth.GET(prx+"/workplace/radar", admin_controller.GetRadar)
	}
	return r
}
