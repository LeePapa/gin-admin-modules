package admin_controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"encoding/json"
	"gin-modules/modules/admin/service"
	"gin-modules/modules/admin/model"
	"gin-modules/middleware"
)

type loginRule struct {
	Username string `form:"username" binding:"required,numeric,len=11"`
	Password string `form:"password" binding:"required"`
}

type loginByCodeRule struct {
	Username string `form:"mobile" binding:"required,numeric,len=11"`
	Code     string `form:"code" binding:"required,numeric,len=6"`
}

type registerRule struct {
	Username string `form:"username" binding:"required,numeric,len=11"`
	Code     string `form:"code" binding:"required,numeric,len=6"`
	Password string `form:"password" binding:"required"`
}

type smsCodeRule struct {
	Mobile  string `form:"mobile" binding:"required,numeric,len=11"`
	Type    int    `form:"type" binding:"required"`
	Ticket  string `form:"ticket" binding:"required"`
	RandStr string `form:"randStr" binding:"required"`
}

//登录操作接口
func Login(ct *gin.Context) {
	var token string
	var err error
	switch ct.PostForm("loginType") {
	case "1":
		var mobileRule loginByCodeRule
		if ct.ShouldBind(&mobileRule) != nil {
			OutputJson(ct, -1, "账号或者验证码格式不正确", struct{}{})
			return
		}
		token, err = admin_service.LoginByCode(mobileRule.Username, mobileRule.Code, ct.ClientIP())
		break
	case "2":
		var usernameRule loginRule
		if ct.ShouldBind(&usernameRule) != nil {
			OutputJson(ct, -1, "账号或者密码格式不正确", struct{}{})
			return
		}
		token, err = admin_service.Login(usernameRule.Username, usernameRule.Password, ct.ClientIP())
		break
	}
	if err != nil {
		OutputJson(ct, -1, err.Error(), struct{}{})
		return
	}
	OutputJson(ct, 200, "登录成功", map[string]string{"token": token})
}

//用户注册接口
func Register(ct *gin.Context) {
	var rule registerRule
	if ct.ShouldBind(&rule) != nil {
		OutputJson(ct, -1, "参数异常", struct{}{})
		return
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(rule.Password), bcrypt.DefaultCost)
	if err != nil {
		OutputJson(ct, -1, "密码加密异常", struct{}{})
		return
	}
	err = admin_service.Register(admin_model.AdminInfo{Username: rule.Username, Password: string(passwordHash)}, rule.Code)
	if err != nil {
		OutputJson(ct, -1, err.Error(), struct{}{})
		return
	}
	OutputJson(ct, 200, "用户注册成功", struct{}{})
}

//发送验证码接口（1注册 2登录）
func SendSmsCode(ct *gin.Context) {
	var rule smsCodeRule
	if ct.ShouldBind(&rule) != nil {
		OutputJson(ct, -1, "参数异常", struct{}{})
		return
	}
	if admin_service.CheckTicket(rule.Ticket, rule.RandStr, ct.ClientIP()) != nil {
		OutputJson(ct, -1, "人机验证错误", struct{}{})
		return
	}
	err := admin_service.SendSmsCode(rule.Type, rule.Mobile, ct.ClientIP())
	if err != nil {
		OutputJson(ct, -1, err.Error(), struct{}{})
		return
	}
	OutputJson(ct, 200, "发送成功", struct{}{})
}

func Auth(ct *gin.Context) {
	tokenInfo, _ := ct.Get("claims")
	print(tokenInfo.(*middleware.CustomClaims).ID)
	OutputJson(ct, 200, "发送成功", tokenInfo)
}

