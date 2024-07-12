package agent

import (
	"github.com/easytech-international-sdn-bhd/core/entities"
	iterator "github.com/ledongthuc/goterators"
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

func (r *CmsCustomerSalespersonRepository) InsertMany(records []*entities.CmsCustomerSalesperson) error {
	_, err := r.db.Insert(iterator.Map(records, func(item *entities.CmsCustomerSalesperson) *entities.CmsCustomerSalesperson {
		item.Validate()
		item.ToUpdate()
		return item
	}))
	if err != nil {
		return err
	}
	return nil
}

func (r *CmsCustomerSalespersonRepository) Update(record *entities.CmsCustomerSalesperson) error {
	_, err := r.db.Where("salesperson_customer_id = ?").Update(record.SalespersonCustomerId)
	if err != nil {
		return err
	}
	return nil
}

func (r *CmsCustomerSalespersonRepository) UpdateMany(records []*entities.CmsCustomerSalesperson) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}
	var sessionErr error
	rollback := false
	for _, record := range records {
		record.Validate()
		record.ToUpdate()
		_, err = session.Where("salesperson_customer_id = ?", record.SalespersonCustomerId).Update(record)
		if err != nil {
			rollback = true
			sessionErr = err
			break
		}
	}
	if rollback {
		err := session.Rollback()
		if err != nil {
			return err
		}
		return sessionErr
	}
	err = session.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (r *CmsCustomerSalespersonRepository) Delete(record *entities.CmsCustomerSalesperson) error {
	record.ActiveStatus = 0
	record.ToUpdate()
	return r.Update(record)
}

func (r *CmsCustomerSalespersonRepository) DeleteMany(records []*entities.CmsCustomerSalesperson) error {
	for _, record := range records {
		record.ActiveStatus = 0
		record.ToUpdate()
	}
	return r.UpdateMany(records)
}
