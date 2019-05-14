package admin_controller

import (
	"github.com/gin-gonic/gin"
	"gin-modules/modules/admin/model"
	"math"
	"gin-modules/middleware"
	"golang.org/x/crypto/bcrypt"
)

type adminSaveRule struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Nickname string `form:"nickname"`
	HeadImg  string `form:"head_img"`
}

type loginRecordRule struct {
	Page      int    `form:"page"`
	Size      int    `form:"size"`
	Username  string `form:"queryUsername"`
	AdminId   int    `form:"queryAdminId"`
	Type      int    `form:"queryType"`
	StartTime string `form:"queryStartTime"`
	EndTime   string `form:"queryEndTime"`
}

type adminListRule struct {
	Page      int    `form:"page"`
	Size      int    `form:"size"`
	Username  string `form:"queryUsername"`
	Nickname  string `form:"queryNickname"`
	Status    int    `form:"queryStatus"`
	AdminId   int    `form:"queryId"`
	StartTime string `form:"queryStartTime"`
	EndTime   string `form:"queryEndTime"`
}

//获取管理员登录记录列表
func GetAdminLoginRecord(ct *gin.Context) {
	var rule loginRecordRule
	ct.ShouldBindQuery(&rule)
	if rule.Size == 0 {
		rule.Size = 100
	}
	if rule.Page == 0 {
		rule.Page = 1
	}
	table := new(admin_model.AdminLoginRecord).TableName()
	countModel := admin_model.AdminDb.Table(table).Joins("LEFT JOIN admin_info ON admin_info.id = admin_login_record.admin_id")
	dataModel := admin_model.AdminDb.Table(table).
		Joins("LEFT JOIN admin_info ON admin_info.id = admin_login_record.admin_id").
		Select("admin_login_record.id,admin_info.username,admin_login_record.type,admin_login_record.ip,admin_login_record.address,admin_login_record.created_time").
		Order("id DESC").
		Limit(rule.Size).
		Offset((rule.Page - 1) * rule.Size)
	if rule.Username != "" {
		dataModel = dataModel.Where("admin_info.username LIKE ?", "%"+rule.Username+"%")
		countModel = countModel.Where("admin_info.username LIKE ?", "%"+rule.Username+"%")
	}
	if rule.Type != 0 {
		dataModel = dataModel.Where("admin_login_record.type = ?", rule.Type)
		countModel = countModel.Where("admin_login_record.type = ?", rule.Type)
	}
	if rule.AdminId != 0 {
		dataModel = dataModel.Where("admin_login_record.admin_id = ?", rule.AdminId)
		countModel = countModel.Where("admin_login_record.admin_id = ?", rule.AdminId)
	}
	if rule.StartTime != "" {
		dataModel = dataModel.Where("admin_login_record.created_time >= ?", rule.StartTime)
		countModel = countModel.Where("admin_login_record.created_time >= ?", rule.StartTime)
	}
	if rule.EndTime != "" {
		dataModel = dataModel.Where("admin_login_record.created_time <= ?", rule.EndTime)
		countModel = countModel.Where("admin_login_record.created_time <= ?", rule.EndTime)
	}
	data := make(map[string]interface{})
	retCount := 0
	data["rows"] = []interface{}{}
	data["count"] = 0
	rows, err := dataModel.Rows()
	defer rows.Close()
	if err == nil && rows.Err() == nil {
		var rowsData []interface{}
		idColumn, usernameColumn, typeColumn, ipColumn, addressColumn, createdTimeColumn := 0, "", 1, "", "", ""
		for rows.Next() {
			rows.Scan(&idColumn, &usernameColumn, &typeColumn, &ipColumn, &addressColumn, &createdTimeColumn)
			rowsData = append(rowsData, map[string]interface{}{
				"id":           idColumn,
				"username":     usernameColumn,
				"type":         typeColumn,
				"ip":           ipColumn,
				"address":      addressColumn,
				"created_time": createdTimeColumn,
			})
		}
		if rowsData != nil {
			data["rows"] = rowsData
		}
	}
	if countModel.Count(&retCount).Error == nil {
		data["count"] = retCount
	}
	data["max_page"] = int(math.Ceil(float64(retCount) / float64(rule.Size)))
	OutputJson(ct, 200, "success", data)
}

