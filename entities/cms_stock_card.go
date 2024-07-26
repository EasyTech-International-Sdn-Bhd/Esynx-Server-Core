package entities

import (
	"time"
)

type CmsStockCard struct {
	Id           uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	StockDtlKey  string    `xorm:"not null unique(unique) VARCHAR(25)" json:"stockDtlKey,omitempty" xml:"stockDtlKey"`
	ProductCode  string    `xorm:"not null VARCHAR(50)" json:"productCode,omitempty" xml:"productCode"`
	Location     string    `xorm:"VARCHAR(60)" json:"location,omitempty" xml:"location"`
	BatchNo      string    `xorm:"VARCHAR(200)" json:"batchNo,omitempty" xml:"batchNo"`
	UnitUom      string    `xorm:"VARCHAR(200)" json:"unitUom,omitempty" xml:"unitUom"`
	DocDate      time.Time `xorm:"DATETIME" json:"docDate,omitempty" xml:"docDate"`
	DocType      string    `xorm:"unique(unique) VARCHAR(10)" json:"docType,omitempty" xml:"docType"`
	DocNo        string    `xorm:"VARCHAR(25)" json:"docNo,omitempty" xml:"docNo"`
	DocKey       string    `xorm:"unique(unique) VARCHAR(100)" json:"docKey,omitempty" xml:"docKey"`
	DtlKey       string    `xorm:"VARCHAR(200)" json:"dtlKey,omitempty" xml:"dtlKey"`
	Quantity     int       `xorm:"INT" json:"quantity,omitempty" xml:"quantity"`
	UnitPrice    float64   `xorm:"DOUBLE" json:"unitPrice,omitempty" xml:"unitPrice"`
	Total        float64   `xorm:"default 0 DOUBLE" json:"total,omitempty" xml:"total"`
	CostType     string    `xorm:"default '0' VARCHAR(200)" json:"costType,omitempty" xml:"costType"`
	ReferTo      string    `xorm:"VARCHAR(200)" json:"referTo,omitempty" xml:"referTo"`
	InputCost    float64   `xorm:"default 0 DOUBLE" json:"inputCost,omitempty" xml:"inputCost"`
	LastModified time.Time `xorm:"DATETIME" json:"lastModified,omitempty" xml:"lastModified"`
	Cancelled    string    `xorm:"default 'F' VARCHAR(10)" json:"cancelled,omitempty" xml:"cancelled"`
	UpdatedAt    time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsStockCard) TableName() string {
	return "cms_stock_card"
}
