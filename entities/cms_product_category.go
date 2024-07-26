package entities

import (
	"time"
)

type CmsProductCategory struct {
	CategoryId           uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"categoryId,omitempty" xml:"categoryId"`
	Categoryidentifierid string    `xorm:"not null unique VARCHAR(20)" json:"categoryidentifierid,omitempty" xml:"categoryidentifierid"`
	CategoryName         string    `xorm:"VARCHAR(400)" json:"categoryName,omitempty" xml:"categoryName"`
	ParentCategoryId     int       `xorm:"default 0 INT" json:"parentCategoryId,omitempty" xml:"parentCategoryId"`
	SequenceNo           int       `xorm:"INT" json:"sequenceNo,omitempty" xml:"sequenceNo"`
	CategoryStatus       int       `xorm:"default 1 comment('1=active, 0=inactive') INT" json:"categoryStatus,omitempty" xml:"categoryStatus"`
	CategoryImageUrl     string    `xorm:"VARCHAR(300)" json:"categoryImageUrl,omitempty" xml:"categoryImageUrl"`
	UpdatedAt            time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
	Moderator            string    `xorm:"VARCHAR(10)" json:"moderator,omitempty" xml:"moderator"`
}

func (m *CmsProductCategory) TableName() string {
	return "cms_product_category"
}
