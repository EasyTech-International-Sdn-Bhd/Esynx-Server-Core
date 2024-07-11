package contracts

import "github.com/easytech-international-sdn-bhd/core/entities"

type ICmsCustomerSalesperson interface {
	GetByAgentId(agentId int64) ([]*entities.CmsCustomerSalesperson, error)
	GetByCustomerId(custId int64) (*entities.CmsCustomerSalesperson, error)
	GetAgentByCustId(custId int64) (*entities.CmsLogin, error)
}
