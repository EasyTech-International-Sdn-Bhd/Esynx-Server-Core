package debitnote

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

// CmsDebitNoteSalesRepository is a repository for managing CMS debit note data.
type CmsDebitNoteSalesRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
	c     *customer.CmsCustomerRepository
	d     *CmsDebitNoteDetailsRepository
}

// NewCmsDebitNoteSalesRepository creates a new instance of CmsDebitNoteSalesRepository with the given IRepository option.
func NewCmsDebitNoteSalesRepository(option *contracts.IRepository) *CmsDebitNoteSalesRepository {
	return &CmsDebitNoteSalesRepository{
		db:    option.Db,
		audit: option.Audit,
		c:     customer.NewCmsCustomerRepository(option),
		d:     NewCmsDebitNoteDetailsRepository(option),
	}
}

// Get retrieves a debit note with the given debitNoteCode from the CmsDebitNoteSalesRepository.
// If the debit note is found, it returns a pointer to the debit note entity and nil error.
// If the debit note is not found, it returns nil and nil error.
// If an error occurs while retrieving the debit note, it returns nil and the error.
func (r *CmsDebitNoteSalesRepository) Get(debitNoteCode string) (*entities.CmsDebitnoteSales, error) {
	var cmsDebitNote entities.CmsDebitnoteSales
	has, err := r.db.Where("dn_code=?", debitNoteCode).Get(&cmsDebitNote)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return &cmsDebitNote, nil
}

// GetWithCustomer retrieves a debit note along with its associated customer by the provided debitNoteCode.
// It first calls the Get method to retrieve the debit note. If the debit note is not found, it returns nil.
// Then, it calls the Get method of the customer repository to retrieve the customer associated with the debit note.
// Finally, it returns an instance of DebitNoteWithCustomer, which contains the debit note and customer, or an error if any occurs.
func (r *CmsDebitNoteSalesRepository) GetWithCustomer(debitNoteCode string) (*models.DebitNoteSalesWithCustomer, error) {
	dn, err := r.Get(debitNoteCode)
	if err != nil {
		return nil, err
	}
	if dn == nil {
		return nil, nil
	}
	c, err := r.c.Get(dn.CustCode)
	if err != nil {
		return nil, err
	}
	return &models.DebitNoteSalesWithCustomer{
		C: c,
		I: dn,
	}, nil
}

// GetWithItems returns the debit note with its associated items.
// It retrieves the debit note by the given debitNoteCode and calls the Get method.
// It then retrieves the items related to the debit note by calling the Get method of the CmsDebitNoteDetailsRepository.
// It returns the debit note with its items as a DebitNoteWithItems struct, or an error if any.
// The DebitNoteWithItems struct contains the debit note object (M) and an array of debit note details (D).
func (r *CmsDebitNoteSalesRepository) GetWithItems(debitNoteCode string) (*models.DebitNoteSalesWithItems, error) {
	cn, err := r.Get(debitNoteCode)
	if err != nil {
		return nil, err
	}
	details, err := r.d.Get(debitNoteCode)
	if err != nil {
		return nil, err
	}
	return &models.DebitNoteSalesWithItems{
		M: cn,
		D: details,
	}, nil
}

// GetByCustomer retrieves a list of debit notes for a specific customer.
// It takes the customer code as a parameter and returns a slice of CmsDebitnote entities and an error.
// The method queries the database using the cust_code and cancelled fields, orders the results by dn_date in descending order,
// and limits the number of records to 100. If an error occurs while querying the database, it will be returned along with a nil slice.
// If the query is successful, the method returns the list of records and a nil error.
func (r *CmsDebitNoteSalesRepository) GetByCustomer(custCode string) ([]*entities.CmsDebitnoteSales, error) {
	var records []*entities.CmsDebitnoteSales
	err := r.db.Where("cust_code = ? AND cancelled = ?", custCode, "F").OrderBy("dn_date DESC").Limit(100).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

// GetByDate retrieves CmsDebitnote records within a specific date range.
// It takes two parameters, from and to, representing the start and end dates.
// It returns a slice of pointers to CmsDebitnote entities and an error, if any.
func (r *CmsDebitNoteSalesRepository) GetByDate(from time.Time, to time.Time) ([]*entities.CmsDebitnoteSales, error) {
	var records []*entities.CmsDebitnoteSales
	err := r.db.Where(builder.Between{Col: "DATE(dn_date)", LessVal: from.Format("2006-01-02"), MoreVal: to.Format("2006-01-02")}.And(builder.Eq{"cancelled": "F"})).OrderBy("dn_date DESC").Limit(100).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsDebitNoteSalesRepository) Find(predicate *builder.Builder) ([]*entities.CmsDebitnoteSales, error) {
	var records []*entities.CmsDebitnoteSales
	var t entities.CmsDebitnoteSales
	err := r.db.SQL(predicate.From(t.TableName())).Find(&records)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, nil
	}
	return records, nil
}

