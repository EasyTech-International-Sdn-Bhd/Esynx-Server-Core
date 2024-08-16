package invoice

import (
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

// CmsInvoiceSalesRepository represents a repository for managing sales invoices in a CMS.
// It contains a database engine, an audit logger, a customer repository, and an invoice details repository.
type CmsInvoiceSalesRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
	c     *customer.CmsCustomerRepository
	d     *CmsInvoiceDetailsRepository
}

// NewCmsInvoiceSalesRepository creates a new instance of CmsInvoiceSalesRepository
// with the provided IRepository option. It initializes the db, audit, c, and d
// fields of the CmsInvoiceSalesRepository struct.
//
// Parameters:
//   - option: Pointer to the IRepository struct that specifies the database engine,
//     user, app name, and audit log.
//
// Returns:
// - Pointer to the created CmsInvoiceSalesRepository instance.
func NewCmsInvoiceSalesRepository(option *contracts.IRepository) *CmsInvoiceSalesRepository {
	return &CmsInvoiceSalesRepository{
		db:    option.Db,
		audit: option.Audit,
		c:     customer.NewCmsCustomerRepository(option),
		d:     NewCmsInvoiceDetailsRepository(option),
	}
}

// Get retrieves an instance of CmsInvoiceSales with the given invoice code.
// It returns a pointer to the retrieved CmsInvoiceSales and an error.
// If the invoice is not found, it returns nil and nil.
func (r *CmsInvoiceSalesRepository) Get(invoiceCode string) (*entities.CmsInvoiceSales, error) {
	var cmsInvoice entities.CmsInvoiceSales
	has, err := r.db.Where("invoice_code = ?", invoiceCode).Get(&cmsInvoice)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return &cmsInvoice, nil
}

// GetWithCustomer retrieves an `InvoiceSalesWithCustomer` object by the given `invoiceCode`.
// It fetches the corresponding `CmsInvoiceSales` record using the `Get` method.
// Then, it retrieves the associated `CmsCustomer` record using the `Get` method.
// Finally, it creates and returns an `InvoiceSalesWithCustomer` object containing the fetched records.
func (r *CmsInvoiceSalesRepository) GetWithCustomer(invoiceCode string) (*models.InvoiceSalesWithCustomer, error) {
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
	return &models.InvoiceSalesWithCustomer{
		I: invoice,
		C: cmsCustomer,
	}, nil
}

// GetWithItems returns an InvoiceSalesWithItems object containing the invoice sales details and the associated invoice items.
// It retrieves the invoice sales by the provided invoice code using the Get method.
// Then, it retrieves the invoice details by the same invoice code using the Get method of the CmsInvoiceDetailsRepository.
// It returns a pointer to the InvoiceSalesWithItems object and error if any occurred.
func (r *CmsInvoiceSalesRepository) GetWithItems(invoiceCode string) (*models.InvoiceSalesWithItems, error) {
	iv, err := r.Get(invoiceCode)
	if err != nil {
		return nil, err
	}
	details, err := r.d.Get(invoiceCode)
	if err != nil {
		return nil, err
	}
	return &models.InvoiceSalesWithItems{
		M: iv,
		D: details,
	}, nil
}

// GetByCustomer retrieves a list of invoices for a specific customer.
// It queries the database using the provided custCode, filtering out cancelled invoices.
// The results are ordered by invoice_date in descending order and limited to 100 records.
// It returns a slice of entities.CmsInvoiceSales and an error if any.
func (r *CmsInvoiceSalesRepository) GetByCustomer(custCode string) ([]*entities.CmsInvoiceSales, error) {
	var records []*entities.CmsInvoiceSales
	err := r.db.Where("cust_code = ? AND cancelled = ?", custCode, "F").OrderBy("invoice_date DESC").Limit(100).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

// GetByDate retrieves a list of CmsInvoiceSales records within a given date range.
// It takes two parameters 'from' and 'to', representing the start and end date respectively.
// It returns a slice of *entities.CmsInvoiceSales and an error if there was a problem retrieving the records.
func (r *CmsInvoiceSalesRepository) GetByDate(from time.Time, to time.Time) ([]*entities.CmsInvoiceSales, error) {
	var records []*entities.CmsInvoiceSales
	err := r.db.Where(builder.Between{Col: "DATE(invoice_date)", LessVal: from.Format("2006-01-02"), MoreVal: to.Format("2006-01-02")}.And(builder.Eq{"cancelled": "F"})).OrderBy("invoice_date DESC").Limit(100).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsInvoiceSalesRepository) Find(predicate *builder.Builder) ([]*entities.CmsInvoiceSales, error) {
	var records []*entities.CmsInvoiceSales
	var t entities.CmsInvoiceSales
	err := r.db.SQL(predicate.From(t.TableName())).Find(&records)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, nil
	}
	return records, nil
}

