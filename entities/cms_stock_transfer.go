package entities

import (
	"time"
)

type CmsStockTransfer struct {
	Id             uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	StCode         string    `xorm:"not null unique VARCHAR(200)"`
	StDate         time.Time `xorm:"not null DATETIME"`
	CustCode       string    `xorm:"comment('can be empty') VARCHAR(200)"`
	FromLocation   string    `xorm:"VARCHAR(200)"`
	ToLocation     string    `xorm:"VARCHAR(200)"`
	Total          float64   `xorm:"DOUBLE"`
	Note           string    `xorm:"VARCHAR(200)"`
	StStatus       int       `xorm:"comment('0 - in app; 1 - confirm; 2 - transferred to ATC') INT"`
	CancelStatus   int       `xorm:"default 0 INT"`
	SalespersonId  int       `xorm:"INT"`
	StFault        int       `xorm:"default 0 INT"`
	StFaultMessage string    `xorm:"VARCHAR(200)"`
}

func (m *CmsStockTransfer) TableName() string {
	return "cms_stock_transfer"
}
