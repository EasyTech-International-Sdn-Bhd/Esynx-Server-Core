package entities

import (
	"time"
)

type CmsStockTmpltBind struct {
	Id            uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	TmpltId       int       `xorm:"not null unique(unq_key) INT" json:"tmpltId,omitempty" xml:"tmpltId"`
	CustCode      string    `xorm:"not null unique(unq_key) VARCHAR(50)" json:"custCode,omitempty" xml:"custCode"`
	SalespersonId int       `xorm:"default 0 unique(unq_key) INT" json:"salespersonId,omitempty" xml:"salespersonId"`
	ActiveStatus  int       `xorm:"default 1 unique(unq_key) INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	UpdatedBy     string    `xorm:"default '' VARCHAR(50)" json:"updatedBy,omitempty" xml:"updatedBy"`
	UpdatedAt     time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsStockTmpltBind) TableName() string {
	return "cms_stock_tmplt_bind"
}
