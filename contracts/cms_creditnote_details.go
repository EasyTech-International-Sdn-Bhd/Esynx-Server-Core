package contracts

import (
	"github.com/easytech-international-sdn-bhd/esynx-server-core/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/models"
)

type ICmsCreditNoteDetails interface {
	Get(creditNoteCode string) ([]*entities.CmsCreditnoteDetails, error)
	GetMany(creditNoteCodes []string) ([]*entities.CmsCreditnoteDetails, error)
	GetWithProduct(creditNoteCode string) ([]*models.CreditNoteDetailsWithProduct, error)
	InsertMany(details []*entities.CmsCreditnoteDetails) error
	Update(details *entities.CmsCreditnoteDetails) error
	UpdateMany(details []*entities.CmsCreditnoteDetails) error
	Delete(details *entities.CmsCreditnoteDetails) error
	DeleteMany(details []*entities.CmsCreditnoteDetails) error
}
