package customer

import (
	"github.com/easytech-international-sdn-bhd/core/entities"
	"github.com/easytech-international-sdn-bhd/core/models"
	"github.com/easytech-international-sdn-bhd/core/repositories/mysql/agent"
	iterator "github.com/ledongthuc/goterators"
	"strings"
	"xorm.io/builder"
	_ "xorm.io/builder"
	"xorm.io/xorm"
)

type CmsCustomerRepository struct {
	db *xorm.Engine
	b  *CmsCustomerBranchRepository
	s  *agent.CmsCustomerSalespersonRepository
}

func NewCmsCustomerRepository(db *xorm.Engine) *CmsCustomerRepository {
	return &CmsCustomerRepository{
		db: db,
		b:  NewCmsCustomerBranchRepository(db),
		s:  agent.NewCmsCustomerSalespersonRepository(db),
	}
}

func (r *CmsCustomerRepository) Get(custCode string) (*entities.CmsCustomer, error) {
	var customer entities.CmsCustomer
	has, err := r.db.Where("cust_code=?", custCode).Get(&customer)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return &customer, nil
}

func (r *CmsCustomerRepository) GetMany(custCodes []string) ([]*entities.CmsCustomer, error) {
	var customers []*entities.CmsCustomer
	err := r.db.In("cust_code", custCodes).Where("customer_status = ?", 1).Find(&customers)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r *CmsCustomerRepository) GetWithBranches(custCode string) (*models.CustomerWithBranches, error) {
	customer, err := r.Get(custCode)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, nil
	}
	branches, err := r.b.GetByCustomerCode(custCode)
	if err != nil {
		return nil, err
	}
	if branches == nil {
		return nil, nil
	}
	return &models.CustomerWithBranches{
		C: customer,
		B: branches,
	}, nil
}

func (r *CmsCustomerRepository) GetWithAgent(custCode string) (*models.CustomerWithAgent, error) {
	customer, err := r.Get(custCode)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, nil
	}
	agentRec, err := r.s.GetAgentByCustId(int64(customer.CustId))
	if err != nil {
		return nil, err
	}
	return &models.CustomerWithAgent{
		C: customer,
		A: agentRec,
	}, nil
}

func (r *CmsCustomerRepository) GetAllStatusByAgentId(agentId int64) ([]*entities.CmsCustomer, error) {
	result, err := r.s.GetByAgentId(agentId)
	if err != nil {
		return nil, err
	}

	var customerIds []int
	for _, record := range result {
		customerIds = append(customerIds, record.CustomerId)
	}

	var records []*entities.CmsCustomer
	err = r.db.In("cust_id", customerIds).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsCustomerRepository) SearchByNameOrCode(predicate string) ([]*entities.CmsCustomer, error) {
	var records []*entities.CmsCustomer
	tokens := strings.Split(predicate, " ")
	var where []builder.Cond
	for _, token := range tokens {
		where = append(where, builder.Like{"cust_code", token})
		where = append(where, builder.Like{"cust_company_name", token})
	}
	err := r.db.Where(builder.Or(where...)).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsCustomerRepository) GetCustomerById(custId string) (*entities.CmsCustomer, error) {
	var record entities.CmsCustomer
	has, err := r.db.Where("cust_id=?", custId).Get(&record)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return &record, nil
}

func (r *CmsCustomerRepository) InsertMany(records []*entities.CmsCustomer) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}

	_, err = session.Insert(iterator.Map(records, func(item *entities.CmsCustomer) *entities.CmsCustomer {
		item.Validate()
		item.ToUpdate()
		return item
	}))
	if err != nil {
		err := session.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	err = session.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (r *CmsCustomerRepository) Update(customer *entities.CmsCustomer) error {
	_, err := r.db.Where("cust_code = ?", customer.CustCode).Update(customer)
	if err != nil {
		return err
	}
	return nil
}

func (r *CmsCustomerRepository) UpdateMany(customers []*entities.CmsCustomer) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}
	var sessionErr error
	rollback := false
	for _, customer := range customers {
		customer.Validate()
		customer.ToUpdate()
		_, err = session.Where("cust_code = ?", customer.CustCode).Update(customer)
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

func (r *CmsCustomerRepository) Delete(customer *entities.CmsCustomer) error {
	customer.CustomerStatus = 0
	customer.ToUpdate()
	return r.Update(customer)
}

func (r *CmsCustomerRepository) DeleteMany(customers []*entities.CmsCustomer) error {
	for _, customer := range customers {
		customer.CustomerStatus = 0
		customer.ToUpdate()
	}
	return r.UpdateMany(customers)
}
