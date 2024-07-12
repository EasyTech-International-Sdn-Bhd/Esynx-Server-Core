package entities

import (
	"time"
)

type CmsCustomerAgeingKo struct {
	Id            uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	DocDate       time.Time `xorm:"DATETIME"`
	DocCode       string    `xorm:"not null unique(unq) VARCHAR(100)"`
	DocKoRef      string    `xorm:"not null unique(unq) VARCHAR(100)"`
	SalespersonId int       `xorm:"comment('doc_ko_type agent id') INT"`
	DocKoType     string    `xorm:"not null unique(unq) VARCHAR(100)"`
	DocAmount     float64   `xorm:"DOUBLE"`
	ActiveStatus  int       `xorm:"not null default 1 INT"`
	DocType       string    `xorm:"VARCHAR(100)"`
	UpdatedAt     time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
}

func (m *CmsCustomerAgeingKo) TableName() string {
	return "cms_customer_ageing_ko"
}
