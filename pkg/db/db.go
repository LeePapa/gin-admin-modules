package pkg_db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"log"
	"gin-modules/pkg/setting"
)

var DbManager *gorm.DB

func Initialize(mysqlConfig pkg_setting.DatabaseIni) {
	var err error
	DbManager, err = gorm.Open(mysqlConfig.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlConfig.User,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Db))
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return mysqlConfig.TablePrefix + defaultTableName
	}
	DbManager.SingularTable(true)
	DbManager.DB().SetMaxIdleConns(mysqlConfig.MaxIdle)
	DbManager.DB().SetMaxOpenConns(mysqlConfig.MaxOpen)
}

func CloseDB() {
	defer DbManager.Close()
}
