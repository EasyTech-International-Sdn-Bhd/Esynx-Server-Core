package entities

import (
	"time"
)

type CmsPaymentGatewayLog struct {
	Id                uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	PaymentId         string    `xorm:"not null VARCHAR(25)" json:"paymentId,omitempty" xml:"paymentId"`
	ActionTaken       string    `xorm:"default 'before_payment' ENUM('after_payment','before_payment')" json:"actionTaken,omitempty" xml:"actionTaken"`
	ActionDescription string    `xorm:"VARCHAR(255)" json:"actionDescription,omitempty" xml:"actionDescription"`
	ActionCode        int       `xorm:"default 0 INT" json:"actionCode,omitempty" xml:"actionCode"`
	ActionTime        time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"actionTime,omitempty" xml:"actionTime"`
}

func (m *CmsPaymentGatewayLog) TableName() string {
	return "cms_payment_gateway_log"
}
