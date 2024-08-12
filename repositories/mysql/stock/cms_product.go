package stock

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/models"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"strconv"
	"strings"
	"xorm.io/builder"
	"xorm.io/xorm"
)

// CmsProductRepository represents a repository for managing CMS products.
//
// It contains references to other repositories for managing related entities such as attachments,
// batches, images, price tags, UOM prices, and warehouse stocks.
//
// Note: This documentation only provides a brief overview of the type and its purpose. Please refer to the
// comments in the code for more detailed information about each method and its usage.
type CmsProductRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
	a     *CmsProductAtchRepository
	b     *CmsProductBatchRepository
	i     *CmsProductImageRepository
	t     *CmsProductPriceTagRepository
	p     *CmsProductUomPriceRepository
	w     *CmsWarehouseStockRepository
}

// NewCmsProductRepository creates a new instance of CmsProductRepository with the given IRepository option.
func NewCmsProductRepository(option *contracts.IRepository) *CmsProductRepository {
	return &CmsProductRepository{
		db:    option.Db,
		audit: option.Audit,
		a:     NewCmsProductAtchRepository(option),
		b:     NewCmsProductBatchRepository(option),
		i:     NewCmsProductImageRepository(option),
		t:     NewCmsProductPriceTagRepository(option),
		p:     NewCmsProductUomPriceRepository(option),
		w:     NewCmsWarehouseStockRepository(option),
	}
}

// Get retrieves a CmsProduct from the database based on the provided productCode.
// If the CmsProduct is found, it is returned along with a nil error.
// If the CmsProduct is not found, nil and nil error are returned.
// If an error occurs during the retrieval process, nil and the error are returned.
// Signature: func (r *CmsProductRepository) Get(productCode string) (*entities.CmsProduct, error)
func (r *CmsProductRepository) Get(productCode string) (*entities.CmsProduct, error) {
	var cmsProduct entities.CmsProduct
	has, err := r.db.Where("product_code = ?", productCode).Get(&cmsProduct)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return &cmsProduct, nil
}

// GetMany retrieves multiple CmsProduct entities based on the provided product codes.
// It executes a database query using the product codes and populates the results into
// a slice of CmsProduct entities. Any error encountered during the query execution
// will be returned along with a nil slice of CmsProduct entities.
func (r *CmsProductRepository) GetMany(productCodes []string) ([]*entities.CmsProduct, error) {
	var cmsProducts []*entities.CmsProduct
	err := r.db.In("product_code", productCodes).Find(&cmsProducts)
	if err != nil {
		return nil, err
	}
	return cmsProducts, nil
}

func (r *CmsProductRepository) Find(predicate *builder.Builder) ([]*entities.CmsProduct, error) {
	var records []*entities.CmsProduct
	var t entities.CmsProduct
	err := r.db.SQL(predicate.From(t.TableName())).Find(&records)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, nil
	}
	return records, nil
}

// Search searches for CmsProduct records matching the given predicate.
// It splits the predicate into tokens and constructs a list of conditions
// for the product_code and product_name columns using the builder.Like operator.
// It then performs a WHERE query with the OR operator on the constructed conditions
// and populates the result in the records slice.
func (r *CmsProductRepository) Search(predicate string) ([]*entities.CmsProduct, error) {
	var records []*entities.CmsProduct
	tokens := strings.Split(predicate, " ")
	var where []builder.Cond
	for _, token := range tokens {
		where = append(where, builder.Like{"product_code", token})
		where = append(where, builder.Like{"product_name", token})
	}
	err := r.db.Where(builder.Or(where...)).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

// InsertMany inserts multiple records into the CmsProductRepository.
func (r *CmsProductRepository) InsertMany(records []*entities.CmsProduct) error {
	_, err := r.db.Insert(iterator.Map(records, func(item *entities.CmsProduct) *entities.CmsProduct {
		return item
	}))
	if err != nil {
		return err
	}

	r.log("INSERT", records)

	return nil
}

// Update updates the specified record in the CmsProductRepository.
// It updates the record in the database based on the provided record's ProductCode.
// If an error occurs during the update process, it returns the error.
// It also logs the update operation with the updated record.
// The logged operation type is "UPDATE" and the payload is an array containing the updated record.
//
// Example usage:
//
//	err := repository.Update(&record)
//	if err != nil {
//	    fmt.Println("Error updating record:", err)
//	}
func (r *CmsProductRepository) Update(record *entities.CmsProduct) error {
	_, err := r.db.Table(record.TableName()).Where("product_code = ?", record.ProductCode).Update(record)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsProduct{record})

	return nil
}

