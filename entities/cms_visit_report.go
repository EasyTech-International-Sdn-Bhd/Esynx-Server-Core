package entities

import (
	"time"
)

type CmsVisitReport struct {
	Id                    uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	CustomerId            int       `xorm:"not null unique(unq) INT" json:"customerId,omitempty" xml:"customerId"`
	BranchCode            string    `xorm:"VARCHAR(20)" json:"branchCode,omitempty" xml:"branchCode"`
	SalespersonId         int       `xorm:"not null INT" json:"salespersonId,omitempty" xml:"salespersonId"`
	PersonMet             string    `xorm:"not null VARCHAR(255)" json:"personMet,omitempty" xml:"personMet"`
	MobileCheckinId       string    `xorm:"not null unique(unq) VARCHAR(255)" json:"mobileCheckinId,omitempty" xml:"mobileCheckinId"`
	CheckinTime           time.Time `xorm:"not null default '1971-01-01 23:01:01' DATETIME" json:"checkinTime,omitempty" xml:"checkinTime"`
	CheckoutTime          time.Time `xorm:"default '1971-01-01 23:01:01' DATETIME" json:"checkoutTime,omitempty" xml:"checkoutTime"`
	CheckinTimeLastUpdate time.Time `xorm:"not null default '1971-01-01 23:01:01' DATETIME" json:"checkinTimeLastUpdate,omitempty" xml:"checkinTimeLastUpdate"`
	LocationLat           string    `xorm:"not null VARCHAR(255)" json:"locationLat,omitempty" xml:"locationLat"`
	LocationLng           string    `xorm:"not null VARCHAR(255)" json:"locationLng,omitempty" xml:"locationLng"`
	CheckoutLat           string    `xorm:"VARCHAR(255)" json:"checkoutLat,omitempty" xml:"checkoutLat"`
	CheckoutLng           string    `xorm:"VARCHAR(255)" json:"checkoutLng,omitempty" xml:"checkoutLng"`
	CheckinLocation       string    `xorm:"VARCHAR(300)" json:"checkinLocation,omitempty" xml:"checkinLocation"`
	CheckoutLocation      string    `xorm:"VARCHAR(300)" json:"checkoutLocation,omitempty" xml:"checkoutLocation"`
	Status                int       `xorm:"not null comment('0 delete 1 active 2 transferred') INT" json:"status,omitempty" xml:"status"`
	CreateBy              string    `xorm:"not null VARCHAR(255)" json:"createBy,omitempty" xml:"createBy"`
	CreateDate            time.Time `xorm:"not null default '1971-01-01 23:01:01' DATETIME" json:"createDate,omitempty" xml:"createDate"`
	Remark1               []byte    `xorm:"BLOB" json:"remark1,omitempty" xml:"remark1"`
	Remark2               []byte    `xorm:"BLOB" json:"remark2,omitempty" xml:"remark2"`
	Remark3               []byte    `xorm:"BLOB" json:"remark3,omitempty" xml:"remark3"`
	UpdateBy              string    `xorm:"not null VARCHAR(255)" json:"updateBy,omitempty" xml:"updateBy"`
	UpdateDate            time.Time `xorm:"not null default '1971-01-01 23:01:01' DATETIME" json:"updateDate,omitempty" xml:"updateDate"`
	Schedule              string    `xorm:"JSON" json:"schedule,omitempty" xml:"schedule"`
	ServiceRating         string    `xorm:"JSON" json:"serviceRating,omitempty" xml:"serviceRating"`
}

func (m *CmsVisitReport) TableName() string {
	return "cms_visit_report"
}
