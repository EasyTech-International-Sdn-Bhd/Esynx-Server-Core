package entities

import (
	"time"
)

type CmsCustomerVisitSchedLog struct {
	Id          uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	SchedId     int       `xorm:"not null INT"`
	SchedSeenBy int       `xorm:"not null INT"`
	SchedSeenAt time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
	SchedNote   []byte    `xorm:"BLOB"`
}

func (m *CmsCustomerVisitSchedLog) TableName() string {
	return "cms_customer_visit_sched_log"
}
