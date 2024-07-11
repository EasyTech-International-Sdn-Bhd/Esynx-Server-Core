package models

import "github.com/easytech-international-sdn-bhd/core/entities"

type DebitNoteDetailsWithProduct struct {
	D *entities.CmsDebitnoteDetails
	P *entities.CmsProduct
}
