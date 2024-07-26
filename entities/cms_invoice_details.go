package entities

import (
	"time"
)

type CmsInvoiceDetails struct {
	Id           uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	InvoiceCode  string    `xorm:"not null unique(invoice_code) VARCHAR(50)" json:"invoiceCode,omitempty" xml:"invoiceCode"`
	ItemCode     string    `xorm:"not null VARCHAR(50)" json:"itemCode,omitempty" xml:"itemCode"`
	ItemName     string    `xorm:"not null VARCHAR(200)" json:"itemName,omitempty" xml:"itemName"`
	ItemPrice    float64   `xorm:"default 0 DOUBLE" json:"itemPrice,omitempty" xml:"itemPrice"`
	Quantity     float64   `xorm:"default 0 DOUBLE" json:"quantity,omitempty" xml:"quantity"`
	TotalPrice   float64   `xorm:"default 0 DOUBLE" json:"totalPrice,omitempty" xml:"totalPrice"`
	Uom          string    `xorm:"not null VARCHAR(20)" json:"uom,omitempty" xml:"uom"`
	Discount     string    `xorm:"default '0' VARCHAR(50)" json:"discount,omitempty" xml:"discount"`
	ActiveStatus int       `xorm:"default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	SequenceNo   int       `xorm:"not null default 0 INT" json:"sequenceNo,omitempty" xml:"sequenceNo"`
	InvDtlUdf    string    `xorm:"JSON" json:"invDtlUdf,omitempty" xml:"invDtlUdf"`
	RefNo        string    `xorm:"unique(invoice_code) VARCHAR(200)" json:"refNo,omitempty" xml:"refNo"`
	UpdatedAt    time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
	SessionId    string    `xorm:"default '' VARCHAR(100)" json:"sessionId,omitempty" xml:"sessionId"`
}

func (m *CmsInvoiceDetails) TableName() string {
	return "cms_invoice_details"
}

func (m *CmsInvoiceDetails) BeforeInsert() {
	m.BeforeUpdate()
}

func (m *CmsInvoiceDetails) BeforeUpdate() {
	m.UpdatedAt = time.Now()
}
