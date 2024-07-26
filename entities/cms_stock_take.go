package entities

import (
	"time"
)

type CmsStockTake struct {
	Id              uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	StId            int       `xorm:"not null unique(unq) INT" json:"stId,omitempty" xml:"stId"`
	CustCode        string    `xorm:"not null VARCHAR(200)" json:"custCode,omitempty" xml:"custCode"`
	CustCompanyName string    `xorm:"not null VARCHAR(200)" json:"custCompanyName,omitempty" xml:"custCompanyName"`
	SalespersonId   int       `xorm:"not null unique(unq) INT" json:"salespersonId,omitempty" xml:"salespersonId"`
	CreatedDate     time.Time `xorm:"not null DATETIME" json:"createdDate,omitempty" xml:"createdDate"`
	DocRefId        string    `xorm:"not null VARCHAR(50)" json:"docRefId,omitempty" xml:"docRefId"`
	SpUpdatedAt     string    `xorm:"not null VARCHAR(50)" json:"spUpdatedAt,omitempty" xml:"spUpdatedAt"`
	ActiveStatus    int       `xorm:"default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	UpdatedAt       time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsStockTake) TableName() string {
	return "cms_stock_take"
}
