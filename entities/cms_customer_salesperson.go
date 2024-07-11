package entities

import (
	"time"
)

type CmsCustomerSalesperson struct {
	SalespersonCustomerId int       `xorm:"not null pk autoincr INT"`
	SalespersonId         int       `xorm:"unique(unique_customer) INT"`
	Sequence              int       `xorm:"default 0 INT"`
	CustomerId            int       `xorm:"unique(unique_customer) INT"`
	ActiveStatus          int       `xorm:"default 1 INT"`
	UpdatedAt             time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	SessionId             string    `xorm:"default '' VARCHAR(100)"`
}

func (m *CmsCustomerSalesperson) TableName() string {
	return "cms_customer_salesperson"
}
