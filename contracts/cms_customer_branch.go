package contracts

import "github.com/easytech-international-sdn-bhd/core/entities"

type ICmsCustomerBranch interface {
	Get(branchCode string) (*entities.CmsCustomerBranch, error)
	GetByCustomerCode(custCode string) ([]*entities.CmsCustomerBranch, error)
	GetByAgentId(agentId int64) ([]*entities.CmsCustomerBranch, error)
	GetAllStatusByAgentId(agentId int64) ([]*entities.CmsCustomerBranch, error)
	InsertBatch(records []*entities.CmsCustomerBranch) error
	Update(record *entities.CmsCustomerBranch) error
	UpdateBatch(records []*entities.CmsCustomerBranch) error
	Delete(record *entities.CmsCustomerBranch) error
	DeleteBatch(records []*entities.CmsCustomerBranch) error
}
