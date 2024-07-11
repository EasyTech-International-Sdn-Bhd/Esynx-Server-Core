package entities

import (
	"time"
)

type CmsDoDetails struct {
	Id                  int       `xorm:"not null pk autoincr INT"`
	DoCode              string    `xorm:"not null default '' unique(unq_key) VARCHAR(50)"`
	ItemCode            string    `xorm:"not null VARCHAR(20)"`
	ItemName            string    `xorm:"not null VARCHAR(200)"`
	ItemPrice           float64   `xorm:"default 0 DOUBLE"`
	Quantity            float64   `xorm:"default 0 DOUBLE"`
	TotalPrice          float64   `xorm:"default 0 DOUBLE"`
	Uom                 string    `xorm:"not null VARCHAR(20)"`
	Discount            string    `xorm:"default '0' VARCHAR(50)"`
	PackingStatus       int       `xorm:"default 0 INT"`
	QrCode              string    `xorm:"VARCHAR(200)"`
	PackedQty           int       `xorm:"default 0 INT"`
	CheckedQty          int       `xorm:"default 0 INT"`
	PackConfirmedStatus int       `xorm:"default 0 INT"`
	PackConfirmedBy     string    `xorm:"default '' VARCHAR(20)"`
	PackedBy            string    `xorm:"VARCHAR(200)"`
	PickerNote          string    `xorm:"VARCHAR(200)"`
	Location            string    `xorm:"VARCHAR(200)"`
	ActiveStatus        int       `xorm:"default 1 INT"`
	RefNo               string    `xorm:"comment('dtlkey') unique(unq_key) VARCHAR(50)"`
	UpdatedAt           time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	SequenceNo          int       `xorm:"default 0 INT"`
	DoDtlUdf            string    `xorm:"JSON"`
}

func (m *CmsDoDetails) TableName() string {
	return "cms_do_details"
}
