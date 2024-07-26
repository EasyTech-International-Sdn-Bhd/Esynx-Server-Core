package entities

import (
	"time"
)

type CmsCustomerSalesperson struct {
	SalespersonCustomerId uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"salespersonCustomerId,omitempty" xml:"salespersonCustomerId"`
	SalespersonId         int       `xorm:"unique(unique_customer) INT" json:"salespersonId,omitempty" xml:"salespersonId"`
	Sequence              int       `xorm:"default 0 INT" json:"sequence,omitempty" xml:"sequence"`
	CustomerId            int       `xorm:"unique(unique_customer) INT" json:"customerId,omitempty" xml:"customerId"`
	ActiveStatus          int       `xorm:"default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	UpdatedAt             time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
	SessionId             string    `xorm:"default '' VARCHAR(100)" json:"sessionId,omitempty" xml:"sessionId"`
}

func (m *CmsCustomerSalesperson) TableName() string {
	return "cms_customer_salesperson"
}

func (m *CmsCustomerSalesperson) Validate() {

}

func (m *CmsCustomerSalesperson) BeforeUpdate() {
	m.UpdatedAt = time.Now()
}
