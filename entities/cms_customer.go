package entities

import (
	"time"
)

type CmsCustomer struct {
	CustId             uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"custId,omitempty" xml:"custId"`
	CreatedDate        time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP" json:"createdDate,omitempty" xml:"createdDate"`
	CustCode           string    `xorm:"unique VARCHAR(200)" json:"custCode,omitempty" xml:"custCode"`
	CustCompanyName    string    `xorm:"VARCHAR(400)" json:"custCompanyName,omitempty" xml:"custCompanyName"`
	CustInchargePerson string    `xorm:"VARCHAR(400)" json:"custInchargePerson,omitempty" xml:"custInchargePerson"`
	CustRemark         string    `xorm:"not null default '' VARCHAR(100)" json:"custRemark,omitempty" xml:"custRemark"`
	CustReference      string    `xorm:"VARCHAR(300)" json:"custReference,omitempty" xml:"custReference"`
	CustEmail          string    `xorm:"VARCHAR(300)" json:"custEmail,omitempty" xml:"custEmail"`
	CustTel            string    `xorm:"VARCHAR(100)" json:"custTel,omitempty" xml:"custTel"`
	CustFax            string    `xorm:"VARCHAR(100)" json:"custFax,omitempty" xml:"custFax"`
	BillingAddress1    string    `xorm:"VARCHAR(200)" json:"billingAddress1,omitempty" xml:"billingAddress1"`
	BillingAddress2    string    `xorm:"VARCHAR(200)" json:"billingAddress2,omitempty" xml:"billingAddress2"`
	BillingAddress3    string    `xorm:"VARCHAR(200)" json:"billingAddress3,omitempty" xml:"billingAddress3"`
	BillingAddress4    string    `xorm:"VARCHAR(200)" json:"billingAddress4,omitempty" xml:"billingAddress4"`
	BillingCity        string    `xorm:"VARCHAR(150)" json:"billingCity,omitempty" xml:"billingCity"`
	BillingState       string    `xorm:"VARCHAR(150)" json:"billingState,omitempty" xml:"billingState"`
	BillingZipcode     string    `xorm:"VARCHAR(150)" json:"billingZipcode,omitempty" xml:"billingZipcode"`
	BillingCountry     string    `xorm:"VARCHAR(150)" json:"billingCountry,omitempty" xml:"billingCountry"`
	ShippingAddress1   string    `xorm:"VARCHAR(200)" json:"shippingAddress1,omitempty" xml:"shippingAddress1"`
	ShippingAddress2   string    `xorm:"VARCHAR(200)" json:"shippingAddress2,omitempty" xml:"shippingAddress2"`
	ShippingAddress3   string    `xorm:"VARCHAR(200)" json:"shippingAddress3,omitempty" xml:"shippingAddress3"`
	ShippingAddress4   string    `xorm:"VARCHAR(200)" json:"shippingAddress4,omitempty" xml:"shippingAddress4"`
	ShippingCity       string    `xorm:"VARCHAR(150)" json:"shippingCity,omitempty" xml:"shippingCity"`
	ShippingState      string    `xorm:"VARCHAR(150)" json:"shippingState,omitempty" xml:"shippingState"`
	ShippingZipcode    string    `xorm:"VARCHAR(150)" json:"shippingZipcode,omitempty" xml:"shippingZipcode"`
	ShippingCountry    string    `xorm:"VARCHAR(150)" json:"shippingCountry,omitempty" xml:"shippingCountry"`
	Termcode           string    `xorm:"VARCHAR(20)" json:"termcode,omitempty" xml:"termcode"`
	SellingPriceType   string    `xorm:"VARCHAR(20)" json:"sellingPriceType,omitempty" xml:"sellingPriceType"`
	CustomerZone       int       `xorm:"INT" json:"customerZone,omitempty" xml:"customerZone"`
	CustomerStatus     int       `xorm:"default 1 comment('0=inactive, 1=active') INT" json:"customerStatus,omitempty" xml:"customerStatus"`
	CreditLimit        float64   `xorm:"default 0 DOUBLE" json:"creditLimit,omitempty" xml:"creditLimit"`
	CurrentBalance     float64   `xorm:"default 0 DOUBLE" json:"currentBalance,omitempty" xml:"currentBalance"`
	Currency           string    `xorm:"VARCHAR(20)" json:"currency,omitempty" xml:"currency"`
	CurrencyRate       float64   `xorm:"default 0 DOUBLE" json:"currencyRate,omitempty" xml:"currencyRate"`
	Latitude           float64   `xorm:"DOUBLE" json:"latitude,omitempty" xml:"latitude"`
	Longitude          float64   `xorm:"DOUBLE" json:"longitude,omitempty" xml:"longitude"`
	UpdatedAt          time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
	CustUdf            string    `xorm:"JSON" json:"custUdf,omitempty" xml:"custUdf"`
	SessionId          string    `xorm:"default '' VARCHAR(100)" json:"sessionId,omitempty" xml:"sessionId"`
	Company            string    `xorm:"default '' VARCHAR(50)" json:"company,omitempty" xml:"company"`
}

func (m *CmsCustomer) TableName() string {
	return "cms_customer"
}

func (m *CmsCustomer) BeforeInsert() {
	m.BeforeUpdate()
}

func (m *CmsCustomer) BeforeUpdate() {
	m.UpdatedAt = time.Now()
	if m.CustUdf == "" {
		m.CustUdf = "{}"
	}
	if m.CreatedDate.IsZero() {
		m.CreatedDate = time.Now()
	}
}
