package entities

import (
	"time"
)

type CmsCustomerAgeingKo struct {
	Id            uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	DocDate       time.Time `xorm:"DATETIME" json:"docDate,omitempty" xml:"docDate"`
	DocCode       string    `xorm:"not null unique(unq) VARCHAR(100)" json:"docCode,omitempty" xml:"docCode"`
	DocKoRef      string    `xorm:"not null unique(unq) VARCHAR(100)" json:"docKoRef,omitempty" xml:"docKoRef"`
	SalespersonId int       `xorm:"comment('doc_ko_type agent id') INT" json:"salespersonId,omitempty" xml:"salespersonId"`
	DocKoType     string    `xorm:"not null unique(unq) VARCHAR(100)" json:"docKoType,omitempty" xml:"docKoType"`
	DocAmount     float64   `xorm:"DOUBLE" json:"docAmount,omitempty" xml:"docAmount"`
	ActiveStatus  int       `xorm:"not null default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	DocType       string    `xorm:"VARCHAR(100)" json:"docType,omitempty" xml:"docType"`
	UpdatedAt     time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsCustomerAgeingKo) TableName() string {
	return "cms_customer_ageing_ko"
}
