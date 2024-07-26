package entities

import (
	"time"
)

type CmsProductAtch struct {
	Id           uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	ProductCode  string    `xorm:"not null VARCHAR(50)"`
	ContentType  string    `xorm:"not null VARCHAR(50)"`
	Link         string    `xorm:"not null VARCHAR(500)"`
	ActiveStatus int       `xorm:"not null default 1 INT"`
	CreatedDate  time.Time `xorm:"DATETIME"`
	UpdatedAt    time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
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
