package agent

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"xorm.io/xorm"
)

// CmsLoginRepository is a repository for managing CMS logins.
type CmsLoginRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
}

// NewCmsLoginRepository returns a new instance of CmsLoginRepository with the given IRepository option.
// The option should have a valid db connection and audit log implementation.
func NewCmsLoginRepository(option *contracts.IRepository) *CmsLoginRepository {
	return &CmsLoginRepository{
		db:    option.Db,
		audit: option.Audit,
	}
}

// Get retrieves a CmsLogin record from the repository based on the given agentId.
// If the record is found, it is returned along with nil error. If the record is not found, both the return value
// will be nil. If an error occurs during the retrieval process, nil record and the corresponding error is returned.
func (r *CmsLoginRepository) Get(agentId int64) (*entities.CmsLogin, error) {
	var record entities.CmsLogin
	has, err := r.db.Where("login_id = ?", agentId).Get(&record)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return &record, nil
}

// GetByAgentCode retrieves a CMS login record from the database based on the given agent code.
// It returns a pointer to the retrieved CmsLogin entity and an error, if any.
// If the record is not found, it returns nil for the entity and nil for the error.
//
// Parameters:
// - agentCode: The agent code to search for in the database.
//
// Returns:
// - *entities.CmsLogin: A pointer to the retrieved CmsLogin entity, or nil if not found.
// - error: An error if any occurred during the retrieval process, or nil if successful.
func (r *CmsLoginRepository) GetByAgentCode(agentCode string) (*entities.CmsLogin, error) {
	var record entities.CmsLogin
	has, err := r.db.Where("staff_code = ?", agentCode).Get(&record)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return &record, nil
}

// GetAll retrieves all records from the CmsLogin table and returns them as a slice of CmsLogin entities.
// If there are no records found, it returns nil.
// If there is an error fetching the records, it returns nil and the error.
func (r *CmsLoginRepository) GetAll() ([]*entities.CmsLogin, error) {
	var records []*entities.CmsLogin
	err := r.db.Find(&records)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, nil
	}
	return records, nil
}

// Find retrieves CmsLogin records from the repository based on the given predicate. `builder`
// If records are found, they are returned along with nil error. If no records are found,
// both the return value will be nil. If an error occurs during the retrieval process,
// nil records and the corresponding error is returned.
func (r *CmsLoginRepository) Find(predicate interface{}) ([]*entities.CmsLogin, error) {
	var records []*entities.CmsLogin
	err := r.db.Where(predicate).Find(&records)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, nil
	}
	return records, nil
}

// InsertMany inserts multiple records into the CmsLoginRepository table.
// It accepts a slice of CmsLogin records as a parameter. It uses the db.Insert
// method to insert the records into the table. After the insertion, it calls the log method
// to log the operation. If any error occurs during the insertion, it returns the error.
// It returns nil if the insertion is successful.
func (r *CmsLoginRepository) InsertMany(records []*entities.CmsLogin) error {
	_, err := r.db.Insert(iterator.Map(records, func(item *entities.CmsLogin) *entities.CmsLogin {
		return item
	}))
	if err != nil {
		return err
	}

	r.log("INSERT", records)

	return nil
}

// Update updates the given record in the CmsLoginRepository's database table.
//
// It updates the record with the provided login_id with the values in the given record.
// If an error occurs during the update, it returns the error.
//
// After updating the record, it logs the "UPDATE" operation with the updated record
// using the log method.
//
// The log method generates an AuditLog record based on the provided payload and logs it
// using the contracts.IAuditLog interface.
//
// Example usage:
//
//	record := &entities.CmsLogin{}
//	err := cmsLoginRepository.Update(record)
//	if err != nil {
//		// handle error
//	}
//
// The CmsLoginRepository struct should have a db field of type *xorm.Engine and an audit
// field of type contracts.IAuditLog, as well as other methods like Get, GetAll, InsertMany, etc.
//
// The contracts.IAuditLog interface should have a Log method that takes a slice of
// entities.AuditLog as its parameter.
func (r *CmsLoginRepository) Update(record *entities.CmsLogin) error {
	_, err := r.db.Where("staff_code = ?", record.StaffCode).Update(record)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsLogin{record})

	return nil
}

// UpdateMany updates multiple records in the CmsLogin table.
// It takes a slice of CmsLogin records as input and applies the updates within a transaction.
// If any error occurs during the update process, it rolls back the transaction and returns the error.
// Otherwise, it commits the transaction and returns nil.
// After updating the records, it logs the operation as "UPDATE" along with the updated records.
func (r *CmsLoginRepository) UpdateMany(records []*entities.CmsLogin) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}
	var sessionErr error
	rollback := false
	for _, record := range records {
		_, err = session.Where("staff_code = ?", record.StaffCode).Update(record)
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

// Delete sets the login status of the given record to 0 and updates it using the Update method.
func (r *CmsLoginRepository) Delete(record *entities.CmsLogin) error {
	record.LoginStatus = 0
	return r.Update(record)
}

// DeleteMany deletes multiple records by setting their LoginStatus to 0 and calling UpdateMany.
func (r *CmsLoginRepository) DeleteMany(records []*entities.CmsLogin) error {
	for _, record := range records {
		record.LoginStatus = 0
	}
	return r.UpdateMany(records)
}

// log logs the operation and payload to the audit log.
// It marshals the payload into JSON, creates an AuditLog instance for each item in the payload, and logs them using the audit.Log method.
//
// Parameters:
// - op: the operation type as a string
// - payload: a slice of CmsLogin objects representing the data to be logged
//
// The method maps the payload to AuditLog instances and logs them using the audit.Log method.
//
// Example usage:
// ```
// r.log("INSERT", records)
// ```
func (r *CmsLoginRepository) log(op string, payload []*entities.CmsLogin) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsLogin) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      item.StaffCode,
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
