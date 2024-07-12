package entities

import (
	"time"
)

type CmsStockTmpltDtl struct {
	Id           uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	TmpltId      int       `xorm:"not null unique(unq) unique(unq_key) INT"`
	DtlCode      string    `xorm:"unique(unq) unique(unq_key) VARCHAR(50)"`
	DtlName      string    `xorm:"VARCHAR(150)"`
	DtlType      string    `xorm:"comment('PACKAGE/ITEM/CATEGORY') unique(unq) unique(unq_key) VARCHAR(50)"`
	ActiveStatus int       `xorm:"default 1 unique(unq) unique(unq_key) INT"`
	UpdatedBy    string    `xorm:"VARCHAR(50)"`
	UpdatedAt    time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
}

func (m *CmsStockTmpltDtl) TableName() string {
	return "cms_stock_tmplt_dtl"
}
