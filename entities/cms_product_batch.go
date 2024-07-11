package entities

import (
	"time"
)

type CmsProductBatch struct {
	Id           uint64    `xorm:"not null pk autoincr UNSIGNED BIGINT"`
	ProductCode  string    `xorm:"not null index(product_code_batch_code_wh_code) unique(unq_key) VARCHAR(50)"`
	Quantity     float64   `xorm:"not null DOUBLE"`
	WhCode       string    `xorm:"not null index(product_code_batch_code_wh_code) unique(unq_key) VARCHAR(50)"`
	BatchCode    string    `xorm:"not null index(product_code_batch_code_wh_code) unique(unq_key) VARCHAR(100)"`
	BatchDesc    string    `xorm:"not null VARCHAR(100)"`
	ExpDate      time.Time `xorm:"DATETIME"`
	MfgDate      time.Time `xorm:"DATETIME"`
	Remark1      string    `xorm:"not null VARCHAR(100)"`
	Remark2      string    `xorm:"not null VARCHAR(100)"`
	ActiveStatus int       `xorm:"not null default 1 INT"`
	UpdatedAt    time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
}

func (m *CmsProductBatch) TableName() string {
	return "cms_product_batch"
}

func (m *CmsProductBatch) Validate() {
	if m.ExpDate.IsZero() {
		m.ExpDate = time.Now().AddDate(10, 1, 1)
	}
}

func (m *CmsProductBatch) ToUpdate() {
	m.UpdatedAt = time.Now()
}
