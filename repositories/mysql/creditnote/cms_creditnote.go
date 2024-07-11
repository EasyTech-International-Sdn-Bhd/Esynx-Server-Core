package creditnote

import (
	"github.com/easytech-international-sdn-bhd/core/entities"
	"github.com/easytech-international-sdn-bhd/core/models"
	"xorm.io/xorm"
)

type CmsCreditNoteRepository struct {
	db *xorm.Engine
}

func NewCmsCreditNoteRepository(db *xorm.Engine) *CmsCreditNoteRepository {
	return &CmsCreditNoteRepository{
		db: db,
	}
}

func Get(creditNoteCode string) (*entities.CmsCreditnote, error) {

}

func GetWithCustomer(creditNoteCode string) (*models.CreditNoteWithCustomer, error) {

}
