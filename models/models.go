package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 驱动一定要导入进去
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"myblog/settings"
)

var Db *gorm.DB
var err error

func SetUp() {
	mysqlConf := settings.ServerConf.MysqlConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", mysqlConf.User, mysqlConf.Pass, mysqlConf.Host, mysqlConf.Port, mysqlConf.Db)
	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置

	}), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("连接数据库失败，错误信息：%s\n", err))
	}
}
