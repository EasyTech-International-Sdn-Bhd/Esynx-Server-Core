package stock

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"xorm.io/builder"
	"xorm.io/xorm"
)

// CmsWarehouseStockRepository represents a repository for managing CMS warehouse stock records.
type CmsWarehouseStockRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
}

// NewCmsWarehouseStockRepository creates a new instance of CmsWarehouseStockRepository
func NewCmsWarehouseStockRepository(option *contracts.IRepository) *CmsWarehouseStockRepository {
	return &CmsWarehouseStockRepository{
		db:    option.Db,
		audit: option.Audit,
	}
}

// Get retrieves the warehouse stock records for a given product code.
// It queries the database to find records where the product code matches the given code
// and the active status is set to 1.
// If successful, it returns a slice of CmsWarehouseStock records. Otherwise, it returns an error.
func (r *CmsWarehouseStockRepository) Get(productCode string) ([]*entities.CmsWarehouseStock, error) {
	var records []*entities.CmsWarehouseStock
	err := r.db.Where("product_code = ? AND active_status = ?", productCode, 1).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsWarehouseStockRepository) Find(predicate *builder.Builder) ([]*entities.CmsWarehouseStock, error) {
	var records []*entities.CmsWarehouseStock
	var t entities.CmsWarehouseStock
	err := r.db.SQL(predicate.From(t.TableName())).Find(&records)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, nil
	}
	return records, nil
}

// InsertMany inserts multiple records into the CmsWarehouseStock table.
// It takes a slice of pointers to CmsWarehouseStock entities as input.
// The method maps each entity to itself and then inserts the mapped records
// into the database using the database engine db.
// If the insert operation fails, it returns an error. Otherwise, it logs the
// insert operation and returns nil.
func (r *CmsWarehouseStockRepository) InsertMany(records []*entities.CmsWarehouseStock) error {
	_, err := r.db.Insert(iterator.Map(records, func(item *entities.CmsWarehouseStock) *entities.CmsWarehouseStock {
		return item
	}))
	if err != nil {
		return err
	}

	// r.log("INSERT", records)

	return nil
}

// Update updates a record in the CmsWarehouseStockRepository with the provided data.
// It updates the record in the database by matching the record's id and updates it with the provided data.
// If the update operation encounters an error, it returns the error.
// After successful update, it logs the "UPDATE" operation and the updated record using the log method.
// It returns nil if the update operation is successful.
func (r *CmsWarehouseStockRepository) Update(record *entities.CmsWarehouseStock) error {
	_, err := r.db.Where("id = ?", record.Id).Omit("id").Update(record)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsWarehouseStock{record})

	return nil
}

// Delete sets the ActiveStatus of the given CmsWarehouseStock record to 0
// and updates it using the Update method. It returns an error if the update operation fails.
func (r *CmsWarehouseStockRepository) Delete(record *entities.CmsWarehouseStock) error {
	record.ActiveStatus = 0
	_, err := r.db.Where("id = ?", record.Id).Cols("active_status").Update(record)
	if err == nil {
		r.log("DELETE", []*entities.CmsWarehouseStock{record})
	}
	return err
}

// UpdateMany updates multiple CmsWarehouseStock records in the database.
func (r *CmsWarehouseStockRepository) UpdateMany(records []*entities.CmsWarehouseStock) error {
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
// and updates them using a database session. It returns an error if the update operation fails.
func (r *CmsWarehouseStockRepository) DeleteMany(records []*entities.CmsWarehouseStock) error {
	ids := iterator.Map(records, func(item *entities.CmsWarehouseStock) uint64 {
		return item.Id
	})

	_, err := r.db.In("id", ids).Cols("active_status").Update(&entities.CmsWarehouseStock{
		ActiveStatus: 0,
	})
	if err != nil {
		return err
	}

	r.log("DELETE", records)
	return nil
}

// log method logs the operation and payload to the audit log using the provided audit logger.
// It takes an operation string (op) and an array of CmsWarehouseStock objects (payload).
// It marshals the payload to JSON and creates an AuditLog object for each item in the payload.
// The AuditLog object contains the operation type, table name, product code, and the marshalled payload.
// The array of AuditLog objects is then passed to the audit logger for logging.
func (r *CmsWarehouseStockRepository) log(op string, payload []*entities.CmsWarehouseStock) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsWarehouseStock) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      item.ProductCode,
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
