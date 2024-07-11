package entities

import (
	"time"
)

type CmsWarehouseStock struct {
	Id             int       `xorm:"not null pk autoincr INT"`
	ProductCode    string    `xorm:"not null unique(unq) VARCHAR(200)"`
	WhCode         string    `xorm:"not null unique(unq) VARCHAR(200)"`
	ReadyStQty     float64   `xorm:"default 0 DOUBLE"`
	AvailableStQty float64   `xorm:"default 0 DOUBLE"`
	PoStQty        float64   `xorm:"default 0 DOUBLE"`
	SoStQty        float64   `xorm:"default 0 DOUBLE"`
	CloudQty       float64   `xorm:"default 0 DOUBLE"`
	UomName        string    `xorm:"not null default '' unique(unq) VARCHAR(200)"`
	ItemLocation   string    `xorm:"default '' VARCHAR(200)"`
	ActiveStatus   int       `xorm:"default 1 INT"`
	RefNo          string    `xorm:"comment('uniquekey -- wh_code + product_code + uom') VARCHAR(200)"`
	UpdatedAt      time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
	SessionId      string    `xorm:"default '' VARCHAR(100)"`
}

func (m *CmsWarehouseStock) TableName() string {
	return "cms_warehouse_stock"
}

func (m *CmsWarehouseStock) Validate() {
	// NOTE: nothing for now
}

func (m *CmsWarehouseStock) ToUpdate() {
	m.UpdatedAt = time.Now()
}