// InsertMany inserts multiple invoices into the database. It takes a slice of CmsInvoiceSales
// and inserts each item into the database using the db.Insert() method. If any error occurs during the insertion,
// it returns the error. It then calls the log() method with the operation type "INSERT" and the inserted invoices.
// Finally, it returns nil to indicate successful insertion.
func (r *CmsInvoiceSalesRepository) InsertMany(invoices []*entities.CmsInvoiceSales) error {
	_, err := r.db.Insert(iterator.Map(invoices, func(item *entities.CmsInvoiceSales) *entities.CmsInvoiceSales {
		return item
	}))
	if err != nil {
		return err
	}

	r.log("INSERT", invoices)

	return nil
}

// Update updates the given `invoice` in the `CmsInvoiceSalesRepository`.
// It updates the corresponding record in the database table by matching the `invoice`'s `InvoiceCode` field.
// If any error occurs during the update operation, it returns that error.
// After the update, it logs the update operation with the `"UPDATE"` operation type and the updated invoice in the repository.
// The logging is performed by calling the `log` method of the repository, passing in the `"UPDATE"` operation type and
// a slice containing only the `invoice` as its payload.
// The `log` method converts the payload into a JSON string and creates an `AuditLog` record
// with the operation type, corresponding record table name, record ID (`InvoiceCode`), and the record body (JSON string).
// Then, it logs this `AuditLog` record using the `Log` method of the `audit` field of the repository.
func (r *CmsInvoiceSalesRepository) Update(invoice *entities.CmsInvoiceSales) error {
	_, err := r.db.Where("invoice_code = ?", invoice.InvoiceCode).Update(invoice)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsInvoiceSales{invoice})

	return nil
}

// Delete sets the "Cancelled" field of the provided invoice to "T" and updates
// it using the Update method of the CmsInvoiceSalesRepository. It returns an error
// if the update operation fails.
func (r *CmsInvoiceSalesRepository) Delete(invoice *entities.CmsInvoiceSales) error {
	invoice.Cancelled = "T"
	_, err := r.db.Where("invoice_code = ?", invoice.InvoiceCode).Cols("cancelled").Update(invoice)
	if err == nil {
		r.log("DELETE", []*entities.CmsInvoiceSales{invoice})
	}
	return err
}

// UpdateMany updates multiple invoices in the database.
// It loops through the invoices and updates each one.
// If an error occurs during the update, it returns the error.
// Finally, the updated invoices are logged.
func (r *CmsInvoiceSalesRepository) UpdateMany(invoices []*entities.CmsInvoiceSales) error {
	for _, inv := range invoices {
		_, err := r.db.Where("invoice_code = ?", inv.InvoiceCode).Update(inv)
		if err != nil {
			return err
		}
	}

	r.log("UPDATE", invoices)

	return nil
}

// DeleteMany sets the "Cancelled" field to "T" for each invoice in the provided slice
// and updates them in the database using a single bulk operation. It returns an error if the update
// operation fails.
func (r *CmsInvoiceSalesRepository) DeleteMany(invoices []*entities.CmsInvoiceSales) error {
	ids := iterator.Map(invoices, func(item *entities.CmsInvoiceSales) string {
		return item.InvoiceCode
	})

	_, err := r.db.In("invoice_code", ids).Cols("cancelled").Update(&entities.CmsInvoiceSales{
		Cancelled: "T",
	})
	if err != nil {
		return err
	}

	r.log("DELETE", invoices)
	return nil
}

// log logs a given operation and payload to the audit log using the audit interface.
// The payload is transformed into an array of AuditLog objects with each item containing information about the operation,
// the record table, record ID, and record body. The record body is the JSON representation of the payload.
// The log method is called by various methods in the CmsInvoiceSalesRepository to log operations such as INSERT and UPDATE.
// The audit.Log method is responsible for sending the logs to the audit log storage.
//
// Parameters:
// - op: The operation type to be logged (e.g., "INSERT", "UPDATE").
// - payload: The array of CMS invoice sales entities to be logged.
//
// Note: This method does not return any values.
func (r *CmsInvoiceSalesRepository) log(op string, payload []*entities.CmsInvoiceSales) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsInvoiceSales) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      item.InvoiceCode,
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
