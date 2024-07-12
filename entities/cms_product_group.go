package entities

import (
	"time"
)

type CmsProductGroup struct {
	Id           uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	Name         string    `xorm:"unique VARCHAR(100)"`
	Description  string    `xorm:"BLOB"`
	ProductCode  string    `xorm:"VARCHAR(250)"`
	DateCreated  time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
	CategoryId   int       `xorm:"default 0 INT"`
	ActiveStatus int       `xorm:"default 1 INT"`
}

func (m *CmsProductGroup) TableName() string {
	return "cms_product_group"
}
