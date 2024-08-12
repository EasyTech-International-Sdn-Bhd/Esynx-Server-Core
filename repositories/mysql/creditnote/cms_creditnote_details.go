package creditnote

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/models"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/mysql/stock"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"xorm.io/builder"
	"xorm.io/xorm"
)

// CmsCreditNoteDetailsRepository represents a repository for managing CMS credit note details.
//
// It contains references to the database engine, an audit log interface, and a product repository.
type CmsCreditNoteDetailsRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
	p     *stock.CmsProductRepository
}

// NewCmsCreditNoteDetailsRepository creates a new instance of CmsCreditNoteDetailsRepository
// with the given IRepository option. It initializes the db, audit, and p fields of the repository.
// The db field holds the xorm.Engine instance of the IRepository. The audit field holds the IAuditLog instance
// of the IRepository. The p field holds the CmsProductRepository instance created using NewCmsProductRepository function
// with the given IRepository option.
func NewCmsCreditNoteDetailsRepository(option *contracts.IRepository) *CmsCreditNoteDetailsRepository {
	return &CmsCreditNoteDetailsRepository{
		db:    option.Db,
		audit: option.Audit,
		p:     stock.NewCmsProductRepository(option),
	}
}

// Get retrieves credit note details by the given creditNoteCode.
//
// It returns a slice of *entities.CmsCreditnoteDetails and an error.
// If the credit note details are not found, it returns nil as the result with a non-nil error.
// If there is an error when retrieving the credit note details, it returns nil as the result
// with the non-nil error.
//
// The creditNoteCode parameter is the code of the credit note.
func (r *CmsCreditNoteDetailsRepository) Get(creditNoteCode string) ([]*entities.CmsCreditnoteDetails, error) {
	var details []*entities.CmsCreditnoteDetails
	err := r.db.Where("cn_code = ? AND active_status = ?", creditNoteCode, 1).Find(&details)
	if err != nil {
		return nil, err
	}
	return details, nil
}

// GetMany retrieves multiple credit note details based on the given credit note codes.
// It returns an array of CmsCreditnoteDetails and an error if any.
// The credit note codes are used to filter the results, and only active details are returned.
func (r *CmsCreditNoteDetailsRepository) GetMany(creditNoteCodes []string) ([]*entities.CmsCreditnoteDetails, error) {
	var details []*entities.CmsCreditnoteDetails
	err := r.db.In("cn_code", creditNoteCodes).Where("active_status = ?", 1).Find(&details)
	if err != nil {
		return nil, err
	}
	return details, nil
}

// GetWithProduct retrieves credit note details with associated products based on the credit note code.
// It first calls the Get method to obtain the details of the credit note.
// Then, it extracts the item codes from the details and calls the GetMany method to retrieve the associated products.
// Finally, it creates instances of CreditNoteDetailsWithProduct by combining each detail with its corresponding product.
// It returns the results as a slice of CreditNoteDetailsWithProduct instances.
// An error is returned if there is an issue with obtaining the credit note details or the associated products.
func (r *CmsCreditNoteDetailsRepository) GetWithProduct(creditNoteCode string) ([]*models.CreditNoteDetailsWithProduct, error) {
	details, err := r.Get(creditNoteCode)
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
	var results []*models.CreditNoteDetailsWithProduct
	for _, detail := range details {
		for _, product := range products {
			if detail.ItemCode == product.ProductCode {
				results = append(results, &models.CreditNoteDetailsWithProduct{
					D: detail,
					P: product,
				})
			}
		}
	}
	return results, nil
}

func (r *CmsCreditNoteDetailsRepository) Find(predicate *builder.Builder) ([]*entities.CmsCreditnoteDetails, error) {
	var records []*entities.CmsCreditnoteDetails
	var t entities.CmsCreditnoteDetails
	err := r.db.SQL(predicate.From(t.TableName())).Find(&records)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, nil
	}
	return records, nil
}

