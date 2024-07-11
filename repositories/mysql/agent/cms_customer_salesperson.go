package agent

import (
	"github.com/easytech-international-sdn-bhd/core/entities"
	"xorm.io/xorm"
)

type CmsCustomerSalespersonRepository struct {
	db *xorm.Engine
	l  *CmsLoginRepository
}

func NewCmsCustomerSalespersonRepository(db *xorm.Engine) *CmsCustomerSalespersonRepository {
	return &CmsCustomerSalespersonRepository{
		db: db,
		l:  NewCmsLoginRepository(db),
	}
}

func (r *CmsCustomerSalespersonRepository) GetByAgentId(agentId int64) ([]*entities.CmsCustomerSalesperson, error) {
	var record []*entities.CmsCustomerSalesperson
	err := r.db.Where("salesperson_id = ? AND active_status = ?", agentId, 1).Find(&record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (r *CmsCustomerSalespersonRepository) GetByCustomerId(custId int64) (*entities.CmsCustomerSalesperson, error) {
	var record entities.CmsCustomerSalesperson
	has, err := r.db.Where("customer_id = ? AND active_status = ?", custId, 1).Get(&record)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return &record, nil
}

func (r *CmsCustomerSalespersonRepository) GetAgentByCustId(custId int64) (*entities.CmsLogin, error) {
	a, err := r.GetByCustomerId(custId)
	if err != nil {
		return nil, err
	}
	c, err := r.l.Get(int64(a.SalespersonId))
	if err != nil {
		return nil, err
	}
	return c, nil
}
