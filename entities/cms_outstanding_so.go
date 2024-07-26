package entities

import (
	"time"
)

type CmsOutstandingSo struct {
	Id              uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	SoDocno         string    `xorm:"not null unique(so_docno) VARCHAR(20)" json:"soDocno,omitempty" xml:"soDocno"`
	SoDockey        string    `xorm:"not null VARCHAR(30)" json:"soDockey,omitempty" xml:"soDockey"`
	SoProductCode   string    `xorm:"not null default '' unique(so_docno) VARCHAR(50)" json:"soProductCode,omitempty" xml:"soProductCode"`
	SoUnitUom       string    `xorm:"unique(so_docno) VARCHAR(50)" json:"soUnitUom,omitempty" xml:"soUnitUom"`
	SoOriQty        float64   `xorm:"not null unique(so_docno) DOUBLE" json:"soOriQty,omitempty" xml:"soOriQty"`
	SoOutQty        float64   `xorm:"not null DOUBLE" json:"soOutQty,omitempty" xml:"soOutQty"`
	SoTransQty      float64   `xorm:"not null DOUBLE" json:"soTransQty,omitempty" xml:"soTransQty"`
	SoDocDate       time.Time `xorm:"not null DATETIME" json:"soDocDate,omitempty" xml:"soDocDate"`
	SoSalespersonId int       `xorm:"not null INT" json:"soSalespersonId,omitempty" xml:"soSalespersonId"`
	SoCustCode      string    `xorm:"not null VARCHAR(50)" json:"soCustCode,omitempty" xml:"soCustCode"`
	SoBranchCode    string    `xorm:"VARCHAR(100)" json:"soBranchCode,omitempty" xml:"soBranchCode"`
	UpdatedAt       time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
	ActiveStatus    int       `xorm:"default 0 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	RefNo           string    `xorm:"VARCHAR(100)" json:"refNo,omitempty" xml:"refNo"`
}

func (m *CmsOutstandingSo) TableName() string {
	return "cms_outstanding_so"
}
