package contracts

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/models"
	"xorm.io/builder"
)

type ICmsDebitNoteDetails interface {
	Get(debitNoteCode string) ([]*entities.CmsDebitnoteDetails, error)
	GetMany(debitNoteCodes []string) ([]*entities.CmsDebitnoteDetails, error)
	GetWithProduct(debitNoteCode string) ([]*models.DebitNoteDetailsWithProduct, error)
	InsertMany(details []*entities.CmsDebitnoteDetails) error
	Update(details *entities.CmsDebitnoteDetails) error
	UpdateMany(details []*entities.CmsDebitnoteDetails) error
	Delete(details *entities.CmsDebitnoteDetails) error
	DeleteMany(details []*entities.CmsDebitnoteDetails) error
	DeleteByAny(predicate *builder.Builder) ([]*entities.CmsDebitnoteDetails, error)
	Find(predicate *builder.Builder) ([]*entities.CmsDebitnoteDetails, error)
}
