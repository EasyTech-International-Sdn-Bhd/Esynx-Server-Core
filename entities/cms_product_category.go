package entities

import (
	"time"
)

type CmsProductCategory struct {
	CategoryId           int       `xorm:"not null pk autoincr index INT"`
	Categoryidentifierid string    `xorm:"not null unique VARCHAR(20)"`
	CategoryName         string    `xorm:"VARCHAR(400)"`
	ParentCategoryId     int       `xorm:"default 0 INT"`
	SequenceNo           int       `xorm:"INT"`
	CategoryStatus       int       `xorm:"default 1 comment('1=active, 0=inactive') INT"`
	CategoryImageUrl     string    `xorm:"VARCHAR(300)"`
	UpdatedAt            time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	Moderator            string    `xorm:"VARCHAR(10)"`
}

func (m *CmsProductCategory) TableName() string {
	return "cms_product_category"
}