// InsertMany inserts multiple CmsCreditnoteDetails into the repository's database.
// It returns an error if the insertion operation fails. After successfully inserting
// the details, it logs the operation as "INSERT" along with the inserted details.
func (r *CmsCreditNoteDetailsRepository) InsertMany(details []*entities.CmsCreditnoteDetails) error {
	_, err := r.db.Insert(iterator.Map(details, func(item *entities.CmsCreditnoteDetails) *entities.CmsCreditnoteDetails {
		return item
	}))
	if err != nil {
		return err
	}

	r.log("INSERT", details)

	return nil
}

// Update updates the given `CmsCreditnoteDetails` object in the repository.
// It updates the corresponding record in the database based on the `id` field of the `details` object.
// If the update operation fails, it returns an error. Otherwise, it logs the UPDATE operation.
// Preconditions: `details` is not nil.
// Postconditions: The `details` object is updated in the repository.
func (r *CmsCreditNoteDetailsRepository) Update(details *entities.CmsCreditnoteDetails) error {
	_, err := r.db.Table(details.TableName()).Where("id = ?", details.Id).Update(details)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsCreditnoteDetails{details})

	return nil
}

// UpdateMany updates multiple CmsCreditnoteDetails records in the database.
// It starts a new database session before beginning a transaction.
// It iterates over the provided details, updating each record individually using its Id.
// If an error occurs during the update, the transaction is rolled back
// and the error is returned.
// Otherwise, the transaction is committed.
// Finally, the updated details are logged using the "UPDATE" operation type.
func (r *CmsCreditNoteDetailsRepository) UpdateMany(details []*entities.CmsCreditnoteDetails) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}
	var sessionErr error
	rollback := false
	for _, detail := range details {
		_, err = session.Table(detail.TableName()).Where("id = ?", detail.Id).Update(detail)
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

// Delete sets the active status of the given CmsCreditnoteDetails record to 0
// and updates it using the Update method of the CmsCreditNoteDetailsRepository.
// It returns an error if the update operation fails.
func (r *CmsCreditNoteDetailsRepository) Delete(details *entities.CmsCreditnoteDetails) error {
	details.ActiveStatus = 0
	_, err := r.db.Where("id = ?", details.Id).Cols("active_status").Update(details)
	if err == nil {
		r.log("DELETE", []*entities.CmsCreditnoteDetails{details})
	}
	return err
}

// DeleteMany sets the active status of each record in the input slice to 0
// and updates them using the UpdateMany method. It returns an error if
// the update operation fails.
func (r *CmsCreditNoteDetailsRepository) DeleteMany(details []*entities.CmsCreditnoteDetails) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}
	var sessionErr error
	rollback := false
	for _, detail := range details {
		detail.ActiveStatus = 0
		_, err = session.Where("id = ?", detail.Id).Cols("active_status").Update(detail)
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

	r.log("DELETE", details)

	return nil
}

// log is a method that is used to log an operation and its payload to the audit log.
//
// Parameters:
//   - op: a string representing the operation type.
//   - payload: a slice of *entities.CmsCreditnoteDetails representing the payload of the operation.
//
// Description:
//   - The method serializes the payload into a JSON record and creates an audit log entry for each item in the payload.
//   - Each audit log entry contains the operation type, record table, record ID, and record body.
//   - The created audit log entries are then passed to the r.audit.Log method for logging.
//
// Example:
//
//   - The log method can be used to log an "INSERT" operation with a slice of *entities.CmsCreditnoteDetails, as shown here:
//
//     err := r.log("INSERT", details)
//
//   - Another example is logging an "UPDATE" operation with a slice of *entities.CmsCreditnoteDetails, as shown here:
//
//     err := r.log("UPDATE", []*entities.CmsCreditnoteDetails{details})
func (r *CmsCreditNoteDetailsRepository) log(op string, payload []*entities.CmsCreditnoteDetails) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsCreditnoteDetails) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      fmt.Sprintf("%s.%s", item.ItemCode, item.CnCode),
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
