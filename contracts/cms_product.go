package contracts

import (
	"github.com/easytech-international-sdn-bhd/core/entities"
	"github.com/easytech-international-sdn-bhd/core/models"
)

type ICmsProduct interface {
	Get(productCode string) (*entities.CmsProduct, error)
	GetMany(productCodes []string) ([]*entities.CmsProduct, error)
	GetWithDetails(productCode string) (*models.ProductWithDetails, error)
	Search(predicate string) ([]*entities.CmsProduct, error)
	InsertMany(records []*entities.CmsProduct) error
	Update(record *entities.CmsProduct) error
	UpdateMany(records []*entities.CmsProduct) error
	Delete(record *entities.CmsProduct) error
	DeleteMany(records []*entities.CmsProduct) error
}
