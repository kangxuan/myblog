package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"myblog/settings"
)

var db *sqlx.DB

func SetUp() {
	var err error
	mysqlConf := settings.ServerConf.MysqlConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", mysqlConf.User, mysqlConf.Pass, mysqlConf.Host, mysqlConf.Port, mysqlConf.Db)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}
	db.SetMaxOpenConns(mysqlConf.MaxOpenConn)
	db.SetMaxIdleConns(mysqlConf.MaxIdleConn)
}
