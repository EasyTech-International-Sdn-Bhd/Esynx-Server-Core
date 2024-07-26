package entities

import (
	"time"
)

type CmsAppAnnouncement struct {
	Id           uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	AlertName    string    `xorm:"not null unique VARCHAR(200)" json:"alertName,omitempty" xml:"alertName"`
	AlertContent []byte    `xorm:"not null BLOB" json:"alertContent,omitempty" xml:"alertContent"`
	AlertImage   string    `xorm:"not null VARCHAR(255)" json:"alertImage,omitempty" xml:"alertImage"`
	AlertAction  []byte    `xorm:"not null BLOB" json:"alertAction,omitempty" xml:"alertAction"`
	DateFrom     time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"dateFrom,omitempty" xml:"dateFrom"`
	DateTo       time.Time `xorm:"not null default '1971-01-01 23:01:01' DATETIME" json:"dateTo,omitempty" xml:"dateTo"`
	ActiveStatus int       `xorm:"not null default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	UpdatedAt    time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsAppAnnouncement) TableName() string {
	return "cms_app_announcement"
}
