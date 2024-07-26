package entities

import (
	"time"
)

type CmsProductAtch struct {
	Id           uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	ProductCode  string    `xorm:"not null VARCHAR(50)" json:"productCode,omitempty" xml:"productCode"`
	ContentType  string    `xorm:"not null VARCHAR(50)" json:"contentType,omitempty" xml:"contentType"`
	Link         string    `xorm:"not null VARCHAR(500)" json:"link,omitempty" xml:"link"`
	ActiveStatus int       `xorm:"not null default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	CreatedDate  time.Time `xorm:"DATETIME" json:"createdDate,omitempty" xml:"createdDate"`
	UpdatedAt    time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsProductAtch) TableName() string {
	return "cms_product_atch"
}

func (m *CmsProductAtch) BeforeInsert() {
	m.BeforeUpdate()
}

func (m *CmsProductAtch) BeforeUpdate() {
	m.UpdatedAt = time.Now()
}
