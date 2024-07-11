package entities

import (
	"time"
)

type CmsProductPriceV2 struct {
	ProductPriceId int64     `xorm:"not null pk autoincr BIGINT"`
	ProductCode    string    `xorm:"unique(unique) VARCHAR(200)"`
	PriceCat       string    `xorm:"unique(unique) VARCHAR(50)"`
	CustCode       string    `xorm:"VARCHAR(30)"`
	ProductPrice   float64   `xorm:"default 0 DOUBLE"`
	Disc1          float64   `xorm:"default 0 DOUBLE"`
	Disc2          float64   `xorm:"default 0 DOUBLE"`
	Disc3          float64   `xorm:"default 0 DOUBLE"`
	DateFrom       time.Time `xorm:"DATETIME"`
	DateTo         time.Time `xorm:"DATETIME"`
	ProductUom     string    `xorm:"unique(unique) VARCHAR(50)"`
	PriceRemark    string    `xorm:"VARCHAR(50)"`
	ActiveStatus   int       `xorm:"default 1 INT"`
	Quantity       float64   `xorm:"default 1 DOUBLE"`
	UpdatedAt      time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	RefNo          string    `xorm:"VARCHAR(100)"`
	MixDisc        string    `xorm:"VARCHAR(100)"`
}

func (m *CmsProductPriceV2) TableName() string {
	return "cms_product_price_v2"
}

func (m *CmsProductPriceV2) Validate() {
	if m.DateFrom.IsZero() {
		m.DateFrom = time.Now().AddDate(-1, 0, 0)
	}
	if m.DateTo.IsZero() {
		m.DateTo = time.Now().AddDate(10, 0, 0)
	}
}

func (m *CmsProductPriceV2) ToUpdate() {
	m.UpdatedAt = time.Now()
}
