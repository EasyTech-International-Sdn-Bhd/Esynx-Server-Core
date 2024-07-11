package entities

import (
	"time"
)

type CmsStockTake struct {
	Id              int       `xorm:"not null pk autoincr INT"`
	StId            int       `xorm:"not null unique(unq) INT"`
	CustCode        string    `xorm:"not null VARCHAR(200)"`
	CustCompanyName string    `xorm:"not null VARCHAR(200)"`
	SalespersonId   int       `xorm:"not null unique(unq) INT"`
	CreatedDate     time.Time `xorm:"not null DATETIME"`
	DocRefId        string    `xorm:"not null VARCHAR(50)"`
	SpUpdatedAt     string    `xorm:"not null VARCHAR(50)"`
	ActiveStatus    int       `xorm:"default 1 INT"`
	UpdatedAt       time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
}

func (m *CmsStockTake) TableName() string {
	return "cms_stock_take"
}
