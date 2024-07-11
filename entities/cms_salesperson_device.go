package entities

import (
	"time"
)

type CmsSalespersonDevice struct {
	Id           int64     `xorm:"pk autoincr BIGINT"`
	Prefix       string    `xorm:"not null pk default 'D' CHAR(30)"`
	DeviceNo     int       `xorm:"default 1 INT"`
	DeviceToken  string    `xorm:"not null unique(unq) VARCHAR(200)"`
	LoginId      string    `xorm:"not null unique(unq) VARCHAR(100)"`
	AppVersion   string    `xorm:"VARCHAR(100)"`
	Os           string    `xorm:"not null VARCHAR(100)"`
	RegisteredAt time.Time `xorm:"not null default '1971-01-01 23:01:01' DATETIME"`
	ActiveStatus int       `xorm:"not null default 1 INT"`
	LastLoggedIn time.Time `xorm:"not null default CURRENT_TIMESTAMP DATETIME"`
}

func (m *CmsSalespersonDevice) TableName() string {
	return "cms_salesperson_device"
}
