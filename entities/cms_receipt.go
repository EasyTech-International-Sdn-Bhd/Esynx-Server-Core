package entities

import (
	"time"
)

type CmsReceipt struct {
	ReceiptId             uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	ReceiptCode           string    `xorm:"index unique VARCHAR(20)"`
	CustCode              string    `xorm:"index VARCHAR(20)"`
	ReceiptDate           time.Time `xorm:"TIMESTAMP"`
	ReceiptAmount         float64   `xorm:"DOUBLE"`
	ReceiptKnockoffAmount float64   `xorm:"DOUBLE"`
	ReceiptDesc           []byte    `xorm:"LONGBLOB"`
	ChequeNo              string    `xorm:"VARCHAR(100)"`
	SalespersonId         int       `xorm:"default 0 INT"`
	ReceiptUdf            string    `xorm:"not null JSON"`
	Approved              int       `xorm:"default 0 INT"`
	Approver              string    `xorm:"VARCHAR(200)"`
	ApprovedAt            time.Time `xorm:"DATETIME"`
	Cancelled             string    `xorm:"CHAR(1)"`
	UpdatedAt             time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	RefNo                 string    `xorm:"VARCHAR(20)"`
}

func (m *CmsReceipt) TableName() string {
	return "cms_receipt"
}
