package entities

import (
	"time"
)

type CmsCustomerContraLocal struct {
	CtId              uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"ctId,omitempty" xml:"ctId"`
	CtCode            string    `xorm:"index unique VARCHAR(20)" json:"ctCode,omitempty" xml:"ctCode"`
	CustCode          string    `xorm:"index VARCHAR(20)" json:"custCode,omitempty" xml:"custCode"`
	CtDate            time.Time `xorm:"TIMESTAMP" json:"ctDate,omitempty" xml:"ctDate"`
	CtAmount          float64   `xorm:"DOUBLE" json:"ctAmount,omitempty" xml:"ctAmount"`
	Cancelled         string    `xorm:"CHAR(1)" json:"cancelled,omitempty" xml:"cancelled"`
	CtUnappliedAmount float64   `xorm:"DOUBLE" json:"ctUnappliedAmount,omitempty" xml:"ctUnappliedAmount"`
	SalespersonId     int       `xorm:"INT" json:"salespersonId,omitempty" xml:"salespersonId"`
	UpdatedAt         time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
	RefNo             string    `xorm:"VARCHAR(20)" json:"refNo,omitempty" xml:"refNo"`
}

func (m *CmsCustomerContraLocal) TableName() string {
	return "cms_customer_contra_local"
}
