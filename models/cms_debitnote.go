package models

import "github.com/easytech-international-sdn-bhd/esynx-server-core/entities"

type DebitNoteWithItems struct {
	M *entities.CmsDebitnote
	D []*entities.CmsDebitnoteDetails
}

type DebitNoteWithCustomer struct {
	I *entities.CmsDebitnote
	C *entities.CmsCustomer
}
