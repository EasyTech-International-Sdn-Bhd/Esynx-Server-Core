package entities

import (
	"time"
)

type CmsPackage struct {
	PkgId            uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"pkgId,omitempty" xml:"pkgId"`
	PkgCode          string    `xorm:"not null default '' index unique VARCHAR(100)" json:"pkgCode,omitempty" xml:"pkgCode"`
	PkgName          string    `xorm:"default '' VARCHAR(100)" json:"pkgName,omitempty" xml:"pkgName"`
	PkgDesc          string    `xorm:"default '' VARCHAR(200)" json:"pkgDesc,omitempty" xml:"pkgDesc"`
	PkgExpiryDate    time.Time `xorm:"DATETIME" json:"pkgExpiryDate,omitempty" xml:"pkgExpiryDate"`
	PkgOpeningQty    float64   `xorm:"default 0 DOUBLE" json:"pkgOpeningQty,omitempty" xml:"pkgOpeningQty"`
	PkgUnitPrice     float64   `xorm:"default 0 DOUBLE" json:"pkgUnitPrice,omitempty" xml:"pkgUnitPrice"`
	PkgLimitedQty    float64   `xorm:"default 0 DOUBLE" json:"pkgLimitedQty,omitempty" xml:"pkgLimitedQty"`
	PkgSoldoutQty    float64   `xorm:"default 0 DOUBLE" json:"pkgSoldoutQty,omitempty" xml:"pkgSoldoutQty"`
	PkgTotal         float64   `xorm:"default 0 DOUBLE" json:"pkgTotal,omitempty" xml:"pkgTotal"`
	PkgRemark        string    `xorm:"default '' VARCHAR(200)" json:"pkgRemark,omitempty" xml:"pkgRemark"`
	PkgUom           string    `xorm:"default '' VARCHAR(20)" json:"pkgUom,omitempty" xml:"pkgUom"`
	PkgPurchaseTotal float64   `xorm:"default 0 DOUBLE" json:"pkgPurchaseTotal,omitempty" xml:"pkgPurchaseTotal"`
	PkgPurchaseQty   float64   `xorm:"default 0 DOUBLE" json:"pkgPurchaseQty,omitempty" xml:"pkgPurchaseQty"`
	PkgQrcode        string    `xorm:"default '' VARCHAR(200)" json:"pkgQrcode,omitempty" xml:"pkgQrcode"`
	PkgStatus        int       `xorm:"default 1 INT" json:"pkgStatus,omitempty" xml:"pkgStatus"`
	UpdatedAt        time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsPackage) TableName() string {
	return "cms_package"
}
