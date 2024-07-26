package entities

import (
	"time"
)

type CmsPurchaseReturn struct {
	Id                     string    `xorm:"not null pk index VARCHAR(20)" json:"id,omitempty" xml:"id"`
	PrDate                 time.Time `xorm:"DATETIME" json:"prDate,omitempty" xml:"prDate"`
	DeliveryDate           time.Time `xorm:"not null DATE" json:"deliveryDate,omitempty" xml:"deliveryDate"`
	TransferReceivedDate   time.Time `xorm:"comment('date transfer from ipad to CMS') DATETIME" json:"transferReceivedDate,omitempty" xml:"transferReceivedDate"`
	TotalDiscount          float64   `xorm:"default 0 DOUBLE" json:"totalDiscount,omitempty" xml:"totalDiscount"`
	DiscountMethod         string    `xorm:"VARCHAR(50)" json:"discountMethod,omitempty" xml:"discountMethod"`
	Tax                    float64   `xorm:"default 0 DOUBLE" json:"tax,omitempty" xml:"tax"`
	Shippingfee            float64   `xorm:"default 0 DOUBLE" json:"shippingfee,omitempty" xml:"shippingfee"`
	GrandTotal             float64   `xorm:"default 0 DOUBLE" json:"grandTotal,omitempty" xml:"grandTotal"`
	SstAmount              float64   `xorm:"default 0 DOUBLE" json:"sstAmount,omitempty" xml:"sstAmount"`
	SstTaxAmount           float64   `xorm:"default 0 DOUBLE" json:"sstTaxAmount,omitempty" xml:"sstTaxAmount"`
	CustCode               string    `xorm:"VARCHAR(200)" json:"custCode,omitempty" xml:"custCode"`
	CustCompanyName        string    `xorm:"VARCHAR(400)" json:"custCompanyName,omitempty" xml:"custCompanyName"`
	CustInchargePerson     string    `xorm:"VARCHAR(400)" json:"custInchargePerson,omitempty" xml:"custInchargePerson"`
	CustReference          string    `xorm:"VARCHAR(300)" json:"custReference,omitempty" xml:"custReference"`
	CustEmail              string    `xorm:"VARCHAR(300)" json:"custEmail,omitempty" xml:"custEmail"`
	CustTel                string    `xorm:"VARCHAR(100)" json:"custTel,omitempty" xml:"custTel"`
	CustFax                string    `xorm:"VARCHAR(100)" json:"custFax,omitempty" xml:"custFax"`
	BillingAddress1        string    `xorm:"VARCHAR(200)" json:"billingAddress1,omitempty" xml:"billingAddress1"`
	BillingAddress2        string    `xorm:"VARCHAR(200)" json:"billingAddress2,omitempty" xml:"billingAddress2"`
	BillingAddress3        string    `xorm:"VARCHAR(200)" json:"billingAddress3,omitempty" xml:"billingAddress3"`
	BillingAddress4        string    `xorm:"VARCHAR(200)" json:"billingAddress4,omitempty" xml:"billingAddress4"`
	BillingCity            string    `xorm:"VARCHAR(150)" json:"billingCity,omitempty" xml:"billingCity"`
	BillingState           string    `xorm:"VARCHAR(150)" json:"billingState,omitempty" xml:"billingState"`
	BillingZipcode         string    `xorm:"VARCHAR(150)" json:"billingZipcode,omitempty" xml:"billingZipcode"`
	BillingCountry         string    `xorm:"VARCHAR(150)" json:"billingCountry,omitempty" xml:"billingCountry"`
	ShippingAddress1       string    `xorm:"VARCHAR(200)" json:"shippingAddress1,omitempty" xml:"shippingAddress1"`
	ShippingAddress2       string    `xorm:"VARCHAR(200)" json:"shippingAddress2,omitempty" xml:"shippingAddress2"`
	ShippingAddress3       string    `xorm:"VARCHAR(200)" json:"shippingAddress3,omitempty" xml:"shippingAddress3"`
	ShippingAddress4       string    `xorm:"VARCHAR(200)" json:"shippingAddress4,omitempty" xml:"shippingAddress4"`
	ShippingCity           string    `xorm:"VARCHAR(150)" json:"shippingCity,omitempty" xml:"shippingCity"`
	ShippingState          string    `xorm:"VARCHAR(150)" json:"shippingState,omitempty" xml:"shippingState"`
	ShippingZipcode        string    `xorm:"VARCHAR(150)" json:"shippingZipcode,omitempty" xml:"shippingZipcode"`
	ShippingCountry        string    `xorm:"VARCHAR(150)" json:"shippingCountry,omitempty" xml:"shippingCountry"`
	Termcode               string    `xorm:"VARCHAR(20)" json:"termcode,omitempty" xml:"termcode"`
	Agent                  string    `xorm:"VARCHAR(50)" json:"agent,omitempty" xml:"agent"`
	ActiveStatus           string    `xorm:"ENUM('ACTIVE','CONFIRMED','POSTED','TRANSFERRED')" json:"activeStatus,omitempty" xml:"activeStatus"`
	OthersPrStatus         string    `xorm:"VARCHAR(150)" json:"othersPrStatus,omitempty" xml:"othersPrStatus"`
	PrStatusLastUpdateDate time.Time `xorm:"comment('sales update time') DATETIME" json:"prStatusLastUpdateDate,omitempty" xml:"prStatusLastUpdateDate"`
	PrStatusLastUpdateBy   string    `xorm:"comment('user name, user who accept the payment') VARCHAR(200)" json:"prStatusLastUpdateBy,omitempty" xml:"prStatusLastUpdateBy"`
	PrRemark               string    `xorm:"VARCHAR(500)" json:"prRemark,omitempty" xml:"prRemark"`
	PrValidity             string    `xorm:"not null default '2' VARCHAR(250)" json:"prValidity,omitempty" xml:"prValidity"`
	PrPaymentType          string    `xorm:"not null VARCHAR(250)" json:"prPaymentType,omitempty" xml:"prPaymentType"`
	PrReference            string    `xorm:"not null VARCHAR(250)" json:"prReference,omitempty" xml:"prReference"`
	PrDeliveryNote         string    `xorm:"not null default '' VARCHAR(200)" json:"prDeliveryNote,omitempty" xml:"prDeliveryNote"`
	PickerNote             string    `xorm:"VARCHAR(500)" json:"pickerNote,omitempty" xml:"pickerNote"`
	PackConfirmedBy        string    `xorm:"VARCHAR(100)" json:"packConfirmedBy,omitempty" xml:"packConfirmedBy"`
	BasketId               string    `xorm:"VARCHAR(100)" json:"basketId,omitempty" xml:"basketId"`
	PackingStatus          string    `xorm:"ENUM('IN_PROGRESS','NO_STOCK','PACKED','PENDING')" json:"packingStatus,omitempty" xml:"packingStatus"`
	CancelStatus           string    `xorm:"ENUM('ADMIN_CANCELLED','NOT_CANCELLED','SYSTEM_CANCELLED','USER_CANCELLED')" json:"cancelStatus,omitempty" xml:"cancelStatus"`
	WarehouseCode          string    `xorm:"VARCHAR(100)" json:"warehouseCode,omitempty" xml:"warehouseCode"`
	PackedBy               string    `xorm:"VARCHAR(100)" json:"packedBy,omitempty" xml:"packedBy"`
	PackConfirmed          string    `xorm:"ENUM('CONFIRMED','PENDING')" json:"packConfirmed,omitempty" xml:"packConfirmed"`
	LastPrint              time.Time `xorm:"DATETIME" json:"lastPrint,omitempty" xml:"lastPrint"`
	BranchCode             string    `xorm:"VARCHAR(100)" json:"branchCode,omitempty" xml:"branchCode"`
	Latitude               string    `xorm:"VARCHAR(100)" json:"latitude,omitempty" xml:"latitude"`
	Longitude              string    `xorm:"VARCHAR(100)" json:"longitude,omitempty" xml:"longitude"`
	PrUdf                  string    `xorm:"JSON" json:"prUdf,omitempty" xml:"prUdf"`
	DocType                string    `xorm:"VARCHAR(20)" json:"docType,omitempty" xml:"docType"`
	PrFault                string    `xorm:"ENUM('NO','YES')" json:"prFault,omitempty" xml:"prFault"`
	PrFaultMessage         string    `xorm:"VARCHAR(200)" json:"prFaultMessage,omitempty" xml:"prFaultMessage"`
	ZoneName               string    `xorm:"not null VARCHAR(100)" json:"zoneName,omitempty" xml:"zoneName"`
	PrApproved             string    `xorm:"ENUM('APPROVED','PENDING','REJECTED')" json:"prApproved,omitempty" xml:"prApproved"`
	PrApprover             string    `xorm:"VARCHAR(100)" json:"prApprover,omitempty" xml:"prApprover"`
	PrComment              string    `xorm:"VARCHAR(1000)" json:"prComment,omitempty" xml:"prComment"`
	PrFrom                 string    `xorm:"default 'S' VARCHAR(20)" json:"prFrom,omitempty" xml:"prFrom"`
}

func (m *CmsPurchaseReturn) TableName() string {
	return "cms_purchase_return"
}
