package entities

import (
	"time"
)

type CmsPaymentGatewayCollections struct {
	Id               uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	Agent            string    `xorm:"unique(gateway_col_unq) VARCHAR(50)"`
	CollectionName   string    `xorm:"VARCHAR(50)"`
	CollectionId     string    `xorm:"unique(gateway_col_unq) VARCHAR(50)"`
	CollectionPeriod time.Time `xorm:"DATE"`
	CreatedAt        time.Time `xorm:"DATETIME"`
	UpdatedAt        time.Time `xorm:"DATETIME"`
}

func (m *CmsPaymentGatewayCollections) TableName() string {
	return "cms_payment_gateway_collections"
}
