package entities

import (
	"time"
)

type CmsSerialNo struct {
	Id           int       `xorm:"not null pk autoincr UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	ProductCode  string    `xorm:"VARCHAR(200)" json:"productCode,omitempty" xml:"productCode"`
	WhCode       string    `xorm:"VARCHAR(200)" json:"whCode,omitempty" xml:"whCode"`
	SerialNo     string    `xorm:"VARCHAR(200)" json:"serialNo,omitempty" xml:"serialNo"`
	BatchNo      string    `xorm:"VARCHAR(200)" json:"batchNo,omitempty" xml:"batchNo"`
	Qty          int       `xorm:"default 0 INT" json:"qty,omitempty" xml:"qty"`
	ActiveStatus int       `xorm:"default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	UpdatedAt    time.Time `xorm:"not null default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsSerialNo) TableName() string {
	return "cms_serial_no"
}