func GetUserInfo(ct *gin.Context) {
	var ret interface{}
	json.Unmarshal([]byte(`{"message":"","result":{"id":"4291d7da9005377ec9aec4a71ea837f","name":"天野远子","username":"admin","password":"","avatar":"/avatar2.jpg","status":1,"telephone":"","lastLoginIp":"27.154.74.117","lastLoginTime":1534837621348,"creatorId":"admin","createTime":1497160610259,"merchantCode":"TLif2btpzg079h15bk","deleted":0,"roleId":"admin","role":{"id":"admin","name":"管理员","describe":"拥有所有权限","status":1,"creatorId":"system","createTime":1497160610259,"deleted":0,"permissions":[{"roleId":"admin","permissionId":"dashboard","permissionName":"仪表盘","actions":"[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"query\",\"defaultCheck\":false,\"describe\":\"查询\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"},{\"action\":\"delete\",\"defaultCheck\":false,\"describe\":\"删除\"}]","actionEntitySet":[{"action":"add","describe":"新增","defaultCheck":false},{"action":"query","describe":"查询","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false},{"action":"update","describe":"修改","defaultCheck":false},{"action":"delete","describe":"删除","defaultCheck":false}],"actionList":null,"dataAccess":null},{"roleId":"admin","permissionId":"exception","permissionName":"异常页面权限","actions":"[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"query\",\"defaultCheck\":false,\"describe\":\"查询\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"},{\"action\":\"delete\",\"defaultCheck\":false,\"describe\":\"删除\"}]","actionEntitySet":[{"action":"add","describe":"新增","defaultCheck":false},{"action":"query","describe":"查询","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false},{"action":"update","describe":"修改","defaultCheck":false},{"action":"delete","describe":"删除","defaultCheck":false}],"actionList":null,"dataAccess":null},{"roleId":"admin","permissionId":"result","permissionName":"结果权限","actions":"[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"query\",\"defaultCheck\":false,\"describe\":\"查询\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"},{\"action\":\"delete\",\"defaultCheck\":false,\"describe\":\"删除\"}]","actionEntitySet":[{"action":"add","describe":"新增","defaultCheck":false},{"action":"query","describe":"查询","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false},{"action":"update","describe":"修改","defaultCheck":false},{"action":"delete","describe":"删除","defaultCheck":false}],"actionList":null,"dataAccess":null},{"roleId":"admin","permissionId":"profile","permissionName":"详细页权限","actions":"[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"query\",\"defaultCheck\":false,\"describe\":\"查询\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"},{\"action\":\"delete\",\"defaultCheck\":false,\"describe\":\"删除\"}]","actionEntitySet":[{"action":"add","describe":"新增","defaultCheck":false},{"action":"query","describe":"查询","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false},{"action":"update","describe":"修改","defaultCheck":false},{"action":"delete","describe":"删除","defaultCheck":false}],"actionList":null,"dataAccess":null},{"roleId":"admin","permissionId":"table","permissionName":"表格权限","actions":"[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"import\",\"defaultCheck\":false,\"describe\":\"导入\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"}]","actionEntitySet":[{"action":"add","describe":"新增","defaultCheck":false},{"action":"import","describe":"导入","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false},{"action":"update","describe":"修改","defaultCheck":false}],"actionList":null,"dataAccess":null},{"roleId":"admin","permissionId":"form","permissionName":"表单权限","actions":"[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"query\",\"defaultCheck\":false,\"describe\":\"查询\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"},{\"action\":\"delete\",\"defaultCheck\":false,\"describe\":\"删除\"}]","actionEntitySet":[{"action":"add","describe":"新增","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false},{"action":"query","describe":"查询","defaultCheck":false},{"action":"update","describe":"修改","defaultCheck":false},{"action":"delete","describe":"删除","defaultCheck":false}],"actionList":null,"dataAccess":null},{"roleId":"admin","permissionId":"order","permissionName":"订单管理","actions":"[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"query\",\"defaultCheck\":false,\"describe\":\"查询\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"},{\"action\":\"delete\",\"defaultCheck\":false,\"describe\":\"删除\"}]","actionEntitySet":[{"action":"add","describe":"新增","defaultCheck":false},{"action":"query","describe":"查询","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false},{"action":"update","describe":"修改","defaultCheck":false},{"action":"delete","describe":"删除","defaultCheck":false}],"actionList":null,"dataAccess":null},{"roleId":"admin","permissionId":"permission","permissionName":"权限管理","actions":"[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"},{\"action\":\"delete\",\"defaultCheck\":false,\"describe\":\"删除\"}]","actionEntitySet":[{"action":"add","describe":"新增","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false},{"action":"update","describe":"修改","defaultCheck":false},{"action":"delete","describe":"删除","defaultCheck":false}],"actionList":null,"dataAccess":null},{"roleId":"admin","permissionId":"role","permissionName":"角色管理","actions":"[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"},{\"action\":\"delete\",\"defaultCheck\":false,\"describe\":\"删除\"}]","actionEntitySet":[{"action":"add","describe":"新增","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false},{"action":"update","describe":"修改","defaultCheck":false},{"action":"delete","describe":"删除","defaultCheck":false}],"actionList":null,"dataAccess":null},{"roleId":"admin","permissionId":"table","permissionName":"桌子管理","actions":"[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"query\",\"defaultCheck\":false,\"describe\":\"查询\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"},{\"action\":\"delete\",\"defaultCheck\":false,\"describe\":\"删除\"}]","actionEntitySet":[{"action":"add","describe":"新增","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false},{"action":"query","describe":"查询","defaultCheck":false},{"action":"update","describe":"修改","defaultCheck":false},{"action":"delete","describe":"删除","defaultCheck":false}],"actionList":null,"dataAccess":null},{"roleId":"admin","permissionId":"user","permissionName":"用户管理","actions":"[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"import\",\"defaultCheck\":false,\"describe\":\"导入\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"},{\"action\":\"delete\",\"defaultCheck\":false,\"describe\":\"删除\"},{\"action\":\"export\",\"defaultCheck\":false,\"describe\":\"导出\"}]","actionEntitySet":[{"action":"add","describe":"新增","defaultCheck":false},{"action":"import","describe":"导入","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false},{"action":"update","describe":"修改","defaultCheck":false},{"action":"delete","describe":"删除","defaultCheck":false},{"action":"export","describe":"导出","defaultCheck":false}],"actionList":null,"dataAccess":null}]}},"status":200,"timestamp":1534844188679}`), &ret)
	ct.JSON(200, ret)
}

