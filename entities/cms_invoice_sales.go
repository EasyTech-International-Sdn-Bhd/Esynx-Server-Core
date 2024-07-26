package entities

import (
	"time"
)

type CmsInvoiceSales struct {
	InvoiceId         uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"invoiceId,omitempty" xml:"invoiceId"`
	InvoiceCode       string    `xorm:"not null default '' unique VARCHAR(200)" json:"invoiceCode,omitempty" xml:"invoiceCode"`
	CustCode          string    `xorm:"VARCHAR(20)" json:"custCode,omitempty" xml:"custCode"`
	InvoiceDate       time.Time `xorm:"TIMESTAMP" json:"invoiceDate,omitempty" xml:"invoiceDate"`
	InvoiceDueDate    time.Time `xorm:"TIMESTAMP" json:"invoiceDueDate,omitempty" xml:"invoiceDueDate"`
	InvoiceAmount     float64   `xorm:"DOUBLE" json:"invoiceAmount,omitempty" xml:"invoiceAmount"`
	OutstandingAmount float64   `xorm:"DOUBLE" json:"outstandingAmount,omitempty" xml:"outstandingAmount"`
	Approved          int       `xorm:"default 0 INT" json:"approved,omitempty" xml:"approved"`
	Approver          string    `xorm:"VARCHAR(800)" json:"approver,omitempty" xml:"approver"`
	ApprovedAt        time.Time `xorm:"DATETIME" json:"approvedAt,omitempty" xml:"approvedAt"`
	SalespersonId     int       `xorm:"default 0 INT" json:"salespersonId,omitempty" xml:"salespersonId"`
	InvUdf            string    `xorm:"not null JSON" json:"invUdf,omitempty" xml:"invUdf"`
	Cancelled         string    `xorm:"CHAR(1)" json:"cancelled,omitempty" xml:"cancelled"`
	UpdatedAt         time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
	CreatedBy         string    `xorm:"VARCHAR(20)" json:"createdBy,omitempty" xml:"createdBy"`
	RefNo             int       `xorm:"INT" json:"refNo,omitempty" xml:"refNo"`
	Termcode          string    `xorm:"VARCHAR(20)" json:"termcode,omitempty" xml:"termcode"`
}

func (m *CmsInvoiceSales) TableName() string {
	return "cms_invoice_sales"
}

func (m *CmsInvoiceSales) BeforeInsert() {
	m.BeforeUpdate()
}

func (m *CmsInvoiceSales) BeforeUpdate() {
	m.UpdatedAt = time.Now()
}
