package entities

import (
	"time"
)

type CmsCustomer struct {
	CustId             int       `xorm:"not null pk autoincr INT"`
	CreatedDate        time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
	CustCode           string    `xorm:"unique VARCHAR(200)"`
	CustCompanyName    string    `xorm:"VARCHAR(400)"`
	CustInchargePerson string    `xorm:"VARCHAR(400)"`
	CustRemark         string    `xorm:"not null default '' VARCHAR(100)"`
	CustReference      string    `xorm:"VARCHAR(300)"`
	CustEmail          string    `xorm:"VARCHAR(300)"`
	CustTel            string    `xorm:"VARCHAR(100)"`
	CustFax            string    `xorm:"VARCHAR(100)"`
	BillingAddress1    string    `xorm:"VARCHAR(200)"`
	BillingAddress2    string    `xorm:"VARCHAR(200)"`
	BillingAddress3    string    `xorm:"VARCHAR(200)"`
	BillingAddress4    string    `xorm:"VARCHAR(200)"`
	BillingCity        string    `xorm:"VARCHAR(150)"`
	BillingState       string    `xorm:"VARCHAR(150)"`
	BillingZipcode     string    `xorm:"VARCHAR(150)"`
	BillingCountry     string    `xorm:"VARCHAR(150)"`
	ShippingAddress1   string    `xorm:"VARCHAR(200)"`
	ShippingAddress2   string    `xorm:"VARCHAR(200)"`
	ShippingAddress3   string    `xorm:"VARCHAR(200)"`
	ShippingAddress4   string    `xorm:"VARCHAR(200)"`
	ShippingCity       string    `xorm:"VARCHAR(150)"`
	ShippingState      string    `xorm:"VARCHAR(150)"`
	ShippingZipcode    string    `xorm:"VARCHAR(150)"`
	ShippingCountry    string    `xorm:"VARCHAR(150)"`
	Termcode           string    `xorm:"VARCHAR(20)"`
	SellingPriceType   string    `xorm:"VARCHAR(20)"`
	CustomerZone       int       `xorm:"INT"`
	CustomerStatus     int       `xorm:"default 1 comment('0=inactive, 1=active') INT"`
	CreditLimit        float64   `xorm:"default 0 DOUBLE"`
	CurrentBalance     float64   `xorm:"default 0 DOUBLE"`
	Currency           string    `xorm:"VARCHAR(20)"`
	CurrencyRate       float64   `xorm:"default 0 DOUBLE"`
	Latitude           float64   `xorm:"DOUBLE"`
	Longitude          float64   `xorm:"DOUBLE"`
	UpdatedAt          time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	CustUdf            string    `xorm:"JSON"`
	SessionId          string    `xorm:"default '' VARCHAR(100)"`
	Company            string    `xorm:"default '' VARCHAR(50)"`
}

func (m *CmsCustomer) TableName() string {
	return "cms_customer"
}

func (m *CmsCustomer) Validate() {
	if m.CustUdf == "" {
		m.CustUdf = "{}"
	}
	if m.CreatedDate.IsZero() {
		m.CreatedDate = time.Now()
	}
}

func (m *CmsCustomer) ToUpdate() {
	m.UpdatedAt = time.Now()
}
