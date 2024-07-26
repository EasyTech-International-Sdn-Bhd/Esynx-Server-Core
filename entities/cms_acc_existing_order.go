package entities

import (
	"time"
)

type CmsAccExistingOrder struct {
	RefNo                     string    `xorm:"VARCHAR(30)" json:"refNo,omitempty" xml:"refNo"`
	OrderId                   string    `xorm:"not null pk index VARCHAR(30)" json:"orderId,omitempty" xml:"orderId"`
	OrderDate                 time.Time `xorm:"DATETIME" json:"orderDate,omitempty" xml:"orderDate"`
	DeliveryDate              time.Time `xorm:"not null DATE" json:"deliveryDate,omitempty" xml:"deliveryDate"`
	TransferReceivedDate      time.Time `xorm:"comment('date transfer from ipad to CMS') DATETIME" json:"transferReceivedDate,omitempty" xml:"transferReceivedDate"`
	TotalDiscount             float64   `xorm:"default 0 DOUBLE" json:"totalDiscount,omitempty" xml:"totalDiscount"`
	DiscountMethod            string    `xorm:"VARCHAR(50)" json:"discountMethod,omitempty" xml:"discountMethod"`
	Tax                       float64   `xorm:"default 0 DOUBLE" json:"tax,omitempty" xml:"tax"`
	Shippingfee               float64   `xorm:"default 0 DOUBLE" json:"shippingfee,omitempty" xml:"shippingfee"`
	GrandTotal                float64   `xorm:"DOUBLE" json:"grandTotal,omitempty" xml:"grandTotal"`
	GstAmount                 float64   `xorm:"DOUBLE" json:"gstAmount,omitempty" xml:"gstAmount"`
	GstTaxAmount              float64   `xorm:"DOUBLE" json:"gstTaxAmount,omitempty" xml:"gstTaxAmount"`
	CustId                    int64     `xorm:"comment('note : customer id can be blank as salesperson is allowed to created manual customer via ipad. However, if they select from their customer list, then cust_id should be stored.') BIGINT" json:"custId,omitempty" xml:"custId"`
	CustCode                  string    `xorm:"VARCHAR(200)" json:"custCode,omitempty" xml:"custCode"`
	CustCompanyName           string    `xorm:"VARCHAR(400)" json:"custCompanyName,omitempty" xml:"custCompanyName"`
	CustInchargePerson        string    `xorm:"VARCHAR(400)" json:"custInchargePerson,omitempty" xml:"custInchargePerson"`
	CustReference             string    `xorm:"VARCHAR(300)" json:"custReference,omitempty" xml:"custReference"`
	CustEmail                 string    `xorm:"VARCHAR(300)" json:"custEmail,omitempty" xml:"custEmail"`
	CustTel                   string    `xorm:"VARCHAR(100)" json:"custTel,omitempty" xml:"custTel"`
	CustFax                   string    `xorm:"VARCHAR(100)" json:"custFax,omitempty" xml:"custFax"`
	BillingAddress1           string    `xorm:"VARCHAR(200)" json:"billingAddress1,omitempty" xml:"billingAddress1"`
	BillingAddress2           string    `xorm:"VARCHAR(200)" json:"billingAddress2,omitempty" xml:"billingAddress2"`
	BillingAddress3           string    `xorm:"VARCHAR(200)" json:"billingAddress3,omitempty" xml:"billingAddress3"`
	BillingAddress4           string    `xorm:"VARCHAR(200)" json:"billingAddress4,omitempty" xml:"billingAddress4"`
	BillingCity               string    `xorm:"VARCHAR(150)" json:"billingCity,omitempty" xml:"billingCity"`
	BillingState              string    `xorm:"VARCHAR(150)" json:"billingState,omitempty" xml:"billingState"`
	BillingZipcode            string    `xorm:"VARCHAR(150)" json:"billingZipcode,omitempty" xml:"billingZipcode"`
	BillingCountry            string    `xorm:"VARCHAR(150)" json:"billingCountry,omitempty" xml:"billingCountry"`
	ShippingAddress1          string    `xorm:"VARCHAR(200)" json:"shippingAddress1,omitempty" xml:"shippingAddress1"`
	ShippingAddress2          string    `xorm:"VARCHAR(200)" json:"shippingAddress2,omitempty" xml:"shippingAddress2"`
	ShippingAddress3          string    `xorm:"VARCHAR(200)" json:"shippingAddress3,omitempty" xml:"shippingAddress3"`
	ShippingAddress4          string    `xorm:"VARCHAR(200)" json:"shippingAddress4,omitempty" xml:"shippingAddress4"`
	ShippingCity              string    `xorm:"VARCHAR(150)" json:"shippingCity,omitempty" xml:"shippingCity"`
	ShippingState             string    `xorm:"VARCHAR(150)" json:"shippingState,omitempty" xml:"shippingState"`
	ShippingZipcode           string    `xorm:"VARCHAR(150)" json:"shippingZipcode,omitempty" xml:"shippingZipcode"`
	ShippingCountry           string    `xorm:"VARCHAR(150)" json:"shippingCountry,omitempty" xml:"shippingCountry"`
	Termcode                  string    `xorm:"VARCHAR(20)" json:"termcode,omitempty" xml:"termcode"`
	SalespersonId             int       `xorm:"comment('0 means no id, it is manual member') INT" json:"salespersonId,omitempty" xml:"salespersonId"`
	OrderStatus               int       `xorm:"default 1 comment('0=in_ipad,1=in_backoffice,2=in_QNE') INT" json:"orderStatus,omitempty" xml:"orderStatus"`
	OthersOrderStatus         string    `xorm:"VARCHAR(150)" json:"othersOrderStatus,omitempty" xml:"othersOrderStatus"`
	OrderStatusLastUpdateDate time.Time `xorm:"comment('sales update time') DATETIME" json:"orderStatusLastUpdateDate,omitempty" xml:"orderStatusLastUpdateDate"`
	OrderStatusLastUpdateBy   string    `xorm:"comment('user name, user who accept the payment') VARCHAR(200)" json:"orderStatusLastUpdateBy,omitempty" xml:"orderStatusLastUpdateBy"`
	OrderRemark               string    `xorm:"VARCHAR(500)" json:"orderRemark,omitempty" xml:"orderRemark"`
	OrderValidity             string    `xorm:"not null default '2' VARCHAR(250)" json:"orderValidity,omitempty" xml:"orderValidity"`
	OrderPaymentType          string    `xorm:"not null VARCHAR(250)" json:"orderPaymentType,omitempty" xml:"orderPaymentType"`
	OrderReference            string    `xorm:"not null VARCHAR(250)" json:"orderReference,omitempty" xml:"orderReference"`
	OrderDeliveryNote         string    `xorm:"not null default '' VARCHAR(200)" json:"orderDeliveryNote,omitempty" xml:"orderDeliveryNote"`
	PickerNote                string    `xorm:"VARCHAR(500)" json:"pickerNote,omitempty" xml:"pickerNote"`
	OrderFrom                 string    `xorm:"default 'S' comment('S = salesperson C = Customer') VARCHAR(10)" json:"orderFrom,omitempty" xml:"orderFrom"`
	CancelStatus              int       `xorm:"default 0 comment('0=not_cancel, 1=cancelled by agent, 2=cancelled by admin') INT" json:"cancelStatus,omitempty" xml:"cancelStatus"`
	WarehouseCode             string    `xorm:"VARCHAR(100)" json:"warehouseCode,omitempty" xml:"warehouseCode"`
	PackedBy                  string    `xorm:"VARCHAR(100)" json:"packedBy,omitempty" xml:"packedBy"`
	PackConfirmed             int       `xorm:"default 0 comment('0 pending 1 confirmed') INT" json:"packConfirmed,omitempty" xml:"packConfirmed"`
	PackConfirmedBy           string    `xorm:"VARCHAR(100)" json:"packConfirmedBy,omitempty" xml:"packConfirmedBy"`
	BasketId                  string    `xorm:"VARCHAR(100)" json:"basketId,omitempty" xml:"basketId"`
	PackingStatus             int       `xorm:"default 0 comment('0=not packed, 1=packed, 2=no stock') INT" json:"packingStatus,omitempty" xml:"packingStatus"`
	InternalUpdatedAt         time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"internalUpdatedAt,omitempty" xml:"internalUpdatedAt"`
	LastPrint                 time.Time `xorm:"DATETIME" json:"lastPrint,omitempty" xml:"lastPrint"`
	BranchCode                string    `xorm:"VARCHAR(100)" json:"branchCode,omitempty" xml:"branchCode"`
	Latitude                  string    `xorm:"VARCHAR(100)" json:"latitude,omitempty" xml:"latitude"`
	Longitude                 string    `xorm:"VARCHAR(100)" json:"longitude,omitempty" xml:"longitude"`
	OrderUdf                  string    `xorm:"JSON" json:"orderUdf,omitempty" xml:"orderUdf"`
	DocType                   string    `xorm:"VARCHAR(20)" json:"docType,omitempty" xml:"docType"`
	OrderFault                int       `xorm:"default 0 comment('0=no error, 1=cust_error, 2=product_error, 3=order_duplicate, 4=cust_limit, 5=detect another company, 6=not enough user') INT" json:"orderFault,omitempty" xml:"orderFault"`
	OrderFaultMessage         string    `xorm:"VARCHAR(200)" json:"orderFaultMessage,omitempty" xml:"orderFaultMessage"`
	ZoneName                  string    `xorm:"not null VARCHAR(100)" json:"zoneName,omitempty" xml:"zoneName"`
	SessionId                 string    `xorm:"default '' VARCHAR(100)" json:"sessionId,omitempty" xml:"sessionId"`
}

func (m *CmsAccExistingOrder) TableName() string {
	return "cms_acc_existing_order"
}
