package entities

type CmsArCreditnoteDetails struct {
	Id              uint64  `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	CnCode          string  `xorm:"not null VARCHAR(20)" json:"cnCode,omitempty" xml:"cnCode"`
	AccNo           string  `xorm:"not null VARCHAR(50)" json:"accNo,omitempty" xml:"accNo"`
	Description     string  `xorm:"default '' VARCHAR(150)" json:"description,omitempty" xml:"description"`
	NetAmount       float64 `xorm:"default 0 DOUBLE" json:"netAmount,omitempty" xml:"netAmount"`
	UnappliedAmount float64 `xorm:"default 0 DOUBLE" json:"unappliedAmount,omitempty" xml:"unappliedAmount"`
	KnockoffAmount  float64 `xorm:"default 0 DOUBLE" json:"knockoffAmount,omitempty" xml:"knockoffAmount"`
	Tax             float64 `xorm:"default 0 DOUBLE" json:"tax,omitempty" xml:"tax"`
	SubTotal        float64 `xorm:"default 0 DOUBLE" json:"subTotal,omitempty" xml:"subTotal"`
	RefNo           string  `xorm:"not null unique VARCHAR(50)" json:"refNo,omitempty" xml:"refNo"`
	ActiveStatus    int     `xorm:"not null default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
}

func (m *CmsArCreditnoteDetails) TableName() string {
	return "cms_ar_creditnote_details"
}
