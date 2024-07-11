package entities

import (
	"time"
)

type CmsCreditnoteSales struct {
	CnId             int       `xorm:"not null pk autoincr INT"`
	CnCode           string    `xorm:"unique(UNIQUE) VARCHAR(20)"`
	CustCode         string    `xorm:"unique(UNIQUE) VARCHAR(20)"`
	CnDate           time.Time `xorm:"TIMESTAMP"`
	CnUdf            string    `xorm:"not null JSON"`
	CnAmount         float64   `xorm:"DOUBLE"`
	SalespersonId    int       `xorm:"INT"`
	Cancelled        string    `xorm:"CHAR(1)"`
	UpdatedAt        time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	CnKnockoffAmount float64   `xorm:"DOUBLE"`
	Approved         int       `xorm:"default 0 INT"`
	Approver         string    `xorm:"VARCHAR(200)"`
	ApprovedAt       time.Time `xorm:"DATETIME"`
}

func (m *CmsCreditnoteSales) TableName() string {
	return "cms_creditnote_sales"
}
