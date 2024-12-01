package invoice

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/models"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/sql/customer"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"time"
	"xorm.io/builder"
	"xorm.io/xorm"
)

// CmsInvoiceRepository is a type that represents a repository for CMS invoices.
// It contains a database engine (*xorm.Engine), an audit log contract (contracts.IAuditLog),
// a customer repository (*customer.CmsCustomerRepository), and an invoice details repository
// (*CmsInvoiceDetailsRepository).
type CmsInvoiceRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
	c     *customer.CmsCustomerRepository
	d     *CmsInvoiceDetailsRepository
}

// NewCmsInvoiceRepository creates a new instance of CmsInvoiceRepository with the given IRepository option.
// It initializes the db, audit, c, and d fields of the CmsInvoiceRepository struct.
//
// Parameters:
//   - option: Pointer to the IRepository struct that specifies the database engine,
//     user, app name, and audit log.
//
// Returns:
// - Pointer to the created CmsInvoiceRepository instance.
func NewCmsInvoiceRepository(option *contracts.IRepository) *CmsInvoiceRepository {
	return &CmsInvoiceRepository{
		db:    option.Db,
		audit: option.Audit,
		c:     customer.NewCmsCustomerRepository(option),
		d:     NewCmsInvoiceDetailsRepository(option),
	}
}

// Get retrieves the CmsInvoice with the given invoice code from the database.
// It returns a pointer to the CmsInvoice and an error if any.
func (r *CmsInvoiceRepository) Get(invoiceCode string) (*entities.CmsInvoice, error) {
	var cmsInvoice entities.CmsInvoice
	has, err := r.db.Where("invoice_code = ?", invoiceCode).Get(&cmsInvoice)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return &cmsInvoice, nil
}

// GetWithCustomer retrieves an invoice with its associated customer by invoice code.
// It first calls the Get method to retrieve the invoice by invoice code. If the invoice is not found,
// it returns nil. Otherwise, it calls the Get method of the CmsCustomerRepository to retrieve the customer
// by the customer code associated with the invoice. If the customer is not found, it returns an error.
// Otherwise, it creates a new InvoiceWithCustomer struct and assigns the retrieved invoice and customer to it,
// then returns the struct along with nil error.
func (r *CmsInvoiceRepository) GetWithCustomer(invoiceCode string) (*models.InvoiceWithCustomer, error) {
	invoice, err := r.Get(invoiceCode)
	if err != nil {
		return nil, err
	}
	if invoice == nil {
		return nil, nil
	}
	cmsCustomer, err := r.c.Get(invoice.CustCode)
	if err != nil {
		return nil, err
	}
	return &models.InvoiceWithCustomer{
		I: invoice,
		C: cmsCustomer,
	}, nil
}

// GetWithItems returns an InvoiceWithItems object containing an invoice and its associated details.
// It retrieves the invoice using the Get method, and then retrieves the details using the Get method of
// CmsInvoiceDetailsRepository. Finally, it returns an InvoiceWithItems object containing the retrieved
// invoice and details.
//
// Parameters:
// - invoiceCode: the code of the invoice to retrieve
//
// Returns:
// - InvoiceWithItems: the retrieved invoice and its associated details
// - error: an error if the retrieval process encounters any issues
func (r *CmsInvoiceRepository) GetWithItems(invoiceCode string) (*models.InvoiceWithItems, error) {
	iv, err := r.Get(invoiceCode)
	if err != nil {
		return nil, err
	}
	details, err := r.d.Get(invoiceCode)
	if err != nil {
		return nil, err
	}
	return &models.InvoiceWithItems{
		M: iv,
		D: details,
	}, nil
}

// GetByCustomer retrieves a list of CmsInvoice records by customer code.
// It queries the database for records where the cust_code equals the given custCode,
// and the cancelled field is set to "F". The results are ordered by invoice_date
// in descending order and limited to a maximum of 100 records.
// It returns a slice of CmsInvoice entities and an error if any occurred.
func (r *CmsInvoiceRepository) GetByCustomer(custCode string) ([]*entities.CmsInvoice, error) {
	var records []*entities.CmsInvoice
	err := r.db.Where("cust_code = ? AND cancelled = ?", custCode, "F").OrderBy("invoice_date DESC").Limit(100).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

// GetByDate retrieves a list of CmsInvoice records within the specified date range.
// It queries the database using the given 'from' and 'to' dates, filtering by the 'invoice_date' column.
// The result is ordered in descending order based on invoice date and limited to 100 records.
// Only invoices with the 'cancelled' field set to 'F' are included in the results.
// Returns the list of CmsInvoice records and any error encountered during the operation.
func (r *CmsInvoiceRepository) GetByDate(from time.Time, to time.Time) ([]*entities.CmsInvoice, error) {
	var records []*entities.CmsInvoice
	err := r.db.Where(builder.Between{Col: "DATE(invoice_date)", LessVal: from.Format("2006-01-02"), MoreVal: to.Format("2006-01-02")}.And(builder.Eq{"cancelled": "F"})).OrderBy("invoice_date DESC").Limit(100).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsInvoiceRepository) Find(predicate *builder.Builder) ([]*entities.CmsInvoice, error) {
	var records []*entities.CmsInvoice
	var t entities.CmsInvoice
	err := r.db.SQL(predicate.From(t.TableName())).Find(&records)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, nil
	}
	return records, nil
}

