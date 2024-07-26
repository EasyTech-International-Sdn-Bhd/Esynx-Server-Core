package entities

import (
	"time"
)

type CmsPaymentGatewayBills struct {
	Id              uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	CollectionId    string    `xorm:"unique(cms_payment_gateway_bills_unq) VARCHAR(50)" json:"collectionId,omitempty" xml:"collectionId"`
	BillId          string    `xorm:"unique(cms_payment_gateway_bills_unq) VARCHAR(50)" json:"billId,omitempty" xml:"billId"`
	BillTempRef     string    `xorm:"unique(cms_payment_gateway_bills_unq) VARCHAR(50)" json:"billTempRef,omitempty" xml:"billTempRef"`
	BillAction      string    `xorm:"default 'before_payment' unique(cms_payment_gateway_bills_unq) ENUM('after_payment','before_payment')" json:"billAction,omitempty" xml:"billAction"`
	BillIsPaid      int       `xorm:"default 0 SMALLINT" json:"billIsPaid,omitempty" xml:"billIsPaid"`
	BillState       string    `xorm:"ENUM('deleted','due','paid')" json:"billState,omitempty" xml:"billState"`
	BillAmount      float64   `xorm:"default 0 DOUBLE" json:"billAmount,omitempty" xml:"billAmount"`
	BillPaidAmount  float64   `xorm:"default 0 DOUBLE" json:"billPaidAmount,omitempty" xml:"billPaidAmount"`
	BillDueAt       time.Time `xorm:"DATE" json:"billDueAt,omitempty" xml:"billDueAt"`
	BillEmail       string    `xorm:"VARCHAR(50)" json:"billEmail,omitempty" xml:"billEmail"`
	BillMobile      string    `xorm:"VARCHAR(20)" json:"billMobile,omitempty" xml:"billMobile"`
	BillUser        string    `xorm:"VARCHAR(70)" json:"billUser,omitempty" xml:"billUser"`
	BillDescription []byte    `xorm:"BLOB" json:"billDescription,omitempty" xml:"billDescription"`
	BillUrl         string    `xorm:"VARCHAR(100)" json:"billUrl,omitempty" xml:"billUrl"`
	BillReference   string    `xorm:"JSON" json:"billReference,omitempty" xml:"billReference"`
	RedirectUrl     string    `xorm:"VARCHAR(100)" json:"redirectUrl,omitempty" xml:"redirectUrl"`
	CallbackUrl     string    `xorm:"VARCHAR(100)" json:"callbackUrl,omitempty" xml:"callbackUrl"`
	CreatedAt       time.Time `xorm:"DATETIME" json:"createdAt,omitempty" xml:"createdAt"`
	UpdatedAt       time.Time `xorm:"DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsPaymentGatewayBills) TableName() string {
	return "cms_payment_gateway_bills"
}
