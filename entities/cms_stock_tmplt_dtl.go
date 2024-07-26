package entities

import (
	"time"
)

type CmsStockTmpltDtl struct {
	Id           uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	TmpltId      int       `xorm:"not null unique(unq) unique(unq_key) INT" json:"tmpltId,omitempty" xml:"tmpltId"`
	DtlCode      string    `xorm:"unique(unq) unique(unq_key) VARCHAR(50)" json:"dtlCode,omitempty" xml:"dtlCode"`
	DtlName      string    `xorm:"VARCHAR(150)" json:"dtlName,omitempty" xml:"dtlName"`
	DtlType      string    `xorm:"comment('PACKAGE/ITEM/CATEGORY') unique(unq) unique(unq_key) VARCHAR(50)" json:"dtlType,omitempty" xml:"dtlType"`
	ActiveStatus int       `xorm:"default 1 unique(unq) unique(unq_key) INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	UpdatedBy    string    `xorm:"VARCHAR(50)" json:"updatedBy,omitempty" xml:"updatedBy"`
	UpdatedAt    time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsStockTmpltDtl) TableName() string {
	return "cms_stock_tmplt_dtl"
}