//获取管理员列表
func GetAdminList(ct *gin.Context) {
	var rule adminListRule
	ct.ShouldBindQuery(&rule)
	if rule.Size == 0 {
		rule.Size = 15
	}
	if rule.Page == 0 {
		rule.Size = 1
	}

	var ret []admin_model.AdminInfo
	var retCount int
	table := new(admin_model.AdminInfo).TableName()
	adminCountModel := admin_model.AdminDb.Table(table)
	adminInfoModel := admin_model.AdminDb.Table(table).Select("id,role_id,username,nickname,head_img,status,created_time").Order("id DESC").Limit(rule.Size).Offset((rule.Page - 1) * rule.Size)
	if rule.Username != "" {
		adminInfoModel = adminInfoModel.Where("username LIKE ?", "%"+rule.Username+"%")
		adminCountModel = adminCountModel.Where("username LIKE ?", "%"+rule.Username+"%")
	}
	if rule.Nickname != "" {
		adminInfoModel = adminInfoModel.Where("nickname LIKE ?", "%"+rule.Nickname+"%")
		adminCountModel = adminCountModel.Where("nickname LIKE ?", "%"+rule.Nickname+"%")
	}
	if rule.AdminId != 0 {
		adminInfoModel = adminInfoModel.Where("id = ?", rule.AdminId)
		adminCountModel = adminCountModel.Where("id = ?", rule.AdminId)
	}
	if rule.Status != -1 {
		adminInfoModel = adminInfoModel.Where("status = ?", rule.Status)
		adminCountModel = adminCountModel.Where("status = ?", rule.Status)
	}
	if rule.StartTime != "" {
		adminInfoModel = adminInfoModel.Where("created_time >= ?", rule.StartTime)
		adminCountModel = adminCountModel.Where("created_time >= ?", rule.StartTime)
	}
	if rule.EndTime != "" {
		adminInfoModel = adminInfoModel.Where("created_time <= ?", rule.EndTime)
		adminCountModel = adminCountModel.Where("created_time <= ?", rule.EndTime)
	}
	data := make(map[string]interface{})
	data["rows"] = []interface{}{}
	data["count"] = 0
	if adminInfoModel.Find(&ret).Error == nil {
		//var admins []map[string]interface{}
		for _, adminInfo := range ret {
			print(adminInfo.ID)
		}
		data["rows"] = ret
	}
	if adminCountModel.Count(&retCount).Error == nil {
		data["count"] = retCount
	}
	data["max_page"] = int(math.Ceil(float64(retCount) / float64(rule.Size)))
	OutputJson(ct, 200, "success", data)
}

// 保存管理员更改信息
func SaveAdmin(ct *gin.Context) {
	var rule adminSaveRule
	if ct.ShouldBind(&rule) != nil {
		OutputJson(ct, -1, "参数异常", struct{}{})
		return
	}
	data := make(map[string]interface{})
	if rule.Nickname != "" {
		data["nickname"] = rule.Nickname
	}
	if rule.Password != "" {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(rule.Password), bcrypt.DefaultCost)
		if err != nil {
			OutputJson(ct, -1, "密码加密异常", struct{}{})
			return
		}
		data["password"] = passwordHash
	}
	if rule.Username != "" {
		data["username"] = rule.Username
	}
	if rule.HeadImg != "" {
		data["head_img"] = rule.HeadImg
	}
	tokenInfo, _ := ct.Get("claims")
	if !admin_model.UpdateAdminInfo(data, "id = ?", tokenInfo.(*middleware.CustomClaims).ID) {
		OutputJson(ct, -1, "管理员信息更新失败", struct{}{})
		return
	}
	OutputJson(ct, 200, "管理员信息更新成功", struct{}{})
}

// 设置管理员权限
func SetAdminRole(ct *gin.Context) {
	roleId := ct.Query("role_id")
	tokenInfo, _ := ct.Get("claims")
	if !admin_model.UpdateAdminInfo(map[string]interface{}{"role_id": roleId}, "id = ?", tokenInfo.(*middleware.CustomClaims).ID) {
		OutputJson(ct, -1, "更新管理员权限失败", struct{}{})
		return
	}
	OutputJson(ct, 200, "更新管理员权限成功", struct{}{})
}

// 设置管理员登录状态（锁定/正常）
func SetAdminStatus(ct *gin.Context) {
	status := ct.Query("status")
	tokenInfo, _ := ct.Get("claims")
	if !admin_model.UpdateAdminInfo(map[string]interface{}{"status": status}, "id = ?", tokenInfo.(*middleware.CustomClaims).ID) {
		OutputJson(ct, -1, "设置管理员状态失败", struct{}{})
		return
	}
	OutputJson(ct, 200, "设置管理员状态成功", struct{}{})
}

// 删除管理员（软删除）
func DeleteAdmin(ct *gin.Context) {
	tokenInfo, _ := ct.Get("claims")
	if !admin_model.UpdateAdminInfo(map[string]interface{}{"is_del": 1}, "id = ?", tokenInfo.(*middleware.CustomClaims).ID) {
		OutputJson(ct, -1, "删除管理员失败", struct{}{})
		return
	}
	OutputJson(ct, 200, "删除管理员成功", struct{}{})
}
