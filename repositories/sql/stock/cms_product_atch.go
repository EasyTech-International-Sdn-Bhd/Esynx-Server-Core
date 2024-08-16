package stock

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"xorm.io/builder"
	"xorm.io/xorm"
)

// CmsProductAtchRepository represents a repository for managing CMS product attachments.
type CmsProductAtchRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
}

// NewCmsProductAtchRepository creates a new instance of CmsProductAtchRepository with the given IRepository option.
func NewCmsProductAtchRepository(option *contracts.IRepository) *CmsProductAtchRepository {
	return &CmsProductAtchRepository{
		db:    option.Db,
		audit: option.Audit,
	}
}

// Get retrieves the CMS product attachments for a given product code.
// It queries the database for records where the product code matches the provided string,
// and the active status is set to 1. The method returns a slice of *entities.CmsProductAtch
// and an error. If an error occurs during the query, the method returns nil and the error.
// Example usage:
//
//	productAtchs, err := repository.Get("PRODUCT_CODE")
func (r *CmsProductAtchRepository) Get(productCode string) ([]*entities.CmsProductAtch, error) {
	var record []*entities.CmsProductAtch
	err := r.db.Where("product_code = ? AND active_status = ?", productCode, 1).Find(&record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (r *CmsProductAtchRepository) Find(predicate *builder.Builder) ([]*entities.CmsProductAtch, error) {
	var records []*entities.CmsProductAtch
	var t entities.CmsProductAtch
	err := r.db.SQL(predicate.From(t.TableName())).Find(&records)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, nil
	}
	return records, nil
}

// InsertMany inserts multiple records into the database.
// It takes a slice of *entities.CmsProductAtch records as input.
// It uses a map function to convert each record to the appropriate type.
// It then calls the Insert method of the database connection and passes the converted records.
// If there is an error during the insert operation, it returns the error.
// After the insert operation, it logs the "INSERT" operation and the inserted records.
// It returns nil if the operations are successful.
func (r *CmsProductAtchRepository) InsertMany(records []*entities.CmsProductAtch) error {
	_, err := r.db.Insert(iterator.Map(records, func(item *entities.CmsProductAtch) *entities.CmsProductAtch {
		return item
	}))
	if err != nil {
		return err
	}

	r.log("INSERT", records)

	return nil
}

// Update updates the given record in the CmsProductAtchRepository.
// It updates the record in the database by matching its ID.
// If an error occurs during the update process, it returns the error.
// After the update, it logs the operation type as "UPDATE" and the updated record.
func (r *CmsProductAtchRepository) Update(record *entities.CmsProductAtch) error {
	_, err := r.db.Where("id = ?", record.Id).Update(record)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsProductAtch{record})

	return nil
}

// Delete sets the active status of the given CmsProductAtch record to 0
// and updates it using the Update method of the CmsProductAtchRepository.
// It returns an error if the update operation fails.
func (r *CmsProductAtchRepository) Delete(record *entities.CmsProductAtch) error {
	record.ActiveStatus = 0
	_, err := r.db.Where("id = ?", record.Id).Cols("active_status").Update(record)
	if err == nil {
		r.log("DELETE", []*entities.CmsProductAtch{record})
	}
	return err
}

func (r *CmsProductAtchRepository) UpdateMany(records []*entities.CmsProductAtch) error {
	for _, record := range records {
		_, err := r.db.Where("id = ?", record.Id).Update(record)
		if err != nil {
			return err
		}
	}

	r.log("UPDATE", records)

	return nil
}

func (r *CmsProductAtchRepository) DeleteMany(records []*entities.CmsProductAtch) error {
	ids := iterator.Map(records, func(item *entities.CmsProductAtch) uint64 {
		return item.Id
	})

	_, err := r.db.In("id", ids).Cols("active_status").Update(&entities.CmsProductAtch{
		ActiveStatus: 0,
	})
	if err != nil {
		return err
	}

	r.log("DELETE", records)
	return nil
}

// log is a method of the CmsProductAtchRepository struct that is used to log audit information.
// It takes two parameters: op, which represents the operation type, and payload, which is a slice of CmsProductAtch entities.
// It converts the payload into a JSON string and maps each item in the payload to an AuditLog entity.
// The AuditLog entity contains information about the operation type, table name, product code, and the JSON representation of the payload.
// The mapped AuditLog entities are passed to the audit.Log method to log the audit information.
//
// Usage Example:
// ```go
//
//	func (r *CmsProductAtchRepository) InsertMany(records []*entities.CmsProductAtch) error {
//	  _, err := r.db.Insert(iterator.Map(records, func(item *entities.CmsProductAtch) *entities.CmsProductAtch {
//	    return item
//	  }))
//	  if err != nil {
//	    return err
//	  }
//
//	  r.log("INSERT", records)
//
//	  return nil
//	}
//
// ```
// Usage Example:
// ```go
//
//	func (r *CmsProductAtchRepository) Update(record *entities.CmsProductAtch) error {
//	  _, err := r.db.Where("id = ?", record.Id).Update(record)
//	  if err != nil {
//	    return err
//	  }
//
//	  r.log("UPDATE", []*entities.CmsProductAtch{record})
//
//	  return nil
//	}
//
// ```
// Usage Example:
// ```go
//
//	func (r *CmsProductAtchRepository) UpdateMany(records []*entities.CmsProductAtch) error {
//	  session := r.db.NewSession()
//	  defer session.Close()
//	  err := session.Begin()
//	  if err != nil {
//	    return err
//	  }
//	  var sessionErr error
//	  rollback := false
//	  for _, product := range records {
//	    _, err = session.Where("id = ?", product.Id).Update(product)
//	    if err != nil {
//	      rollback = true
//	      sessionErr = err
//	      break
//	    }
//	  }
//	  if rollback {
//	    err := session.Rollback()
//	    if err != nil {
//	      return err
//	    }
//	    return sessionErr
//	  }
//	  err = session.Commit()
//	  if err != nil {
//	    return err
//	  }
//
//	  r.log("UPDATE", records)
//
//	  return nil
//	}
//
// IAuditLog is an interface that defines the Log method used by the log method:
// ```go
//
//	type IAuditLog interface {
//	  Log(audits []*entities.AuditLog)
//	}
//
// ```
//
// Note: This method is not directly called outside of the CmsProductAtchRepository struct.
func (r *CmsProductAtchRepository) log(op string, payload []*entities.CmsProductAtch) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsProductAtch) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      item.ProductCode,
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
