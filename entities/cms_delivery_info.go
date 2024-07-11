package entities

import (
	"time"
)

type CmsDeliveryInfo struct {
	Id        int       `xorm:"not null pk autoincr INT"`
	OrderId   string    `xorm:"unique VARCHAR(20)"`
	Message   string    `xorm:"VARCHAR(500)"`
	UpdatedAt time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (m *CmsDeliveryInfo) TableName() string {
	return "cms_delivery_info"
}
