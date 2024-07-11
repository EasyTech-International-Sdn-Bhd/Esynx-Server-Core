package entities

type CmsStockTransferDtl struct {
	Id                int     `xorm:"not null pk autoincr unique(unq) INT"`
	StCode            string  `xorm:"not null unique(unq) VARCHAR(200)"`
	ProductCode       string  `xorm:"not null VARCHAR(200)"`
	ProductName       string  `xorm:"VARCHAR(200)"`
	FromLocation      string  `xorm:"comment('sqlacc item basis') VARCHAR(200)"`
	ToLocation        string  `xorm:"comment('sqlacc item basis') VARCHAR(200)"`
	Quantity          int     `xorm:"INT"`
	UnitUom           string  `xorm:"VARCHAR(200)"`
	UnitPrice         float64 `xorm:"comment('unit cost') DOUBLE"`
	SubTotal          float64 `xorm:"DOUBLE"`
	SalespersonRemark string  `xorm:"VARCHAR(500)"`
	CancelStatus      int     `xorm:"default 0 INT"`
}

func (m *CmsStockTransferDtl) TableName() string {
	return "cms_stock_transfer_dtl"
}
