package entities

import (
	"time"
)

type CmsProductGroup struct {
	Id           int       `xorm:"not null pk autoincr INT"`
	Name         string    `xorm:"unique VARCHAR(100)"`
	Description  string    `xorm:"LONGTEXT(4294967295)"`
	ProductCode  string    `xorm:"VARCHAR(250)"`
	DateCreated  time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
	CategoryId   int       `xorm:"default 0 INT"`
	ActiveStatus int       `xorm:"default 1 INT"`
}

func (m *CmsProductGroup) TableName() string {
	return "cms_product_group"
}
