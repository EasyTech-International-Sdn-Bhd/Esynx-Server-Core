package stock

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"strconv"
	"xorm.io/builder"
	"xorm.io/xorm"
)

// CmsProductImageRepository represents a repository for managing CMS product images.
type CmsProductImageRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
}

// NewCmsProductImageRepository creates a new instance of CmsProductImageRepository with the provided IRepository option.
func NewCmsProductImageRepository(option *contracts.IRepository) *CmsProductImageRepository {
	return &CmsProductImageRepository{
		db:    option.Db,
		audit: option.Audit,
	}
}

// Get retrieves the CmsProductImage records for a given productId.
func (r *CmsProductImageRepository) Get(productId string) ([]*entities.CmsProductImage, error) {
	var records []*entities.CmsProductImage
	err := r.db.Where("product_id = ? AND active_status = ?", productId, 1).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

// GetByProductCode retrieves product images by product code.
// It returns a slice of entities.CmsProductImage and an error.
// The first parameter is the product code used to filter the records.
// The function queries the database for records with the specified product code
// and active status of 1. If an error occurs during the query, it is returned.
// Otherwise, the records are returned.
func (r *CmsProductImageRepository) GetByProductCode(productCode string) ([]*entities.CmsProductImage, error) {
	var records []*entities.CmsProductImage
	err := r.db.Where("product_code = ? AND active_status = ?", productCode, 1).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsProductImageRepository) Find(predicate *builder.Builder) ([]*entities.CmsProductImage, error) {
	var records []*entities.CmsProductImage
	var t entities.CmsProductImage
	err := r.db.SQL(predicate.From(t.TableName())).Find(&records)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, nil
	}
	return records, nil
}

// InsertMany inserts multiple records into the CmsProductImageRepository.
// It takes a slice of *entities.CmsProductImage as input and returns an error if any.
// The records are inserted into the database using the underlying ORM engine r.db.
// It also logs the "INSERT" operation and the inserted records.
func (r *CmsProductImageRepository) InsertMany(records []*entities.CmsProductImage) error {
	_, err := r.db.Insert(iterator.Map(records, func(item *entities.CmsProductImage) *entities.CmsProductImage {
		return item
	}))
	if err != nil {
		return err
	}

	r.log("INSERT", records)

	return nil
}

// Update updates a record in the CmsProductImage table based on the provided record's
// ProductImageId. It returns an error if the update operation fails.
// It also logs the update operation with the provided record.
// Example usage:
//
//	  record := &entities.CmsProductImage{}
//	  err := repository.Update(record)
//	  if err != nil {
//		   // handle error
//	  }
func (r *CmsProductImageRepository) Update(record *entities.CmsProductImage) error {
	_, err := r.db.Table(record.TableName()).Where("product_image_id = ?", record.ProductImageId).Update(record)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsProductImage{record})

	return nil
}

// UpdateMany updates multiple records in the CmsProductImageRepository.
// It takes a slice of records to update and returns an error if any update operation fails.
func (r *CmsProductImageRepository) UpdateMany(records []*entities.CmsProductImage) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}
	var sessionErr error
	rollback := false
	for _, record := range records {
		_, err = session.Table(record.TableName()).Where("product_image_id = ?", record.ProductImageId).Update(record)
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

// Delete sets the ActiveStatus of the given CmsProductImage record to 0
// and updates it using the Update method of the CmsProductImageRepository.
// It returns an error if the update operation fails.
func (r *CmsProductImageRepository) Delete(record *entities.CmsProductImage) error {
	record.ActiveStatus = 0
	_, err := r.db.Where("product_image_id = ?", record.ProductImageId).Cols("active_status").Update(record)
	if err == nil {
		r.log("DELETE", []*entities.CmsProductImage{record})
	}
	return err
}

// DeleteMany sets the ActiveStatus of each record in the input slice to 0
// and updates them using the UpdateMany method. It returns an error if
// the update operation fails.
func (r *CmsProductImageRepository) DeleteMany(records []*entities.CmsProductImage) error {
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
		_, err = session.Where("product_image_id = ?", record.ProductImageId).Cols("active_status").Update(record)
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

// log logs the given operation and payload to the audit log.
// The payload is serialized to JSON and mapped to an entity of type AuditLog,
// which is then passed to the audit.Log method.
// The operation indicates the type of operation being logged (e.g., "INSERT" or "UPDATE").
// The payload is a slice of entities of type CmsProductImage.
func (r *CmsProductImageRepository) log(op string, payload []*entities.CmsProductImage) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsProductImage) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      strconv.Itoa(item.ProductId),
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
