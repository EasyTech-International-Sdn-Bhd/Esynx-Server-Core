package contracts

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"xorm.io/builder"
)

type ICmsWarehouseStock interface {
	Get(productCode string) ([]*entities.CmsWarehouseStock, error)
	InsertMany(records []*entities.CmsWarehouseStock) error
	Update(record *entities.CmsWarehouseStock) error
	UpdateMany(records []*entities.CmsWarehouseStock) error
	Delete(record *entities.CmsWarehouseStock) error
	DeleteMany(records []*entities.CmsWarehouseStock) error
	Find(predicate *builder.Builder) ([]*entities.CmsWarehouseStock, error)
}
