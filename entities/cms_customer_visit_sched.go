package entities

import (
	"time"
)

type CmsCustomerVisitSched struct {
	Id                  uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	SchedCode           string    `xorm:"VARCHAR(50)" json:"schedCode,omitempty" xml:"schedCode"`
	CustCode            string    `xorm:"not null unique(cms_customer_visit_sched_unq) VARCHAR(100)" json:"custCode,omitempty" xml:"custCode"`
	BranchCode          string    `xorm:"default '' VARCHAR(50)" json:"branchCode,omitempty" xml:"branchCode"`
	PeriodStart         time.Time `xorm:"DATE" json:"periodStart,omitempty" xml:"periodStart"`
	PeriodEnd           time.Time `xorm:"DATE" json:"periodEnd,omitempty" xml:"periodEnd"`
	SchedDatetime       time.Time `xorm:"not null unique(cms_customer_visit_sched_unq) DATETIME" json:"schedDatetime,omitempty" xml:"schedDatetime"`
	SchedNote           string    `xorm:"VARCHAR(200)" json:"schedNote,omitempty" xml:"schedNote"`
	SalespersonId       int       `xorm:"not null default 0 unique(cms_customer_visit_sched_unq) INT" json:"salespersonId,omitempty" xml:"salespersonId"`
	TechAssignee        int       `xorm:"default 0 INT" json:"techAssignee,omitempty" xml:"techAssignee"`
	TechAssigned        int       `xorm:"default 0 INT" json:"techAssigned,omitempty" xml:"techAssigned"`
	SiteInCharge        string    `xorm:"default '' VARCHAR(100)" json:"siteInCharge,omitempty" xml:"siteInCharge"`
	SiteInChargeContact string    `xorm:"default '' VARCHAR(20)" json:"siteInChargeContact,omitempty" xml:"siteInChargeContact"`
	SiteLocation        string    `xorm:"default '' VARCHAR(200)" json:"siteLocation,omitempty" xml:"siteLocation"`
	SiteLocationDesc    string    `xorm:"JSON" json:"siteLocationDesc,omitempty" xml:"siteLocationDesc"`
	Tags                string    `xorm:"JSON" json:"tags,omitempty" xml:"tags"`
	ProjectNo           string    `xorm:"JSON" json:"projectNo,omitempty" xml:"projectNo"`
	SchedStatus         int       `xorm:"default 0 unique(cms_customer_visit_sched_unq) INT" json:"schedStatus,omitempty" xml:"schedStatus"`
	ActiveStatus        int       `xorm:"default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	CreatedBy           string    `xorm:"default '0' VARCHAR(30)" json:"createdBy,omitempty" xml:"createdBy"`
	InProgress          int       `xorm:"default 0 comment('0 - default 1 - in progress 2 - done / completed') INT" json:"inProgress,omitempty" xml:"inProgress"`
	CheckedinBy         int       `xorm:"default 0 INT" json:"checkedinBy,omitempty" xml:"checkedinBy"`
	CheckedinAt         time.Time `xorm:"DATETIME" json:"checkedinAt,omitempty" xml:"checkedinAt"`
	SchedSyncStatus     string    `xorm:"default '-1' VARCHAR(10)" json:"schedSyncStatus,omitempty" xml:"schedSyncStatus"`
	SchedFault          string    `xorm:"default '0' VARCHAR(10)" json:"schedFault,omitempty" xml:"schedFault"`
	SchedFaultMessage   string    `xorm:"VARCHAR(500)" json:"schedFaultMessage,omitempty" xml:"schedFaultMessage"`
	UpdatedBy           int       `xorm:"default 0 INT" json:"updatedBy,omitempty" xml:"updatedBy"`
	UpdatedAt           time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
	CreatedAt           time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"createdAt,omitempty" xml:"createdAt"`
}

func (m *CmsCustomerVisitSched) TableName() string {
	return "cms_customer_visit_sched"
}
