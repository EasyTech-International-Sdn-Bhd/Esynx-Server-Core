package entities

import (
	"time"
)

type CmsStockAdjustmentDtl struct {
	Id            uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	AdjId         string    `xorm:"not null unique(unq) VARCHAR(25)"`
	DeviceItemId  int64     `xorm:"default 0 unique(unq) BIGINT"`
	ProductCode   string    `xorm:"VARCHAR(25)"`
	WarehouseCode string    `xorm:"VARCHAR(25)"`
	BatchNo       string    `xorm:"JSON"`
	ProjectNo     string    `xorm:"JSON"`
	Quantity      float64   `xorm:"default 0 DOUBLE"`
	ProductUom    string    `xorm:"VARCHAR(25)"`
	CancelStatus  string    `xorm:"ENUM('ADMIN_CANCELLED','NOT_CANCELLED','SYSTEM_CANCELLED','USER_CANCELLED')"`
	UpdatedAt     time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
}

func (m *CmsStockAdjustmentDtl) TableName() string {
	return "cms_stock_adjustment_dtl"
}
