package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetMySqlBb() *gorm.DB {
	_db, err := gorm.Open("mysql", "root:wyx123@(127.0.0.1:3306)/sf?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	return _db
}
