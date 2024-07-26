package entities

import (
	"time"
)

type CmsStockAdjustmentDtl struct {
	Id            uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	AdjId         string    `xorm:"not null unique(unq) VARCHAR(25)" json:"adjId,omitempty" xml:"adjId"`
	DeviceItemId  int64     `xorm:"default 0 unique(unq) BIGINT" json:"deviceItemId,omitempty" xml:"deviceItemId"`
	ProductCode   string    `xorm:"VARCHAR(25)" json:"productCode,omitempty" xml:"productCode"`
	WarehouseCode string    `xorm:"VARCHAR(25)" json:"warehouseCode,omitempty" xml:"warehouseCode"`
	BatchNo       string    `xorm:"JSON" json:"batchNo,omitempty" xml:"batchNo"`
	ProjectNo     string    `xorm:"JSON" json:"projectNo,omitempty" xml:"projectNo"`
	Quantity      float64   `xorm:"default 0 DOUBLE" json:"quantity,omitempty" xml:"quantity"`
	ProductUom    string    `xorm:"VARCHAR(25)" json:"productUom,omitempty" xml:"productUom"`
	CancelStatus  string    `xorm:"ENUM('ADMIN_CANCELLED','NOT_CANCELLED','SYSTEM_CANCELLED','USER_CANCELLED')" json:"cancelStatus,omitempty" xml:"cancelStatus"`
	UpdatedAt     time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsStockAdjustmentDtl) TableName() string {
	return "cms_stock_adjustment_dtl"
}
