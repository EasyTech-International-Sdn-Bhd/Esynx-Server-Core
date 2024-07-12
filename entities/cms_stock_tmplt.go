package entities

import (
	"time"
)

type CmsStockTmplt struct {
	Id           uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	TmpltName    string    `xorm:"VARCHAR(200)"`
	ActiveStatus int       `xorm:"default 1 INT"`
	CreatedAt    time.Time `xorm:"DATETIME"`
	CreatedBy    string    `xorm:"default '' VARCHAR(50)"`
	UpdatedBy    string    `xorm:"default '' VARCHAR(50)"`
	UpdatedAt    time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
}

func (m *CmsStockTmplt) TableName() string {
	return "cms_stock_tmplt"
}
