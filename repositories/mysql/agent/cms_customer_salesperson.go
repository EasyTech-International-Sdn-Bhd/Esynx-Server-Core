package agent

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"strconv"
	"xorm.io/xorm"
)

// CmsCustomerSalespersonRepository is a repository for managing CMS customer salespersons.
// It has a database connection, an audit log, and an instance of CmsLoginRepository.
// The repository allows for querying and manipulating CMS customer salesperson data.
// Usage examples: creating a new CmsCustomerSalespersonRepository, querying salespersons by agent ID, customer ID, and accessing the corresponding agent login information.
type CmsCustomerSalespersonRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
	l     *CmsLoginRepository
}

// NewCmsCustomerSalespersonRepository returns a new instance of CmsCustomerSalespersonRepository with the given IRepository option.
// The option should have a valid db connection and audit log implementation.
func NewCmsCustomerSalespersonRepository(option *contracts.IRepository) *CmsCustomerSalespersonRepository {
	return &CmsCustomerSalespersonRepository{
		db:    option.Db,
		audit: option.Audit,
		l:     NewCmsLoginRepository(option),
	}
}

// GetByAgentId retrieves customer salesperson records by agent ID.
//
// It queries the database for customer salesperson records where the salesperson ID
// matches the given agent ID and the active status is set to 1 (active).
//
// If successful, it returns a slice of *entities.CmsCustomerSalesperson representing the records
// and a nil error. If an error occurs during the retrieval process, it returns nil and the corresponding error.
func (r *CmsCustomerSalespersonRepository) GetByAgentId(agentId int64) ([]*entities.CmsCustomerSalesperson, error) {
	var record []*entities.CmsCustomerSalesperson
	err := r.db.Where("salesperson_id = ? AND active_status = ?", agentId, 1).Find(&record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// GetByCustomerId retrieves a record from the database based on the customer ID.
// It returns a pointer to the retrieved record of type CmsCustomerSalesperson and an error.
// If the record exists, the returned error is nil. If the record does not exist, both the record and the error are nil.
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

// GetAgentByCustId returns the agent for a given customer ID.
// It retrieves the associated CmsCustomerSalesperson record using GetByCustomerId,
// and then retrieves the CmsLogin record using the salesperson ID from the CmsCustomerSalesperson record.
// If any error occurs during the retrieval process, it returns nil and the error.
// Otherwise, it returns the CmsLogin record and nil.
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

// InsertMany inserts multiple records into the CmsCustomerSalesperson table.
// It iterates over the records, validates each item, and inserts them into the database.
// If an error occurs during the insertion process, the error is returned.
// After the insertion, a log entry is created with the operation type and the inserted records.
// The log entry is passed to the audit log implementation for logging.
func (r *CmsCustomerSalespersonRepository) InsertMany(records []*entities.CmsCustomerSalesperson) error {
	_, err := r.db.Insert(iterator.Map(records, func(item *entities.CmsCustomerSalesperson) *entities.CmsCustomerSalesperson {
		item.Validate()
		return item
	}))
	if err != nil {
		return err
	}

	r.log("INSERT", records)

	return nil
}

// Update updates a record in the CmsCustomerSalesperson table.
// It takes a pointer to a CmsCustomerSalesperson object as a parameter.
// It updates the record in the database based on the SalespersonCustomerId field of the object.
// If an error occurs while updating the record, it returns that error.
// Otherwise, it logs the update operation with the "UPDATE" operation type and the updated records by calling the log method.
// Returns nil if the update is successful.
func (r *CmsCustomerSalespersonRepository) Update(record *entities.CmsCustomerSalesperson) error {
	_, err := r.db.Table(record.TableName()).Where("salesperson_customer_id = ?").Update(record.SalespersonCustomerId)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsCustomerSalesperson{record})

	return nil
}

// UpdateMany updates multiple `CmsCustomerSalesperson` records in the database.
// It validates each record, begins a new session, and iterates over the records
// to update them individually. If any update fails, it triggers a rollback of the session
// and returns the error. Otherwise, the updates are committed, and the method logs
// the updated records.
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
		_, err = session.Table(record.TableName()).Where("salesperson_customer_id = ?", record.SalespersonCustomerId).Update(record)
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

	r.log("UPDATE", records)

	return nil
}

// Delete sets the active status of the given CmsCustomerSalesperson record to 0
// and updates it using the Update method of the CmsCustomerSalespersonRepository.
// It returns an error if the update operation fails.
func (r *CmsCustomerSalespersonRepository) Delete(record *entities.CmsCustomerSalesperson) error {
	record.ActiveStatus = 0
	_, err := r.db.Where("salesperson_customer_id = ?", record.SalespersonCustomerId).Cols("active_status").Update(record)
	if err == nil {
		r.log("DELETE", []*entities.CmsCustomerSalesperson{record})
	}
	return err
}

// DeleteMany sets the ActiveStatus of each record in the input slice to 0
// and updates them using the UpdateMany method. It returns an error if
// the update operation fails.
func (r *CmsCustomerSalespersonRepository) DeleteMany(records []*entities.CmsCustomerSalesperson) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}
	var sessionErr error
	rollback := false
	for _, record := range records {
		record.ActiveStatus = 0
		_, err = session.Where("salesperson_customer_id = ?", record.SalespersonCustomerId).Cols("active_status").Update(record)
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

	r.log("DELETE", records)

	return nil
}

// log logs the operation and its payload to the audit log.
func (r *CmsCustomerSalespersonRepository) log(op string, payload []*entities.CmsCustomerSalesperson) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsCustomerSalesperson) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      strconv.FormatUint(item.SalespersonCustomerId, 10),
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
