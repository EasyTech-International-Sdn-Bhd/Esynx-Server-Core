package entities

import (
	"time"
)

type CmsReportCollection struct {
	Id           uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	UniqueId     string    `xorm:"not null unique VARCHAR(30)" json:"uniqueId,omitempty" xml:"uniqueId"`
	Data         string    `xorm:"JSON" json:"data,omitempty" xml:"data"`
	ActiveStatus int       `xorm:"not null default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	UpdatedAt    time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
	RangeId      string    `xorm:"VARCHAR(80)" json:"rangeId,omitempty" xml:"rangeId"`
}

func (m *CmsReportCollection) TableName() string {
	return "cms_report_collection"
}
