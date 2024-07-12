package entities

import (
	"time"
)

type CmsPackage struct {
	PkgId            uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	PkgCode          string    `xorm:"not null default '' index unique VARCHAR(100)"`
	PkgName          string    `xorm:"default '' VARCHAR(100)"`
	PkgDesc          string    `xorm:"default '' VARCHAR(200)"`
	PkgExpiryDate    time.Time `xorm:"DATETIME"`
	PkgOpeningQty    float64   `xorm:"default 0 DOUBLE"`
	PkgUnitPrice     float64   `xorm:"default 0 DOUBLE"`
	PkgLimitedQty    float64   `xorm:"default 0 DOUBLE"`
	PkgSoldoutQty    float64   `xorm:"default 0 DOUBLE"`
	PkgTotal         float64   `xorm:"default 0 DOUBLE"`
	PkgRemark        string    `xorm:"default '' VARCHAR(200)"`
	PkgUom           string    `xorm:"default '' VARCHAR(20)"`
	PkgPurchaseTotal float64   `xorm:"default 0 DOUBLE"`
	PkgPurchaseQty   float64   `xorm:"default 0 DOUBLE"`
	PkgQrcode        string    `xorm:"default '' VARCHAR(200)"`
	PkgStatus        int       `xorm:"default 1 INT"`
	UpdatedAt        time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
}

func (m *CmsPackage) TableName() string {
	return "cms_package"
}
