package agent

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"strconv"
	"xorm.io/builder"
	"xorm.io/xorm"
)

// CmsCustomerAgentRepository is a repository for managing CMS customer salespersons.
// It has a database connection, an audit log, and an instance of CmsLoginRepository.
// The repository allows for querying and manipulating CMS customer salesperson data.
// Usage examples: creating a new CmsCustomerAgentRepository, querying salespersons by agent ID, customer ID, and accessing the corresponding agent login information.
type CmsCustomerAgentRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
	l     *CmsLoginRepository
}

// NewCmsCustomerAgentRepository returns a new instance of CmsCustomerAgentRepository with the given IRepository option.
// The option should have a valid db connection and audit log implementation.
func NewCmsCustomerAgentRepository(option *contracts.IRepository) *CmsCustomerAgentRepository {
	return &CmsCustomerAgentRepository{
		db:    option.Db,
		audit: option.Audit,
		l:     NewCmsLoginRepository(option),
	}
}

func (r *CmsCustomerAgentRepository) GetByManyCustomers(custCodes []string) ([]*entities.CmsCustomerAgent, error) {
	var res []*entities.CmsCustomerAgent
	var t entities.CmsCustomerAgent
	err := r.db.SQL(builder.Select("*").From(t.TableName()).Where(builder.In("cust_code", custCodes))).Find(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *CmsCustomerAgentRepository) GetByManyAgents(agentCodes []string) ([]*entities.CmsCustomerAgent, error) {
	var res []*entities.CmsCustomerAgent
	var t entities.CmsCustomerAgent
	err := r.db.SQL(builder.Select("*").From(t.TableName()).Where(builder.In("agent_code", agentCodes))).Find(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetByAgentCode retrieves customer salesperson records by agent ID.
//
// It queries the database for customer salesperson records where the salesperson ID
// matches the given agent ID and the active status is set to 1 (active).
//
// If successful, it returns a slice of *entities.CmsCustomerAgent representing the records
// and a nil error. If an error occurs during the retrieval process, it returns nil and the corresponding error.
func (r *CmsCustomerAgentRepository) GetByAgentCode(agentId string) ([]*entities.CmsCustomerAgent, error) {
	var record []*entities.CmsCustomerAgent
	err := r.db.Where("agent_code = ? AND active_status = ?", agentId, 1).Find(&record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// GetByCustomerCode retrieves a record from the database based on the customer ID.
// It returns a pointer to the retrieved record of type CmsCustomerAgent and an error.
// If the record exists, the returned error is nil. If the record does not exist, both the record and the error are nil.
func (r *CmsCustomerAgentRepository) GetByCustomerCode(custCode string) ([]*entities.CmsCustomerAgent, error) {
	var record []*entities.CmsCustomerAgent
	err := r.db.Where("cust_code = ? AND active_status = ?", custCode, 1).Find(&record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// GetAgentsByCustCode returns the agent for a given customer ID.
// It retrieves the associated CmsCustomerAgent record using GetByCustomerId,
// and then retrieves the CmsLogin record using the salesperson ID from the CmsCustomerAgent record.
// If any error occurs during the retrieval process, it returns nil and the error.
// Otherwise, it returns the CmsLogin record and nil.
func (r *CmsCustomerAgentRepository) GetAgentsByCustCode(custCode string) ([]*entities.CmsLogin, error) {
	a, err := r.GetByCustomerCode(custCode)
	if err != nil {
		return nil, err
	}

	agentCodes := iterator.Map(a, func(item *entities.CmsCustomerAgent) string {
		return item.AgentCode
	})

	c, err := r.l.Find(builder.Select("*").Where(builder.In("agent_code", agentCodes)))
	if err != nil {
		return nil, err
	}
	return c, nil
}

// InsertMany inserts multiple records into the CmsCustomerAgent table.
// It iterates over the records, validates each item, and inserts them into the database.
// If an error occurs during the insertion process, the error is returned.
// After the insertion, a log entry is created with the operation type and the inserted records.
// The log entry is passed to the audit log implementation for logging.
func (r *CmsCustomerAgentRepository) InsertMany(records []*entities.CmsCustomerAgent) error {
	_, err := r.db.Insert(iterator.Map(records, func(item *entities.CmsCustomerAgent) *entities.CmsCustomerAgent {
		item.Validate()
		return item
	}))
	if err != nil {
		return err
	}

	r.log("INSERT", records)

	return nil
}

// Update updates a record in the CmsCustomerAgent table.
// It takes a pointer to a CmsCustomerAgent object as a parameter.
// It updates the record in the database based on the.Id field of the object.
// If an error occurs while updating the record, it returns that error.
// Otherwise, it logs the update operation with the "UPDATE" operation type and the updated records by calling the log method.
// Returns nil if the update is successful.
func (r *CmsCustomerAgentRepository) Update(record *entities.CmsCustomerAgent) error {
	_, err := r.db.Where("id = ?").Omit("id").Update(record.Id)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsCustomerAgent{record})

	return nil
}

// Delete sets the active status of the given CmsCustomerAgent record to 0
// and updates it using the Update method of the CmsCustomerAgentRepository.
// It returns an error if the update operation fails.
func (r *CmsCustomerAgentRepository) Delete(record *entities.CmsCustomerAgent) error {
	record.ActiveStatus = 0
	_, err := r.db.Where("id = ?", record.Id).Cols("active_status").Update(record)
	if err == nil {
		r.log("DELETE", []*entities.CmsCustomerAgent{record})
	}
	return err
}

// UpdateMany updates multiple `CmsCustomerAgent` records in the database.
// It validates each record and iterates over them to update individually.
// If any update fails, it returns the error. Otherwise, it logs the updated records.
func (r *CmsCustomerAgentRepository) UpdateMany(records []*entities.CmsCustomerAgent) error {
	for _, record := range records {
		_, err := r.db.Where("id = ?", record.Id).Omit("id").Update(record)
		if err != nil {
			return err
		}
	}

	r.log("UPDATE", records)

	return nil
}

// DeleteMany deletes multiple `CmsCustomerAgent` records by setting their ActiveStatus to 0 in a bulk update operation,
// and logs the operation as "DELETE".
func (r *CmsCustomerAgentRepository) DeleteMany(records []*entities.CmsCustomerAgent) error {
	ids := iterator.Map(records, func(item *entities.CmsCustomerAgent) interface{} {
		return item.Id
	})

	_, err := r.db.In("id", ids).Cols("active_status").Update(&entities.CmsCustomerAgent{
		ActiveStatus: 0,
	})
	if err != nil {
		return err
	}

	r.log("DELETE", records)

	return nil
}

// log logs the operation and its payload to the audit log.
func (r *CmsCustomerAgentRepository) log(op string, payload []*entities.CmsCustomerAgent) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsCustomerAgent) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      strconv.FormatUint(item.Id, 10),
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
