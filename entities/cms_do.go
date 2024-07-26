package entities

import (
	"time"
)

type CmsDo struct {
	DoId              uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"doId,omitempty" xml:"doId"`
	DoCode            string    `xorm:"not null default '' unique VARCHAR(50)" json:"doCode,omitempty" xml:"doCode"`
	CustCode          string    `xorm:"not null VARCHAR(20)" json:"custCode,omitempty" xml:"custCode"`
	DoDate            time.Time `xorm:"not null DATETIME" json:"doDate,omitempty" xml:"doDate"`
	DoAmount          float64   `xorm:"default 0 DOUBLE" json:"doAmount,omitempty" xml:"doAmount"`
	Cancelled         string    `xorm:"not null default 'F' VARCHAR(2)" json:"cancelled,omitempty" xml:"cancelled"`
	Salesperson       string    `xorm:"VARCHAR(200)" json:"salesperson,omitempty" xml:"salesperson"`
	SalespersonId     int       `xorm:"not null INT" json:"salespersonId,omitempty" xml:"salespersonId"`
	Remarks           string    `xorm:"VARCHAR(200)" json:"remarks,omitempty" xml:"remarks"`
	DoReference       string    `xorm:"VARCHAR(200)" json:"doReference,omitempty" xml:"doReference"`
	DeliveryLocation  string    `xorm:"VARCHAR(200)" json:"deliveryLocation,omitempty" xml:"deliveryLocation"`
	SelfCollect       int       `xorm:"not null default 0 INT" json:"selfCollect,omitempty" xml:"selfCollect"`
	TransferStatus    int       `xorm:"default 0 comment('0 - havent converted; 1 - converted to INV already') INT" json:"transferStatus,omitempty" xml:"transferStatus"`
	PackingStatus     int       `xorm:"default 0 INT" json:"packingStatus,omitempty" xml:"packingStatus"`
	PackedBy          string    `xorm:"VARCHAR(200)" json:"packedBy,omitempty" xml:"packedBy"`
	PickerNote        string    `xorm:"VARCHAR(200)" json:"pickerNote,omitempty" xml:"pickerNote"`
	PackConfirmed     int       `xorm:"default 0 INT" json:"packConfirmed,omitempty" xml:"packConfirmed"`
	PackConfirmedBy   string    `xorm:"VARCHAR(20)" json:"packConfirmedBy,omitempty" xml:"packConfirmedBy"`
	BasketId          string    `xorm:"VARCHAR(100)" json:"basketId,omitempty" xml:"basketId"`
	OrderFault        int       `xorm:"default 0 INT" json:"orderFault,omitempty" xml:"orderFault"`
	OrderFaultMessage string    `xorm:"VARCHAR(200)" json:"orderFaultMessage,omitempty" xml:"orderFaultMessage"`
	RefNo             string    `xorm:"comment('dockey') VARCHAR(200)" json:"refNo,omitempty" xml:"refNo"`
	BranchName        string    `xorm:"VARCHAR(200)" json:"branchName,omitempty" xml:"branchName"`
	UpdatedAt         time.Time `xorm:"not null default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
	ScRemark          string    `xorm:"VARCHAR(200)" json:"scRemark,omitempty" xml:"scRemark"`
	ActiveStatus      int       `xorm:"default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	DoUdf             string    `xorm:"JSON" json:"doUdf,omitempty" xml:"doUdf"`
	Approved          int       `xorm:"default 0 INT" json:"approved,omitempty" xml:"approved"`
	Approver          string    `xorm:"VARCHAR(100)" json:"approver,omitempty" xml:"approver"`
	ApprovedAt        time.Time `xorm:"DATETIME" json:"approvedAt,omitempty" xml:"approvedAt"`
}

func (m *CmsDo) TableName() string {
	return "cms_do"
}

func (m *CmsDo) BeforeInsert() {
	m.BeforeUpdate()
}

func (m *CmsDo) BeforeUpdate() {
	m.UpdatedAt = time.Now()
}
