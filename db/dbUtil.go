package init

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//func init() {
//	maxIdle := 30
//	maxConn := 30
//	orm.RegisterDriver("mysql", orm.DRMySQL)
//	orm.RegisterDataBase("default", "mysql", "root:wyx123@tcp(localhost:3306)/sf?charset=utf8",maxIdle,maxConn)
//	orm.SetMaxIdleConns("default", 30)
//	orm.SetMaxOpenConns("default", 30)
//	orm.RegisterModel(new(models.SvInfo))
//}

func init() {
	db, _ := gorm.Open("mysql", "root:wyx123@(127.0.0.1:3306)/sf?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
}