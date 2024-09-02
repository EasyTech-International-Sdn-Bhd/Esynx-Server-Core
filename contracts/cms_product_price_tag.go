package contracts

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"xorm.io/builder"
)

type ICmsProductPriceTag interface {
	Get(productCode string) ([]*entities.CmsProductPrice, error)
	GetByCustCode(productCode string, custCode string) ([]*entities.CmsProductPrice, error)
	GetByPriceType(productCode string, priceType string) ([]*entities.CmsProductPrice, error)
	InsertMany(records []*entities.CmsProductPrice) error
	Update(record *entities.CmsProductPrice) error
	UpdateMany(records []*entities.CmsProductPrice) error
	Delete(record *entities.CmsProductPrice) error
	DeleteMany(records []*entities.CmsProductPrice) error
	Find(predicate *builder.Builder) ([]*entities.CmsProductPrice, error)
}
