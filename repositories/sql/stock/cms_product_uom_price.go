package stock

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"xorm.io/builder"
	"xorm.io/xorm"
)

// CmsProductUomPriceRepository represents a repository for managing CMS product unit of measure prices.
type CmsProductUomPriceRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
}

// NewCmsProductUomPriceRepository returns a new instance of CmsProductUomPriceRepository,
// which implements the IRepository interface. It takes an option object of type IRepository
// as a parameter and initializes the fields db and audit of the CmsProductUomPriceRepository
// instance with the values from the option object.
// The db field is of type *xorm.Engine and represents the database engine.
// The audit field is of type contracts.IAuditLog and represents the audit log repository.
// This function returns a pointer to the newly created CmsProductUomPriceRepository instance.
func NewCmsProductUomPriceRepository(option *contracts.IRepository) *CmsProductUomPriceRepository {
	return &CmsProductUomPriceRepository{
		db:    option.Db,
		audit: option.Audit,
	}
}

// Get retrieves a list of CmsProductUomPrice entities based on the given product code.
// The method uses the CmsProductUomPriceRepository's database connection to perform the query.
// The function filters the records where the product_code is equal to the provided productCode
// and the active_status equals 1. If an error occurs during the query, it returns nil and the error.
// Otherwise, it returns the list of records and nil as the error.
// Note that the Get method is a member of the CmsProductUomPriceRepository struct.
func (r *CmsProductUomPriceRepository) Get(productCode string) ([]*entities.CmsProductUomPrice, error) {
	var records []*entities.CmsProductUomPrice
	err := r.db.Where("product_code = ? AND active_status = ?", productCode, 1).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsProductUomPriceRepository) Find(predicate *builder.Builder) ([]*entities.CmsProductUomPrice, error) {
	var records []*entities.CmsProductUomPrice
	var t entities.CmsProductUomPrice
	err := r.db.SQL(predicate.From(t.TableName())).Find(&records)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, nil
	}
	return records, nil
}

// InsertMany inserts multiple records into the CmsProductUomPrice table.
// It takes a slice of records as input and returns an error if there was a problem during insertion.
// Each record is inserted using the database engine's Insert function.
// After insertion, the method logs the operation and returns nil if successful.
func (r *CmsProductUomPriceRepository) InsertMany(records []*entities.CmsProductUomPrice) error {
	_, err := r.db.Insert(iterator.Map(records, func(item *entities.CmsProductUomPrice) *entities.CmsProductUomPrice {
		return item
	}))
	if err != nil {
		return err
	}

	r.log("INSERT", records)

	return nil
}

// Update updates a record in the CmsProductUomPriceRepository database table.
//
// It takes a pointer to a CmsProductUomPrice record as input and updates the corresponding record
// in the database table. If any error occurs during the update operation, it returns the error.
//
// After updating the record, it logs the "UPDATE" operation along with the updated record using the
// log method of CmsProductUomPriceRepository.
func (r *CmsProductUomPriceRepository) Update(record *entities.CmsProductUomPrice) error {
	_, err := r.db.Where("product_uom_price_id = ?", record.ProductUomPriceId).Omit("product_uom_price_id").Update(record)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsProductUomPrice{record})

	return nil
}

// Delete sets the ActiveStatus of the given CmsProductUomPrice record to 0
// and updates it using the Update method of the CmsProductUomPriceRepository.
// It returns an error if the update operation fails.
func (r *CmsProductUomPriceRepository) Delete(record *entities.CmsProductUomPrice) error {
	record.ActiveStatus = 0
	_, err := r.db.Where("product_uom_price_id = ?", record.ProductUomPriceId).Cols("active_status").Update(record)
	if err == nil {
		r.log("DELETE", []*entities.CmsProductUomPrice{record})
	}
	return err
}

// UpdateMany updates multiple records in the database.
func (r *CmsProductUomPriceRepository) UpdateMany(records []*entities.CmsProductUomPrice) error {
	for _, record := range records {
		_, err := r.db.Where("product_uom_price_id = ?", record.ProductUomPriceId).Omit("product_uom_price_id").Update(record)
		if err != nil {
			return err
		}
	}

	r.log("UPDATE", records)

	return nil
}

// DeleteMany sets the ActiveStatus of each record in the input slice to 0
// and updates them using the UpdateMany method. It returns an error if
// the update operation fails.
func (r *CmsProductUomPriceRepository) DeleteMany(records []*entities.CmsProductUomPrice) error {
	ids := iterator.Map(records, func(item *entities.CmsProductUomPrice) uint64 {
		return item.ProductUomPriceId
	})

	_, err := r.db.In("product_uom_price_id", ids).Cols("active_status").Update(&entities.CmsProductUomPrice{
		ActiveStatus: 0,
	})
	if err != nil {
		return err
	}

	r.log("DELETE", records)
	return nil
}

// log logs the provided operation and payload to the audit log.
//
// op is the operation type.
// payload is the array of entities to be logged.
// It marshals the payload into JSON and creates an AuditLog record for each item in the payload.
// The AuditLog record contains the operation type, table name, record id, and record body.
// It then logs the AuditLog records to the IAuditLog implementation.
// The function does not handle the error returned by the IAuditLog implementation.
func (r *CmsProductUomPriceRepository) log(op string, payload []*entities.CmsProductUomPrice) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsProductUomPrice) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      item.ProductCode,
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
