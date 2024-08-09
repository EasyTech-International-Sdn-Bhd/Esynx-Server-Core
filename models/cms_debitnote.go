package models

import "github.com/easytech-international-sdn-bhd/esynx-common/entities"

type DebitNoteWithItems struct {
	M *entities.CmsDebitnote
	D []*entities.CmsDebitnoteDetails
}

type DebitNoteWithCustomer struct {
	I *entities.CmsDebitnote
	C *entities.CmsCustomer
}

type DebitNoteSalesWithItems struct {
	M *entities.CmsDebitnoteSales
	D []*entities.CmsDebitnoteDetails
}

type DebitNoteSalesWithCustomer struct {
	I *entities.CmsDebitnoteSales
	C *entities.CmsCustomer
}
