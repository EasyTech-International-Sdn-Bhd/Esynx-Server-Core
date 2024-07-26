package entities

type CmsStockTakeDtl struct {
	Id                uint64  `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	StId              int     `xorm:"not null unique(unq) INT" json:"stId,omitempty" xml:"stId"`
	ProductCode       string  `xorm:"not null unique(unq) VARCHAR(200)" json:"productCode,omitempty" xml:"productCode"`
	ProductName       string  `xorm:"not null VARCHAR(200)" json:"productName,omitempty" xml:"productName"`
	CurrentQuantity   float64 `xorm:"not null DOUBLE" json:"currentQuantity,omitempty" xml:"currentQuantity"`
	SuggestedQuantity float64 `xorm:"not null DOUBLE" json:"suggestedQuantity,omitempty" xml:"suggestedQuantity"`
	SpRemark          string  `xorm:"VARCHAR(500)" json:"spRemark,omitempty" xml:"spRemark"`
	UnitUom           string  `xorm:"not null VARCHAR(20)" json:"unitUom,omitempty" xml:"unitUom"`
	ActiveStatus      int     `xorm:"default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
}

func (m *CmsStockTakeDtl) TableName() string {
	return "cms_stock_take_dtl"
}
