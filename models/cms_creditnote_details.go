package models

import "github.com/easytech-international-sdn-bhd/esynx-common/entities"

type CreditNoteDetailsWithProduct struct {
	D *entities.CmsCreditnoteDetails
	P *entities.CmsProduct
}
