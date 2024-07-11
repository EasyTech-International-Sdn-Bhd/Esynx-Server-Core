package entities

import (
	"time"
)

type CmsOrderItem struct {
	OrderItemId         uint64    `xorm:"not null pk autoincr unique UNSIGNED BIGINT"`
	OrderId             string    `xorm:"comment('this reference is unique and link to order table. it cannot use order id because the order are sending from different ipad, the order id which is generated from ipad might be the same when reach to CMS.') unique(unique_key) VARCHAR(20)"`
	IpadItemId          int64     `xorm:"default 0 unique(unique_key) BIGINT"`
	ProductCode         string    `xorm:"VARCHAR(50)"`
	ProductId           int       `xorm:"unique(unique_key) INT"`
	ProductName         string    `xorm:"VARCHAR(400)"`
	SalespersonRemark   []byte    `xorm:"not null BLOB"`
	Quantity            float64   `xorm:"DOUBLE"`
	EdittedQuantity     int       `xorm:"default 0 INT"`
	UnitPrice           float64   `xorm:"DOUBLE"`
	Disc1               float64   `xorm:"default 0 DOUBLE"`
	Disc2               float64   `xorm:"default 0 DOUBLE"`
	Disc3               float64   `xorm:"default 0 DOUBLE"`
	UnitUom             string    `xorm:"VARCHAR(100)"`
	AttributeRemark     string    `xorm:"comment('The chosen attribute name and value will be stored here, e.g. Size=L, Colour=Red') LONGTEXT(4294967295)"`
	OptionalRemark      string    `xorm:"comment('The selected optional item will store here, e.g. Sport Rim, Leather Seat.') LONGTEXT(4294967295)"`
	DiscountMethod      string    `xorm:"comment('Percentage or Fixed (Amount)') VARCHAR(100)"`
	PickerNote          string    `xorm:"not null VARCHAR(500)"`
	DiscountAmount      string    `xorm:"default '' VARCHAR(50)"`
	SubTotal            float64   `xorm:"comment('the optional item price will affect the sub-total') DOUBLE"`
	SequenceNo          int       `xorm:"INT"`
	UomId               int       `xorm:"not null INT"`
	PackingStatus       int       `xorm:"default 0 comment('0=not packed, 1=packed, 2=no stock, 3=no stock but informed') INT"`
	PackedQty           int       `xorm:"default 0 INT"`
	PackConfirmedBy     string    `xorm:"VARCHAR(45)"`
	PackConfirmedStatus int       `xorm:"default 0 INT"`
	Isparent            int       `xorm:"INT"`
	ParentCode          string    `xorm:"VARCHAR(100)"`
	PackedBy            string    `xorm:"VARCHAR(200)"`
	OrderItemValidity   int       `xorm:"default 2 comment('0 = reject 1=pending 2 = approved') INT"`
	CancelStatus        int       `xorm:"default 0 comment('0=not canceled, 1=canceld') INT"`
	UpdatedAt           time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	WarehouseCode       string    `xorm:"VARCHAR(100)"`
	ProjNo              []byte    `xorm:"BLOB"`
	IsExchange          int       `xorm:"not null default 0 INT"`
	UnitLength          []byte    `xorm:"BLOB"`
	IsReturn            int       `xorm:"not null default 0 TINYINT(1)"`
	ReturnItem          []byte    `xorm:"BLOB"`
	LastUpdatedBy       string    `xorm:"VARCHAR(100)"`
}

func (m *CmsOrderItem) TableName() string {
	return "cms_order_item"
}
