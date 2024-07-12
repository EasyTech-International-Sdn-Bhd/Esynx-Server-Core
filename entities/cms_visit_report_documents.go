package entities

import (
	"time"
)

type CmsVisitReportDocuments struct {
	Id           uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	CheckInId    string    `xorm:"not null unique(unq) VARCHAR(25)"`
	DocId        string    `xorm:"not null unique(unq) VARCHAR(25)"`
	ActiveStatus int       `xorm:"default 1 INT"`
	UpdatedAt    time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
}

func (m *CmsVisitReportDocuments) TableName() string {
	return "cms_visit_report_documents"
}
