package entities

import (
	"time"
)

type CmsAppAdvertisement struct {
	Id           uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	AdverCode    string    `xorm:"default '' unique VARCHAR(50)"`
	AdverName    string    `xorm:"not null VARCHAR(50)"`
	AdverLink    string    `xorm:"not null VARCHAR(255)"`
	AdverType    string    `xorm:"default 'IMAGE' ENUM('IMAGE','VIDEO')"`
	AdverContent []byte    `xorm:"BLOB"`
	DtlCode      string    `xorm:"default '0' VARCHAR(50)"`
	DtlType      string    `xorm:"default '0' VARCHAR(50)"`
	SequenceNo   int       `xorm:"default 0 INT"`
	ActiveStatus int       `xorm:"not null default 1 INT"`
	UpdatedAt    time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
}

func (m *CmsAppAdvertisement) TableName() string {
	return "cms_app_advertisement"
}
