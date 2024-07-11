package entities

import (
	"time"
)

type CmsStockAdjustment struct {
	Id                   []byte    `xorm:"not null pk default uuid_to_bin(uuid()) BINARY(16)"`
	AdjId                string    `xorm:"not null unique VARCHAR(25)"`
	AdjDate              time.Time `xorm:"DATE"`
	AdjDescription       string    `xorm:"VARCHAR(255)"`
	AdjNote              []byte    `xorm:"BLOB"`
	ActiveStatus         string    `xorm:"ENUM('ACTIVE','CANCELLED')"`
	CancelStatus         string    `xorm:"ENUM('ADMIN_CANCELLED','NOT_CANCELLED','SYSTEM_CANCELLED','USER_CANCELLED')"`
	TransferReceivedDate time.Time `xorm:"DATETIME"`
	LastUpdatedBy        int       `xorm:"default 0 INT"`
	Agent                string    `xorm:"VARCHAR(50)"`
	UpdatedAt            time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
}

func (m *CmsStockAdjustment) TableName() string {
	return "cms_stock_adjustment"
}
