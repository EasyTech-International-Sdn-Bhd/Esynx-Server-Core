package entities

import (
	"time"
)

type CmsPayment struct {
	PaymentId                   string    `xorm:"unique VARCHAR(50)"`
	PaymentDate                 time.Time `xorm:"DATETIME"`
	PaymentTransferReceivedDate time.Time `xorm:"DATETIME"`
	CustCode                    string    `xorm:"VARCHAR(50)"`
	Description                 string    `xorm:"VARCHAR(500)"`
	PaymentAmount               float64   `xorm:"DOUBLE"`
	PaymentStatus               int       `xorm:"comment('0=in_ipad,1=in_backoffice,2=in_Accounting') INT"`
	CancelStatus                int       `xorm:"INT"`
	PaymentReference            string    `xorm:"VARCHAR(50)"`
	SalespersonId               int       `xorm:"INT"`
	SalespersonPaymentRemark    string    `xorm:"VARCHAR(500)"`
	LastPrint                   time.Time `xorm:"DATETIME"`
	Checked                     int       `xorm:"default 0 INT"`
	PaymentFault                int       `xorm:"default 0 INT"`
	PaymentFaultMessage         string    `xorm:"VARCHAR(200)"`
	PaymentStatusLastUpdateDate time.Time `xorm:"DATETIME"`
	PaymentStatusLastUpdateBy   string    `xorm:"VARCHAR(50)"`
	KnockoffInv                 string    `xorm:"JSON"`
	DocType                     string    `xorm:"default 'payment' VARCHAR(10)"`
	DocId                       string    `xorm:"VARCHAR(20)"`
}

func (m *CmsPayment) TableName() string {
	return "cms_payment"
}
