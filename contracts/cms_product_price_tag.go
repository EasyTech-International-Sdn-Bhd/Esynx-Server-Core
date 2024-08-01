package contracts

import "github.com/easytech-international-sdn-bhd/esynx-common/entities"

type ICmsProductPriceTag interface {
	Get(productCode string) ([]*entities.CmsProductPriceV2, error)
	GetByCustCode(productCode string, custCode string) ([]*entities.CmsProductPriceV2, error)
	GetByPriceType(productCode string, priceType string) ([]*entities.CmsProductPriceV2, error)
	InsertMany(records []*entities.CmsProductPriceV2) error
	Update(record *entities.CmsProductPriceV2) error
	UpdateMany(records []*entities.CmsProductPriceV2) error
	Delete(record *entities.CmsProductPriceV2) error
	DeleteMany(records []*entities.CmsProductPriceV2) error
}
