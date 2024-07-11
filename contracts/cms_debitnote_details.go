package contracts

import (
	"github.com/easytech-international-sdn-bhd/core/entities"
	"github.com/easytech-international-sdn-bhd/core/models"
)

type ICmsDebitNoteDetails interface {
	Get(debitNoteCode string) ([]*entities.CmsDebitnoteDetails, error)
	GetMany(debitNoteCodes []string) ([]*entities.CmsDebitnoteDetails, error)
	GetByProductCode(productCode string) ([]*entities.CmsDebitnoteDetails, error)
	GetWithProduct(debitNoteCode string) ([]*models.DebitNoteDetailsWithProduct, error)
	InsertMany(details []*entities.CmsDebitnoteDetails) error
	Update(details *entities.CmsDebitnoteDetails) error
	UpdateMany(details []*entities.CmsDebitnoteDetails) error
	Delete(details *entities.CmsDebitnoteDetails) error
	DeleteMany(details []*entities.CmsDebitnoteDetails) error
}
