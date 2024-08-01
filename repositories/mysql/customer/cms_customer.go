package customer

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/models"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/mysql/agent"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"strings"
	"xorm.io/builder"
	_ "xorm.io/builder"
	"xorm.io/xorm"
)

// CmsCustomerRepository is a repository for managing CMS customers in the database.
// It provides methods for retrieving and manipulating customer data.
type CmsCustomerRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
	b     *CmsCustomerBranchRepository
	s     *agent.CmsCustomerSalespersonRepository
}

// NewCmsCustomerRepository is a function that creates a new instance of CmsCustomerRepository.
// It takes an option parameter of type *contracts.IRepository and returns a pointer to CmsCustomerRepository.
func NewCmsCustomerRepository(option *contracts.IRepository) *CmsCustomerRepository {
	return &CmsCustomerRepository{
		db:    option.Db,
		audit: option.Audit,
		b:     NewCmsCustomerBranchRepository(option),
		s:     agent.NewCmsCustomerSalespersonRepository(option),
	}
}

// Get retrieves a CmsCustomer entity by the provided custCode.
// It queries the database with the custCode and stores the result in the customer variable.
// If an error occurs during the query, it returns nil and the error.
// If the customer does not exist in the database, it returns nil and nil for the error.
// Otherwise, it returns a pointer to the customer and nil for the error.
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

// GetMany retrieves multiple CmsCustomer entities based on the provided customer codes.
// It filters the customers by customer code and customer status using the db connection.
// If any error occurred during the retrieval, it returns nil and the error.
// Otherwise, it returns the retrieved customers and nil error.
func (r *CmsCustomerRepository) GetMany(custCodes []string) ([]*entities.CmsCustomer, error) {
	var customers []*entities.CmsCustomer
	err := r.db.In("cust_code", custCodes).Where("customer_status = ?", 1).Find(&customers)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

// GetWithBranches retrieves a customer with its branches by the customer code.
// It calls the Get method to retrieve the customer by the customer code.
// If the customer is not found, returns nil, nil.
// Otherwise, it calls the GetByCustomerCode method to retrieve the branches by the customer code.
// If the branches are not found, returns nil, nil.
// Finally, it returns a CustomerWithBranches object containing the customer and branches.
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

// GetWithAgent retrieves a customer along with its associated agent record.
// If the customer does not exist, it returns nil.
// If there is an error during the retrieval process, it returns the error.
// The returned CustomerWithAgent struct contains the customer and agent information.
//
// Signature:
// func (r *CmsCustomerRepository) GetWithAgent(custCode string) (*models.CustomerWithAgent, error)
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

// GetAllStatusByAgentId retrieves all customer records with the given agent ID.
//
// It first calls the GetByAgentId method of the CmsCustomerSalespersonRepository to get the customer IDs
// associated with the given agent ID. Then, it uses the obtained customer IDs to fetch the corresponding
// customer records from the database using the In operator.
//
// If successful, it returns a slice of *entities.CmsCustomer representing the customer records and nil error.
// If an error occurs during the retrieval process, it returns nil and the corresponding error.
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

// SearchByNameOrCode searches for customers by name or code based on the given predicate.
// It splits the predicate into tokens and checks for a match in the cust_code and cust_company_name fields.
// Returns a list of matched customers or an error if retrieval fails.
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

// GetCustomerById retrieves a customer record from the database based on the specified customer ID.
//
// If the customer ID is not found, nil will be returned along with nil error.
// If there is an error during the retrieval process, nil customer record will be returned along with the error.
//
// Parameters:
// - custId: The customer ID used to identify the customer record.
//
// Returns:
// - *entities.CmsCustomer: The retrieved customer record.
// - error: The error occurred during the retrieval process, if any.
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

// InsertMany inserts multiple records into the CmsCustomerRepository database table.
// It takes a slice of *entities.CmsCustomer records as parameter and maps them to a new slice.
// The new slice is then passed to the Insert method of the database engine.
// If there is an error during insertion, it is returned. Otherwise, the method logs the "INSERT" operation and returns nil.
func (r *CmsCustomerRepository) InsertMany(records []*entities.CmsCustomer) error {
	_, err := r.db.Insert(iterator.Map(records, func(item *entities.CmsCustomer) *entities.CmsCustomer {
		return item
	}))
	if err != nil {
		return err
	}

	r.log("INSERT", records)

	return nil
}

// Update updates a customer record in the CMS Customer Repository.
// It takes a pointer to a CmsCustomer object as a parameter and returns an error.
// The function updates the record in the repository using the `db` connection and the provided customer object.
// If the update operation fails, an error is returned.
// After updating the record, the function logs the "UPDATE" operation with the updated customer object.
// The log is sent to the `audit` object for storing.
func (r *CmsCustomerRepository) Update(customer *entities.CmsCustomer) error {
	_, err := r.db.Where("cust_code = ?", customer.CustCode).Update(customer)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsCustomer{customer})

	return nil
}

// UpdateMany updates multiple customer records in the database.
// It takes a slice of customer entities as input and updates each record based on the cust_code field.
// If any update fails, it rolls back the transaction and returns the error.
// Otherwise, it commits the transaction and logs the update operation.
// This method returns an error if the transaction fails or encounters any other error during the update process.
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

	r.log("UPDATE", customers)

	return nil
}

// Delete sets the customer's `CustomerStatus` to 0 and calls the `Update` method.
// This effectively marks the customer as deleted in the database.
//
// Parameters:
// - customer: A pointer to the `entities.CmsCustomer` object representing the customer to delete.
//
// Returns:
// - error: An error if any occurred during the deletion process.
func (r *CmsCustomerRepository) Delete(customer *entities.CmsCustomer) error {
	customer.CustomerStatus = 0
	return r.Update(customer)
}

// DeleteMany sets the CustomerStatus field to 0 for each customer in the provided slice,
// and then calls the UpdateMany method to update the customers in the database.
func (r *CmsCustomerRepository) DeleteMany(customers []*entities.CmsCustomer) error {
	for _, customer := range customers {
		customer.CustomerStatus = 0
	}
	return r.UpdateMany(customers)
}

// log logs the given operation and payload to the audit log.
func (r *CmsCustomerRepository) log(op string, payload []*entities.CmsCustomer) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsCustomer) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      item.CustCode,
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