func GetUserRole(ct *gin.Context) {
	var ret interface{}
	json.Unmarshal([]byte(`{"message":"","result":{"data":[{"id":"admin","name":"管理员","describe":"拥有所有权限","status":1,"creatorId":"system","createTime":1497160610259,"deleted":0,"permissions":[{"roleId":"admin","permissionId":"comment","permissionName":"评论管理","actions":"[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"query\",\"defaultCheck\":false,\"describe\":\"查询\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"},{\"action\":\"delete\",\"defaultCheck\":false,\"describe\":\"删除\"}]","actionEntitySet":[{"action":"add","describe":"新增","defaultCheck":false},{"action":"query","describe":"查询","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false},{"action":"update","describe":"修改","defaultCheck":false},{"action":"delete","describe":"删除","defaultCheck":false}],"actionList":null,"dataAccess":null},{"roleId":"admin","permissionId":"member","permissionName":"会员管理","actions":"[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"query\",\"defaultCheck\":false,\"describe\":\"查询\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"},{\"action\":\"delete\",\"defaultCheck\":false,\"describe\":\"删除\"}]","actionEntitySet":[{"action":"add","describe":"新增","defaultCheck":false},{"action":"query","describe":"查询","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false},{"action":"update","describe":"修改","defaultCheck":false},{"action":"delete","describe":"删除","defaultCheck":false}],"actionList":null,"dataAccess":null},{"roleId":"admin","permissionId":"menu","permissionName":"菜单管理","actions":"[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"import\",\"defaultCheck\":false,\"describe\":\"导入\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"}]","actionEntitySet":[{"action":"add","describe":"新增","defaultCheck":false},{"action":"import","describe":"导入","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false},{"action":"update","describe":"修改","defaultCheck":false}],"actionList":null,"dataAccess":null},{"roleId":"admin","permissionId":"order","permissionName":"订单管理","actions":"[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"query\",\"defaultCheck\":false,\"describe\":\"查询\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"},{\"action\":\"delete\",\"defaultCheck\":false,\"describe\":\"删除\"}]","actionEntitySet":[{"action":"add","describe":"新增","defaultCheck":false},{"action":"query","describe":"查询","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false},{"action":"update","describe":"修改","defaultCheck":false},{"action":"delete","describe":"删除","defaultCheck":false}],"actionList":null,"dataAccess":null},{"roleId":"admin","permissionId":"permission","permissionName":"权限管理","actions":"[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"},{\"action\":\"delete\",\"defaultCheck\":false,\"describe\":\"删除\"}]","actionEntitySet":[{"action":"add","describe":"新增","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false},{"action":"update","describe":"修改","defaultCheck":false},{"action":"delete","describe":"删除","defaultCheck":false}],"actionList":null,"dataAccess":null},{"roleId":"admin","permissionId":"role","permissionName":"角色管理","actions":"[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"},{\"action\":\"delete\",\"defaultCheck\":false,\"describe\":\"删除\"}]","actionEntitySet":[{"action":"add","describe":"新增","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false},{"action":"update","describe":"修改","defaultCheck":false},{"action":"delete","describe":"删除","defaultCheck":false}],"actionList":null,"dataAccess":null},{"roleId":"admin","permissionId":"test","permissionName":"测试权限","actions":"[]","actionEntitySet":[],"actionList":null,"dataAccess":null},{"roleId":"admin","permissionId":"user","permissionName":"用户管理","actions":"[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"import\",\"defaultCheck\":false,\"describe\":\"导入\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"},{\"action\":\"delete\",\"defaultCheck\":false,\"describe\":\"删除\"},{\"action\":\"export\",\"defaultCheck\":false,\"describe\":\"导出\"}]","actionEntitySet":[{"action":"add","describe":"新增","defaultCheck":false},{"action":"import","describe":"导入","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false},{"action":"update","describe":"修改","defaultCheck":false},{"action":"delete","describe":"删除","defaultCheck":false},{"action":"export","describe":"导出","defaultCheck":false}],"actionList":null,"dataAccess":null}]},{"id":"svip","name":"SVIP","describe":"超级会员","status":1,"creatorId":"system","createTime":1532417744846,"deleted":0,"permissions":[{"roleId":"admin","permissionId":"comment","permissionName":"评论管理","actions":"[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"query\",\"defaultCheck\":false,\"describe\":\"查询\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"},{\"action\":\"delete\",\"defaultCheck\":false,\"describe\":\"删除\"}]","actionEntitySet":[{"action":"add","describe":"新增","defaultCheck":false},{"action":"query","describe":"查询","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false},{"action":"update","describe":"修改","defaultCheck":false},{"action":"delete","describe":"删除","defaultCheck":false}],"actionList":null,"dataAccess":null},{"roleId":"admin","permissionId":"member","permissionName":"会员管理","actions":"[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"query\",\"defaultCheck\":false,\"describe\":\"查询\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"},{\"action\":\"delete\",\"defaultCheck\":false,\"describe\":\"删除\"}]","actionEntitySet":[{"action":"add","describe":"新增","defaultCheck":false},{"action":"query","describe":"查询","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false}],"actionList":null,"dataAccess":null},{"roleId":"admin","permissionId":"menu","permissionName":"菜单管理","actions":"[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"import\",\"defaultCheck\":false,\"describe\":\"导入\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"}]","actionEntitySet":[{"action":"add","describe":"新增","defaultCheck":false},{"action":"import","describe":"导入","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false}],"actionList":null,"dataAccess":null},{"roleId":"admin","permissionId":"order","permissionName":"订单管理","actions":"[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"query\",\"defaultCheck\":false,\"describe\":\"查询\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"},{\"action\":\"delete\",\"defaultCheck\":false,\"describe\":\"删除\"}]","actionEntitySet":[{"action":"add","describe":"新增","defaultCheck":false},{"action":"query","describe":"查询","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false},{"action":"update","describe":"修改","defaultCheck":false}],"actionList":null,"dataAccess":null},{"roleId":"admin","permissionId":"permission","permissionName":"权限管理","actions":"[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"},{\"action\":\"delete\",\"defaultCheck\":false,\"describe\":\"删除\"}]","actionEntitySet":[{"action":"add","describe":"新增","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false},{"action":"update","describe":"修改","defaultCheck":false}],"actionList":null,"dataAccess":null},{"roleId":"admin","permissionId":"role","permissionName":"角色管理","actions":"[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"},{\"action\":\"delete\",\"defaultCheck\":false,\"describe\":\"删除\"}]","actionEntitySet":[{"action":"add","describe":"新增","defaultCheck":false},{"action":"update","describe":"修改","defaultCheck":false},{"action":"delete","describe":"删除","defaultCheck":false}],"actionList":null,"dataAccess":null},{"roleId":"admin","permissionId":"test","permissionName":"测试权限","actions":"[]","actionEntitySet":[],"actionList":null,"dataAccess":null},{"roleId":"admin","permissionId":"user","permissionName":"用户管理","actions":"[{\"action\":\"add\",\"defaultCheck\":false,\"describe\":\"新增\"},{\"action\":\"import\",\"defaultCheck\":false,\"describe\":\"导入\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"},{\"action\":\"update\",\"defaultCheck\":false,\"describe\":\"修改\"},{\"action\":\"delete\",\"defaultCheck\":false,\"describe\":\"删除\"},{\"action\":\"export\",\"defaultCheck\":false,\"describe\":\"导出\"}]","actionEntitySet":[{"action":"add","describe":"新增","defaultCheck":false},{"action":"import","describe":"导入","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false},{"action":"update","describe":"修改","defaultCheck":false}],"actionList":null,"dataAccess":null}]},{"id":"user","name":"普通会员","describe":"普通用户，只能查询","status":1,"creatorId":"system","createTime":1497160610259,"deleted":0,"permissions":[{"roleId":"user","permissionId":"comment","permissionName":"评论管理","actions":"[{\"action\":\"query\",\"defaultCheck\":false,\"describe\":\"查询\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"}]","actionEntitySet":[{"action":"query","describe":"查询","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false}],"actionList":null,"dataAccess":null},{"roleId":"user","permissionId":"marketing","permissionName":"营销管理","actions":"[]","actionEntitySet":[],"actionList":null,"dataAccess":null},{"roleId":"user","permissionId":"member","permissionName":"会员管理","actions":"[{\"action\":\"query\",\"defaultCheck\":false,\"describe\":\"查询\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"}]","actionEntitySet":[{"action":"query","describe":"查询","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false}],"actionList":null,"dataAccess":null},{"roleId":"user","permissionId":"menu","permissionName":"菜单管理","actions":"[]","actionEntitySet":[],"actionList":null,"dataAccess":null},{"roleId":"user","permissionId":"order","permissionName":"订单管理","actions":"[{\"action\":\"query\",\"defaultCheck\":false,\"describe\":\"查询\"},{\"action\":\"get\",\"defaultCheck\":false,\"describe\":\"详情\"}]","actionEntitySet":[{"action":"query","describe":"查询","defaultCheck":false},{"action":"get","describe":"详情","defaultCheck":false}],"actionList":null,"dataAccess":null},{"roleId":"user","permissionId":"permission","permissionName":"权限管理","actions":"[]","actionEntitySet":[],"actionList":null,"dataAccess":null},{"roleId":"user","permissionId":"role","permissionName":"角色管理","actions":"[]","actionEntitySet":[],"actionList":null,"dataAccess":null},{"roleId":"user","permissionId":"test","permissionName":"测试权限","actions":"[]","actionEntitySet":[],"actionList":null,"dataAccess":null},{"roleId":"user","permissionId":"user","permissionName":"用户管理","actions":"[]","actionEntitySet":[],"actionList":null,"dataAccess":null}]}],"pageSize":10,"pageNo":0,"totalPage":1,"totalCount":5},"status":200,"timestamp":1537079497645}`), &ret)
	ct.JSON(200, ret)
}

