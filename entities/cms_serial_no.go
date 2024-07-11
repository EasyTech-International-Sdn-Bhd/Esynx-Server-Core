package entities

import (
	"time"
)

type CmsSerialNo struct {
	Id           int       `xorm:"not null autoincr index INT"`
	ProductCode  string    `xorm:"VARCHAR(200)"`
	WhCode       string    `xorm:"VARCHAR(200)"`
	SerialNo     string    `xorm:"VARCHAR(200)"`
	BatchNo      string    `xorm:"VARCHAR(200)"`
	Qty          int       `xorm:"default 0 INT"`
	ActiveStatus int       `xorm:"default 1 INT"`
	UpdatedAt    time.Time `xorm:"not null default CURRENT_TIMESTAMP DATETIME"`
}

func (m *CmsSerialNo) TableName() string {
	return "cms_serial_no"
}
