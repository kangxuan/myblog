package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 驱动一定要导入进去
	"github.com/jmoiron/sqlx"
	"myblog/settings"
)

var Db *sqlx.DB

func SetUp() {
	var err error
	mysqlConf := settings.ServerConf.MysqlConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", mysqlConf.User, mysqlConf.Pass, mysqlConf.Host, mysqlConf.Port, mysqlConf.Db)
	Db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}
	Db.SetMaxOpenConns(mysqlConf.MaxOpenConn)
	Db.SetMaxIdleConns(mysqlConf.MaxIdleConn)
}
