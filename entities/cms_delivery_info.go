package entities

import (
	"time"
)

type CmsDeliveryInfo struct {
	Id        uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	OrderId   string    `xorm:"unique VARCHAR(20)" json:"orderId,omitempty" xml:"orderId"`
	Message   string    `xorm:"VARCHAR(500)" json:"message,omitempty" xml:"message"`
	UpdatedAt time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsDeliveryInfo) TableName() string {
	return "cms_delivery_info"
}
