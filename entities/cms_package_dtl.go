package entities

import (
	"time"
)

type CmsPackageDtl struct {
	DtlId            uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"dtlId,omitempty" xml:"dtlId"`
	DtlParent        string    `xorm:"not null default '' index(dtl_idx) unique(dtl_unique) VARCHAR(100)" json:"dtlParent,omitempty" xml:"dtlParent"`
	DtlCode          string    `xorm:"not null default '' index(dtl_idx) unique(dtl_unique) VARCHAR(100)" json:"dtlCode,omitempty" xml:"dtlCode"`
	DtlName          string    `xorm:"default '' VARCHAR(100)" json:"dtlName,omitempty" xml:"dtlName"`
	DtlRemark        string    `xorm:"default '' VARCHAR(200)" json:"dtlRemark,omitempty" xml:"dtlRemark"`
	DtlQty           float64   `xorm:"default 0 unique(dtl_unique) DOUBLE" json:"dtlQty,omitempty" xml:"dtlQty"`
	DtlUnitPrice     float64   `xorm:"default 0 DOUBLE" json:"dtlUnitPrice,omitempty" xml:"dtlUnitPrice"`
	DtlTotal         float64   `xorm:"default 0 DOUBLE" json:"dtlTotal,omitempty" xml:"dtlTotal"`
	DtlUom           string    `xorm:"default '' VARCHAR(20)" json:"dtlUom,omitempty" xml:"dtlUom"`
	DtlPurchaseTotal float64   `xorm:"default 0 DOUBLE" json:"dtlPurchaseTotal,omitempty" xml:"dtlPurchaseTotal"`
	DtlPurchasePrice float64   `xorm:"default 0 DOUBLE" json:"dtlPurchasePrice,omitempty" xml:"dtlPurchasePrice"`
	DtlStatus        int       `xorm:"default 1 INT" json:"dtlStatus,omitempty" xml:"dtlStatus"`
	UpdatedAt        time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsPackageDtl) TableName() string {
	return "cms_package_dtl"
}
