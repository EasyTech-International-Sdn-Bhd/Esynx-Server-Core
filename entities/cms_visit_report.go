package entities

import (
	"time"
)

type CmsVisitReport struct {
	Id                    uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	CustomerId            int       `xorm:"not null unique(unq) INT"`
	BranchCode            string    `xorm:"VARCHAR(20)"`
	SalespersonId         int       `xorm:"not null INT"`
	PersonMet             string    `xorm:"not null VARCHAR(255)"`
	MobileCheckinId       string    `xorm:"not null unique(unq) VARCHAR(255)"`
	CheckinTime           time.Time `xorm:"not null default '1971-01-01 23:01:01' DATETIME"`
	CheckoutTime          time.Time `xorm:"default '1971-01-01 23:01:01' DATETIME"`
	CheckinTimeLastUpdate time.Time `xorm:"not null default '1971-01-01 23:01:01' DATETIME"`
	LocationLat           string    `xorm:"not null VARCHAR(255)"`
	LocationLng           string    `xorm:"not null VARCHAR(255)"`
	CheckoutLat           string    `xorm:"VARCHAR(255)"`
	CheckoutLng           string    `xorm:"VARCHAR(255)"`
	CheckinLocation       string    `xorm:"VARCHAR(300)"`
	CheckoutLocation      string    `xorm:"VARCHAR(300)"`
	Status                int       `xorm:"not null comment('0 delete 1 active 2 transferred') INT"`
	CreateBy              string    `xorm:"not null VARCHAR(255)"`
	CreateDate            time.Time `xorm:"not null default '1971-01-01 23:01:01' DATETIME"`
	Remark1               []byte    `xorm:"BLOB"`
	Remark2               []byte    `xorm:"BLOB"`
	Remark3               []byte    `xorm:"BLOB"`
	UpdateBy              string    `xorm:"not null VARCHAR(255)"`
	UpdateDate            time.Time `xorm:"not null default '1971-01-01 23:01:01' DATETIME"`
	Schedule              string    `xorm:"JSON"`
	ServiceRating         string    `xorm:"JSON"`
}

func (m *CmsVisitReport) TableName() string {
	return "cms_visit_report"
}
