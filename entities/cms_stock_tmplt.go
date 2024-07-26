package entities

import (
	"time"
)

type CmsStockTmplt struct {
	Id           uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	TmpltName    string    `xorm:"VARCHAR(200)" json:"tmpltName,omitempty" xml:"tmpltName"`
	ActiveStatus int       `xorm:"default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	CreatedAt    time.Time `xorm:"DATETIME" json:"createdAt,omitempty" xml:"createdAt"`
	CreatedBy    string    `xorm:"default '' VARCHAR(50)" json:"createdBy,omitempty" xml:"createdBy"`
	UpdatedBy    string    `xorm:"default '' VARCHAR(50)" json:"updatedBy,omitempty" xml:"updatedBy"`
	UpdatedAt    time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsStockTmplt) TableName() string {
	return "cms_stock_tmplt"
}