func GetService(ct *gin.Context) {
	var ret interface{}
	json.Unmarshal([]byte(`{"message":"","result":{"pageSize":null,"pageNo":null,"totalCount":57,"totalPage":null,"data":[]},"status":200,"timestamp":1534955098193}`), &ret)
	ct.JSON(200, ret)
}
func GetProject(ct *gin.Context) {
	var ret interface{}
	json.Unmarshal([]byte(`{"message":"","result":{"data":[{"id":1,"cover":"https://gw.alipayobjects.com/zos/rmsportal/WdGqmHpayyMjiEhcKoVE.png","title":"Alipay","description":"那是一种内在的东西， 他们到达不了，也无法触及的","status":1,"updatedAt":"2018-07-26 00:00:00"},{"id":2,"cover":"https://gw.alipayobjects.com/zos/rmsportal/zOsKZmFRdUtvpqCImOVY.png","title":"Angular","description":"希望是一个好东西，也许是最好的，好东西是不会消亡的","status":1,"updatedAt":"2018-07-26 00:00:00"},{"id":3,"cover":"https://gw.alipayobjects.com/zos/rmsportal/dURIMkkrRFpPgTuzkwnB.png","title":"Ant Design","description":"城镇中有那么多的酒馆，她却偏偏走进了我的酒馆","status":1,"updatedAt":"2018-07-26 00:00:00"},{"id":4,"cover":"https://gw.alipayobjects.com/zos/rmsportal/sfjbOqnsXXJgNCjCzDBL.png","title":"Ant Design Pro","description":"那时候我只会想自己想要什么，从不想自己拥有什么","status":1,"updatedAt":"2018-07-26 00:00:00"},{"id":5,"cover":"https://gw.alipayobjects.com/zos/rmsportal/siCrBXXhmvTQGWPNLBow.png","title":"Bootstrap","description":"凛冬将至","status":1,"updatedAt":"2018-07-26 00:00:00"},{"id":6,"cover":"https://gw.alipayobjects.com/zos/rmsportal/ComBAopevLwENQdKWiIn.png","title":"Vue","description":"生命就像一盒巧克力，结果往往出人意料","status":1,"updatedAt":"2018-07-26 00:00:00"}],"pageSize":10,"pageNo":0,"totalPage":6,"totalCount":57},"status":200,"timestamp":1534955098193}`), &ret)
	ct.JSON(200, ret)
}
func GetActivity(ct *gin.Context) {
	var ret interface{}
	json.Unmarshal([]byte(`{"message":"","result":[{"id":1,"user":{"nickname":"Melissa Martin","avatar":"https://gw.alipayobjects.com/zos/rmsportal/BiazfanxmamNRoxxVxka.png"},"project":{"name":"白鹭酱油开发组","action":"更新","event":"番组计划"},"time":"2018-08-23 14:47:00"},{"id":1,"user":{"nickname":"蓝莓酱","avatar":"https://gw.alipayobjects.com/zos/rmsportal/jZUIxmJycoymBprLOUbT.png"},"project":{"name":"白鹭酱油开发组","action":"更新","event":"番组计划"},"time":"2018-08-23 09:35:37"},{"id":1,"user":{"nickname":"Joseph Garcia","avatar":"http://dummyimage.com/64x64"},"project":{"name":"白鹭酱油开发组","action":"创建","event":"番组计划"},"time":"2017-05-27 00:00:00"},{"id":1,"user":{"nickname":"曲丽丽","avatar":"http://dummyimage.com/64x64"},"project":{"name":"高逼格设计天团","action":"更新","event":"六月迭代"},"time":"2018-08-23 14:47:00"},{"id":1,"user":{"nickname":"Deborah Moore","avatar":"http://dummyimage.com/64x64"},"project":{"name":"高逼格设计天团","action":"created","event":"六月迭代"},"time":"2018-08-23 14:47:00"},{"id":1,"user":{"nickname":"曲丽丽","avatar":"https://gw.alipayobjects.com/zos/rmsportal/BiazfanxmamNRoxxVxka.png"},"project":{"name":"高逼格设计天团","action":"created","event":"六月迭代"},"time":"2018-08-23 14:47:00"}],"status":200,"timestamp":0}`), &ret)
	ct.JSON(200, ret)
}
func GetTeams(ct *gin.Context) {
	var ret interface{}
	json.Unmarshal([]byte(`{"message":"","result":[{"id":1,"name":"科学搬砖组","avatar":"https://gw.alipayobjects.com/zos/rmsportal/BiazfanxmamNRoxxVxka.png"},{"id":2,"name":"程序员日常","avatar":"https://gw.alipayobjects.com/zos/rmsportal/cnrhVkzwxjPwAaCfPbdc.png"},{"id":1,"name":"设计天团","avatar":"https://gw.alipayobjects.com/zos/rmsportal/gaOngJwsRYRaVAuXXcmB.png"},{"id":1,"name":"中二少女团","avatar":"https://gw.alipayobjects.com/zos/rmsportal/ubnKSIfAJTxIgXOKlciN.png"},{"id":1,"name":"骗你学计算机","avatar":"https://gw.alipayobjects.com/zos/rmsportal/WhxKECPNujWoWEFNdnJE.png"}],"status":200,"timestamp":0}`), &ret)
	ct.JSON(200, ret)
}
func GetRadar(ct *gin.Context) {
	var ret interface{}
	json.Unmarshal([]byte(`{"message":"","result":[{"item":"引用","个人":70,"团队":30,"部门":40},{"item":"口碑","个人":60,"团队":70,"部门":40},{"item":"产量","个人":50,"团队":60,"部门":40},{"item":"贡献","个人":40,"团队":50,"部门":40},{"item":"热度","个人":60,"团队":70,"部门":40},{"item":"引用","个人":70,"团队":50,"部门":40}],"status":200,"timestamp":1534955098193}`), &ret)
	ct.JSON(200, ret)
}
