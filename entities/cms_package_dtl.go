package entities

import (
	"time"
)

type CmsPackageDtl struct {
	DtlId            uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	DtlParent        string    `xorm:"not null default '' index(dtl_idx) unique(dtl_unique) VARCHAR(100)"`
	DtlCode          string    `xorm:"not null default '' index(dtl_idx) unique(dtl_unique) VARCHAR(100)"`
	DtlName          string    `xorm:"default '' VARCHAR(100)"`
	DtlRemark        string    `xorm:"default '' VARCHAR(200)"`
	DtlQty           float64   `xorm:"default 0 unique(dtl_unique) DOUBLE"`
	DtlUnitPrice     float64   `xorm:"default 0 DOUBLE"`
	DtlTotal         float64   `xorm:"default 0 DOUBLE"`
	DtlUom           string    `xorm:"default '' VARCHAR(20)"`
	DtlPurchaseTotal float64   `xorm:"default 0 DOUBLE"`
	DtlPurchasePrice float64   `xorm:"default 0 DOUBLE"`
	DtlStatus        int       `xorm:"default 1 INT"`
	UpdatedAt        time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
}

func (m *CmsPackageDtl) TableName() string {
	return "cms_package_dtl"
}
