package contracts

import "github.com/easytech-international-sdn-bhd/esynx-common/entities"

type ICmsCustomerAgent interface {
	GetByAgentCode(agentCode string) ([]*entities.CmsCustomerAgent, error)
	GetByCustomerCode(custCode string) ([]*entities.CmsCustomerAgent, error)
	GetByManyCustomers(custCodes []string) ([]*entities.CmsCustomerAgent, error)
	GetByManyAgents(agentCodes []string) ([]*entities.CmsCustomerAgent, error)
	GetAgentsByCustCode(custCode string) ([]*entities.CmsLogin, error)
	InsertMany(records []*entities.CmsCustomerAgent) error
	Update(record *entities.CmsCustomerAgent) error
	UpdateMany(records []*entities.CmsCustomerAgent) error
	Delete(record *entities.CmsCustomerAgent) error
	DeleteMany(records []*entities.CmsCustomerAgent) error
}
