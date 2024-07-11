package entities

import (
	"time"
)

type CmsAccExistingOrder struct {
	RefNo                     string    `xorm:"VARCHAR(30)"`
	OrderId                   string    `xorm:"not null pk index VARCHAR(30)"`
	OrderDate                 time.Time `xorm:"DATETIME"`
	DeliveryDate              time.Time `xorm:"not null DATE"`
	TransferReceivedDate      time.Time `xorm:"comment('date transfer from ipad to CMS') DATETIME"`
	TotalDiscount             float64   `xorm:"default 0 DOUBLE"`
	DiscountMethod            string    `xorm:"VARCHAR(50)"`
	Tax                       float64   `xorm:"default 0 DOUBLE"`
	Shippingfee               float64   `xorm:"default 0 DOUBLE"`
	GrandTotal                float64   `xorm:"DOUBLE"`
	GstAmount                 float64   `xorm:"DOUBLE"`
	GstTaxAmount              float64   `xorm:"DOUBLE"`
	CustId                    int64     `xorm:"comment('note : customer id can be blank as salesperson is allowed to created manual customer via ipad. However, if they select from their customer list, then cust_id should be stored.') BIGINT"`
	CustCode                  string    `xorm:"VARCHAR(200)"`
	CustCompanyName           string    `xorm:"VARCHAR(400)"`
	CustInchargePerson        string    `xorm:"VARCHAR(400)"`
	CustReference             string    `xorm:"VARCHAR(300)"`
	CustEmail                 string    `xorm:"VARCHAR(300)"`
	CustTel                   string    `xorm:"VARCHAR(100)"`
	CustFax                   string    `xorm:"VARCHAR(100)"`
	BillingAddress1           string    `xorm:"VARCHAR(200)"`
	BillingAddress2           string    `xorm:"VARCHAR(200)"`
	BillingAddress3           string    `xorm:"VARCHAR(200)"`
	BillingAddress4           string    `xorm:"VARCHAR(200)"`
	BillingCity               string    `xorm:"VARCHAR(150)"`
	BillingState              string    `xorm:"VARCHAR(150)"`
	BillingZipcode            string    `xorm:"VARCHAR(150)"`
	BillingCountry            string    `xorm:"VARCHAR(150)"`
	ShippingAddress1          string    `xorm:"VARCHAR(200)"`
	ShippingAddress2          string    `xorm:"VARCHAR(200)"`
	ShippingAddress3          string    `xorm:"VARCHAR(200)"`
	ShippingAddress4          string    `xorm:"VARCHAR(200)"`
	ShippingCity              string    `xorm:"VARCHAR(150)"`
	ShippingState             string    `xorm:"VARCHAR(150)"`
	ShippingZipcode           string    `xorm:"VARCHAR(150)"`
	ShippingCountry           string    `xorm:"VARCHAR(150)"`
	Termcode                  string    `xorm:"VARCHAR(20)"`
	SalespersonId             int       `xorm:"comment('0 means no id, it is manual member') INT"`
	OrderStatus               int       `xorm:"default 1 comment('0=in_ipad,1=in_backoffice,2=in_QNE') INT"`
	OthersOrderStatus         string    `xorm:"VARCHAR(150)"`
	OrderStatusLastUpdateDate time.Time `xorm:"comment('sales update time') DATETIME"`
	OrderStatusLastUpdateBy   string    `xorm:"comment('user name, user who accept the payment') VARCHAR(200)"`
	OrderRemark               string    `xorm:"VARCHAR(500)"`
	OrderValidity             string    `xorm:"not null default '2' VARCHAR(250)"`
	OrderPaymentType          string    `xorm:"not null VARCHAR(250)"`
	OrderReference            string    `xorm:"not null VARCHAR(250)"`
	OrderDeliveryNote         string    `xorm:"not null default '' VARCHAR(200)"`
	PickerNote                string    `xorm:"VARCHAR(500)"`
	OrderFrom                 string    `xorm:"default 'S' comment('S = salesperson C = Customer') VARCHAR(10)"`
	CancelStatus              int       `xorm:"default 0 comment('0=not_cancel, 1=cancelled by agent, 2=cancelled by admin') INT"`
	WarehouseCode             string    `xorm:"VARCHAR(100)"`
	PackedBy                  string    `xorm:"VARCHAR(100)"`
	PackConfirmed             int       `xorm:"default 0 comment('0 pending 1 confirmed') INT"`
	PackConfirmedBy           string    `xorm:"VARCHAR(100)"`
	BasketId                  string    `xorm:"VARCHAR(100)"`
	PackingStatus             int       `xorm:"default 0 comment('0=not packed, 1=packed, 2=no stock') INT"`
	InternalUpdatedAt         time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	LastPrint                 time.Time `xorm:"DATETIME"`
	BranchCode                string    `xorm:"VARCHAR(100)"`
	Latitude                  string    `xorm:"VARCHAR(100)"`
	Longitude                 string    `xorm:"VARCHAR(100)"`
	OrderUdf                  string    `xorm:"JSON"`
	DocType                   string    `xorm:"VARCHAR(20)"`
	OrderFault                int       `xorm:"default 0 comment('0=no error, 1=cust_error, 2=product_error, 3=order_duplicate, 4=cust_limit, 5=detect another company, 6=not enough user') INT"`
	OrderFaultMessage         string    `xorm:"VARCHAR(200)"`
	ZoneName                  string    `xorm:"not null VARCHAR(100)"`
	SessionId                 string    `xorm:"default '' VARCHAR(100)"`
}

func (m *CmsAccExistingOrder) TableName() string {
	return "cms_acc_existing_order"
}
