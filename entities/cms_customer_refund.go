package entities

import (
	"time"
)

type CmsCustomerRefund struct {
	CfId             uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	CfCode           string    `xorm:"index unique VARCHAR(20)"`
	CustCode         string    `xorm:"index VARCHAR(20)"`
	CfDate           time.Time `xorm:"TIMESTAMP"`
	CfAmount         float64   `xorm:"DOUBLE"`
	Cancelled        string    `xorm:"CHAR(1)"`
	CfKnockoffAmount float64   `xorm:"DOUBLE"`
	SalespersonId    int       `xorm:"INT"`
	RefNo            int       `xorm:"INT"`
	UpdatedAt        time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (m *CmsCustomerRefund) TableName() string {
	return "cms_customer_refund"
}
