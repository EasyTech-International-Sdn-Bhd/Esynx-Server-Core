package models

import "github.com/easytech-international-sdn-bhd/esynx-server-core/entities"

type DebitNoteDetailsWithProduct struct {
	D *entities.CmsDebitnoteDetails
	P *entities.CmsProduct
}
