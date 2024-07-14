package contracts

import (
	"github.com/easytech-international-sdn-bhd/esynx-server-core/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/models"
	"time"
)

type ICmsDebitNote interface {
	Get(debitNoteCode string) (*entities.CmsDebitnote, error)
	GetWithCustomer(debitNoteCode string) (*models.DebitNoteWithCustomer, error)
	GetWithItems(debitNoteCode string) (*models.DebitNoteWithItems, error)
	GetByCustomer(custCode string) ([]*entities.CmsDebitnote, error)
	GetByDate(from time.Time, to time.Time) ([]*entities.CmsDebitnote, error)
	InsertMany(debitNotes []*entities.CmsDebitnote) error
	Update(debitNote *entities.CmsDebitnote) error
	UpdateMany(debitNotes []*entities.CmsDebitnote) error
	Delete(debitNote *entities.CmsDebitnote) error
	DeleteMany(debitNotes []*entities.CmsDebitnote) error
}
