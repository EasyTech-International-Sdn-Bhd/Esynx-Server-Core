package entities

import (
	"time"
)

type CmsProductUomPriceV2 struct {
	ProductUomPriceId   uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"productUomPriceId,omitempty" xml:"productUomPriceId"`
	ProductCode         string    `xorm:"unique(unique_uom) VARCHAR(50)" json:"productCode,omitempty" xml:"productCode"`
	ProductUom          string    `xorm:"unique(unique_uom) VARCHAR(50)" json:"productUom,omitempty" xml:"productUom"`
	ProductUomRate      float64   `xorm:"default 0 unique(unique_uom) DOUBLE" json:"productUomRate,omitempty" xml:"productUomRate"`
	ProductStdPrice     float64   `xorm:"default 0 DOUBLE" json:"productStdPrice,omitempty" xml:"productStdPrice"`
	ProductMinPrice     float64   `xorm:"default 0 DOUBLE" json:"productMinPrice,omitempty" xml:"productMinPrice"`
	ProductDefaultPrice int       `xorm:"default 0 comment('Each product only can select 1 default price, 0=not default, 1=default') INT" json:"productDefaultPrice,omitempty" xml:"productDefaultPrice"`
	ActiveStatus        int       `xorm:"default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	UpdatedAt           time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
	QrCode              string    `xorm:"VARCHAR(30)" json:"qrCode,omitempty" xml:"qrCode"`
	RefNo               string    `xorm:"VARCHAR(200)" json:"refNo,omitempty" xml:"refNo"`
}

func (m *CmsProductUomPriceV2) TableName() string {
	return "cms_product_uom_price_v2"
}

func (m *CmsProductUomPriceV2) BeforeInsert() {
	m.BeforeUpdate()
}

func (m *CmsProductUomPriceV2) BeforeUpdate() {
	m.UpdatedAt = time.Now()
}
