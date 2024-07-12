package entities

import (
	"time"
)

type CmsProductImage struct {
	ProductImageId          uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	ProductId               int       `xorm:"unique(unique_key) INT"`
	ImageUrl                string    `xorm:"unique(unique_key) VARCHAR(400)"`
	SequenceNo              int       `xorm:"INT"`
	ProductDefaultImage     int       `xorm:"default 0 comment('Each product only can select 1 default price, 0=not default image, 1=default') INT"`
	ActiveStatus            int       `xorm:"default 1 comment('1=active, 0 =inactive') INT"`
	ProductImageCreatedDate time.Time `xorm:"comment('When ipad download latest product list, the download function checks and compare the last update date with product image created date, only the newly uploaded image to be included in the download.') DATETIME"`
	UpdatedAt               time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	ProductCode             string    `xorm:"default '' VARCHAR(100)"`
}

func (m *CmsProductImage) TableName() string {
	return "cms_product_image"
}

func (m *CmsProductImage) Validate() {
	if m.ProductImageCreatedDate.IsZero() {
		m.ProductImageCreatedDate = time.Now()
	}
}

func (m *CmsProductImage) ToUpdate() {
	m.UpdatedAt = time.Now()
}
