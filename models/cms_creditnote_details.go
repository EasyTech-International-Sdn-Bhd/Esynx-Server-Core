package models

import "github.com/easytech-international-sdn-bhd/esynx-server-core/entities"

type CreditNoteDetailsWithProduct struct {
	D *entities.CmsCreditnoteDetails
	P *entities.CmsProduct
}
