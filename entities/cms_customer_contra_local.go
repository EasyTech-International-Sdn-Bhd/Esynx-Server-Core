package entities

import (
	"time"
)

type CmsCustomerContraLocal struct {
	CtId              int       `xorm:"not null pk autoincr INT"`
	CtCode            string    `xorm:"index unique VARCHAR(20)"`
	CustCode          string    `xorm:"index VARCHAR(20)"`
	CtDate            time.Time `xorm:"TIMESTAMP"`
	CtAmount          float64   `xorm:"DOUBLE"`
	Cancelled         string    `xorm:"CHAR(1)"`
	CtUnappliedAmount float64   `xorm:"DOUBLE"`
	SalespersonId     int       `xorm:"INT"`
	UpdatedAt         time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	RefNo             string    `xorm:"VARCHAR(20)"`
}

func (m *CmsCustomerContraLocal) TableName() string {
	return "cms_customer_contra_local"
}
