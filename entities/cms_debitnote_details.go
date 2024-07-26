package entities

import "time"

type CmsDebitnoteDetails struct {
	Id           uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	DnCode       string    `xorm:"VARCHAR(100)" json:"dnCode,omitempty" xml:"dnCode"`
	ItemCode     string    `xorm:"VARCHAR(200)" json:"itemCode,omitempty" xml:"itemCode"`
	ItemName     string    `xorm:"VARCHAR(200)" json:"itemName,omitempty" xml:"itemName"`
	ItemPrice    string    `xorm:"VARCHAR(200)" json:"itemPrice,omitempty" xml:"itemPrice"`
	Quantity     string    `xorm:"VARCHAR(200)" json:"quantity,omitempty" xml:"quantity"`
	Uom          string    `xorm:"VARCHAR(200)" json:"uom,omitempty" xml:"uom"`
	TotalPrice   string    `xorm:"VARCHAR(200)" json:"totalPrice,omitempty" xml:"totalPrice"`
	Discount     string    `xorm:"comment('0%+10+50%') VARCHAR(100)" json:"discount,omitempty" xml:"discount"`
	ActiveStatus int       `xorm:"default 1 comment('0=inactive, 1=active') INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	UpdatedAt    time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
	RefNo        string    `xorm:"unique VARCHAR(200)" json:"refNo,omitempty" xml:"refNo"`
}

func (m *CmsDebitnoteDetails) TableName() string {
	return "cms_debitnote_details"
}

func (m *CmsDebitnoteDetails) BeforeInsert() {
	m.BeforeUpdate()
}

func (m *CmsDebitnoteDetails) BeforeUpdate() {
	m.UpdatedAt = time.Now()
}
