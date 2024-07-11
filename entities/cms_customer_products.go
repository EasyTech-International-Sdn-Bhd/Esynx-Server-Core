package entities

import (
	"time"
)

type CmsCustomerProducts struct {
	Id           int       `xorm:"not null pk autoincr INT"`
	CustCode     string    `xorm:"not null unique(unq) VARCHAR(50)"`
	ProductCode  string    `xorm:"not null unique(unq) VARCHAR(50)"`
	VipQty       float64   `xorm:"not null default 0 DOUBLE"`
	ActiveStatus int       `xorm:"not null default 1 INT"`
	UpdatedAt    time.Time `xorm:"not null default CURRENT_TIMESTAMP DATETIME"`
}

func (m *CmsCustomerProducts) TableName() string {
	return "cms_customer_products"
}
