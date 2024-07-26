package entities

import (
	"time"
)

type CmsProductGroup struct {
	Id           uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	Name         string    `xorm:"unique VARCHAR(100)" json:"name,omitempty" xml:"name"`
	Description  string    `xorm:"BLOB" json:"description,omitempty" xml:"description"`
	ProductCode  string    `xorm:"VARCHAR(250)" json:"productCode,omitempty" xml:"productCode"`
	DateCreated  time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP" json:"dateCreated,omitempty" xml:"dateCreated"`
	CategoryId   int       `xorm:"default 0 INT" json:"categoryId,omitempty" xml:"categoryId"`
	ActiveStatus int       `xorm:"default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
}

func (m *CmsProductGroup) TableName() string {
	return "cms_product_group"
}
