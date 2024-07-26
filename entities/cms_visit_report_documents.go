package entities

import (
	"time"
)

type CmsVisitReportDocuments struct {
	Id           uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	CheckInId    string    `xorm:"not null unique(unq) VARCHAR(25)" json:"checkInId,omitempty" xml:"checkInId"`
	DocId        string    `xorm:"not null unique(unq) VARCHAR(25)" json:"docId,omitempty" xml:"docId"`
	ActiveStatus int       `xorm:"default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	UpdatedAt    time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsVisitReportDocuments) TableName() string {
	return "cms_visit_report_documents"
}
