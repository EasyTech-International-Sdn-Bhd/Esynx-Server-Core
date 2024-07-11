package entities

import (
	"time"
)

type CmsProductImage struct {
	ProductImageId          int       `xorm:"not null pk autoincr INT"`
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
