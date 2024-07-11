package models

import "github.com/easytech-international-sdn-bhd/core/entities"

type CreditNoteWithItems struct {
	M *entities.CmsCreditnote
	D []*entities.CmsCreditnoteDetails
}

type CreditNoteWithCustomer struct {
	I *entities.CmsCreditnote
	C *entities.CmsCustomer
}
