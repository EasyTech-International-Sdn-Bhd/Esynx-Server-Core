package entities

import (
	"time"
)

type CmsProductBatch struct {
	Id           uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	ProductCode  string    `xorm:"not null index(product_code_batch_code_wh_code) unique(unq_key) VARCHAR(50)" json:"productCode,omitempty" xml:"productCode"`
	Quantity     float64   `xorm:"not null DOUBLE" json:"quantity,omitempty" xml:"quantity"`
	WhCode       string    `xorm:"not null index(product_code_batch_code_wh_code) unique(unq_key) VARCHAR(50)" json:"whCode,omitempty" xml:"whCode"`
	BatchCode    string    `xorm:"not null index(product_code_batch_code_wh_code) unique(unq_key) VARCHAR(100)" json:"batchCode,omitempty" xml:"batchCode"`
	BatchDesc    string    `xorm:"not null VARCHAR(100)" json:"batchDesc,omitempty" xml:"batchDesc"`
	ExpDate      time.Time `xorm:"DATETIME" json:"expDate,omitempty" xml:"expDate"`
	MfgDate      time.Time `xorm:"DATETIME" json:"mfgDate,omitempty" xml:"mfgDate"`
	Remark1      string    `xorm:"not null VARCHAR(100)" json:"remark1,omitempty" xml:"remark1"`
	Remark2      string    `xorm:"not null VARCHAR(100)" json:"remark2,omitempty" xml:"remark2"`
	ActiveStatus int       `xorm:"not null default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	UpdatedAt    time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsProductBatch) TableName() string {
	return "cms_product_batch"
}

func (m *CmsProductBatch) BeforeInsert() {
	m.BeforeUpdate()
}

func (m *CmsProductBatch) BeforeUpdate() {
	m.UpdatedAt = time.Now()
	if m.ExpDate.IsZero() {
		m.ExpDate = time.Now().AddDate(10, 1, 1)
	}
}
