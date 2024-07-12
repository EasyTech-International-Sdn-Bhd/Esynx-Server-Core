package entities

import (
	"time"
)

type CmsCustomerMerchandSched struct {
	Id             uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	AppmntCode     string    `xorm:"not null unique(unq) VARCHAR(55)"`
	AppmntNote     []byte    `xorm:"BLOB"`
	AppmntDocType  string    `xorm:"not null JSON"`
	AppmntUdf      string    `xorm:"not null JSON"`
	CustCode       string    `xorm:"not null unique(unq) VARCHAR(50)"`
	BranchCode     string    `xorm:"default '' unique(unq) VARCHAR(50)"`
	SalespersonId  int       `xorm:"not null INT"`
	StartDate      time.Time `xorm:"not null DATETIME"`
	EndDate        time.Time `xorm:"not null DATETIME"`
	Repetitive     int       `xorm:"default 0 INT"`
	AppmntInterval string    `xorm:"default '' VARCHAR(20)"`
	AppmntStatus   int       `xorm:"default 0 INT"`
	ActiveStatus   int       `xorm:"default 1 INT"`
	CreatedBy      string    `xorm:"not null VARCHAR(50)"`
	CreatedAt      time.Time `xorm:"not null DATETIME"`
	UpdatedBy      string    `xorm:"not null VARCHAR(50)"`
	UpdatedAt      time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
}

func (m *CmsCustomerMerchandSched) TableName() string {
	return "cms_customer_merchand_sched"
}
