package contracts

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"xorm.io/builder"
)

type ICmsCustomerBranch interface {
	Get(branchCode string) (*entities.CmsCustomerBranch, error)
	GetByCustomerCode(custCode string) ([]*entities.CmsCustomerBranch, error)
	GetByAgentCode(agentCode string) ([]*entities.CmsCustomerBranch, error)
	GetAllStatusByAgentCode(agentCode string) ([]*entities.CmsCustomerBranch, error)
	InsertMany(records []*entities.CmsCustomerBranch) error
	Update(record *entities.CmsCustomerBranch) error
	UpdateMany(records []*entities.CmsCustomerBranch) error
	Delete(record *entities.CmsCustomerBranch) error
	DeleteMany(records []*entities.CmsCustomerBranch) error
	Find(predicate *builder.Builder) ([]*entities.CmsCustomerBranch, error)
}
