package entities

import (
	"time"
)

type CmsReceipt struct {
	ReceiptId             uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"receiptId,omitempty" xml:"receiptId"`
	ReceiptCode           string    `xorm:"index unique VARCHAR(20)" json:"receiptCode,omitempty" xml:"receiptCode"`
	CustCode              string    `xorm:"index VARCHAR(20)" json:"custCode,omitempty" xml:"custCode"`
	ReceiptDate           time.Time `xorm:"TIMESTAMP" json:"receiptDate,omitempty" xml:"receiptDate"`
	ReceiptAmount         float64   `xorm:"DOUBLE" json:"receiptAmount,omitempty" xml:"receiptAmount"`
	ReceiptKnockoffAmount float64   `xorm:"DOUBLE" json:"receiptKnockoffAmount,omitempty" xml:"receiptKnockoffAmount"`
	ReceiptDesc           []byte    `xorm:"LONGBLOB" json:"receiptDesc,omitempty" xml:"receiptDesc"`
	ChequeNo              string    `xorm:"VARCHAR(100)" json:"chequeNo,omitempty" xml:"chequeNo"`
	SalespersonId         int       `xorm:"default 0 INT" json:"salespersonId,omitempty" xml:"salespersonId"`
	ReceiptUdf            string    `xorm:"not null JSON" json:"receiptUdf,omitempty" xml:"receiptUdf"`
	Approved              int       `xorm:"default 0 INT" json:"approved,omitempty" xml:"approved"`
	Approver              string    `xorm:"VARCHAR(200)" json:"approver,omitempty" xml:"approver"`
	ApprovedAt            time.Time `xorm:"DATETIME" json:"approvedAt,omitempty" xml:"approvedAt"`
	Cancelled             string    `xorm:"CHAR(1)" json:"cancelled,omitempty" xml:"cancelled"`
	UpdatedAt             time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
	RefNo                 string    `xorm:"VARCHAR(20)" json:"refNo,omitempty" xml:"refNo"`
}

func (m *CmsReceipt) TableName() string {
	return "cms_receipt"
}
