package contracts

import (
	"github.com/easytech-international-sdn-bhd/core/entities"
	"github.com/easytech-international-sdn-bhd/core/models"
)

type ICmsCustomer interface {
	Get(custCode string) (*entities.CmsCustomer, error)
	GetMany(custCodes []string) ([]*entities.CmsCustomer, error)
	GetWithBranches(custCode string) (*models.CustomerWithBranches, error)
	GetWithAgents(custCode string) (*models.CustomerWithAgents, error)
	GetCustomerById(custId string) (*entities.CmsCustomer, error)
	GetAllStatusByAgentId(agentId int64) ([]*entities.CmsCustomer, error)
	SearchByNameOrCode(predicate string) ([]*entities.CmsCustomer, error)
	InsertBatch(records []entities.CmsCustomer) error
	Update(customer *entities.CmsCustomer) error
	UpdateBatch(customers []*entities.CmsCustomer) error
	Delete(customer *entities.CmsCustomer) error
	DeleteBatch(customer []*entities.CmsCustomer) error
}
