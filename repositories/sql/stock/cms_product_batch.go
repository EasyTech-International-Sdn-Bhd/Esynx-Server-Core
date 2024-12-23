package stock

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"xorm.io/builder"
	"xorm.io/xorm"
)

// Get retrieves product batches by product code.
type CmsProductBatchRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
}

// NewCmsProductBatchRepository creates a new instance of CmsProductBatchRepository
// with the specified option. It initializes the repository with the 'db' and 'audit' fields
// from the 'option' parameter.
func NewCmsProductBatchRepository(option *contracts.IRepository) *CmsProductBatchRepository {
	return &CmsProductBatchRepository{
		db:    option.Db,
		audit: option.Audit,
	}
}

// Get retrieves a list of CmsProductBatch entities from the database based on the provided productCode.
// It filters the records based on the productCode and active_status = 1.
// If an error occurs during the database operation, it returns nil and the error.
// Otherwise, it returns the retrieved records and nil.
func (r *CmsProductBatchRepository) Get(productCode string) ([]*entities.CmsProductBatch, error) {
	var records []*entities.CmsProductBatch
	err := r.db.Where("product_code = ? AND active_status = ?", productCode, 1).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

// GetByWarehouse retrieves a list of CmsProductBatch records based on the given product code and warehouse code.
// It filters the records by the active status (1) and returns them as []*entities.CmsProductBatch.
// If there is an error during the database query, it returns nil and the error.
func (r *CmsProductBatchRepository) GetByWarehouse(productCode string, warehouse string) ([]*entities.CmsProductBatch, error) {
	var records []*entities.CmsProductBatch
	err := r.db.Where("product_code = ? AND wh_code = ? AND active_status = ?", productCode, warehouse, 1).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsProductBatchRepository) Find(predicate *builder.Builder) ([]*entities.CmsProductBatch, error) {
	var records []*entities.CmsProductBatch
	var t entities.CmsProductBatch
	err := r.db.SQL(predicate.From(t.TableName())).Find(&records)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, nil
	}
	return records, nil
}

// InsertMany inserts multiple records of type *entities.CmsProductBatch into the repository.
// It calls the Insert method of the underlying database engine to insert the records.
// If an error occurs during the insertion, it returns the error.
// After the insertion, it logs the "INSERT" operation and the inserted records.
//
// Parameters:
// - records: a slice of *entities.CmsProductBatch representing the records to be inserted.
//
// Returns:
// - error: an error if the insertion fails, nil otherwise.
func (r *CmsProductBatchRepository) InsertMany(records []*entities.CmsProductBatch) error {
	_, err := r.db.Insert(iterator.Map(records, func(item *entities.CmsProductBatch) *entities.CmsProductBatch {
		return item
	}))
	if err != nil {
		return err
	}

	// r.log("INSERT", records)

	return nil
}

// Update updates a record of CmsProductBatch in the database.
// It takes a pointer to an entities.CmsProductBatch as input,
// and returns an error if the update fails.
// The method first calls the Update function of the database engine,
// passing the record's id as the condition and the record itself as the new values.
// If there is an error during the update, the method returns the error.
// Afterwards, the method calls the log method, passing "UPDATE" as the operation
// and an array containing the updated record as the payload.
// Finally, the method returns nil if the update is successful.
func (r *CmsProductBatchRepository) Update(record *entities.CmsProductBatch) error {
	_, err := r.db.Where("id = ?", record.Id).Omit("id").Update(record)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsProductBatch{record})

	return nil
}

// Delete sets the active status of the given CmsProductBatch record to 0
// and updates it using the Update method of the CmsProductBatchRepository.
// It returns an error if the update operation fails.
func (r *CmsProductBatchRepository) Delete(record *entities.CmsProductBatch) error {
	record.ActiveStatus = 0
	_, err := r.db.Where("id = ?", record.Id).Cols("active_status").Update(record)
	if err == nil {
		r.log("DELETE", []*entities.CmsProductBatch{record})
	}
	return err
}

// UpdateMany updates multiple records in the CmsProductBatch table.
// It takes an array of CmsProductBatch records as input.
// It iterates over each record and performs the update operation.
// If any update fails, it returns the error.
// Finally, it logs the UPDATE operation and returns nil if all updates are successful.
func (r *CmsProductBatchRepository) UpdateMany(records []*entities.CmsProductBatch) error {
	for _, record := range records {
		_, err := r.db.Where("id = ?", record.Id).Omit("id").Update(record)
		if err != nil {
			return err
		}
	}

	r.log("UPDATE", records)

	return nil
}

// DeleteMany sets the ActiveStatus of each record in the input slice to 0
// and updates them in the CmsProductBatch table. It returns an error if
// the update operation fails.
func (r *CmsProductBatchRepository) DeleteMany(records []*entities.CmsProductBatch) error {
	ids := iterator.Map(records, func(item *entities.CmsProductBatch) uint64 {
		return item.Id
	})

	_, err := r.db.In("id", ids).Cols("active_status").Update(&entities.CmsProductBatch{
		ActiveStatus: 0,
	})
	if err != nil {
		return err
	}

	r.log("DELETE", records)
	return nil
}

// log logs the operation and payload to an audit log.
//
// It takes in the operation type as a string and the payload as a slice of *entities.CmsProductBatch.
// It marshals the payload into a JSON string and creates an *entities.AuditLog object
// for each *entities.CmsProductBatch in the payload, with the operation type, table name,
// record ID, and record body.
// The resulting audit logs are then passed to the r.audit.Log method to handle the logging.
// There is no return value.
//
// Example usage:
//
//	// r.log("INSERT", records)
//	r.log("UPDATE", []*entities.CmsProductBatch{record})
//	r.log("UPDATE", records)
//
// Preconditions:
// - CmsProductBatchRepository must have an instance of IAuditLog accessible through the r.audit field.
// - Entities in the payload must have the TableName method that returns the name of the table of the entity.
// - The payload must be a non-empty slice.
// Postconditions:
// - The operation and payload are logged to the audit log.
func (r *CmsProductBatchRepository) log(op string, payload []*entities.CmsProductBatch) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsProductBatch) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      fmt.Sprintf("%s.%s", item.BatchCode, item.ProductCode),
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
