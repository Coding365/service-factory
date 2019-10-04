package db

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetMySqlBb() *gorm.DB {
	dbType := beego.AppConfig.String("dbType")
	dbConn := beego.AppConfig.String("dbConn")
	_db, err := gorm.Open(dbType, dbConn)
	if err != nil {
		panic("连接数据库失败")
	}
	return _db
}
