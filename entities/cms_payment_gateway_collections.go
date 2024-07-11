package entities

import (
	"time"
)

type CmsPaymentGatewayCollections struct {
	Id               []byte    `xorm:"not null pk default uuid_to_bin(uuid()) unique BINARY(16)"`
	Agent            string    `xorm:"unique(cms_payment_gateway_collections_unq) VARCHAR(50)"`
	CollectionName   string    `xorm:"VARCHAR(50)"`
	CollectionId     string    `xorm:"unique(cms_payment_gateway_collections_unq) VARCHAR(50)"`
	CollectionPeriod time.Time `xorm:"DATE"`
	CreatedAt        time.Time `xorm:"DATETIME"`
	UpdatedAt        time.Time `xorm:"DATETIME"`
}

func (m *CmsPaymentGatewayCollections) TableName() string {
	return "cms_payment_gateway_collections"
}
