package debitnote

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/models"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/mysql/customer"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"slices"
	"time"
	"xorm.io/builder"
	"xorm.io/xorm"
)

// CmsDebitNoteRepository is a repository for managing CMS debit note data.
type CmsDebitNoteRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
	c     *customer.CmsCustomerRepository
	d     *CmsDebitNoteDetailsRepository
	s     *CmsDebitNoteSalesRepository
}

// NewCmsDebitNoteRepository creates a new instance of CmsDebitNoteRepository with the given IRepository option.
func NewCmsDebitNoteRepository(option *contracts.IRepository) *CmsDebitNoteRepository {
	return &CmsDebitNoteRepository{
		db:    option.Db,
		audit: option.Audit,
		c:     customer.NewCmsCustomerRepository(option),
		d:     NewCmsDebitNoteDetailsRepository(option),
		s:     NewCmsDebitNoteSalesRepository(option),
	}
}

// Get retrieves a debit note with the given debitNoteCode from the CmsDebitNoteRepository.
// If the debit note is found, it returns a pointer to the debit note entity and nil error.
// If the debit note is not found, it returns nil and nil error.
// If an error occurs while retrieving the debit note, it returns nil and the error.
func (r *CmsDebitNoteRepository) Get(debitNoteCode string) (*entities.CmsDebitnote, error) {
	var cmsDebitNote entities.CmsDebitnote
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
func (r *CmsDebitNoteRepository) GetWithCustomer(debitNoteCode string) (*models.DebitNoteWithCustomer, error) {
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
	return &models.DebitNoteWithCustomer{
		C: c,
		I: dn,
	}, nil
}

// GetWithItems returns the debit note with its associated items.
// It retrieves the debit note by the given debitNoteCode and calls the Get method.
// It then retrieves the items related to the debit note by calling the Get method of the CmsDebitNoteDetailsRepository.
// It returns the debit note with its items as a DebitNoteWithItems struct, or an error if any.
// The DebitNoteWithItems struct contains the debit note object (M) and an array of debit note details (D).
func (r *CmsDebitNoteRepository) GetWithItems(debitNoteCode string) (*models.DebitNoteWithItems, error) {
	cn, err := r.Get(debitNoteCode)
	if err != nil {
		return nil, err
	}
	details, err := r.d.Get(debitNoteCode)
	if err != nil {
		return nil, err
	}
	return &models.DebitNoteWithItems{
		M: cn,
		D: details,
	}, nil
}

// GetByCustomer retrieves a list of debit notes for a specific customer.
// It takes the customer code as a parameter and returns a slice of CmsDebitnote entities and an error.
// The method queries the database using the cust_code and cancelled fields, orders the results by dn_date in descending order,
// and limits the number of records to 100. If an error occurs while querying the database, it will be returned along with a nil slice.
// If the query is successful, the method returns the list of records and a nil error.
func (r *CmsDebitNoteRepository) GetByCustomer(custCode string) ([]*entities.CmsDebitnote, error) {
	var records []*entities.CmsDebitnote
	err := r.db.Where("cust_code = ? AND cancelled = ?", custCode, "F").OrderBy("dn_date DESC").Limit(100).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

// GetByDate retrieves CmsDebitnote records within a specific date range.
// It takes two parameters, from and to, representing the start and end dates.
// It returns a slice of pointers to CmsDebitnote entities and an error, if any.
func (r *CmsDebitNoteRepository) GetByDate(from time.Time, to time.Time) ([]*entities.CmsDebitnote, error) {
	var records []*entities.CmsDebitnote
	err := r.db.Where(builder.Between{Col: "DATE(dn_date)", LessVal: from.Format("2006-01-02"), MoreVal: to.Format("2006-01-02")}.And(builder.Eq{"cancelled": "F"})).OrderBy("dn_date DESC").Limit(100).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsDebitNoteRepository) Find(predicate *builder.Builder) ([]*entities.CmsDebitnote, error) {
	var records []*entities.CmsDebitnote
	var t entities.CmsDebitnote
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
func (r *CmsDebitNoteRepository) InsertMany(debitNotes []*entities.CmsDebitnote) error {
	_, err := r.db.Insert(iterator.Map(debitNotes, func(item *entities.CmsDebitnote) *entities.CmsDebitnote {
		return item
	}))
	if err != nil {
		return err
	}

	dt := r.mapToDebitNoteSales(debitNotes)
	if len(dt) > 0 {
		err = r.s.InsertMany(dt)
		if err != nil {
			return err
		}
	}

	r.log("INSERT", debitNotes)

	return nil
}

// Update updates a CmsDebitnote in the repository.
// It updates the record in the database with the matching dn_code.
// If there is an error during the update, it returns the error.
// It also logs the "UPDATE" operation with the updated CmsDebitnote.
func (r *CmsDebitNoteRepository) Update(debitNote *entities.CmsDebitnote) error {
	_, err := r.db.Where("dn_code = ?", debitNote.DnCode).Update(debitNote)
	if err != nil {
		return err
	}

	dt := r.mapToDebitNoteSales([]*entities.CmsDebitnote{debitNote})
	if len(dt) > 0 {
		err = r.s.Update(dt[0])
		if err != nil {
			return err
		}
	}

	r.log("UPDATE", []*entities.CmsDebitnote{debitNote})

	return nil
}

// UpdateMany updates multiple debit notes in the database.
func (r *CmsDebitNoteRepository) UpdateMany(debitNotes []*entities.CmsDebitnote) error {
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

	dt := r.mapToDebitNoteSales(debitNotes)
	if len(dt) > 0 {
		err = r.s.UpdateMany(dt)
		if err != nil {
			return err
		}
	}

	r.log("UPDATE", debitNotes)

	return nil
}

// Delete sets the "Cancelled" attribute of the debitNote to "T" and then calls the Update method to save the changes.
// If the update operation fails, an error is returned.
func (r *CmsDebitNoteRepository) Delete(debitNote *entities.CmsDebitnote) error {
	debitNote.Cancelled = "T"
	return r.Update(debitNote)
}

// DeleteMany sets the `Cancelled` field of each debit note in the given slice to "T".
// Then, it calls the `UpdateMany` method to update the changed debit notes in the database.
// If any error occurs during the update process, it rolls back the session and returns the error.
// Otherwise, it commits the session and returns nil.
func (r *CmsDebitNoteRepository) DeleteMany(debitNotes []*entities.CmsDebitnote) error {
	for _, debitNote := range debitNotes {
		debitNote.Cancelled = "T"
	}
	return r.UpdateMany(debitNotes)
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
//	r.log("UPDATE", []*entities.CmsDebitnote{debitNote})
//	r.log("UPDATE", debitNotes)
//
// Requires:
// - entities.CmsDebitnote declaration
// - entities.AuditLog declaration
// - iterator.Map function
// - IAuditLog interface implementation
func (r *CmsDebitNoteRepository) log(op string, payload []*entities.CmsDebitnote) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsDebitnote) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      item.DnCode,
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}

func (r *CmsDebitNoteRepository) mapToDebitNoteSales(invoices []*entities.CmsDebitnote) []*entities.CmsDebitnoteSales {
	return iterator.Map(iterator.Filter(invoices, func(item *entities.CmsDebitnote) bool {
		if slices.Contains([]string{"SL"}, item.FromDoc) {
			return true
		}
		return false
	}), func(i *entities.CmsDebitnote) *entities.CmsDebitnoteSales {
		return &entities.CmsDebitnoteSales{
			DnCode:            i.DnCode,
			CustCode:          i.CustCode,
			DnDate:            i.DnDate,
			DnAmount:          i.DnAmount,
			OutstandingAmount: i.OutstandingAmount,
			Approved:          i.Approved,
			Approver:          i.Approver,
			ApprovedAt:        i.ApprovedAt,
			SalespersonId:     i.SalespersonId,
			DnUdf:             i.DnUdf,
			Cancelled:         i.Cancelled,
			RefNo:             i.RefNo,
		}
	})
}
