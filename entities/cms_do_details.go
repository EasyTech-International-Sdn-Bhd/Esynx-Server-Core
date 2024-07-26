package entities

import (
	"time"
)

type CmsDoDetails struct {
	Id                  uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	DoCode              string    `xorm:"not null default '' unique(unq_key) VARCHAR(50)" json:"doCode,omitempty" xml:"doCode"`
	ItemCode            string    `xorm:"not null VARCHAR(20)" json:"itemCode,omitempty" xml:"itemCode"`
	ItemName            string    `xorm:"not null VARCHAR(200)" json:"itemName,omitempty" xml:"itemName"`
	ItemPrice           float64   `xorm:"default 0 DOUBLE" json:"itemPrice,omitempty" xml:"itemPrice"`
	Quantity            float64   `xorm:"default 0 DOUBLE" json:"quantity,omitempty" xml:"quantity"`
	TotalPrice          float64   `xorm:"default 0 DOUBLE" json:"totalPrice,omitempty" xml:"totalPrice"`
	Uom                 string    `xorm:"not null VARCHAR(20)" json:"uom,omitempty" xml:"uom"`
	Discount            string    `xorm:"default '0' VARCHAR(50)" json:"discount,omitempty" xml:"discount"`
	PackingStatus       int       `xorm:"default 0 INT" json:"packingStatus,omitempty" xml:"packingStatus"`
	QrCode              string    `xorm:"VARCHAR(200)" json:"qrCode,omitempty" xml:"qrCode"`
	PackedQty           int       `xorm:"default 0 INT" json:"packedQty,omitempty" xml:"packedQty"`
	CheckedQty          int       `xorm:"default 0 INT" json:"checkedQty,omitempty" xml:"checkedQty"`
	PackConfirmedStatus int       `xorm:"default 0 INT" json:"packConfirmedStatus,omitempty" xml:"packConfirmedStatus"`
	PackConfirmedBy     string    `xorm:"default '' VARCHAR(20)" json:"packConfirmedBy,omitempty" xml:"packConfirmedBy"`
	PackedBy            string    `xorm:"VARCHAR(200)" json:"packedBy,omitempty" xml:"packedBy"`
	PickerNote          string    `xorm:"VARCHAR(200)" json:"pickerNote,omitempty" xml:"pickerNote"`
	Location            string    `xorm:"VARCHAR(200)" json:"location,omitempty" xml:"location"`
	ActiveStatus        int       `xorm:"default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	RefNo               string    `xorm:"comment('dtlkey') unique(unq_key) VARCHAR(50)" json:"refNo,omitempty" xml:"refNo"`
	UpdatedAt           time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
	SequenceNo          int       `xorm:"default 0 INT" json:"sequenceNo,omitempty" xml:"sequenceNo"`
	DoDtlUdf            string    `xorm:"JSON" json:"doDtlUdf,omitempty" xml:"doDtlUdf"`
}

func (m *CmsDoDetails) TableName() string {
	return "cms_do_details"
}

func (m *CmsDoDetails) BeforeInsert() {
	m.BeforeUpdate()
}

func (m *CmsDoDetails) BeforeUpdate() {
	m.UpdatedAt = time.Now()
}
