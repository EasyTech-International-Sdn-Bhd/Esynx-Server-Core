package entities

import (
	"time"
)

type CmsWarehouseStock struct {
	Id             uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	ProductCode    string    `xorm:"not null unique(unq) VARCHAR(200)" json:"productCode,omitempty" xml:"productCode"`
	WhCode         string    `xorm:"not null unique(unq) VARCHAR(200)" json:"whCode,omitempty" xml:"whCode"`
	ReadyStQty     float64   `xorm:"default 0 DOUBLE" json:"readyStQty,omitempty" xml:"readyStQty"`
	AvailableStQty float64   `xorm:"default 0 DOUBLE" json:"availableStQty,omitempty" xml:"availableStQty"`
	PoStQty        float64   `xorm:"default 0 DOUBLE" json:"poStQty,omitempty" xml:"poStQty"`
	SoStQty        float64   `xorm:"default 0 DOUBLE" json:"soStQty,omitempty" xml:"soStQty"`
	CloudQty       float64   `xorm:"default 0 DOUBLE" json:"cloudQty,omitempty" xml:"cloudQty"`
	UomName        string    `xorm:"not null default '' unique(unq) VARCHAR(200)" json:"uomName,omitempty" xml:"uomName"`
	ItemLocation   string    `xorm:"default '' VARCHAR(200)" json:"itemLocation,omitempty" xml:"itemLocation"`
	ActiveStatus   int       `xorm:"default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	RefNo          string    `xorm:"comment('uniquekey -- wh_code + product_code + uom') VARCHAR(200)" json:"refNo,omitempty" xml:"refNo"`
	UpdatedAt      time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
	SessionId      string    `xorm:"default '' VARCHAR(100)" json:"sessionId,omitempty" xml:"sessionId"`
}

func (m *CmsWarehouseStock) TableName() string {
	return "cms_warehouse_stock"
}

func (m *CmsWarehouseStock) BeforeInsert() {
	m.BeforeUpdate()
}

func (m *CmsWarehouseStock) BeforeUpdate() {
	m.UpdatedAt = time.Now()
}
