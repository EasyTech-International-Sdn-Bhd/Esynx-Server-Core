package contracts

import "github.com/easytech-international-sdn-bhd/core/entities"

type ICmsLogin interface {
	Get(agentId int64) (*entities.CmsLogin, error)
	GetByAgentCode(agentCode string) (*entities.CmsLogin, error)
	GetAll() ([]*entities.CmsLogin, error)
	InsertBatch(records []*entities.CmsLogin) error
	Update(record *entities.CmsLogin) error
	UpdateBatch(records []*entities.CmsLogin) error
	Delete(record *entities.CmsLogin) error
	DeleteBatch(records []*entities.CmsLogin) error
}
