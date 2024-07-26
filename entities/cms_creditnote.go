package entities

import (
	"time"
)

type CmsCreditnote struct {
	CnId             uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"cnId,omitempty" xml:"cnId"`
	CnCode           string    `xorm:"index unique VARCHAR(20)" json:"cnCode,omitempty" xml:"cnCode"`
	CustCode         string    `xorm:"index VARCHAR(20)" json:"custCode,omitempty" xml:"custCode"`
	CnDate           time.Time `xorm:"TIMESTAMP" json:"cnDate,omitempty" xml:"cnDate"`
	CnUdf            string    `xorm:"not null JSON" json:"cnUdf,omitempty" xml:"cnUdf"`
	CnAmount         float64   `xorm:"DOUBLE" json:"cnAmount,omitempty" xml:"cnAmount"`
	SalespersonId    int       `xorm:"INT" json:"salespersonId,omitempty" xml:"salespersonId"`
	Cancelled        string    `xorm:"CHAR(1)" json:"cancelled,omitempty" xml:"cancelled"`
	UpdatedAt        time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
	CnKnockoffAmount float64   `xorm:"DOUBLE" json:"cnKnockoffAmount,omitempty" xml:"cnKnockoffAmount"`
	Approved         int       `xorm:"default 0 INT" json:"approved,omitempty" xml:"approved"`
	Approver         string    `xorm:"VARCHAR(200)" json:"approver,omitempty" xml:"approver"`
	ApprovedAt       time.Time `xorm:"DATETIME" json:"approvedAt,omitempty" xml:"approvedAt"`
	FromDoc          string    `xorm:"default 'SL' ENUM('AR','SL')" json:"fromDoc,omitempty" xml:"fromDoc"`
	RefNo            string    `xorm:"VARCHAR(20)" json:"refNo,omitempty" xml:"refNo"`
}

func (m *CmsCreditnote) TableName() string {
	return "cms_creditnote"
}

func (m *CmsCreditnote) BeforeInsert() {
	m.BeforeUpdate()
}

func (m *CmsCreditnote) BeforeUpdate() {
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = time.Now()
	}
}
