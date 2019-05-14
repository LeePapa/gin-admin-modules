package admin_model

// 通过用户名获取管理员信息
func GetAdminByUsername(username string, fields string) (AdminInfo, bool) {
	var ret AdminInfo
	err := AdminDb.Select(fields).Where("username = ? AND is_del = ?", username, 0).First(&ret).Error
	if err != nil {
		return AdminInfo{}, false
	}
	return ret, true
}

// 添加管理员
func AddAdmin(data AdminInfo) int {
	err := AdminDb.Omit("CreatedTime", "UpdatedTime").Create(&data).Error
	if err != nil {
		return 0
	}
	return data.ID
}

// 更新管理员信息
func UpdateAdminInfo(data map[string]interface{}, query interface{}, args ...interface{}) bool {
	err := AdminDb.Model(AdminInfo{}).Where(query, args...).Updates(data).Error
	if err != nil {
		return false
	}
	return true
}

// 通过手机号码获取短信信息
func GetSmsByMobile(codeType int, mobile string, fields string) (SmsCode, bool) {
	var ret SmsCode
	err := AdminDb.Select(fields).Order("id DESC").Where("mobile = ? AND type = ?", mobile, codeType).First(&ret).Error
	if err != nil {
		return SmsCode{}, false
	}
	return ret, true
}

// 添加手机验证码信息
func AddSmsCode(data SmsCode) int {
	err := AdminDb.Omit("CreatedTime", "UpdatedTime").Create(&data).Error
	if err != nil {
		return 0
	}
	return data.ID
}

// 更新手机验证码信息状态
func UpdateSmsCode(data map[string]interface{}, query interface{}, args ...interface{}) bool {
	err := AdminDb.Model(SmsCode{}).Where(query, args...).Updates(data).Error
	if err != nil {
		return false
	}
	return true
}

// 通过排序添加菜单
func AddMenuBySort(data AdminMenu) int {
	var ret AdminMenu
	err := AdminDb.Select("sort").Where("pid = ? AND is_del = ?", data.Pid, 0).Order("`sort` ASC").First(&ret).Error
	data.Sort = 99
	if err == nil {
		data.Sort = ret.Sort - 1
	}
	err = AdminDb.Omit("CreatedTime", "UpdatedTime").Create(&data).Error
	if err != nil {
		return 0
	}
	return data.ID
}

// 通过排序获取菜单信息
func GetMenuByOrder(fields string, order string, query interface{}, args ...interface{}) (AdminMenu, bool) {
	var ret AdminMenu
	err := AdminDb.Select(fields).Where(query, args...).Order(order).First(&ret).Error
	if err != nil {
		return AdminMenu{}, false
	}
	return ret, true
}

// 获取菜单列表
func GetMenuList(fields string, order string, page int, nums int, query interface{}, args ...interface{}) ([]AdminMenu, bool) {
	var ret []AdminMenu
	err := AdminDb.Select(fields).Where(query, args...).Order(order).Limit(nums).Offset((page - 1) * nums).Find(&ret).Error
	if err != nil {
		return []AdminMenu{}, false
	}
	return ret, true
}

// 获取菜单信息
func GetMenuInfo(fields string, query interface{}, args ...interface{}) (AdminMenu, bool) {
	var ret AdminMenu
	err := AdminDb.Select(fields).Where(query, args...).First(&ret).Error
	if err != nil {
		return AdminMenu{}, false
	}
	return ret, true
}

// 更新菜单信息
func UpdateMenu(data map[string]interface{}, query interface{}, args ...interface{}) bool {
	err := AdminDb.Model(AdminMenu{}).Where(query, args...).Updates(data).Error
	if err != nil {
		return false
	}
	return true
}

// 添加管理员登录记录
func AddLoginRecord(data AdminLoginRecord) int {
	err := AdminDb.Omit("CreatedTime", "UpdatedTime").Create(&data).Error
	if err != nil {
		return 0
	}
	return data.ID
}

// 获取权限信息
func GetRoleInfo(fields string, query interface{}, args ...interface{}) (AdminRole, bool) {
	var ret AdminRole
	err := AdminDb.Select(fields).Where(query, args...).First(&ret).Error
	if err != nil {
		return AdminRole{}, false
	}
	return ret, true
}

// 更新手机验证码信息状态
func UpdateRoleInfo(data map[string]interface{}, query interface{}, args ...interface{}) bool {
	err := AdminDb.Model(AdminRole{}).Where(query, args...).Updates(data).Error
	if err != nil {
		return false
	}
	return true
}
