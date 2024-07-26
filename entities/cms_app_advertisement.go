package entities

import (
	"time"
)

type CmsAppAdvertisement struct {
	Id           uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	AdverCode    string    `xorm:"default '' unique VARCHAR(50)" json:"adverCode,omitempty" xml:"adverCode"`
	AdverName    string    `xorm:"not null VARCHAR(50)" json:"adverName,omitempty" xml:"adverName"`
	AdverLink    string    `xorm:"not null VARCHAR(255)" json:"adverLink,omitempty" xml:"adverLink"`
	AdverType    string    `xorm:"default 'IMAGE' ENUM('IMAGE','VIDEO')" json:"adverType,omitempty" xml:"adverType"`
	AdverContent []byte    `xorm:"BLOB" json:"adverContent,omitempty" xml:"adverContent"`
	DtlCode      string    `xorm:"default '0' VARCHAR(50)" json:"dtlCode,omitempty" xml:"dtlCode"`
	DtlType      string    `xorm:"default '0' VARCHAR(50)" json:"dtlType,omitempty" xml:"dtlType"`
	SequenceNo   int       `xorm:"default 0 INT" json:"sequenceNo,omitempty" xml:"sequenceNo"`
	ActiveStatus int       `xorm:"not null default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	UpdatedAt    time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsAppAdvertisement) TableName() string {
	return "cms_app_advertisement"
}
