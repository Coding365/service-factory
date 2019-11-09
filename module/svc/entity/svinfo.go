package entity

import (
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
