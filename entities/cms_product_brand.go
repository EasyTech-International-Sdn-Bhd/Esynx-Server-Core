package entities

import (
	"time"
)

type CmsProductBrand struct {
	BrandId       uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"brandId,omitempty" xml:"brandId"`
	BrandCode     string    `xorm:"not null unique VARCHAR(20)" json:"brandCode,omitempty" xml:"brandCode"`
	BrandName     string    `xorm:"VARCHAR(400)" json:"brandName,omitempty" xml:"brandName"`
	ParentBrandId int       `xorm:"default 0 INT" json:"parentBrandId,omitempty" xml:"parentBrandId"`
	SequenceNo    int       `xorm:"INT" json:"sequenceNo,omitempty" xml:"sequenceNo"`
	BrandStatus   int       `xorm:"default 1 comment('1=active, 0=inactive') INT" json:"brandStatus,omitempty" xml:"brandStatus"`
	BrandImageUrl string    `xorm:"VARCHAR(300)" json:"brandImageUrl,omitempty" xml:"brandImageUrl"`
	UpdatedAt     time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
	Moderator     string    `xorm:"not null VARCHAR(10)" json:"moderator,omitempty" xml:"moderator"`
}

func (m *CmsProductBrand) TableName() string {
	return "cms_product_brand"
}
