package entities

import (
	"time"
)

type CmsCustomerVisitSched struct {
	Id                  uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	SchedCode           string    `xorm:"VARCHAR(50)"`
	CustCode            string    `xorm:"not null unique(cms_customer_visit_sched_unq) VARCHAR(100)"`
	BranchCode          string    `xorm:"default '' VARCHAR(50)"`
	PeriodStart         time.Time `xorm:"DATE"`
	PeriodEnd           time.Time `xorm:"DATE"`
	SchedDatetime       time.Time `xorm:"not null unique(cms_customer_visit_sched_unq) DATETIME"`
	SchedNote           string    `xorm:"VARCHAR(200)"`
	SalespersonId       int       `xorm:"not null default 0 unique(cms_customer_visit_sched_unq) INT"`
	TechAssignee        int       `xorm:"default 0 INT"`
	TechAssigned        int       `xorm:"default 0 INT"`
	SiteInCharge        string    `xorm:"default '' VARCHAR(100)"`
	SiteInChargeContact string    `xorm:"default '' VARCHAR(20)"`
	SiteLocation        string    `xorm:"default '' VARCHAR(200)"`
	SiteLocationDesc    string    `xorm:"JSON"`
	Tags                string    `xorm:"JSON"`
	ProjectNo           string    `xorm:"JSON"`
	SchedStatus         int       `xorm:"default 0 unique(cms_customer_visit_sched_unq) INT"`
	ActiveStatus        int       `xorm:"default 1 INT"`
	CreatedBy           string    `xorm:"default '0' VARCHAR(30)"`
	InProgress          int       `xorm:"default 0 comment('0 - default 1 - in progress 2 - done / completed') INT"`
	CheckedinBy         int       `xorm:"default 0 INT"`
	CheckedinAt         time.Time `xorm:"DATETIME"`
	SchedSyncStatus     string    `xorm:"default '-1' VARCHAR(10)"`
	SchedFault          string    `xorm:"default '0' VARCHAR(10)"`
	SchedFaultMessage   string    `xorm:"VARCHAR(500)"`
	UpdatedBy           int       `xorm:"default 0 INT"`
	UpdatedAt           time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
	CreatedAt           time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
}

func (m *CmsCustomerVisitSched) TableName() string {
	return "cms_customer_visit_sched"
}
