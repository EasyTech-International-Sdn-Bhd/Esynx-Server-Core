package stock

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"xorm.io/builder"
	"xorm.io/xorm"
)

// CmsProductPriceTagRepository is a repository for managing CmsProductPriceTag data.
type CmsProductPriceTagRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
}

// NewCmsProductPriceTagRepository initializes a new instance of CmsProductPriceTagRepository
// using the provided IRepository option. It sets the db and audit fields of the repository.
func NewCmsProductPriceTagRepository(option *contracts.IRepository) *CmsProductPriceTagRepository {
	return &CmsProductPriceTagRepository{
		db:    option.Db,
		audit: option.Audit,
	}
}

// Get retrieves product price tags by product code
// It returns an array of CmsProductPriceV2 entities and an error, if any.
// The method queries the database for records with matching product code and active status.
// The matching records are stored in the 'records' variable.
// If an error occurs during the query, it is returned.
// Otherwise, the 'records' variable is returned.
func (r *CmsProductPriceTagRepository) Get(productCode string) ([]*entities.CmsProductPriceV2, error) {
	var records []*entities.CmsProductPriceV2
	err := r.db.Where("product_code = ? AND active_status = ?", productCode, 1).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

// GetByCustCode retrieves the records of products by customer code.
// It queries the database for records that match the provided product code,
// active status, and customer code. It returns a slice of
// CmsProductPriceV2 entities and an error if any occurred during the database operation.
func (r *CmsProductPriceTagRepository) GetByCustCode(productCode string, custCode string) ([]*entities.CmsProductPriceV2, error) {
	var records []*entities.CmsProductPriceV2
	err := r.db.Where("product_code = ? AND active_status = ? AND cust_code = ?", productCode, 1, custCode).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

// GetByPriceType retrieves a list of CmsProductPriceV2 records based on the given product code and price type.
// It queries the database by matching the product code, active status (1), and price category.
// It returns the matching records if found, otherwise it returns an error.
func (r *CmsProductPriceTagRepository) GetByPriceType(productCode string, priceType string) ([]*entities.CmsProductPriceV2, error) {
	var records []*entities.CmsProductPriceV2
	err := r.db.Where("product_code = ? AND active_status = ? AND price_cat = ?", productCode, 1, priceType).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsProductPriceTagRepository) Find(predicate *builder.Builder) ([]*entities.CmsProductPriceV2, error) {
	var records []*entities.CmsProductPriceV2
	var t entities.CmsProductPriceV2
	err := r.db.SQL(predicate.From(t.TableName())).Find(&records)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, nil
	}
	return records, nil
}

// InsertMany inserts multiple records into the CmsProductPriceTagRepository database table. It takes an array
// of CmsProductPriceV2 entities as a parameter and returns an error if the database operation fails. The
// method first maps the records using the iterator.Map function, which returns the same records. Then,
// it uses the Insert function of the db field to insert the mapped records into the table. After the insertion,
// the method logs the operation with the "INSERT" operation type and the inserted records.
func (r *CmsProductPriceTagRepository) InsertMany(records []*entities.CmsProductPriceV2) error {
	_, err := r.db.Insert(iterator.Map(records, func(item *entities.CmsProductPriceV2) *entities.CmsProductPriceV2 {
		return item
	}))
	if err != nil {
		return err
	}

	r.log("INSERT", records)

	return nil
}

// Update updates a record in the CmsProductPriceTagRepository.
// It updates the record in the database with the given ProductPriceId.
// If the update operation fails, it returns an error.
// After the update, it logs the "UPDATE" operation with the updated record.
//
// Example usage:
//
//	record := &entities.CmsProductPriceV2{
//	  ProductPriceId: 1,
//	  // set other fields
//	}
//	err := repository.Update(record)
//	if err != nil {
//	  log.Println("Update failed:", err)
//	}
//
// Parameters:
//   - record: The record to be updated in the database.
//
// Returns:
//   - error: The error if the update operation fails, nil otherwise.
func (r *CmsProductPriceTagRepository) Update(record *entities.CmsProductPriceV2) error {
	_, err := r.db.Where("product_price_id = ?", record.ProductPriceId).Update(record)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsProductPriceV2{record})

	return nil
}

// Delete sets the active status of the given record to 0 and updates it in the database
// with the Update method. It returns an error if the update operation fails.
func (r *CmsProductPriceTagRepository) Delete(record *entities.CmsProductPriceV2) error {
	record.ActiveStatus = 0
	_, err := r.db.Where("product_price_id = ?", record.ProductPriceId).Cols("active_status").Update(record)
	if err == nil {
		r.log("DELETE", []*entities.CmsProductPriceV2{record})
	}
	return err
}

func (r *CmsProductPriceTagRepository) UpdateMany(records []*entities.CmsProductPriceV2) error {
	for _, record := range records {
		_, err := r.db.Where("product_price_id = ?", record.ProductPriceId).Update(record)
		if err != nil {
			return err
		}
	}

	r.log("UPDATE", records)

	return nil
}

func (r *CmsProductPriceTagRepository) DeleteMany(records []*entities.CmsProductPriceV2) error {
	ids := iterator.Map(records, func(item *entities.CmsProductPriceV2) uint64 {
		return item.ProductPriceId
	})

	_, err := r.db.In("product_price_id", ids).Cols("active_status").Update(&entities.CmsProductPriceV2{
		ActiveStatus: 0,
	})
	if err != nil {
		return err
	}

	r.log("DELETE", records)
	return nil
}

// log logs an audit entry with the specified operation type and payload.
// It marshals the payload into a JSON string and creates an AuditLog object
// for each item in the payload. It then logs the generated AuditLog objects
// using the audit logger specified in the CmsProductPriceTagRepository.
func (r *CmsProductPriceTagRepository) log(op string, payload []*entities.CmsProductPriceV2) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsProductPriceV2) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      item.ProductCode,
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
