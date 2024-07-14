package contracts

import (
	"github.com/easytech-international-sdn-bhd/esynx-server-core/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/models"
)

type ICmsCustomer interface {
	Get(custCode string) (*entities.CmsCustomer, error)
	GetMany(custCodes []string) ([]*entities.CmsCustomer, error)
	GetWithBranches(custCode string) (*models.CustomerWithBranches, error)
	GetWithAgent(custCode string) (*models.CustomerWithAgent, error)
	GetCustomerById(custId string) (*entities.CmsCustomer, error)
	GetAllStatusByAgentId(agentId int64) ([]*entities.CmsCustomer, error)
	SearchByNameOrCode(predicate string) ([]*entities.CmsCustomer, error)
	InsertMany(records []*entities.CmsCustomer) error
	Update(customer *entities.CmsCustomer) error
	UpdateMany(customers []*entities.CmsCustomer) error
	Delete(customer *entities.CmsCustomer) error
	DeleteMany(customer []*entities.CmsCustomer) error
}
