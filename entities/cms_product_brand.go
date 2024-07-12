package entities

import (
	"time"
)

type CmsProductBrand struct {
	BrandId       uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	BrandCode     string    `xorm:"not null unique VARCHAR(20)"`
	BrandName     string    `xorm:"VARCHAR(400)"`
	ParentBrandId int       `xorm:"default 0 INT"`
	SequenceNo    int       `xorm:"INT"`
	BrandStatus   int       `xorm:"default 1 comment('1=active, 0=inactive') INT"`
	BrandImageUrl string    `xorm:"VARCHAR(300)"`
	UpdatedAt     time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	Moderator     string    `xorm:"not null VARCHAR(10)"`
}

func (m *CmsProductBrand) TableName() string {
	return "cms_product_brand"
}
