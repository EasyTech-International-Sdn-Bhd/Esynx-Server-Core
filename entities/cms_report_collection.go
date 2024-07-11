package entities

import (
	"time"
)

type CmsReportCollection struct {
	Id           uint64    `xorm:"not null pk autoincr unique UNSIGNED BIGINT"`
	UniqueId     string    `xorm:"not null unique VARCHAR(30)"`
	Data         string    `xorm:"JSON"`
	ActiveStatus int       `xorm:"not null default 1 INT"`
	UpdatedAt    time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
	RangeId      string    `xorm:"VARCHAR(80)"`
}

func (m *CmsReportCollection) TableName() string {
	return "cms_report_collection"
}
