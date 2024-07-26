package entities

import (
	"time"
)

type CmsDebitnoteSales struct {
	DnId              uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"dnId,omitempty" xml:"dnId"`
	DnCode            string    `xorm:"index unique VARCHAR(20)" json:"dnCode,omitempty" xml:"dnCode"`
	CustCode          string    `xorm:"index VARCHAR(20)" json:"custCode,omitempty" xml:"custCode"`
	DnDate            time.Time `xorm:"TIMESTAMP" json:"dnDate,omitempty" xml:"dnDate"`
	DnAmount          float64   `xorm:"DOUBLE" json:"dnAmount,omitempty" xml:"dnAmount"`
	OutstandingAmount float64   `xorm:"DOUBLE" json:"outstandingAmount,omitempty" xml:"outstandingAmount"`
	Cancelled         string    `xorm:"CHAR(1)" json:"cancelled,omitempty" xml:"cancelled"`
	Approved          int       `xorm:"default 0 INT" json:"approved,omitempty" xml:"approved"`
	Approver          string    `xorm:"VARCHAR(200)" json:"approver,omitempty" xml:"approver"`
	ApprovedAt        time.Time `xorm:"DATETIME" json:"approvedAt,omitempty" xml:"approvedAt"`
	UpdatedAt         time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsDebitnoteSales) TableName() string {
	return "cms_debitnote_sales"
}

func (m *CmsDebitnoteSales) BeforeInsert() {
	m.BeforeUpdate()
}

func (m *CmsDebitnoteSales) BeforeUpdate() {
	m.UpdatedAt = time.Now()
}
