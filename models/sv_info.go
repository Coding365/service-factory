package models

import (
	"fmt"
	"time"
)

package models

type SvInfo struct {
	Code        string `gorm:"column:code"`
	ContentPath string `gorm:"column:content_path"`
	CreateTime  time.Time   `gorm:"column:createTime"`
	Creator     string `gorm:"column:creator"`
	ID          string `gorm:"column:id;primary_key"`
	Image       string `gorm:"column:image"`
	JSONStr     string `gorm:"column:jsonStr"`
	Name        string `gorm:"column:name"`
	Namespace   string `gorm:"column:namespace"`
	UpdateTime  time.Time   `gorm:"column:updateTime"`
}

// TableName sets the insert table name for this struct type
func (s *SvInfo) TableName() string {
	return "sv_info"
}


func Insert(info *SvInfo)  {
	o := orm.NewOrm()
	id, err := o.Insert(info)
	if err == nil {
		fmt.Println(id)
	}
}