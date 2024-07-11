package entities

import (
	"time"
)

type CmsCustomerBranch struct {
	BranchId         int       `xorm:"not null pk autoincr INT"`
	CustId           int       `xorm:"not null INT"`
	AgentId          int       `xorm:"not null INT"`
	CustCode         string    `xorm:"not null unique(unique_branch) VARCHAR(100)"`
	BranchCode       string    `xorm:"not null unique(unique_branch) VARCHAR(100)"`
	BranchName       string    `xorm:"not null VARCHAR(200)"`
	BranchAttn       string    `xorm:"not null VARCHAR(100)"`
	BranchPhone      string    `xorm:"not null VARCHAR(100)"`
	BranchFax        string    `xorm:"not null VARCHAR(100)"`
	BillingAddress1  string    `xorm:"not null VARCHAR(200)"`
	BillingAddress2  string    `xorm:"not null VARCHAR(200)"`
	BillingAddress3  string    `xorm:"not null VARCHAR(200)"`
	BillingAddress4  string    `xorm:"not null VARCHAR(200)"`
	BillingState     string    `xorm:"not null VARCHAR(100)"`
	BillingPostcode  string    `xorm:"not null VARCHAR(100)"`
	BillingCountry   string    `xorm:"not null VARCHAR(100)"`
	ShippingAddress1 string    `xorm:"not null VARCHAR(200)"`
	ShippingAddress2 string    `xorm:"not null VARCHAR(200)"`
	ShippingAddress3 string    `xorm:"not null VARCHAR(200)"`
	ShippingAddress4 string    `xorm:"not null VARCHAR(200)"`
	ShippingState    string    `xorm:"not null VARCHAR(100)"`
	ShippingPostcode string    `xorm:"not null VARCHAR(100)"`
	ShippingCountry  string    `xorm:"not null VARCHAR(100)"`
	BranchArea       string    `xorm:"not null VARCHAR(500)"`
	BranchRemark     string    `xorm:"not null VARCHAR(500)"`
	BranchActive     int       `xorm:"not null default 1 comment('1=active, 0=not active') INT"`
	SessionId        string    `xorm:"default '' VARCHAR(100)"`
	Company          string    `xorm:"default '' VARCHAR(50)"`
	UpdatedAt        time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (m *CmsCustomerBranch) TableName() string {
	return "cms_customer_branch"
}
