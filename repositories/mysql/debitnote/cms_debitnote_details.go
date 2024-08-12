package debitnote

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

// CmsDebitNoteDetailsRepository represents a repository for managing CMS debit note details.
type CmsDebitNoteDetailsRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
	p     *stock.CmsProductRepository
}

// NewCmsDebitNoteDetailsRepository creates a new instance of CmsDebitNoteDetailsRepository with the given IRepository option.
func NewCmsDebitNoteDetailsRepository(option *contracts.IRepository) *CmsDebitNoteDetailsRepository {
	return &CmsDebitNoteDetailsRepository{
		db:    option.Db,
		audit: option.Audit,
		p:     stock.NewCmsProductRepository(option),
	}
}

// Get retrieves the debit note details associated with the given debitNoteCode.
// It returns a slice of entities.CmsDebitnoteDetails and an error. If an error occurs,
// the slice will be nil and the error will be non-nil. Otherwise, the slice will contain
// the retrieved details and the error will be nil.
func (r *CmsDebitNoteDetailsRepository) Get(debitNoteCode string) ([]*entities.CmsDebitnoteDetails, error) {
	var details []*entities.CmsDebitnoteDetails
	err := r.db.Where("dn_code = ? AND active_status = ?", debitNoteCode, 1).Find(&details)
	if err != nil {
		return nil, err
	}
	return details, nil
}

// GetMany retrieves multiple records of CmsDebitnoteDetails based on debitNoteCodes.
func (r *CmsDebitNoteDetailsRepository) GetMany(debitNoteCodes []string) ([]*entities.CmsDebitnoteDetails, error) {
	var details []*entities.CmsDebitnoteDetails
	err := r.db.In("dn_code", debitNoteCodes).Where("active_status = ?", 1).Find(&details)
	if err != nil {
		return nil, err
	}
	return details, nil
}

// GetWithProduct retrieves the debit note details with associated products.
//
// It first calls the Get method to retrieve the debit note details based on the debitNoteCode.
// Then, it collects the item codes from the details and calls the GetMany method to
// retrieve the associated products.
//
// Finally, it constructs DebitNoteDetailsWithProduct objects by matching the item codes
// with the product codes and returns the results.
//
// The method returns []*models.DebitNoteDetailsWithProduct and error.
func (r *CmsDebitNoteDetailsRepository) GetWithProduct(debitNoteCode string) ([]*models.DebitNoteDetailsWithProduct, error) {
	details, err := r.Get(debitNoteCode)
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
	var results []*models.DebitNoteDetailsWithProduct
	for _, detail := range details {
		for _, product := range products {
			if detail.ItemCode == product.ProductCode {
				results = append(results, &models.DebitNoteDetailsWithProduct{
					D: detail,
					P: product,
				})
			}
		}
	}
	return results, nil
}

func (r *CmsDebitNoteDetailsRepository) Find(predicate *builder.Builder) ([]*entities.CmsDebitnoteDetails, error) {
	var records []*entities.CmsDebitnoteDetails
	var t entities.CmsDebitnoteDetails
	err := r.db.SQL(predicate.From(t.TableName())).Find(&records)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, nil
	}
	return records, nil
}

// InsertMany inserts multiple CmsDebitnoteDetails into the database.
// It takes an array of CmsDebitnoteDetails as input and maps each item using iterator.Map.
// The mapped items are then inserted into the database using the Insert function of the db object.
// If there is any error during the insertion, it is returned.
// After the successful insertion, the log function is called to log the operation.
// Finally, it returns nil to indicate success.
func (r *CmsDebitNoteDetailsRepository) InsertMany(details []*entities.CmsDebitnoteDetails) error {
	_, err := r.db.Insert(iterator.Map(details, func(item *entities.CmsDebitnoteDetails) *entities.CmsDebitnoteDetails {
		return item
	}))
	if err != nil {
		return err
	}

	r.log("INSERT", details)

	return nil
}

// Update updates the details of a CmsDebitnoteDetails entity in the database.
// It takes a pointer to the entity as input and returns an error if any operation fails.
// The details are updated based on the "id" field of the entity.
// The method updates the corresponding row in the database table using the "Update" function of the database engine.
// If the update is successful, the method logs the operation as "UPDATE" with the updated details using the "log" method.
// If any error occurs during the update or logging process, it is returned as an error.
func (r *CmsDebitNoteDetailsRepository) Update(details *entities.CmsDebitnoteDetails) error {
	_, err := r.db.Table(details.TableName()).Where("id = ?", details.Id).Update(details)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsDebitnoteDetails{details})

	return nil
}

// UpdateMany updates multiple `CmsDebitnoteDetails` records in the database.
func (r *CmsDebitNoteDetailsRepository) UpdateMany(details []*entities.CmsDebitnoteDetails) error {
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

// Delete sets the ActiveStatus of the given CmsDebitnoteDetails object to 0 and updates it in the repository database. It returns an error if the update fails.
func (r *CmsDebitNoteDetailsRepository) Delete(details *entities.CmsDebitnoteDetails) error {
	details.ActiveStatus = 0
	return r.Update(details)
}

// DeleteMany deletes multiple CmsDebitnoteDetails by setting their ActiveStatus to 0.
// It then calls the UpdateMany method to update the changes in the database.
// Returns an error if there is any error during the update process.
func (r *CmsDebitNoteDetailsRepository) DeleteMany(details []*entities.CmsDebitnoteDetails) error {
	for _, detail := range details {
		detail.ActiveStatus = 0
	}
	return r.UpdateMany(details)
}

// log logs the specified operation and payload to the audit log. It marshals the payload into JSON format
// and creates an AuditLog object for each item in the payload. It then calls the Log method of the IAuditLog interface
// to save the audit log.
func (r *CmsDebitNoteDetailsRepository) log(op string, payload []*entities.CmsDebitnoteDetails) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsDebitnoteDetails) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      fmt.Sprintf("%s.%s", item.DnCode, item.ItemCode),
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
