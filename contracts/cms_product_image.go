package contracts

import "github.com/easytech-international-sdn-bhd/esynx-common/entities"

type ICmsProductImage interface {
	Get(productId string) ([]*entities.CmsProductImage, error)
	GetByProductCode(productCode string) ([]*entities.CmsProductImage, error)
	InsertMany(records []*entities.CmsProductImage) error
	Update(record *entities.CmsProductImage) error
	UpdateMany(records []*entities.CmsProductImage) error
	Delete(record *entities.CmsProductImage) error
	DeleteMany(records []*entities.CmsProductImage) error
}
