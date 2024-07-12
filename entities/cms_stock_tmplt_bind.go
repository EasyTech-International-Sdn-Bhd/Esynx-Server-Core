package entities

import (
	"time"
)

type CmsStockTmpltBind struct {
	Id            uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	TmpltId       int       `xorm:"not null unique(unq_key) INT"`
	CustCode      string    `xorm:"not null unique(unq_key) VARCHAR(50)"`
	SalespersonId int       `xorm:"default 0 unique(unq_key) INT"`
	ActiveStatus  int       `xorm:"default 1 unique(unq_key) INT"`
	UpdatedBy     string    `xorm:"default '' VARCHAR(50)"`
	UpdatedAt     time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
}

func (m *CmsStockTmpltBind) TableName() string {
	return "cms_stock_tmplt_bind"
}
