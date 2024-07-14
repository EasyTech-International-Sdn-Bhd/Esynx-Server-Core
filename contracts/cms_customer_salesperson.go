package contracts

import "github.com/easytech-international-sdn-bhd/esynx-server-core/entities"

type ICmsCustomerSalesperson interface {
	GetByAgentId(agentId int64) ([]*entities.CmsCustomerSalesperson, error)
	GetByCustomerId(custId int64) (*entities.CmsCustomerSalesperson, error)
	GetAgentByCustId(custId int64) (*entities.CmsLogin, error)
	InsertMany(records []*entities.CmsCustomerSalesperson) error
	Update(record *entities.CmsCustomerSalesperson) error
	UpdateMany(records []*entities.CmsCustomerSalesperson) error
	Delete(record *entities.CmsCustomerSalesperson) error
	DeleteMany(records []*entities.CmsCustomerSalesperson) error
}
