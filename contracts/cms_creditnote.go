package contracts

import (
	"github.com/easytech-international-sdn-bhd/core/entities"
	"github.com/easytech-international-sdn-bhd/core/models"
	"time"
)

type ICmsCreditNote interface {
	Get(creditNoteCode string) (*entities.CmsCreditnote, error)
	GetWithCustomer(creditNoteCode string) (*models.CreditNoteWithCustomer, error)
	GetWithItems(creditNoteCode string) (*models.CreditNoteWithItems, error)
	GetByCustomer(custCode string) ([]*entities.CmsCreditnote, error)
	GetByDate(from time.Time, to time.Time) ([]*entities.CmsCreditnote, error)
	InsertMany(creditNotes []*entities.CmsCreditnote) error
	Update(creditNote *entities.CmsCreditnote) error
	UpdateMany(creditNotes []*entities.CmsCreditnote) error
	Delete(creditNote *entities.CmsCreditnote) error
	DeleteMany(creditNotes []*entities.CmsCreditnote) error
}
