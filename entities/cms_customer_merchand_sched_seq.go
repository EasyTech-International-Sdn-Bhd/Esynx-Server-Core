package entities

import (
	"time"
)

type CmsCustomerMerchandSchedSeq struct {
	Id            int       `xorm:"not null pk autoincr INT"`
	AppmntDtlId   int       `xorm:"not null unique INT"`
	SalespersonId int       `xorm:"not null INT"`
	Sequence      int       `xorm:"not null INT"`
	ActiveStatus  int       `xorm:"default 1 INT"`
	UpdatedAt     time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
}

func (m *CmsCustomerMerchandSchedSeq) TableName() string {
	return "cms_customer_merchand_sched_seq"
}
