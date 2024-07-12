package entities

import (
	"time"
)

type CmsInvoice struct {
	InvoiceId         uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
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
	InvUdf            string    `xorm:"JSON"`
	RefNo             int       `xorm:"INT"`
	DocType           string    `xorm:"default 'IV' ENUM('CS','IV')"`
	FromDoc           string    `xorm:"default 'SL' ENUM('AR','SL','CS')"`
	Cancelled         string    `xorm:"CHAR(1)"`
	UpdatedAt         time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	Termcode          string    `xorm:"VARCHAR(20)"`
}

func (m *CmsInvoice) TableName() string {
	return "cms_invoice"
}

func (m *CmsInvoice) Validate() {
}

func (m *CmsInvoice) ToUpdate() {
	m.UpdatedAt = time.Now()
}
