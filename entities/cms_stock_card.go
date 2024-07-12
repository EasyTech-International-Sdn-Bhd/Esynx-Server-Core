package entities

import (
	"time"
)

type CmsStockCard struct {
	Id           uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	StockDtlKey  string    `xorm:"not null unique(unique) VARCHAR(25)"`
	ProductCode  string    `xorm:"not null VARCHAR(50)"`
	Location     string    `xorm:"VARCHAR(60)"`
	BatchNo      string    `xorm:"VARCHAR(200)"`
	UnitUom      string    `xorm:"VARCHAR(200)"`
	DocDate      time.Time `xorm:"DATETIME"`
	DocType      string    `xorm:"unique(unique) VARCHAR(10)"`
	DocNo        string    `xorm:"VARCHAR(25)"`
	DocKey       string    `xorm:"unique(unique) VARCHAR(100)"`
	DtlKey       string    `xorm:"VARCHAR(200)"`
	Quantity     int       `xorm:"INT"`
	UnitPrice    float64   `xorm:"DOUBLE"`
	Total        float64   `xorm:"default 0 DOUBLE"`
	CostType     string    `xorm:"default '0' VARCHAR(200)"`
	ReferTo      string    `xorm:"VARCHAR(200)"`
	InputCost    float64   `xorm:"default 0 DOUBLE"`
	LastModified time.Time `xorm:"DATETIME"`
	Cancelled    string    `xorm:"default 'F' VARCHAR(10)"`
	UpdatedAt    time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (m *CmsStockCard) TableName() string {
	return "cms_stock_card"
}
