package entities

import (
	"time"
)

type CmsSalespersonDevice struct {
	Id           int64     `xorm:"not null pk autoincr UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	Prefix       string    `xorm:"not null pk default 'D' CHAR(30)" json:"prefix,omitempty" xml:"prefix"`
	DeviceNo     int       `xorm:"default 1 INT" json:"deviceNo,omitempty" xml:"deviceNo"`
	DeviceToken  string    `xorm:"not null unique(unq) VARCHAR(200)" json:"deviceToken,omitempty" xml:"deviceToken"`
	LoginId      string    `xorm:"not null unique(unq) VARCHAR(100)" json:"loginId,omitempty" xml:"loginId"`
	AppVersion   string    `xorm:"VARCHAR(100)" json:"appVersion,omitempty" xml:"appVersion"`
	Os           string    `xorm:"not null VARCHAR(100)" json:"os,omitempty" xml:"os"`
	RegisteredAt time.Time `xorm:"not null default '1971-01-01 23:01:01' DATETIME" json:"registeredAt,omitempty" xml:"registeredAt"`
	ActiveStatus int       `xorm:"not null default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	LastLoggedIn time.Time `xorm:"not null default CURRENT_TIMESTAMP DATETIME" json:"lastLoggedIn,omitempty" xml:"lastLoggedIn"`
}

func (m *CmsSalespersonDevice) TableName() string {
	return "cms_salesperson_device"
}
