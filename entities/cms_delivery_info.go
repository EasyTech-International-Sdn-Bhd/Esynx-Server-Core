package entities

import (
	"time"
)

type CmsDeliveryInfo struct {
	Id        uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	OrderId   string    `xorm:"unique VARCHAR(20)"`
	Message   string    `xorm:"VARCHAR(500)"`
	UpdatedAt time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (m *CmsDeliveryInfo) TableName() string {
	return "cms_delivery_info"
}
