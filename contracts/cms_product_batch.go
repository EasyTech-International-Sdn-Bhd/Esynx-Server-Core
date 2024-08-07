package contracts

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"xorm.io/builder"
)

type ICmsProductBatch interface {
	Get(productCode string) ([]*entities.CmsProductBatch, error)
	GetByWarehouse(productCode string, warehouse string) ([]*entities.CmsProductBatch, error)
	InsertMany(records []*entities.CmsProductBatch) error
	Update(record *entities.CmsProductBatch) error
	UpdateMany(records []*entities.CmsProductBatch) error
	Delete(record *entities.CmsProductBatch) error
	DeleteMany(records []*entities.CmsProductBatch) error
	Find(predicate *builder.Builder) ([]*entities.CmsProductBatch, error)
}
