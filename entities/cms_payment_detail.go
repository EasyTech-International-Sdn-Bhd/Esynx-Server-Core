package entities

import (
	"time"
)

type CmsPaymentDetail struct {
	PaymentDetailId     int       `xorm:"not null pk autoincr INT"`
	PaymentId           string    `xorm:"unique VARCHAR(50)"`
	MobileDetailId      int       `xorm:"INT"`
	PaymentMethod       string    `xorm:"VARCHAR(50)"`
	PaymentBy           string    `xorm:"VARCHAR(50)"`
	ChequeNo            string    `xorm:"VARCHAR(50)"`
	PaymentAmount       float64   `xorm:"DOUBLE"`
	PaymentDetailRemark string    `xorm:"VARCHAR(500)"`
	PaymentAttachment   string    `xorm:"VARCHAR(500)"`
	CancelStatus        int       `xorm:"INT"`
	UpdatedAt           time.Time `xorm:"DATETIME"`
}

func (m *CmsPaymentDetail) TableName() string {
	return "cms_payment_detail"
}
