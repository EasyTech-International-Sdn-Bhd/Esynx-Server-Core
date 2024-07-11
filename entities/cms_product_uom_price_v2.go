package entities

import (
	"time"
)

type CmsProductUomPriceV2 struct {
	ProductUomPriceId   int       `xorm:"not null pk autoincr INT"`
	ProductCode         string    `xorm:"unique(unique_uom) VARCHAR(50)"`
	ProductUom          string    `xorm:"unique(unique_uom) VARCHAR(50)"`
	ProductUomRate      float64   `xorm:"default 0 unique(unique_uom) DOUBLE"`
	ProductStdPrice     float64   `xorm:"default 0 DOUBLE"`
	ProductMinPrice     float64   `xorm:"default 0 DOUBLE"`
	ProductDefaultPrice int       `xorm:"default 0 comment('Each product only can select 1 default price, 0=not default, 1=default') INT"`
	ActiveStatus        int       `xorm:"default 1 INT"`
	UpdatedAt           time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	QrCode              string    `xorm:"VARCHAR(30)"`
	RefNo               string    `xorm:"VARCHAR(200)"`
}

func (m *CmsProductUomPriceV2) TableName() string {
	return "cms_product_uom_price_v2"
}
