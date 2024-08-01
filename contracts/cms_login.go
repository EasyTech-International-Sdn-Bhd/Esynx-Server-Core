package contracts

import "github.com/easytech-international-sdn-bhd/esynx-common/entities"

type ICmsLogin interface {
	Get(agentId int64) (*entities.CmsLogin, error)
	GetByAgentCode(agentCode string) (*entities.CmsLogin, error)
	GetAll() ([]*entities.CmsLogin, error)
	InsertMany(records []*entities.CmsLogin) error
	Update(record *entities.CmsLogin) error
	UpdateMany(records []*entities.CmsLogin) error
	Delete(record *entities.CmsLogin) error
	DeleteMany(records []*entities.CmsLogin) error
}
