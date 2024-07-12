package entities

import (
	"time"
)

type CmsAppAnnouncement struct {
	Id           uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	AlertName    string    `xorm:"not null unique VARCHAR(200)"`
	AlertContent []byte    `xorm:"not null BLOB"`
	AlertImage   string    `xorm:"not null VARCHAR(255)"`
	AlertAction  []byte    `xorm:"not null BLOB"`
	DateFrom     time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
	DateTo       time.Time `xorm:"not null default '1971-01-01 23:01:01' DATETIME"`
	ActiveStatus int       `xorm:"not null default 1 INT"`
	UpdatedAt    time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
}

func (m *CmsAppAnnouncement) TableName() string {
	return "cms_app_announcement"
}
