package entities

type CmsPdfImage struct {
	Id                 uint64 `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	PdfHeader          string `xorm:"not null VARCHAR(200)" json:"pdfHeader,omitempty" xml:"pdfHeader"`
	PdfFooter          string `xorm:"not null VARCHAR(200)" json:"pdfFooter,omitempty" xml:"pdfFooter"`
	PdfContent         []byte `xorm:"BLOB" json:"pdfContent,omitempty" xml:"pdfContent"`
	PdfContentRow      []byte `xorm:"BLOB" json:"pdfContentRow,omitempty" xml:"pdfContentRow"`
	SalespersonId      int    `xorm:"not null default 0 INT" json:"salespersonId,omitempty" xml:"salespersonId"`
	CustCode           string `xorm:"not null default '' VARCHAR(50)" json:"custCode,omitempty" xml:"custCode"`
	DocType            string `xorm:"default 'sales' VARCHAR(100)" json:"docType,omitempty" xml:"docType"`
	ParentDocType      string `xorm:"default 'sales' VARCHAR(100)" json:"parentDocType,omitempty" xml:"parentDocType"`
	ActiveStatus       int    `xorm:"default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	ShowInApp          int    `xorm:"default 0 INT" json:"showInApp,omitempty" xml:"showInApp"`
	PdfContentCust     []byte `xorm:"BLOB" json:"pdfContentCust,omitempty" xml:"pdfContentCust"`
	PdfContentBranch   []byte `xorm:"BLOB" json:"pdfContentBranch,omitempty" xml:"pdfContentBranch"`
	FixedContentBody   []byte `xorm:"BLOB" json:"fixedContentBody,omitempty" xml:"fixedContentBody"`
	FixedContentTop    []byte `xorm:"BLOB" json:"fixedContentTop,omitempty" xml:"fixedContentTop"`
	FixedContentBottom []byte `xorm:"BLOB" json:"fixedContentBottom,omitempty" xml:"fixedContentBottom"`
}

func (m *CmsPdfImage) TableName() string {
	return "cms_pdf_image"
}