// InsertMany inserts multiple debit notes into the database.
// It takes a slice of debit notes as parameter and returns an error if any.
// It maps each debit note to itself using the iterator.Map function.
// Then it performs the insertion using the db.Insert function.
// After successful insertion, it logs the operation as "INSERT" along with the inserted debit notes.
func (r *CmsDebitNoteSalesRepository) InsertMany(debitNotes []*entities.CmsDebitnoteSales) error {
	_, err := r.db.Insert(iterator.Map(debitNotes, func(item *entities.CmsDebitnoteSales) *entities.CmsDebitnoteSales {
		return item
	}))
	if err != nil {
		return err
	}

	r.log("INSERT", debitNotes)

	return nil
}

// Update updates a CmsDebitnote in the repository.
// It updates the record in the database with the matching dn_code.
// If there is an error during the update, it returns the error.
// It also logs the "UPDATE" operation with the updated CmsDebitnote.
func (r *CmsDebitNoteSalesRepository) Update(debitNote *entities.CmsDebitnoteSales) error {
	_, err := r.db.Where("dn_code = ?", debitNote.DnCode).Update(debitNote)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsDebitnoteSales{debitNote})

	return nil
}

// UpdateMany updates multiple debit notes in the database.
func (r *CmsDebitNoteSalesRepository) UpdateMany(debitNotes []*entities.CmsDebitnoteSales) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}
	var sessionErr error
	rollback := false
	for _, dn := range debitNotes {
		_, err = session.Where("dn_code = ?", dn.DnCode).Update(dn)
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

	r.log("UPDATE", debitNotes)

	return nil
}

// Delete sets the "Cancelled" attribute of the debitNote to "T" where dn_code matches.
// It logs the operation as "DELETE".
func (r *CmsDebitNoteSalesRepository) Delete(debitNote *entities.CmsDebitnoteSales) error {
	debitNote.Cancelled = "T"
	_, err := r.db.Where("dn_code = ?", debitNote.DnCode).Cols("cancelled").Update(debitNote)
	if err == nil {
		r.log("DELETE", []*entities.CmsDebitnoteSales{debitNote})
	}
	return err
}

// DeleteMany sets the `Cancelled` field of each debit note in the given slice to "T".
// It uses a session to ensure the operation is atomic.
// If any error occurs during the update process, it rolls back the session and returns the error.
// Otherwise, it commits the session and logs the operation as "DELETE".
func (r *CmsDebitNoteSalesRepository) DeleteMany(debitNotes []*entities.CmsDebitnoteSales) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}

	var sessionErr error
	rollback := false
	for _, debitNote := range debitNotes {
		debitNote.Cancelled = "T"
		_, err = session.Where("dn_code = ?", debitNote.DnCode).Cols("cancelled").Update(debitNote)
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

	r.log("DELETE", debitNotes)
	return nil
}

// log logs an audit record for a given operation and payload.
// It marshals the payload to JSON and creates an AuditLog object for each item in the payload.
// The AuditLog objects contain information about the operation, the table name, record ID, and record body.
// The AuditLog objects are then passed to the Log method of the IAuditLog interface for logging.
//
// Parameters:
// - op: The operation type of the audit record.
// - payload: The list of entities to be logged.
//
// Example usage:
//
//	r.log("INSERT", debitNotes)
//	r.log("UPDATE", []*entities.CmsDebitnoteSales{debitNote})
//	r.log("UPDATE", debitNotes)
//
// Requires:
// - entities.CmsDebitnoteSales declaration
// - entities.AuditLog declaration
// - iterator.Map function
// - IAuditLog interface implementation
func (r *CmsDebitNoteSalesRepository) log(op string, payload []*entities.CmsDebitnoteSales) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsDebitnoteSales) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      item.DnCode,
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
