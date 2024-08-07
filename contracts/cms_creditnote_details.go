package contracts

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/models"
	"xorm.io/builder"
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
	Find(predicate *builder.Builder) ([]*entities.CmsCreditnoteDetails, error)
}
