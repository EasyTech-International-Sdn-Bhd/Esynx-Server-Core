package entities

import (
	"time"
)

type CmsInvoiceSales struct {
	InvoiceId         int64     `xorm:"not null pk autoincr BIGINT"`
	InvoiceCode       string    `xorm:"not null default '' unique VARCHAR(200)"`
	CustCode          string    `xorm:"VARCHAR(20)"`
	InvoiceDate       time.Time `xorm:"TIMESTAMP"`
	InvoiceDueDate    time.Time `xorm:"TIMESTAMP"`
	InvoiceAmount     float64   `xorm:"DOUBLE"`
	OutstandingAmount float64   `xorm:"DOUBLE"`
	Approved          int       `xorm:"default 0 INT"`
	Approver          string    `xorm:"VARCHAR(800)"`
	ApprovedAt        time.Time `xorm:"DATETIME"`
	SalespersonId     int       `xorm:"default 0 INT"`
	InvUdf            string    `xorm:"not null JSON"`
	Cancelled         string    `xorm:"CHAR(1)"`
	UpdatedAt         time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	CreatedBy         string    `xorm:"VARCHAR(20)"`
	RefNo             int       `xorm:"INT"`
}

func (m *CmsInvoiceSales) TableName() string {
	return "cms_invoice_sales"
}
