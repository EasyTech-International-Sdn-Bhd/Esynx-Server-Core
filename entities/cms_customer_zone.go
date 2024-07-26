package entities

import (
	"time"
)

type CmsCustomerZone struct {
	ZoneId       uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"zoneId,omitempty" xml:"zoneId"`
	ZoneName     string    `xorm:"not null default '' unique VARCHAR(100)" json:"zoneName,omitempty" xml:"zoneName"`
	ZoneRemark   string    `xorm:"not null default '' VARCHAR(1000)" json:"zoneRemark,omitempty" xml:"zoneRemark"`
	ActiveStatus int       `xorm:"not null default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	UpdatedAt    time.Time `xorm:"not null default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsCustomerZone) TableName() string {
	return "cms_customer_zone"
}
