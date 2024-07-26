package entities

import (
	"time"
)

type CmsCustomerProducts struct {
	Id           uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	CustCode     string    `xorm:"not null unique(unq) VARCHAR(50)" json:"custCode,omitempty" xml:"custCode"`
	ProductCode  string    `xorm:"not null unique(unq) VARCHAR(50)" json:"productCode,omitempty" xml:"productCode"`
	VipQty       float64   `xorm:"not null default 0 DOUBLE" json:"vipQty,omitempty" xml:"vipQty"`
	ActiveStatus int       `xorm:"not null default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	UpdatedAt    time.Time `xorm:"not null default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsCustomerProducts) TableName() string {
	return "cms_customer_products"
}
