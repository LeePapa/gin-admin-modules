package admin_model

import (
	"github.com/jinzhu/gorm"
)

// 获取查询到的第一条
func (a AdminInfo) Get(fields string, query interface{}, args ...interface{}) (AdminInfo, bool) {
	var ret AdminInfo
	err := AdminDb.Select(fields).Where(query, args...).First(&ret).Error
	if err != nil {
		return AdminInfo{}, false
	}
	return ret, true
}

// 获取列表
func (a AdminInfo) GetList(fields string, order string, page int, nums int, query interface{}, args ...interface{}) ([]AdminInfo, bool) {
	var ret []AdminInfo
	err := AdminDb.Select(fields).Where(query, args...).Order(order).Limit(nums).Offset((page - 1) * nums).Find(&ret).Error
	if err != nil {
		return []AdminInfo{}, false
	}
	return ret, true
}

// 获取查询到的最后一条数据
func (a AdminInfo) GetLast(fields string, query interface{}, args ...interface{}) (AdminInfo, bool) {
	var ret AdminInfo
	err := AdminDb.Select(fields).Where(query, args...).Last(&ret).Error
	if err != nil {
		return AdminInfo{}, false
	}
	return ret, true
}

// 写入记录
func (a AdminInfo) Create(data AdminInfo) int {
	err := AdminDb.Omit("CreatedTime", "UpdatedTime").Create(&data).Error
	if err != nil {
		return 0
	}
	return data.ID
}

// 更新数据
func (a AdminInfo) Update(data map[string]interface{}, query interface{}, args ...interface{}) bool {
	err := AdminDb.Model(a).Where(query, args...).Updates(data).Error
	if err != nil {
		return false
	}
	return true
}

// 合计数据
func (a AdminInfo) Sum(field string, query interface{}, args ...interface{}) int {
	var ret int
	err := AdminDb.Table(a.TableName()).Select("sum(" + field + ") as nums").Pluck("nums", &ret).Error
	if err != nil {
		return 0
	}
	return ret
}

// 统计数据
func (a AdminInfo) Count(field string, query interface{}, args ...interface{}) int {
	var ret int
	err := AdminDb.Table(a.TableName()).Select("count(" + field + ") as nums").Pluck("nums", &ret).Error
	if err != nil {
		return 0
	}
	return ret
}

// 自增
func (a AdminInfo) Increment(field string, num int, query interface{}, args ...interface{}) bool {
	err := AdminDb.Model(a).Where(query, args...).UpdateColumn(field, gorm.Expr(field+" + ?", num)).Error
	if err != nil {
		return false
	}
	return true
}

// 自减
func (a AdminInfo) Decrement(field string, num int, query interface{}, args ...interface{}) bool {
	err := AdminDb.Model(a).Where(query, args...).UpdateColumn(field, gorm.Expr(field+" - ?", num)).Error
	if err != nil {
		return false
	}
	return true
}

func (a AdminInfo) GetAdminByUsername(username string, fields string) (AdminInfo, bool) {
	adminInfo, has := a.Get(fields, "username = ? AND is_del = ?", username, 0)
	if !has {
		return AdminInfo{}, false
	}
	return adminInfo, true
}

func (a AdminInfo) GetModel() *gorm.DB {
	return AdminDb
}
