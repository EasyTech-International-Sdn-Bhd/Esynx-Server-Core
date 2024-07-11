package entities

type CmsPdfImage struct {
	Id                 uint64 `xorm:"not null pk autoincr unique UNSIGNED BIGINT"`
	PdfHeader          string `xorm:"not null VARCHAR(200)"`
	PdfFooter          string `xorm:"not null VARCHAR(200)"`
	PdfContent         []byte `xorm:"BLOB"`
	PdfContentRow      []byte `xorm:"BLOB"`
	SalespersonId      int    `xorm:"not null default 0 INT"`
	CustCode           string `xorm:"not null default '' VARCHAR(50)"`
	DocType            string `xorm:"default 'sales' VARCHAR(100)"`
	ParentDocType      string `xorm:"default 'sales' VARCHAR(100)"`
	ActiveStatus       int    `xorm:"default 1 INT"`
	ShowInApp          int    `xorm:"default 0 INT"`
	PdfContentCust     []byte `xorm:"BLOB"`
	PdfContentBranch   []byte `xorm:"BLOB"`
	FixedContentBody   []byte `xorm:"BLOB"`
	FixedContentTop    []byte `xorm:"BLOB"`
	FixedContentBottom []byte `xorm:"BLOB"`
}

func (m *CmsPdfImage) TableName() string {
	return "cms_pdf_image"
}
