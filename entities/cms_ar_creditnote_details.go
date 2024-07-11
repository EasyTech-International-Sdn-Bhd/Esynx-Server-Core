package entities

type CmsArCreditnoteDetails struct {
	Id              int     `xorm:"not null pk autoincr INT"`
	CnCode          string  `xorm:"not null VARCHAR(20)"`
	AccNo           string  `xorm:"not null VARCHAR(50)"`
	Description     string  `xorm:"default '' VARCHAR(150)"`
	NetAmount       float64 `xorm:"default 0 DOUBLE"`
	UnappliedAmount float64 `xorm:"default 0 DOUBLE"`
	KnockoffAmount  float64 `xorm:"default 0 DOUBLE"`
	Tax             float64 `xorm:"default 0 DOUBLE"`
	SubTotal        float64 `xorm:"default 0 DOUBLE"`
	RefNo           string  `xorm:"not null unique VARCHAR(50)"`
	ActiveStatus    int     `xorm:"not null default 1 INT"`
}

func (m *CmsArCreditnoteDetails) TableName() string {
	return "cms_ar_creditnote_details"
}
