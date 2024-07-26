package entities

import (
	"time"
)

type CmsInvoiceDetails struct {
	Id           uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	InvoiceCode  string    `xorm:"not null unique(invoice_code) VARCHAR(50)"`
	ItemCode     string    `xorm:"not null VARCHAR(50)"`
	ItemName     string    `xorm:"not null VARCHAR(200)"`
	ItemPrice    float64   `xorm:"default 0 DOUBLE"`
	Quantity     float64   `xorm:"default 0 DOUBLE"`
	TotalPrice   float64   `xorm:"default 0 DOUBLE"`
	Uom          string    `xorm:"not null VARCHAR(20)"`
	Discount     string    `xorm:"default '0' VARCHAR(50)"`
	ActiveStatus int       `xorm:"default 1 INT"`
	SequenceNo   int       `xorm:"not null default 0 INT"`
	InvDtlUdf    string    `xorm:"JSON"`
	RefNo        string    `xorm:"unique(invoice_code) VARCHAR(200)"`
	UpdatedAt    time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	SessionId    string    `xorm:"default '' VARCHAR(100)"`
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
