package entities

import (
	"time"
)

type CmsProduct struct {
	ProductId                uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	CategoryId               int       `xorm:"default 0 comment('0 means roof.') index INT"`
	Productidentifierid      string    `xorm:"not null default '' VARCHAR(20)"`
	ProductCode              string    `xorm:"unique VARCHAR(50)"`
	QrCode                   string    `xorm:"VARCHAR(100)"`
	ProductName              string    `xorm:"VARCHAR(400)"`
	ProductDesc              string    `xorm:"comment('the product desc is the THML tag format') BLOB"`
	ProductRemark            string    `xorm:"comment('product remark is normal text format') BLOB"`
	ProductPromo             string    `xorm:"BLOB"`
	SequenceNo               int       `xorm:"INT"`
	ProductStatus            int       `xorm:"default 1 comment('1=active, 0=inactive') INT"`
	ProductCurrentQuantity   float64   `xorm:"default 0 comment('it can be N/A, or 9999. the quantity will be deduct when order transfer to CMS') DOUBLE"`
	ProductAvailableQuantity float64   `xorm:"default 0 DOUBLE"`
	ProductVirtualQuantity   float64   `xorm:"not null default 0 DOUBLE"`
	ProductGroupId           int       `xorm:"default 0 INT"`
	ProductCostPrice         float64   `xorm:"default 0 DOUBLE"`
	SearchFilter             string    `xorm:"JSON"`
	IsReplacement            int       `xorm:"not null default 1 INT"`
	SstCode                  string    `xorm:"VARCHAR(10)"`
	SstAmount                float64   `xorm:"not null default 0 DOUBLE"`
	UpdatedAt                time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	Moderator                string    `xorm:"VARCHAR(10)"`
	BrandId                  int       `xorm:"default 0 INT"`
	IsStockLevel             int       `xorm:"default 0 INT"`
	ProductUdf               string    `xorm:"JSON"`
}

func (m *CmsProduct) TableName() string {
	return "cms_product"
}

func (m *CmsProduct) BeforeInsert() {
	m.BeforeUpdate()
}

func (m *CmsProduct) BeforeUpdate() {
	m.UpdatedAt = time.Now()
	if m.SearchFilter == "" {
		m.SearchFilter = "{}"
	}
	if m.ProductUdf == "" {
		m.ProductUdf = "{}"
	}
}
