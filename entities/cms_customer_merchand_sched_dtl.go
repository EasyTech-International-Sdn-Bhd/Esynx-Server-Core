package entities

import (
	"time"
)

type CmsCustomerMerchandSchedDtl struct {
	Id             uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	AppmntId       int       `xorm:"unique(unq) INT"`
	AppmntCode     string    `xorm:"not null unique(unq) VARCHAR(55)"`
	AppmntDate     time.Time `xorm:"unique(unq) DATE"`
	AppmntTime     time.Time `xorm:"TIME"`
	AppmntDuration time.Time `xorm:"TIME"`
	AppmntStatus   int       `xorm:"default 0 INT"`
	ActiveStatus   int       `xorm:"default 1 INT"`
	UpdatedBy      string    `xorm:"not null VARCHAR(50)"`
	UpdatedAt      time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
}

func (m *CmsCustomerMerchandSchedDtl) TableName() string {
	return "cms_customer_merchand_sched_dtl"
}
