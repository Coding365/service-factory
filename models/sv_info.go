package models

import (
	"github.com/satori/go.uuid"
	"service-factory/db"
	"time"
)

type SvInfo struct {
	Code        string    `gorm:"column:code"`
	ContentPath string    `gorm:"column:content_path"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	Creator     string    `gorm:"column:creator"`
	ID          string    `gorm:"column:id;primary_key"`
	Image       string    `gorm:"column:image"`
	JSONStr     string    `gorm:"column:jsonStr"`
	Name        string    `gorm:"column:name"`
	Namespace   string    `gorm:"column:namespace"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

// TableName sets the insert table name for this struct type
func (s *SvInfo) TableName() string {
	return "sv_info"
}

func Insert(info *SvInfo) string {
	mysqldb := db.GetMySqlBb()
	uid, _ := uuid.NewV4()
	info.ID = uid.String()
	mysqldb.Create(info)
	mysqldb.Close()
	return info.ID
}

func Select(id string) *SvInfo {
	mysqldb := db.GetMySqlBb()
	svInfo := new(SvInfo)
	mysqldb.Where("id = ?", id).First(svInfo)
	return svInfo
}
