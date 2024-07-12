package entities

import (
	"time"
)

type CmsOutstandingSo struct {
	Id              uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	SoDocno         string    `xorm:"not null unique(so_docno) VARCHAR(20)"`
	SoDockey        string    `xorm:"not null VARCHAR(30)"`
	SoProductCode   string    `xorm:"not null default '' unique(so_docno) VARCHAR(50)"`
	SoUnitUom       string    `xorm:"unique(so_docno) VARCHAR(50)"`
	SoOriQty        float64   `xorm:"not null unique(so_docno) DOUBLE"`
	SoOutQty        float64   `xorm:"not null DOUBLE"`
	SoTransQty      float64   `xorm:"not null DOUBLE"`
	SoDocDate       time.Time `xorm:"not null DATETIME"`
	SoSalespersonId int       `xorm:"not null INT"`
	SoCustCode      string    `xorm:"not null VARCHAR(50)"`
	SoBranchCode    string    `xorm:"VARCHAR(100)"`
	UpdatedAt       time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
	ActiveStatus    int       `xorm:"default 0 INT"`
	RefNo           string    `xorm:"VARCHAR(100)"`
}

func (m *CmsOutstandingSo) TableName() string {
	return "cms_outstanding_so"
}
