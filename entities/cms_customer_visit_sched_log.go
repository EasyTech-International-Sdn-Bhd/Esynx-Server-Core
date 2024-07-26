package entities

import (
	"time"
)

type CmsCustomerVisitSchedLog struct {
	Id          uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	SchedId     int       `xorm:"not null INT" json:"schedId,omitempty" xml:"schedId"`
	SchedSeenBy int       `xorm:"not null INT" json:"schedSeenBy,omitempty" xml:"schedSeenBy"`
	SchedSeenAt time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"schedSeenAt,omitempty" xml:"schedSeenAt"`
	SchedNote   []byte    `xorm:"BLOB" json:"schedNote,omitempty" xml:"schedNote"`
}

func (m *CmsCustomerVisitSchedLog) TableName() string {
	return "cms_customer_visit_sched_log"
}
