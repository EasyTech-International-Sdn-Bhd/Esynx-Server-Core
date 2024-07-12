package entities

import (
	"time"
)

type CmsPaymentGatewayLog struct {
	Id                uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	PaymentId         string    `xorm:"not null VARCHAR(25)"`
	ActionTaken       string    `xorm:"default 'before_payment' ENUM('after_payment','before_payment')"`
	ActionDescription string    `xorm:"VARCHAR(255)"`
	ActionCode        int       `xorm:"default 0 INT"`
	ActionTime        time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
}

func (m *CmsPaymentGatewayLog) TableName() string {
	return "cms_payment_gateway_log"
}
