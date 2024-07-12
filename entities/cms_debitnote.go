package entities

import (
	"time"
)

type CmsDebitnote struct {
	DnId              uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	DnCode            string    `xorm:"index unique VARCHAR(20)"`
	CustCode          string    `xorm:"index VARCHAR(20)"`
	DnDate            time.Time `xorm:"TIMESTAMP"`
	DnAmount          float64   `xorm:"DOUBLE"`
	OutstandingAmount float64   `xorm:"DOUBLE"`
	SalespersonId     int       `xorm:"INT"`
	Cancelled         string    `xorm:"CHAR(1)"`
	Approved          int       `xorm:"default 0 INT"`
	Approver          string    `xorm:"VARCHAR(200)"`
	ApprovedAt        time.Time `xorm:"DATETIME"`
	UpdatedAt         time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (m *CmsDebitnote) TableName() string {
	return "cms_debitnote"
}

func (m *CmsDebitnote) Validate() {
}

func (m *CmsDebitnote) ToUpdate() {
	m.UpdatedAt = time.Now()
}
