package entities

import (
	"time"
)

type CmsCustomerZone struct {
	ZoneId       uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	ZoneName     string    `xorm:"not null default '' unique VARCHAR(100)"`
	ZoneRemark   string    `xorm:"not null default '' VARCHAR(1000)"`
	ActiveStatus int       `xorm:"not null default 1 INT"`
	UpdatedAt    time.Time `xorm:"not null default CURRENT_TIMESTAMP DATETIME"`
}

func (m *CmsCustomerZone) TableName() string {
	return "cms_customer_zone"
}
