package entities

import (
	"time"
)

type CmsPurchaseReturnDtl struct {
	Id                  uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	PrId                string    `xorm:"comment('this reference is unique and link to order table. it cannot use order id because the order are sending from different ipad, the order id which is generated from ipad might be the same when reach to CMS.') unique(unique_key) VARCHAR(20)"`
	DeviceItemId        int       `xorm:"default 0 unique(unique_key) INT"`
	ProductCode         string    `xorm:"unique(unique_key) VARCHAR(50)"`
	ProductName         string    `xorm:"VARCHAR(400)"`
	UserRemark          string    `xorm:"not null default '' VARCHAR(200)"`
	Quantity            float64   `xorm:"default 0 DOUBLE"`
	EdittedQuantity     int       `xorm:"default 0 INT"`
	UnitPrice           float64   `xorm:"default 0 DOUBLE"`
	Discount            string    `xorm:"VARCHAR(50)"`
	ProductUom          string    `xorm:"VARCHAR(100)"`
	AttributeRemark     string    `xorm:"comment('The chosen attribute name and value will be stored here, e.g. Size=L, Colour=Red') BLOB"`
	OptionalRemark      string    `xorm:"comment('The selected optional item will options here, e.g. Sport Rim, Leather Seat.') BLOB"`
	DiscountMethod      string    `xorm:"comment('Percentage or Fixed (Amount)') VARCHAR(100)"`
	PickerNote          string    `xorm:"not null VARCHAR(500)"`
	DiscountAmount      float64   `xorm:"default 0 DOUBLE"`
	SubTotal            float64   `xorm:"default 0 comment('the optional item price will affect the sub-total') DOUBLE"`
	Sequence            int       `xorm:"default 0 INT"`
	PackingStatus       string    `xorm:"ENUM('IN_PROGRESS','NO_STOCK','PACKED','PENDING')"`
	PackedQty           int       `xorm:"default 0 INT"`
	IsParent            int       `xorm:"default 0 INT"`
	ParentCode          string    `xorm:"VARCHAR(100)"`
	PackedBy            string    `xorm:"VARCHAR(200)"`
	PrItemValidity      string    `xorm:"ENUM('APPROVED','PENDING','REJECTED')"`
	PackConfirmedBy     string    `xorm:"VARCHAR(45)"`
	PackConfirmedStatus string    `xorm:"ENUM('CONFIRMED','PENDING')"`
	PackerNote          string    `xorm:"VARCHAR(200)"`
	CancelStatus        string    `xorm:"ENUM('ADMIN_CANCELLED','NOT_CANCELLED','SYSTEM_CANCELLED','USER_CANCELLED')"`
	UpdatedAt           time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (m *CmsPurchaseReturnDtl) TableName() string {
	return "cms_purchase_return_dtl"
}