// UpdateMany updates multiple records in the CmsProductRepository.
// It starts a new session, updates each record in the given array,
// and commits the changes at the end. If any error occurs during
// the update process, it rolls back the session and returns the error.
// The method returns a nil error if all updates are successful.
// It logs the "UPDATE" operation with the updated records.
func (r *CmsProductRepository) UpdateMany(records []*entities.CmsProduct) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}
	var sessionErr error
	rollback := false
	for _, product := range records {
		_, err = session.Table(product.TableName()).Where("product_code = ?", product.ProductCode).Update(product)
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

// Delete sets the product status of the given CmsProduct record to 0
// and updates it using the Update method of the CmsProductRepository.
// It returns an error if the update operation fails.
func (r *CmsProductRepository) Delete(record *entities.CmsProduct) error {
	record.ProductStatus = 0
	_, err := r.db.Where("product_code = ?", record.ProductCode).Cols("product_status").Update(record)
	if err == nil {
		r.log("DELETE", []*entities.CmsProduct{record})
	}
	return err
}

// DeleteMany sets the ProductStatus of each record in the input slice to 0
// and updates them using the UpdateMany method. It returns an error if
// the update operation fails.
func (r *CmsProductRepository) DeleteMany(records []*entities.CmsProduct) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}
	var sessionErr error
	rollback := false
	for _, record := range records {
		record.ProductStatus = 0
		_, err = session.Where("product_code = ?", record.ProductCode).Cols("product_status").Update(record)
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

// GetWithDetails retrieves a product with all its details based on the provided product code.
// It calls multiple repository methods to retrieve the required details and constructs a
// ProductWithDetails object containing the product and its associated details. If any of the
// repository methods return an error, GetWithDetails returns nil and the error.
func (r *CmsProductRepository) GetWithDetails(productCode string) (*models.ProductWithDetails, error) {
	p, err := r.Get(productCode)
	if err != nil {
		return nil, err
	}
	a, err := r.a.Get(productCode)
	if err != nil {
		return nil, err
	}
	b, err := r.b.Get(productCode)
	if err != nil {
		return nil, err
	}
	i, err := r.i.Get(strconv.FormatUint(p.ProductId, 10))
	if err != nil {
		return nil, err
	}
	t, err := r.t.Get(productCode)
	if err != nil {
		return nil, err
	}
	pr, err := r.p.Get(productCode)
	if err != nil {
		return nil, err
	}
	w, err := r.w.Get(productCode)
	return &models.ProductWithDetails{
		P: p,
		A: a,
		I: i,
		U: pr,
		T: t,
		B: b,
		W: w,
	}, nil
}

// log writes the audit logs for the given operation and payload.
// It marshals the payload into JSON format and creates an AuditLog object
// for each item in the payload. The AuditLog object contains the operation
// type, record table, record ID, and record body. The audit logs are then
// passed to the IAuditLog.Log() method for logging.
//
// Parameters:
//   - op: The operation type (e.g., "INSERT", "UPDATE").
//   - payload: The payload containing the records to be logged.
func (r *CmsProductRepository) log(op string, payload []*entities.CmsProduct) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsProduct) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      item.ProductCode,
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
