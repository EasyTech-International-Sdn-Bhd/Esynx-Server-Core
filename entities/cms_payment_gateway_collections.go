package entities

import (
	"time"
)

type CmsPaymentGatewayCollections struct {
	Id               uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	Agent            string    `xorm:"unique(gateway_col_unq) VARCHAR(50)" json:"agent,omitempty" xml:"agent"`
	CollectionName   string    `xorm:"VARCHAR(50)" json:"collectionName,omitempty" xml:"collectionName"`
	CollectionId     string    `xorm:"unique(gateway_col_unq) VARCHAR(50)" json:"collectionId,omitempty" xml:"collectionId"`
	CollectionPeriod time.Time `xorm:"DATE" json:"collectionPeriod,omitempty" xml:"collectionPeriod"`
	CreatedAt        time.Time `xorm:"DATETIME" json:"createdAt,omitempty" xml:"createdAt"`
	UpdatedAt        time.Time `xorm:"DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
}

func (m *CmsPaymentGatewayCollections) TableName() string {
	return "cms_payment_gateway_collections"
}
