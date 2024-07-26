package entities

import (
	"time"
)

type CmsProduct struct {
	ProductId                uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"productId,omitempty" xml:"productId"`
	CategoryId               int       `xorm:"default 0 comment('0 means roof.') index INT" json:"categoryId,omitempty" xml:"categoryId"`
	Productidentifierid      string    `xorm:"not null default '' VARCHAR(20)" json:"productidentifierid,omitempty" xml:"productidentifierid"`
	ProductCode              string    `xorm:"unique VARCHAR(50)" json:"productCode,omitempty" xml:"productCode"`
	QrCode                   string    `xorm:"VARCHAR(100)" json:"qrCode,omitempty" xml:"qrCode"`
	ProductName              string    `xorm:"VARCHAR(400)" json:"productName,omitempty" xml:"productName"`
	ProductDesc              string    `xorm:"comment('the product desc is the THML tag format') BLOB" json:"productDesc,omitempty" xml:"productDesc"`
	ProductRemark            string    `xorm:"comment('product remark is normal text format') BLOB" json:"productRemark,omitempty" xml:"productRemark"`
	ProductPromo             string    `xorm:"BLOB" json:"productPromo,omitempty" xml:"productPromo"`
	SequenceNo               int       `xorm:"INT" json:"sequenceNo,omitempty" xml:"sequenceNo"`
	ProductStatus            int       `xorm:"default 1 comment('1=active, 0=inactive') INT" json:"productStatus,omitempty" xml:"productStatus"`
	ProductCurrentQuantity   float64   `xorm:"default 0 comment('it can be N/A, or 9999. the quantity will be deduct when order transfer to CMS') DOUBLE" json:"productCurrentQuantity,omitempty" xml:"productCurrentQuantity"`
	ProductAvailableQuantity float64   `xorm:"default 0 DOUBLE" json:"productAvailableQuantity,omitempty" xml:"productAvailableQuantity"`
	ProductVirtualQuantity   float64   `xorm:"not null default 0 DOUBLE" json:"productVirtualQuantity,omitempty" xml:"productVirtualQuantity"`
	ProductGroupId           int       `xorm:"default 0 INT" json:"productGroupId,omitempty" xml:"productGroupId"`
	ProductCostPrice         float64   `xorm:"default 0 DOUBLE" json:"productCostPrice,omitempty" xml:"productCostPrice"`
	SearchFilter             string    `xorm:"JSON" json:"searchFilter,omitempty" xml:"searchFilter"`
	IsReplacement            int       `xorm:"not null default 1 INT" json:"isReplacement,omitempty" xml:"isReplacement"`
	SstCode                  string    `xorm:"VARCHAR(10)" json:"sstCode,omitempty" xml:"sstCode"`
	SstAmount                float64   `xorm:"not null default 0 DOUBLE" json:"sstAmount,omitempty" xml:"sstAmount"`
	UpdatedAt                time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
	Moderator                string    `xorm:"VARCHAR(10)" json:"moderator,omitempty" xml:"moderator"`
	BrandId                  int       `xorm:"default 0 INT" json:"brandId,omitempty" xml:"brandId"`
	IsStockLevel             int       `xorm:"default 0 INT" json:"isStockLevel,omitempty" xml:"isStockLevel"`
	ProductUdf               string    `xorm:"JSON" json:"productUdf,omitempty" xml:"productUdf"`
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
