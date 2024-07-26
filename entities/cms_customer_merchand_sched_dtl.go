package entities

import (
	"time"
)

type CmsCustomerMerchandSchedDtl struct {
	Id             uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	AppmntId       int       `xorm:"unique(unq) INT" json:"appmntId,omitempty" xml:"appmntId"`
	AppmntCode     string    `xorm:"not null unique(unq) VARCHAR(55)" json:"appmntCode,omitempty" xml:"appmntCode"`
	AppmntDate     time.Time `xorm:"unique(unq) DATE" json:"appmntDate,omitempty" xml:"appmntDate"`
	AppmntTime     time.Time `xorm:"TIME" json:"appmntTime,omitempty" xml:"appmntTime"`
	AppmntDuration time.Time `xorm:"TIME" json:"appmntDuration,omitempty" xml:"appmntDuration"`
	AppmntStatus   int       `xorm:"default 0 INT" json:"appmntStatus,omitempty" xml:"appmntStatus"`
	ActiveStatus   int       `xorm:"default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	UpdatedBy      string    `xorm:"not null VARCHAR(50)" json:"updatedBy,omitempty" xml:"updatedBy"`
	UpdatedAt      time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsCustomerMerchandSchedDtl) TableName() string {
	return "cms_customer_merchand_sched_dtl"
}
