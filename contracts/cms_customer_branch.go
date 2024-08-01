package contracts

import "github.com/easytech-international-sdn-bhd/esynx-common/entities"

type ICmsCustomerBranch interface {
	Get(branchCode string) (*entities.CmsCustomerBranch, error)
	GetByCustomerCode(custCode string) ([]*entities.CmsCustomerBranch, error)
	GetByAgentId(agentId int64) ([]*entities.CmsCustomerBranch, error)
	GetAllStatusByAgentId(agentId int64) ([]*entities.CmsCustomerBranch, error)
	InsertMany(records []*entities.CmsCustomerBranch) error
	Update(record *entities.CmsCustomerBranch) error
	UpdateMany(records []*entities.CmsCustomerBranch) error
	Delete(record *entities.CmsCustomerBranch) error
	DeleteMany(records []*entities.CmsCustomerBranch) error
}
