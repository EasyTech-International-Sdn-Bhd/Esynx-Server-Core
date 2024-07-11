package entities

import (
	"time"
)

type CmsPurchaseReturn struct {
	Id                     string    `xorm:"not null pk index VARCHAR(20)"`
	PrDate                 time.Time `xorm:"DATETIME"`
	DeliveryDate           time.Time `xorm:"not null DATE"`
	TransferReceivedDate   time.Time `xorm:"comment('date transfer from ipad to CMS') DATETIME"`
	TotalDiscount          float64   `xorm:"default 0 DOUBLE"`
	DiscountMethod         string    `xorm:"VARCHAR(50)"`
	Tax                    float64   `xorm:"default 0 DOUBLE"`
	Shippingfee            float64   `xorm:"default 0 DOUBLE"`
	GrandTotal             float64   `xorm:"default 0 DOUBLE"`
	SstAmount              float64   `xorm:"default 0 DOUBLE"`
	SstTaxAmount           float64   `xorm:"default 0 DOUBLE"`
	CustCode               string    `xorm:"VARCHAR(200)"`
	CustCompanyName        string    `xorm:"VARCHAR(400)"`
	CustInchargePerson     string    `xorm:"VARCHAR(400)"`
	CustReference          string    `xorm:"VARCHAR(300)"`
	CustEmail              string    `xorm:"VARCHAR(300)"`
	CustTel                string    `xorm:"VARCHAR(100)"`
	CustFax                string    `xorm:"VARCHAR(100)"`
	BillingAddress1        string    `xorm:"VARCHAR(200)"`
	BillingAddress2        string    `xorm:"VARCHAR(200)"`
	BillingAddress3        string    `xorm:"VARCHAR(200)"`
	BillingAddress4        string    `xorm:"VARCHAR(200)"`
	BillingCity            string    `xorm:"VARCHAR(150)"`
	BillingState           string    `xorm:"VARCHAR(150)"`
	BillingZipcode         string    `xorm:"VARCHAR(150)"`
	BillingCountry         string    `xorm:"VARCHAR(150)"`
	ShippingAddress1       string    `xorm:"VARCHAR(200)"`
	ShippingAddress2       string    `xorm:"VARCHAR(200)"`
	ShippingAddress3       string    `xorm:"VARCHAR(200)"`
	ShippingAddress4       string    `xorm:"VARCHAR(200)"`
	ShippingCity           string    `xorm:"VARCHAR(150)"`
	ShippingState          string    `xorm:"VARCHAR(150)"`
	ShippingZipcode        string    `xorm:"VARCHAR(150)"`
	ShippingCountry        string    `xorm:"VARCHAR(150)"`
	Termcode               string    `xorm:"VARCHAR(20)"`
	Agent                  string    `xorm:"VARCHAR(50)"`
	ActiveStatus           string    `xorm:"ENUM('ACTIVE','CONFIRMED','POSTED','TRANSFERRED')"`
	OthersPrStatus         string    `xorm:"VARCHAR(150)"`
	PrStatusLastUpdateDate time.Time `xorm:"comment('sales update time') DATETIME"`
	PrStatusLastUpdateBy   string    `xorm:"comment('user name, user who accept the payment') VARCHAR(200)"`
	PrRemark               string    `xorm:"VARCHAR(500)"`
	PrValidity             string    `xorm:"not null default '2' VARCHAR(250)"`
	PrPaymentType          string    `xorm:"not null VARCHAR(250)"`
	PrReference            string    `xorm:"not null VARCHAR(250)"`
	PrDeliveryNote         string    `xorm:"not null default '' VARCHAR(200)"`
	PickerNote             string    `xorm:"VARCHAR(500)"`
	PackConfirmedBy        string    `xorm:"VARCHAR(100)"`
	BasketId               string    `xorm:"VARCHAR(100)"`
	PackingStatus          string    `xorm:"ENUM('IN_PROGRESS','NO_STOCK','PACKED','PENDING')"`
	CancelStatus           string    `xorm:"ENUM('ADMIN_CANCELLED','NOT_CANCELLED','SYSTEM_CANCELLED','USER_CANCELLED')"`
	WarehouseCode          string    `xorm:"VARCHAR(100)"`
	PackedBy               string    `xorm:"VARCHAR(100)"`
	PackConfirmed          string    `xorm:"ENUM('CONFIRMED','PENDING')"`
	LastPrint              time.Time `xorm:"DATETIME"`
	BranchCode             string    `xorm:"VARCHAR(100)"`
	Latitude               string    `xorm:"VARCHAR(100)"`
	Longitude              string    `xorm:"VARCHAR(100)"`
	PrUdf                  string    `xorm:"JSON"`
	DocType                string    `xorm:"VARCHAR(20)"`
	PrFault                string    `xorm:"ENUM('NO','YES')"`
	PrFaultMessage         string    `xorm:"VARCHAR(200)"`
	ZoneName               string    `xorm:"not null VARCHAR(100)"`
	PrApproved             string    `xorm:"ENUM('APPROVED','PENDING','REJECTED')"`
	PrApprover             string    `xorm:"VARCHAR(100)"`
	PrComment              string    `xorm:"VARCHAR(1000)"`
	PrFrom                 string    `xorm:"default 'S' VARCHAR(20)"`
}

func (m *CmsPurchaseReturn) TableName() string {
	return "cms_purchase_return"
}
