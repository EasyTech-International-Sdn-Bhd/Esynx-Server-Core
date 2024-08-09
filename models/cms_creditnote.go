package models

import "github.com/easytech-international-sdn-bhd/esynx-common/entities"

type CreditNoteWithItems struct {
	M *entities.CmsCreditnote
	D []*entities.CmsCreditnoteDetails
}

type CreditNoteSalesWithItems struct {
	M *entities.CmsCreditnoteSales
	D []*entities.CmsCreditnoteDetails
}

type CreditNoteWithCustomer struct {
	I *entities.CmsCreditnote
	C *entities.CmsCustomer
}

type CreditNoteSalesWithCustomer struct {
	I *entities.CmsCreditnoteSales
	C *entities.CmsCustomer
}
