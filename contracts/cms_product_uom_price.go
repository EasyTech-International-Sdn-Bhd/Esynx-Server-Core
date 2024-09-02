package contracts

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"xorm.io/builder"
)

type ICmsProductUomPrice interface {
	Get(productCode string) ([]*entities.CmsProductUomPrice, error)
	InsertMany(records []*entities.CmsProductUomPrice) error
	Update(record *entities.CmsProductUomPrice) error
	UpdateMany(records []*entities.CmsProductUomPrice) error
	Delete(record *entities.CmsProductUomPrice) error
	DeleteMany(records []*entities.CmsProductUomPrice) error
	Find(predicate *builder.Builder) ([]*entities.CmsProductUomPrice, error)
}
