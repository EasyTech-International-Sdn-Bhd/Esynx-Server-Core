package invoice

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/models"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/mysql/stock"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"xorm.io/xorm"
)

// CmsInvoiceDetailsRepository represents a repository for managing CMS invoice details.
type CmsInvoiceDetailsRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
	p     *stock.CmsProductRepository
}

// NewCmsInvoiceDetailsRepository creates a new instance of CmsInvoiceDetailsRepository with the given IRepository option.
// It initializes the db, audit, and p fields of the CmsInvoiceDetailsRepository struct.
//
// Parameters:
//   - option: Pointer to the IRepository struct that specifies the database engine,
//     user, app name, and audit log.
//
// Returns:
// - Pointer to the created CmsInvoiceDetailsRepository instance.
func NewCmsInvoiceDetailsRepository(option *contracts.IRepository) *CmsInvoiceDetailsRepository {
	return &CmsInvoiceDetailsRepository{
		db:    option.Db,
		audit: option.Audit,
		p:     stock.NewCmsProductRepository(option),
	}
}

// Get retrieves the invoice details associated with the given invoice code.
// It queries the database using the invoice code and active status as filters.
// It returns a slice of *entities.CmsInvoiceDetails and an error if any occurred.
func (r *CmsInvoiceDetailsRepository) Get(invoiceCode string) ([]*entities.CmsInvoiceDetails, error) {
	var details []*entities.CmsInvoiceDetails
	err := r.db.Where("invoice_code = ? AND active_status = ?", invoiceCode, 1).Find(&details)
	if err != nil {
		return nil, err
	}
	return details, nil
}

// GetMany retrieves multiple invoice details based on the given invoice codes.
// It queries the database with the provided invoice codes and active status equal to 1.
// If there are no matching records, it returns an empty slice.
// If there is an error during the query, it returns the error.
// The returned value is a slice of entities.CmsInvoiceDetails.
// Each entities.CmsInvoiceDetails represents an invoice detail entry in the database.
//
// Parameters:
// - invoiceCodes: a slice of strings representing the invoice codes to search for.
//
// Returns:
// - a slice of entities.CmsInvoiceDetails: the found invoice details.
// - error: any error that occurred during the query.
func (r *CmsInvoiceDetailsRepository) GetMany(invoiceCodes []string) ([]*entities.CmsInvoiceDetails, error) {
	var details []*entities.CmsInvoiceDetails
	err := r.db.In("invoice_code", invoiceCodes).Where("active_status = ?", 1).Find(&details)
	if err != nil {
		return nil, err
	}
	return details, nil
}

// GetWithProduct retrieves invoice details with associated product information
// for a given invoice code.
func (r *CmsInvoiceDetailsRepository) GetWithProduct(invoiceCode string) ([]*models.InvoiceDetailsWithProduct, error) {
	details, err := r.Get(invoiceCode)
	if err != nil {
		return nil, err
	}
	var productCodes []string
	for _, detail := range details {
		if detail.ItemCode != "" {
			productCodes = append(productCodes, detail.ItemCode)
		}
	}
	products, err := r.p.GetMany(productCodes)
	if err != nil {
		return nil, err
	}
	var results []*models.InvoiceDetailsWithProduct
	for _, detail := range details {
		for _, product := range products {
			if detail.ItemCode == product.ProductCode {
				results = append(results, &models.InvoiceDetailsWithProduct{
					D: detail,
					P: product,
				})
			}
		}
	}
	return results, nil
}

func (r *CmsInvoiceDetailsRepository) Find(predicate interface{}) ([]*entities.CmsInvoiceDetails, error) {
	var records []*entities.CmsInvoiceDetails
	err := r.db.Where(predicate).Find(&records)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, nil
	}
	return records, nil
}

// InsertMany inserts multiple CmsInvoiceDetails into the database.
// It returns an error if the insertion fails.
func (r *CmsInvoiceDetailsRepository) InsertMany(details []*entities.CmsInvoiceDetails) error {
	_, err := r.db.Insert(iterator.Map(details, func(item *entities.CmsInvoiceDetails) *entities.CmsInvoiceDetails {
		return item
	}))
	if err != nil {
		return err
	}

	r.log("INSERT", details)

	return nil
}

// Update updates the details of a CMS invoice in the repository.
// It takes a pointer to the entities.CmsInvoiceDetails struct as input,
// and returns an error if any.
func (r *CmsInvoiceDetailsRepository) Update(details *entities.CmsInvoiceDetails) error {
	_, err := r.db.Where("id = ?", details.Id).Update(details)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsInvoiceDetails{details})

	return nil
}

// UpdateMany updates multiple CmsInvoiceDetails in a transaction.
// It begins a session, updates each detail one by one using their ID,
// and commits the transaction if all updates are successful.
// If any error occurs during the updates, the transaction is rolled back and
// the first encountered error is returned.
// The method also logs the update operation and the updated details.
// It returns an error if the transaction fails to begin or commit.
func (r *CmsInvoiceDetailsRepository) UpdateMany(details []*entities.CmsInvoiceDetails) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}
	var sessionErr error
	rollback := false
	for _, detail := range details {
		_, err = session.Where("id = ?", detail.Id).Update(detail)
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

	r.log("UPDATE", details)

	return nil
}

// Delete sets the ActiveStatus of the given CmsInvoiceDetails to 0 and calls the Update method to persist the changes.
func (r *CmsInvoiceDetailsRepository) Delete(details *entities.CmsInvoiceDetails) error {
	details.ActiveStatus = 0
	return r.Update(details)
}

// DeleteMany deletes multiple CmsInvoiceDetails by setting their ActiveStatus to 0
// and calling the UpdateMany method to update them in the repository.
func (r *CmsInvoiceDetailsRepository) DeleteMany(details []*entities.CmsInvoiceDetails) error {
	for _, detail := range details {
		detail.ActiveStatus = 0
	}
	return r.UpdateMany(details)
}

// log logs the operation and payload to the audit log.
// op is the operation type (e.g., INSERT, UPDATE).
// The payload is an array of CmsInvoiceDetails entities.
// It marshals the payload to JSON and creates AuditLog entities for each item, which are then logged to the IAuditLog instance.
func (r *CmsInvoiceDetailsRepository) log(op string, payload []*entities.CmsInvoiceDetails) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsInvoiceDetails) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      fmt.Sprintf("%s.%s", item.InvoiceCode, item.ItemCode),
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
