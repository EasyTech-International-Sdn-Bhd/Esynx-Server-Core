package entities

import (
	"time"
)

type CmsAccExistingOrderItem struct {
	OrderItemId         uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	OrderId             string    `xorm:"comment('this reference is unique and link to order table. it cannot use order id because the order are sending from different ipad, the order id which is generated from ipad might be the same when reach to CMS.') unique(unx) VARCHAR(20)"`
	IpadItemId          int       `xorm:"default 0 unique(unx) INT"`
	ProductCode         string    `xorm:"unique(unx) VARCHAR(50)"`
	ProductId           int       `xorm:"INT"`
	ProductName         string    `xorm:"VARCHAR(400)"`
	SalespersonRemark   string    `xorm:"not null default '' VARCHAR(200)"`
	Quantity            float64   `xorm:"DOUBLE"`
	EdittedQuantity     float32   `xorm:"default 0 FLOAT"`
	UnitPrice           float64   `xorm:"DOUBLE"`
	Disc1               float64   `xorm:"DOUBLE"`
	Disc2               float64   `xorm:"DOUBLE"`
	Disc3               float64   `xorm:"DOUBLE"`
	UnitUom             string    `xorm:"VARCHAR(100)"`
	AttributeRemark     string    `xorm:"comment('The chosen attribute name and value will be stored here, e.g. Size=L, Colour=Red') BLOB"`
	OptionalRemark      string    `xorm:"comment('The selected optional item will options here, e.g. Sport Rim, Leather Seat.') VARCHAR(100)"`
	DiscountMethod      string    `xorm:"comment('Percentage or Fixed (Amount)') VARCHAR(100)"`
	PickerNote          string    `xorm:"not null VARCHAR(500)"`
	DiscountAmount      string    `xorm:"VARCHAR(30)"`
	SubTotal            float64   `xorm:"comment('the optional item price will affect the sub-total') DOUBLE"`
	SequenceNo          int       `xorm:"INT"`
	UomId               int       `xorm:"not null INT"`
	PackingStatus       int       `xorm:"default 0 comment('0=not packed, 1=packed, 2=no stock, 3=no stock but informed') INT"`
	PackedQty           float32   `xorm:"default 0 FLOAT"`
	Isparent            int       `xorm:"INT"`
	ParentCode          string    `xorm:"VARCHAR(100)"`
	PackedBy            string    `xorm:"VARCHAR(200)"`
	PackConfirmedBy     string    `xorm:"VARCHAR(50)"`
	PackConfirmedStatus int       `xorm:"default 0 INT"`
	OrderItemValidity   int       `xorm:"default 2 comment('0 = reject 1=pending 2 = approved') INT"`
	CancelStatus        int       `xorm:"default 0 comment('0=not canceled, 1=canceld') INT"`
	UpdatedAt           time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	BatchNo             []byte    `xorm:"BLOB"`
	SerialNo            []byte    `xorm:"BLOB"`
	SessionId           string    `xorm:"default '' VARCHAR(100)"`
	ItemUdf             string    `xorm:"JSON"`
}

func (m *CmsAccExistingOrderItem) TableName() string {
	return "cms_acc_existing_order_item"
}
