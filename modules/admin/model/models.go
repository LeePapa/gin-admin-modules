package admin_model

import (
	"database/sql"
	"github.com/guregu/null"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type AdminInfo struct {
	ID          int       `gorm:"column:id;primary_key" json:"id"`
	RoleID      int       `gorm:"column:role_id" json:"role_id"`
	Username    string    `gorm:"column:username" json:"username"`
	Password    string    `gorm:"column:password" json:"password"`
	Nickname    string    `gorm:"column:nickname" json:"nickname"`
	HeadImg     string    `gorm:"column:head_img" json:"head_img"`
	Status      int       `gorm:"column:status" json:"status"`
	CreatedTime time.Time `gorm:"column:created_time" json:"created_time"`
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updated_time"`
	IsDel       int       `gorm:"column:is_del" json:"is_del"`
}

// TableName sets the insert table name for this struct type
func (a *AdminInfo) TableName() string {
	return "admin_info"
}

type AdminLoginRecord struct {
	ID          int       `gorm:"column:id;primary_key" json:"id"`
	Type        int       `gorm:"column:type" json:"type"`
	AdminID     int       `gorm:"column:admin_id" json:"admin_id"`
	IP          string    `gorm:"column:ip" json:"ip"`
	Address     string    `gorm:"column:address" json:"address"`
	CreatedTime time.Time `gorm:"column:created_time" json:"created_time"`
}

// TableName sets the insert table name for this struct type
func (a *AdminLoginRecord) TableName() string {
	return "admin_login_record"
}

type AdminMenu struct {
	ID          int         `gorm:"column:id;primary_key" json:"id"`
	Pid         int         `gorm:"column:pid" json:"pid"`
	Level       int         `gorm:"column:level" json:"level"`
	Name        string      `gorm:"column:name" json:"name"`
	Icon        string      `gorm:"column:icon" json:"icon"`
	Routers     null.String `gorm:"column:routers" json:"routers"`
	Sort        int         `gorm:"column:sort" json:"sort"`
	Key         string      `gorm:"column:key" json:"key"`
	Path        string      `gorm:"column:path" json:"path"`
	IsDefault   int         `gorm:"column:is_default" json:"is_default"`
	Status      int         `gorm:"column:status" json:"status"`
	CreatedTime time.Time   `gorm:"column:created_time" json:"created_time"`
	UpdatedTime time.Time   `gorm:"column:updated_time" json:"updated_time"`
	IsDel       int         `gorm:"column:is_del" json:"is_del"`
}

// TableName sets the insert table name for this struct type
func (a *AdminMenu) TableName() string {
	return "admin_menu"
}

type AdminRole struct {
	ID          int       `gorm:"column:id;primary_key" json:"id"`
	AdminID     int       `gorm:"column:admin_id" json:"admin_id"`
	Name        string    `gorm:"column:name" json:"name"`
	MenuList    string    `gorm:"column:menu_list" json:"menu_list"`
	IsDefault   int       `gorm:"column:is_default" json:"is_default"`
	IsSuper     int       `gorm:"column:is_super" json:"is_super"`
	Status      int       `gorm:"column:status" json:"status"`
	CreatedTime time.Time `gorm:"column:created_time" json:"created_time"`
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updated_time"`
	IsDel       int       `gorm:"column:is_del" json:"is_del"`
}

// TableName sets the insert table name for this struct type
func (a *AdminRole) TableName() string {
	return "admin_roles"
}

type AdminToken struct {
	ID          int       `gorm:"column:id;primary_key" json:"id"`
	AdminID     int       `gorm:"column:admin_id" json:"admin_id"`
	Token       string    `gorm:"column:token" json:"token"`
	IP          string    `gorm:"column:ip" json:"ip"`
	Status      int       `gorm:"column:status" json:"status"`
	ExpiredTime null.Time `gorm:"column:expired_time" json:"expired_time"`
	CreatedTime time.Time `gorm:"column:created_time" json:"created_time"`
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updated_time"`
}

// TableName sets the insert table name for this struct type
func (a *AdminToken) TableName() string {
	return "admin_token"
}

type SmsCode struct {
	ID          int       `gorm:"column:id;primary_key" json:"id"`
	Type        int       `gorm:"column:type" json:"type"`
	IP          string    `gorm:"column:ip" json:"ip"`
	Mobile      string    `gorm:"column:mobile" json:"mobile"`
	Code        string    `gorm:"column:code" json:"code"`
	ExpireTime  int       `gorm:"column:expire_time" json:"expire_time"`
	Status      int       `gorm:"column:status" json:"status"`
	SendStatus  int       `gorm:"column:send_status" json:"send_status"`
	CreatedTime time.Time `gorm:"column:created_time" json:"created_time"`
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updated_time"`
}

// TableName sets the insert table name for this struct type
func (s *SmsCode) TableName() string {
	return "sms_code"
}






