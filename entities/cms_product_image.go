package entities

import (
	"time"
)

type CmsProductImage struct {
	ProductImageId          uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"productImageId,omitempty" xml:"productImageId"`
	ProductId               int       `xorm:"unique(unique_key) INT" json:"productId,omitempty" xml:"productId"`
	ImageUrl                string    `xorm:"unique(unique_key) VARCHAR(400)" json:"imageUrl,omitempty" xml:"imageUrl"`
	SequenceNo              int       `xorm:"INT" json:"sequenceNo,omitempty" xml:"sequenceNo"`
	ProductDefaultImage     int       `xorm:"default 0 comment('Each product only can select 1 default price, 0=not default image, 1=default') INT" json:"productDefaultImage,omitempty" xml:"productDefaultImage"`
	ActiveStatus            int       `xorm:"default 1 comment('1=active, 0 =inactive') INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	ProductImageCreatedDate time.Time `xorm:"comment('When ipad download latest product list, the download function checks and compare the last update date with product image created date, only the newly uploaded image to be included in the download.') DATETIME" json:"productImageCreatedDate,omitempty" xml:"productImageCreatedDate"`
	UpdatedAt               time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
	ProductCode             string    `xorm:"default '' VARCHAR(100)" json:"productCode,omitempty" xml:"productCode"`
}

func (m *CmsProductImage) TableName() string {
	return "cms_product_image"
}

func (m *CmsProductImage) BeforeInsert() {
	m.BeforeUpdate()
}

func (m *CmsProductImage) BeforeUpdate() {
	if m.ProductImageCreatedDate.IsZero() {
		m.ProductImageCreatedDate = time.Now()
	}
	m.UpdatedAt = time.Now()
}
