package contracts

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"xorm.io/builder"
)

type ICmsProductAtch interface {
	Get(productCode string) ([]*entities.CmsProductAtch, error)
	InsertMany(records []*entities.CmsProductAtch) error
	Update(record *entities.CmsProductAtch) error
	UpdateMany(records []*entities.CmsProductAtch) error
	Delete(record *entities.CmsProductAtch) error
	DeleteMany(records []*entities.CmsProductAtch) error
	Find(predicate *builder.Builder) ([]*entities.CmsProductAtch, error)
}
