package entities

import (
	"time"
)

type CmsDo struct {
	DoId              int       `xorm:"not null pk autoincr INT"`
	DoCode            string    `xorm:"not null default '' unique VARCHAR(50)"`
	CustCode          string    `xorm:"not null VARCHAR(20)"`
	DoDate            time.Time `xorm:"not null DATETIME"`
	DoAmount          float64   `xorm:"default 0 DOUBLE"`
	Cancelled         string    `xorm:"not null default 'F' VARCHAR(2)"`
	Salesperson       string    `xorm:"VARCHAR(200)"`
	SalespersonId     int       `xorm:"not null INT"`
	Remarks           string    `xorm:"VARCHAR(200)"`
	DoReference       string    `xorm:"VARCHAR(200)"`
	DeliveryLocation  string    `xorm:"VARCHAR(200)"`
	SelfCollect       int       `xorm:"not null default 0 INT"`
	TransferStatus    int       `xorm:"default 0 comment('0 - havent converted; 1 - converted to INV already') INT"`
	PackingStatus     int       `xorm:"default 0 INT"`
	PackedBy          string    `xorm:"VARCHAR(200)"`
	PickerNote        string    `xorm:"VARCHAR(200)"`
	PackConfirmed     int       `xorm:"default 0 INT"`
	PackConfirmedBy   string    `xorm:"VARCHAR(20)"`
	BasketId          string    `xorm:"VARCHAR(100)"`
	OrderFault        int       `xorm:"default 0 INT"`
	OrderFaultMessage string    `xorm:"VARCHAR(200)"`
	RefNo             string    `xorm:"comment('dockey') VARCHAR(200)"`
	BranchName        string    `xorm:"VARCHAR(200)"`
	UpdatedAt         time.Time `xorm:"not null default CURRENT_TIMESTAMP DATETIME"`
	ScRemark          string    `xorm:"VARCHAR(200)"`
	ActiveStatus      int       `xorm:"default 1 INT"`
	DoUdf             string    `xorm:"JSON"`
	Approved          int       `xorm:"default 0 INT"`
	Approver          string    `xorm:"VARCHAR(100)"`
	ApprovedAt        time.Time `xorm:"DATETIME"`
}

func (m *CmsDo) TableName() string {
	return "cms_do"
}
