package entities

import (
	"time"
)

type CmsPaymentGatewayBills struct {
	Id              []byte    `xorm:"not null pk default uuid_to_bin(uuid()) unique BINARY(16)"`
	CollectionId    string    `xorm:"unique(cms_payment_gateway_bills_unq) VARCHAR(50)"`
	BillId          string    `xorm:"unique(cms_payment_gateway_bills_unq) VARCHAR(50)"`
	BillTempRef     string    `xorm:"unique(cms_payment_gateway_bills_unq) VARCHAR(50)"`
	BillAction      string    `xorm:"default 'before_payment' unique(cms_payment_gateway_bills_unq) ENUM('after_payment','before_payment')"`
	BillIsPaid      int       `xorm:"default 0 SMALLINT"`
	BillState       string    `xorm:"ENUM('deleted','due','paid')"`
	BillAmount      float64   `xorm:"default 0 DOUBLE"`
	BillPaidAmount  float64   `xorm:"default 0 DOUBLE"`
	BillDueAt       time.Time `xorm:"DATE"`
	BillEmail       string    `xorm:"VARCHAR(50)"`
	BillMobile      string    `xorm:"VARCHAR(20)"`
	BillUser        string    `xorm:"VARCHAR(70)"`
	BillDescription []byte    `xorm:"BLOB"`
	BillUrl         string    `xorm:"VARCHAR(100)"`
	BillReference   string    `xorm:"JSON"`
	RedirectUrl     string    `xorm:"VARCHAR(100)"`
	CallbackUrl     string    `xorm:"VARCHAR(100)"`
	CreatedAt       time.Time `xorm:"DATETIME"`
	UpdatedAt       time.Time `xorm:"DATETIME"`
}

func (m *CmsPaymentGatewayBills) TableName() string {
	return "cms_payment_gateway_bills"
}
