package entities

import (
	"time"
)

type CmsStockAdjustment struct {
	Id                   uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	AdjId                string    `xorm:"not null unique VARCHAR(25)" json:"adjId,omitempty" xml:"adjId"`
	AdjDate              time.Time `xorm:"DATE" json:"adjDate,omitempty" xml:"adjDate"`
	AdjDescription       string    `xorm:"VARCHAR(255)" json:"adjDescription,omitempty" xml:"adjDescription"`
	AdjNote              []byte    `xorm:"BLOB" json:"adjNote,omitempty" xml:"adjNote"`
	ActiveStatus         string    `xorm:"ENUM('ACTIVE','CANCELLED')" json:"activeStatus,omitempty" xml:"activeStatus"`
	CancelStatus         string    `xorm:"ENUM('ADMIN_CANCELLED','NOT_CANCELLED','SYSTEM_CANCELLED','USER_CANCELLED')" json:"cancelStatus,omitempty" xml:"cancelStatus"`
	TransferReceivedDate time.Time `xorm:"DATETIME" json:"transferReceivedDate,omitempty" xml:"transferReceivedDate"`
	LastUpdatedBy        int       `xorm:"default 0 INT" json:"lastUpdatedBy,omitempty" xml:"lastUpdatedBy"`
	Agent                string    `xorm:"VARCHAR(50)" json:"agent,omitempty" xml:"agent"`
	UpdatedAt            time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsStockAdjustment) TableName() string {
	return "cms_stock_adjustment"
}
