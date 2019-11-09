package dao

import (
	uuid "github.com/satori/go.uuid"
	"service-factory/db"
	"service-factory/module/svc/entity"
)

func SvInfoInsert(info *entity.SvInfo) string {
	mysqldb := db.GetMySqlBb()
	uid, _ := uuid.NewV4()
	info.ID = uid.String()
	mysqldb.Create(info)
	mysqldb.Close()
	return info.ID
}

func SvInfoSelect(id string) *entity.SvInfo {
	mysqldb := db.GetMySqlBb()
	svInfo := new(entity.SvInfo)
	mysqldb.Where("id = ?", id).First(svInfo)
	return svInfo
}

func SvInfoSelectAll(pageSize int, pageNo int) []entity.SvInfo {
	if pageSize <= 0 {
		pageSize = 20
	}
	if pageNo <= 0 {
		pageNo = 1
	}
	mysqldb := db.GetMySqlBb()
	var svInfos []entity.SvInfo
	mysqldb.Find(&svInfos).Limit(pageSize).Offset((pageNo-1)*pageSize + 1).Order("created_at desc")
	return svInfos
}
