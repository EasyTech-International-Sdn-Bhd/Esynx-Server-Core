package entities

type CmsDebitnoteDetails struct {
	Id         int    `xorm:"not null pk autoincr INT"`
	DnCode     string `xorm:"VARCHAR(100)"`
	ItemCode   string `xorm:"VARCHAR(200)"`
	ItemName   string `xorm:"VARCHAR(200)"`
	ItemPrice  string `xorm:"VARCHAR(200)"`
	Quantity   string `xorm:"VARCHAR(200)"`
	Uom        string `xorm:"VARCHAR(200)"`
	TotalPrice string `xorm:"VARCHAR(200)"`
	Discount   string `xorm:"comment('0%+10+50%') VARCHAR(100)"`
	RefNo      string `xorm:"unique VARCHAR(200)"`
}

func (m *CmsDebitnoteDetails) TableName() string {
	return "cms_debitnote_details"
}
