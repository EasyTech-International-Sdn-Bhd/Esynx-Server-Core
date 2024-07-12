package entities

type CmsStockTakeDtl struct {
	Id                uint64  `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	StId              int     `xorm:"not null unique(unq) INT"`
	ProductCode       string  `xorm:"not null unique(unq) VARCHAR(200)"`
	ProductName       string  `xorm:"not null VARCHAR(200)"`
	CurrentQuantity   float64 `xorm:"not null DOUBLE"`
	SuggestedQuantity float64 `xorm:"not null DOUBLE"`
	SpRemark          string  `xorm:"VARCHAR(500)"`
	UnitUom           string  `xorm:"not null VARCHAR(20)"`
	ActiveStatus      int     `xorm:"default 1 INT"`
}

func (m *CmsStockTakeDtl) TableName() string {
	return "cms_stock_take_dtl"
}
