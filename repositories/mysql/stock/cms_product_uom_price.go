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

// Get retrieves a list of CmsProductUomPriceV2 entities based on the given product code.
// The method uses the CmsProductUomPriceRepository's database connection to perform the query.
// The function filters the records where the product_code is equal to the provided productCode
// and the active_status equals 1. If an error occurs during the query, it returns nil and the error.
// Otherwise, it returns the list of records and nil as the error.
// Note that the Get method is a member of the CmsProductUomPriceRepository struct.
func (r *CmsProductUomPriceRepository) Get(productCode string) ([]*entities.CmsProductUomPriceV2, error) {
	var records []*entities.CmsProductUomPriceV2
	err := r.db.Where("product_code = ? AND active_status = ?", productCode, 1).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsProductUomPriceRepository) Find(predicate *builder.Builder) ([]*entities.CmsProductUomPriceV2, error) {
	var records []*entities.CmsProductUomPriceV2
	var t entities.CmsProductUomPriceV2
	err := r.db.SQL(predicate.From(t.TableName())).Find(&records)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, nil
	}
	return records, nil
}

// InsertMany inserts multiple records into the CmsProductUomPriceV2 table.
// It takes a slice of records as input and returns an error if there was a problem during insertion.
// Each record is inserted using the database engine's Insert function.
// After insertion, the method logs the operation and returns nil if successful.
func (r *CmsProductUomPriceRepository) InsertMany(records []*entities.CmsProductUomPriceV2) error {
	_, err := r.db.Insert(iterator.Map(records, func(item *entities.CmsProductUomPriceV2) *entities.CmsProductUomPriceV2 {
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
// It takes a pointer to a CmsProductUomPriceV2 record as input and updates the corresponding record
// in the database table. If any error occurs during the update operation, it returns the error.
//
// After updating the record, it logs the "UPDATE" operation along with the updated record using the
// log method of CmsProductUomPriceRepository.
func (r *CmsProductUomPriceRepository) Update(record *entities.CmsProductUomPriceV2) error {
	_, err := r.db.Table(record.TableName()).Where("product_uom_price_id = ?", record.ProductUomPriceId).Update(record)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsProductUomPriceV2{record})

	return nil
}

// UpdateMany updates multiple records in the database.
func (r *CmsProductUomPriceRepository) UpdateMany(records []*entities.CmsProductUomPriceV2) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}
	var sessionErr error
	rollback := false
	for _, record := range records {
		_, err = session.Table(record.TableName()).Where("product_uom_price_id = ?", record.ProductUomPriceId).Update(record)
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

// Delete sets the ActiveStatus of the given record to 0 and updates it using the Update method.
func (r *CmsProductUomPriceRepository) Delete(record *entities.CmsProductUomPriceV2) error {
	record.ActiveStatus = 0
	return r.Update(record)
}

// DeleteMany sets the ActiveStatus of each record in the given slice to 0, and calls UpdateMany to update the records in the repository.
func (r *CmsProductUomPriceRepository) DeleteMany(records []*entities.CmsProductUomPriceV2) error {
	for _, record := range records {
		record.ActiveStatus = 0
	}
	return r.UpdateMany(records)
}

// log logs the provided operation and payload to the audit log.
//
// op is the operation type.
// payload is the array of entities to be logged.
// It marshals the payload into JSON and creates an AuditLog record for each item in the payload.
// The AuditLog record contains the operation type, table name, record id, and record body.
// It then logs the AuditLog records to the IAuditLog implementation.
// The function does not handle the error returned by the IAuditLog implementation.
func (r *CmsProductUomPriceRepository) log(op string, payload []*entities.CmsProductUomPriceV2) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsProductUomPriceV2) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      item.ProductCode,
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
