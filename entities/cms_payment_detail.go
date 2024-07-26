package entities

import (
	"time"
)

type CmsPaymentDetail struct {
	PaymentDetailId     uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"paymentDetailId,omitempty" xml:"paymentDetailId"`
	PaymentId           string    `xorm:"unique VARCHAR(50)" json:"paymentId,omitempty" xml:"paymentId"`
	MobileDetailId      int       `xorm:"INT" json:"mobileDetailId,omitempty" xml:"mobileDetailId"`
	PaymentMethod       string    `xorm:"VARCHAR(50)" json:"paymentMethod,omitempty" xml:"paymentMethod"`
	PaymentBy           string    `xorm:"VARCHAR(50)" json:"paymentBy,omitempty" xml:"paymentBy"`
	ChequeNo            string    `xorm:"VARCHAR(50)" json:"chequeNo,omitempty" xml:"chequeNo"`
	PaymentAmount       float64   `xorm:"DOUBLE" json:"paymentAmount,omitempty" xml:"paymentAmount"`
	PaymentDetailRemark string    `xorm:"VARCHAR(500)" json:"paymentDetailRemark,omitempty" xml:"paymentDetailRemark"`
	PaymentAttachment   string    `xorm:"VARCHAR(500)" json:"paymentAttachment,omitempty" xml:"paymentAttachment"`
	CancelStatus        int       `xorm:"INT" json:"cancelStatus,omitempty" xml:"cancelStatus"`
	UpdatedAt           time.Time `xorm:"DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsPaymentDetail) TableName() string {
	return "cms_payment_detail"
}
