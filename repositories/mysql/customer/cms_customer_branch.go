package customer

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"xorm.io/builder"
	"xorm.io/xorm"
)

// CmsCustomerBranchRepository is a repository for managing CMS customer branches.
type CmsCustomerBranchRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
}

// NewCmsCustomerBranchRepository is a function that creates a new instance of CmsCustomerBranchRepository.
// It takes an option parameter of type *contracts.IRepository and returns a pointer to CmsCustomerBranchRepository.
func NewCmsCustomerBranchRepository(option *contracts.IRepository) *CmsCustomerBranchRepository {
	return &CmsCustomerBranchRepository{
		db:    option.Db,
		audit: option.Audit,
	}
}

// Get retrieves a single CmsCustomerBranch entity from the database based on the branchCode provided.
// If the entity is found, it is returned as a pointer along with a nil error.
// If the entity is not found, a nil pointer and nil error are returned.
// If an error occurs during the database operation, a nil pointer and the error are returned.
func (r *CmsCustomerBranchRepository) Get(branchCode string) (*entities.CmsCustomerBranch, error) {
	var branch entities.CmsCustomerBranch
	has, err := r.db.Where("branch_code = ?", branchCode).Get(&branch)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return &branch, nil
}

// GetByCustomerCode retrieves customer branches by the customer code.
// It queries the database for branches with the provided customer code and active status.
// Returns an array of CmsCustomerBranch entities and any errors that occurred.
func (r *CmsCustomerBranchRepository) GetByCustomerCode(custCode string) ([]*entities.CmsCustomerBranch, error) {
	var branches []*entities.CmsCustomerBranch
	err := r.db.Where("cust_code = ? AND branch_active = ?", custCode, 1).Find(&branches)
	if err != nil {
		return nil, err
	}
	return branches, nil
}

// GetByAgentId returns the customer branches associated with the given agent ID,
// filtered by active branches. If no branches are found, it returns an empty result ([]*entities.CmsCustomerBranch, nil).
func (r *CmsCustomerBranchRepository) GetByAgentId(agentId int64) ([]*entities.CmsCustomerBranch, error) {
	branches, err := r.GetAllStatusByAgentId(agentId)
	if err != nil {
		return nil, err
	}
	if len(branches) == 0 {
		return iterator.Filter(branches, func(item *entities.CmsCustomerBranch) bool {
			return item.BranchActive == 1
		}), nil
	}
	return nil, nil
}

// GetAllStatusByAgentId retrieves all customer branches with a specific agent ID.
// It returns a slice of CmsCustomerBranch entities and an error, if any.
// The agent ID parameter specifies the ID of the agent.
// The returned slice contains all branches associated with the given agent ID.
// If no branches are found, the function returns an empty slice.
// The error value is nil if the query is successful, otherwise it contains the error encountered.
func (r *CmsCustomerBranchRepository) GetAllStatusByAgentId(agentId int64) ([]*entities.CmsCustomerBranch, error) {
	var branches []*entities.CmsCustomerBranch
	err := r.db.Where("agent_id = ?", agentId).Find(&branches)
	if err != nil {
		return nil, err
	}
	return branches, nil
}

func (r *CmsCustomerBranchRepository) Find(predicate *builder.Builder) ([]*entities.CmsCustomerBranch, error) {
	var records []*entities.CmsCustomerBranch
	var t entities.CmsCustomerBranch
	err := r.db.SQL(predicate.From(t.TableName())).Find(&records)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, nil
	}
	return records, nil
}

// InsertMany inserts multiple records into the CmsCustomerBranchRepository table.
// It takes a slice of *entities.CmsCustomerBranch records as input.
// It maps the records using the Map function from the iterator package,
// and then performs the insertion using the db.Insert method.
// If an error occurs during the insertion, it returns the error.
// It then logs the "INSERT" operation with the inserted records,
// and returns nil if everything is successful.
func (r *CmsCustomerBranchRepository) InsertMany(records []*entities.CmsCustomerBranch) error {
	_, err := r.db.Insert(iterator.Map(records, func(item *entities.CmsCustomerBranch) *entities.CmsCustomerBranch {
		return item
	}))
	if err != nil {
		return err
	}

	r.log("INSERT", records)

	return nil
}

// Update updates a CmsCustomerBranch record in the database.
// It updates the record with the given branch_code and uses the provided record values.
// If the update operation fails, it returns an error.
// After a successful update, it logs the "UPDATE" operation along with the updated record.
func (r *CmsCustomerBranchRepository) Update(record *entities.CmsCustomerBranch) error {
	_, err := r.db.Table(record.TableName()).Where("branch_code = ?", record.BranchCode).Update(record)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsCustomerBranch{record})

	return nil
}

// UpdateMany updates multiple records in the CmsCustomerBranchRepository.
// It begins a session, updates each record in the provided slice, and commits
// the changes if all updates are successful. If any update fails, it rolls
// back the session and returns the error. It logs the "UPDATE" operation and
// the updated records using the log method.
//
// If an error occurs when beginning the session, it returns the error.
//
// If an error occurs during an update, it sets the rollback flag to true, sets
// the sessionErr to the error, and breaks out of the loop.
//
// After updating all records, if the rollback flag is set to true, it
// rolls back the session and returns the sessionErr.
//
// If an error occurs when committing the session, it returns the error.
//
// Finally, it returns nil, indicating that the updates were successful.
func (r *CmsCustomerBranchRepository) UpdateMany(records []*entities.CmsCustomerBranch) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}
	var sessionErr error
	rollback := false
	for _, branch := range records {
		_, err = session.Table(branch.TableName()).Where("branch_code = ?", branch.BranchCode).Update(branch)
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

// Delete sets the BranchActive field of the provided CmsCustomerBranch record to 0
// and calls the Update method to persist the changes.
// Returns an error if the Update operation failed.
func (r *CmsCustomerBranchRepository) Delete(record *entities.CmsCustomerBranch) error {
	record.BranchActive = 0
	return r.Update(record)
}

// DeleteMany deletes multiple records by setting their BranchActive field to 0
// and calling the UpdateMany method.
func (r *CmsCustomerBranchRepository) DeleteMany(records []*entities.CmsCustomerBranch) error {
	for _, record := range records {
		record.BranchActive = 0
	}
	return r.UpdateMany(records)
}

// log writes an audit log record for the given operation and payload.
// The payload is marshaled to JSON and transformed into AuditLog objects
// before being logged using the 'audit' instance of the IAuditLog interface.
func (r *CmsCustomerBranchRepository) log(op string, payload []*entities.CmsCustomerBranch) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsCustomerBranch) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      fmt.Sprintf("%s.%s", item.CustCode, item.BranchCode),
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
