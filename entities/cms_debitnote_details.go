package entities

import "time"

type CmsDebitnoteDetails struct {
	Id           uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	DnCode       string    `xorm:"VARCHAR(100)"`
	ItemCode     string    `xorm:"VARCHAR(200)"`
	ItemName     string    `xorm:"VARCHAR(200)"`
	ItemPrice    string    `xorm:"VARCHAR(200)"`
	Quantity     string    `xorm:"VARCHAR(200)"`
	Uom          string    `xorm:"VARCHAR(200)"`
	TotalPrice   string    `xorm:"VARCHAR(200)"`
	Discount     string    `xorm:"comment('0%+10+50%') VARCHAR(100)"`
	ActiveStatus int       `xorm:"default 1 comment('0=inactive, 1=active') INT"`
	UpdatedAt    time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	RefNo        string    `xorm:"unique VARCHAR(200)"`
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