// InsertMany inserts multiple CmsInvoice entities into the database.
// The method takes a slice of pointers to CmsInvoice entities as input.
// It uses the `db` field of the CmsInvoiceRepository to perform the insertion.
// The method maps the input slice to itself using the `iterator.Map` function.
// It then calls the `Insert` method of the `db` field with the mapped slice as input.
// If the insertion is successful, the method logs the operation.
// The method returns an error if the insertion or logging fails.
func (r *CmsInvoiceRepository) InsertMany(invoices []*entities.CmsInvoice) error {
	_, err := r.db.Insert(iterator.Map(invoices, func(item *entities.CmsInvoice) *entities.CmsInvoice {
		return item
	}))
	if err != nil {
		return err
	}

	r.log("INSERT", invoices)

	return nil
}

// Update updates the given invoice in the repository.
//
// It updates the invoice in the database with the matching invoice code.
// If there is an error during the update operation, it returns the error.
// After the update, it logs the operation with the "UPDATE" operation type and the updated invoice.
// The logging is done by calling the `log` method of the repository using the "UPDATE" operation type
// and a slice containing the updated invoice as the payload.
//
// The `Update` method expects a pointer to a `CmsInvoice` struct as the input parameter.
// It does not return any value.
func (r *CmsInvoiceRepository) Update(invoice *entities.CmsInvoice) error {
	_, err := r.db.Where("invoice_code = ?", invoice.InvoiceCode).Omit("invoice_code").Update(invoice)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsInvoice{invoice})

	return nil
}

// Delete sets the Cancelled field of the given CmsInvoice record to "T"
// and updates it using the Update method. It returns an error if the
// update operation fails.
func (r *CmsInvoiceRepository) Delete(invoice *entities.CmsInvoice) error {
	invoice.Cancelled = "T"
	_, err := r.db.Where("invoice_code = ?", invoice.InvoiceCode).Cols("cancelled", "ref_no").Update(&entities.CmsInvoice{
		Cancelled: "T",
		RefNo:     fmt.Sprintf("DELETED-%s", time.Now().Format("20060102")),
	})
	if err == nil {
		r.log("DELETE", []*entities.CmsInvoice{invoice})
	}
	return err
}

// UpdateMany updates multiple invoices in the database.
// It iterates through the invoices to update each one.
// If any update fails, the method returns the error.
// The method also logs the update operation with the invoices.
func (r *CmsInvoiceRepository) UpdateMany(invoices []*entities.CmsInvoice) error {
	for _, inv := range invoices {
		_, err := r.db.Where("invoice_code = ?", inv.InvoiceCode).Omit("invoice_code").Update(inv)
		if err != nil {
			return err
		}
	}

	r.log("UPDATE", invoices)

	return nil
}

// DeleteMany sets the Cancelled field of each record in the input slice to "T"
// and updates them using the UpdateMany method. It returns an error if the
// update operation fails.
func (r *CmsInvoiceRepository) DeleteMany(invoices []*entities.CmsInvoice) error {
	ids := iterator.Map(invoices, func(item *entities.CmsInvoice) string {
		return item.InvoiceCode
	})

	_, err := r.db.In("invoice_code", ids).Cols("cancelled", "ref_no").Update(&entities.CmsInvoice{
		Cancelled: "T",
		RefNo:     fmt.Sprintf("DELETED-%s", time.Now().Format("20060102")),
	})
	if err != nil {
		return err
	}

	r.log("DELETE", invoices)

	return nil
}

func (r *CmsInvoiceRepository) DeleteByAny(predicate *builder.Builder) ([]*entities.CmsInvoice, error) {
	var t entities.CmsInvoice

	var records []*entities.CmsInvoice
	err := r.db.SQL(predicate.From(t.TableName())).Find(&records)
	if err != nil {
		return nil, err
	}

	err = r.DeleteMany(records)
	if err != nil {
		return nil, err
	}

	r.log("DELETE", records)

	return records, nil
}

// log is a method used to record audit logs for operations performed on CmsInvoiceRepository.
//
// It takes an operation type string and a slice of CmsInvoice pointers as input.
// It marshals the payload into a JSON string and constructs AuditLog objects for each CmsInvoice in the payload.
// The AuditLog objects contain information about the operation type, the table name, invoice code, and the record body.
// The constructed AuditLog objects are then logged using the audit.Log method.
//
// Parameters:
//   - op: The operation type string.
//   - payload: A slice of CmsInvoice pointers representing the payload for the audit log.
//     Each CmsInvoice contains information about the invoice.
//
// Example Usage:
//
//	r.log("INSERT", invoices)
//	r.log("UPDATE", []*entities.CmsInvoice{invoice})
//
// Note: The audit.Log method is not shown as it belongs to a different package.
func (r *CmsInvoiceRepository) log(op string, payload []*entities.CmsInvoice) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsInvoice) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      item.InvoiceCode,
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
