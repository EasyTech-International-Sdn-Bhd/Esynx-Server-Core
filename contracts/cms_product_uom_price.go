package contracts

import "github.com/easytech-international-sdn-bhd/esynx-common/entities"

type ICmsProductUomPrice interface {
	Get(productCode string) ([]*entities.CmsProductUomPriceV2, error)
	InsertMany(records []*entities.CmsProductUomPriceV2) error
	Update(record *entities.CmsProductUomPriceV2) error
	UpdateMany(records []*entities.CmsProductUomPriceV2) error
	Delete(record *entities.CmsProductUomPriceV2) error
	DeleteMany(records []*entities.CmsProductUomPriceV2) error
}
