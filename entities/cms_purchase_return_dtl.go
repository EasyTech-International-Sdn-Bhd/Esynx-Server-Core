package entities

import (
	"time"
)

type CmsPurchaseReturnDtl struct {
	Id                  uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	PrId                string    `xorm:"comment('this reference is unique and link to order table. it cannot use order id because the order are sending from different ipad, the order id which is generated from ipad might be the same when reach to CMS.') unique(unique_key) VARCHAR(20)" json:"prId,omitempty" xml:"prId"`
	DeviceItemId        int       `xorm:"default 0 unique(unique_key) INT" json:"deviceItemId,omitempty" xml:"deviceItemId"`
	ProductCode         string    `xorm:"unique(unique_key) VARCHAR(50)" json:"productCode,omitempty" xml:"productCode"`
	ProductName         string    `xorm:"VARCHAR(400)" json:"productName,omitempty" xml:"productName"`
	UserRemark          string    `xorm:"not null default '' VARCHAR(200)" json:"userRemark,omitempty" xml:"userRemark"`
	Quantity            float64   `xorm:"default 0 DOUBLE" json:"quantity,omitempty" xml:"quantity"`
	EdittedQuantity     int       `xorm:"default 0 INT" json:"edittedQuantity,omitempty" xml:"edittedQuantity"`
	UnitPrice           float64   `xorm:"default 0 DOUBLE" json:"unitPrice,omitempty" xml:"unitPrice"`
	Discount            string    `xorm:"VARCHAR(50)" json:"discount,omitempty" xml:"discount"`
	ProductUom          string    `xorm:"VARCHAR(100)" json:"productUom,omitempty" xml:"productUom"`
	AttributeRemark     string    `xorm:"comment('The chosen attribute name and value will be stored here, e.g. Size=L, Colour=Red') BLOB" json:"attributeRemark,omitempty" xml:"attributeRemark"`
	OptionalRemark      string    `xorm:"comment('The selected optional item will options here, e.g. Sport Rim, Leather Seat.') BLOB" json:"optionalRemark,omitempty" xml:"optionalRemark"`
	DiscountMethod      string    `xorm:"comment('Percentage or Fixed (Amount)') VARCHAR(100)" json:"discountMethod,omitempty" xml:"discountMethod"`
	PickerNote          string    `xorm:"not null VARCHAR(500)" json:"pickerNote,omitempty" xml:"pickerNote"`
	DiscountAmount      float64   `xorm:"default 0 DOUBLE" json:"discountAmount,omitempty" xml:"discountAmount"`
	SubTotal            float64   `xorm:"default 0 comment('the optional item price will affect the sub-total') DOUBLE" json:"subTotal,omitempty" xml:"subTotal"`
	Sequence            int       `xorm:"default 0 INT" json:"sequence,omitempty" xml:"sequence"`
	PackingStatus       string    `xorm:"ENUM('IN_PROGRESS','NO_STOCK','PACKED','PENDING')" json:"packingStatus,omitempty" xml:"packingStatus"`
	PackedQty           int       `xorm:"default 0 INT" json:"packedQty,omitempty" xml:"packedQty"`
	IsParent            int       `xorm:"default 0 INT" json:"isParent,omitempty" xml:"isParent"`
	ParentCode          string    `xorm:"VARCHAR(100)" json:"parentCode,omitempty" xml:"parentCode"`
	PackedBy            string    `xorm:"VARCHAR(200)" json:"packedBy,omitempty" xml:"packedBy"`
	PrItemValidity      string    `xorm:"ENUM('APPROVED','PENDING','REJECTED')" json:"prItemValidity,omitempty" xml:"prItemValidity"`
	PackConfirmedBy     string    `xorm:"VARCHAR(45)" json:"packConfirmedBy,omitempty" xml:"packConfirmedBy"`
	PackConfirmedStatus string    `xorm:"ENUM('CONFIRMED','PENDING')" json:"packConfirmedStatus,omitempty" xml:"packConfirmedStatus"`
	PackerNote          string    `xorm:"VARCHAR(200)" json:"packerNote,omitempty" xml:"packerNote"`
	CancelStatus        string    `xorm:"ENUM('ADMIN_CANCELLED','NOT_CANCELLED','SYSTEM_CANCELLED','USER_CANCELLED')" json:"cancelStatus,omitempty" xml:"cancelStatus"`
	UpdatedAt           time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsPurchaseReturnDtl) TableName() string {
	return "cms_purchase_return_dtl"
}
