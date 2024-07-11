package entities

type CmsCreditnoteDetails struct {
	Id           int     `xorm:"not null pk autoincr INT"`
	CnCode       string  `xorm:"unique(cn_code) VARCHAR(100)"`
	ItemCode     string  `xorm:"unique(cn_code) VARCHAR(200)"`
	ItemName     string  `xorm:"VARCHAR(200)"`
	ItemPrice    string  `xorm:"VARCHAR(200)"`
	Quantity     float64 `xorm:"default 0 DOUBLE"`
	Uom          string  `xorm:"VARCHAR(200)"`
	TotalPrice   float64 `xorm:"default 0 DOUBLE"`
	Discount     string  `xorm:"comment('0%+10+50%') VARCHAR(100)"`
	ActiveStatus int     `xorm:"default 1 INT"`
	SequenceNo   int     `xorm:"not null default 0 INT"`
	CnDtlUdf     string  `xorm:"not null JSON"`
	RefNo        string  `xorm:"unique(cn_code) VARCHAR(200)"`
	UpdatedAt    string  `xorm:"VARCHAR(20)"`
}

func (m *CmsCreditnoteDetails) TableName() string {
	return "cms_creditnote_details"
}
