package entities

import (
	"time"
)

type CmsCustomerMerchandSched struct {
	Id             uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	AppmntCode     string    `xorm:"not null unique(unq) VARCHAR(55)" json:"appmntCode,omitempty" xml:"appmntCode"`
	AppmntNote     []byte    `xorm:"BLOB" json:"appmntNote,omitempty" xml:"appmntNote"`
	AppmntDocType  string    `xorm:"not null JSON" json:"appmntDocType,omitempty" xml:"appmntDocType"`
	AppmntUdf      string    `xorm:"not null JSON" json:"appmntUdf,omitempty" xml:"appmntUdf"`
	CustCode       string    `xorm:"not null unique(unq) VARCHAR(50)" json:"custCode,omitempty" xml:"custCode"`
	BranchCode     string    `xorm:"default '' unique(unq) VARCHAR(50)" json:"branchCode,omitempty" xml:"branchCode"`
	SalespersonId  int       `xorm:"not null INT" json:"salespersonId,omitempty" xml:"salespersonId"`
	StartDate      time.Time `xorm:"not null DATETIME" json:"startDate,omitempty" xml:"startDate"`
	EndDate        time.Time `xorm:"not null DATETIME" json:"endDate,omitempty" xml:"endDate"`
	Repetitive     int       `xorm:"default 0 INT" json:"repetitive,omitempty" xml:"repetitive"`
	AppmntInterval string    `xorm:"default '' VARCHAR(20)" json:"appmntInterval,omitempty" xml:"appmntInterval"`
	AppmntStatus   int       `xorm:"default 0 INT" json:"appmntStatus,omitempty" xml:"appmntStatus"`
	ActiveStatus   int       `xorm:"default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	CreatedBy      string    `xorm:"not null VARCHAR(50)" json:"createdBy,omitempty" xml:"createdBy"`
	CreatedAt      time.Time `xorm:"not null DATETIME" json:"createdAt,omitempty" xml:"createdAt"`
	UpdatedBy      string    `xorm:"not null VARCHAR(50)" json:"updatedBy,omitempty" xml:"updatedBy"`
	UpdatedAt      time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsCustomerMerchandSched) TableName() string {
	return "cms_customer_merchand_sched"
}
