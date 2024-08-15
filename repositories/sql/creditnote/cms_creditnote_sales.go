package creditnote

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

// CmsCreditNoteSalesRepository represents a repository for managing credit notes in a CMS system.
// It is responsible for interacting with the database (db) to perform CRUD operations on credit notes,
// using the provided db engine (xorm.Engine).
// It also utilizes the IAuditLog interface for logging audit data and the CmsCustomerRepository and
// CmsCreditNoteDetailsRepository for accessing related customer and credit note details data respectively.
type CmsCreditNoteSalesRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
	c     *customer.CmsCustomerRepository
	d     *CmsCreditNoteDetailsRepository
}

// NewCmsCreditNoteSalesRepository creates a new instance of CmsCreditNoteSalesRepository.
// It takes an option parameter of type *contracts.IRepository and returns a pointer to CmsCreditNoteSalesRepository.
func NewCmsCreditNoteSalesRepository(option *contracts.IRepository) *CmsCreditNoteSalesRepository {
	return &CmsCreditNoteSalesRepository{
		db:    option.Db,
		audit: option.Audit,
		c:     customer.NewCmsCustomerRepository(option),
		d:     NewCmsCreditNoteDetailsRepository(option),
	}
}

// Get retrieves a CmsCreditnote entity from the database based on the creditNoteCode provided.
// If the entity is found, it is returned with a nil error. If the entity is not found, nil is returned.
// If an error occurs during the retrieval process, nil is returned with the error.
func (r *CmsCreditNoteSalesRepository) Get(creditNoteCode string) (*entities.CmsCreditnoteSales, error) {
	var cmsCreditNote entities.CmsCreditnoteSales
	has, err := r.db.Where("cn_code=?", creditNoteCode).Get(&cmsCreditNote)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return &cmsCreditNote, nil
}

// GetWithCustomer retrieves a credit note along with its associated customer information.
// It first calls the Get method to fetch the credit note details using the provided creditNoteCode.
// If the credit note is not found, it returns nil, nil. If there is an error during retrieval, it returns nil, err.
// Otherwise, it calls the Get method of the CmsCustomerRepository to fetch the customer details using CustCode
// from the retrieved credit note. If there is an error during retrieval, it returns nil, err.
// Finally, it returns the CreditNoteWithCustomer struct containing the customer and credit note information.
// Otherwise, it returns a nil pointer to CreditNoteWithCustomer and the error.
// Signature: (creditNoteCode string) -> (*models.CreditNoteWithCustomer, error)
func (r *CmsCreditNoteSalesRepository) GetWithCustomer(creditNoteCode string) (*models.CreditNoteSalesWithCustomer, error) {
	cn, err := r.Get(creditNoteCode)
	if err != nil {
		return nil, err
	}
	if cn == nil {
		return nil, nil
	}
	c, err := r.c.Get(cn.CustCode)
	if err != nil {
		return nil, err
	}
	return &models.CreditNoteSalesWithCustomer{
		C: c,
		I: cn,
	}, nil
}

// GetWithItems retrieves a credit note with its associated items by the creditNoteCode.
//
// If the credit note is not found, it returns nil as the result with a non-nil error.
// If there is an error when retrieving the credit note or its associated items, it returns nil as the result
// with the non-nil error.
//
// The returned *models.CreditNoteWithItems contains the credit note as its 'M' field and the associated items
// as its 'D' field.
func (r *CmsCreditNoteSalesRepository) GetWithItems(creditNoteCode string) (*models.CreditNoteSalesWithItems, error) {
	cn, err := r.Get(creditNoteCode)
	if err != nil {
		return nil, err
	}
	details, err := r.d.Get(creditNoteCode)
	if err != nil {
		return nil, err
	}
	return &models.CreditNoteSalesWithItems{
		M: cn,
		D: details,
	}, nil
}

// GetByCustomer retrieves credit notes for a specific customer identified by custCode.
// It returns a slice of CmsCreditnote entities and an error if any occurred.
func (r *CmsCreditNoteSalesRepository) GetByCustomer(custCode string) ([]*entities.CmsCreditnoteSales, error) {
	var records []*entities.CmsCreditnoteSales
	err := r.db.Where("cust_code = ? AND cancelled = ?", custCode, "F").OrderBy("cn_date DESC").Limit(100).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

// GetByDate retrieves credit notes within a specified date range.
func (r *CmsCreditNoteSalesRepository) GetByDate(from time.Time, to time.Time) ([]*entities.CmsCreditnoteSales, error) {
	var records []*entities.CmsCreditnoteSales
	err := r.db.Where(builder.Between{Col: "DATE(cn_date)", LessVal: from.Format("2006-01-02"), MoreVal: to.Format("2006-01-02")}.And(builder.Eq{"cancelled": "F"})).OrderBy("cn_date DESC").Limit(100).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsCreditNoteSalesRepository) Find(predicate *builder.Builder) ([]*entities.CmsCreditnoteSales, error) {
	var records []*entities.CmsCreditnoteSales
	var t entities.CmsCreditnoteSales
	err := r.db.SQL(predicate.From(t.TableName())).Find(&records)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, nil
	}
	return records, nil
}

// InsertMany inserts multiple credit notes into the database.
// It receives a slice of credit notes and inserts each item using the db.Insert function.
// If any error occurs during the insertion, it returns the error.
// After the insertion, it logs the "INSERT" operation with the inserted credit notes.
// It returns nil if everything is successful.
func (r *CmsCreditNoteSalesRepository) InsertMany(creditNotes []*entities.CmsCreditnoteSales) error {
	_, err := r.db.Insert(iterator.Map(creditNotes, func(item *entities.CmsCreditnoteSales) *entities.CmsCreditnoteSales {
		return item
	}))
	if err != nil {
		return err
	}

	r.log("INSERT", creditNotes)

	return nil
}

// Update updates a credit note in the database.
//
// Parameters:
// - creditNote: The credit note to be updated.
//
// Returns:
// - error: An error if the update operation fails.
func (r *CmsCreditNoteSalesRepository) Update(creditNote *entities.CmsCreditnoteSales) error {
	_, err := r.db.Where("cn_code = ?", creditNote.CnCode).Update(creditNote)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsCreditnoteSales{creditNote})

	return nil
}

// UpdateMany updates multiple credit notes in the database.
// It takes a slice of credit notes as the input and returns an error if any.
func (r *CmsCreditNoteSalesRepository) UpdateMany(creditNotes []*entities.CmsCreditnoteSales) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}
	var sessionErr error
	rollback := false
	for _, cn := range creditNotes {
		_, err = session.Where("cn_code = ?", cn.CnCode).Update(cn)
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

	r.log("UPDATE", creditNotes)

	return nil
}

// Delete sets the "Cancelled" attribute of the given creditNote to "T"
// and updates it using the Update method. It returns an error if the update operation fails.
func (r *CmsCreditNoteSalesRepository) Delete(creditNote *entities.CmsCreditnoteSales) error {
	creditNote.Cancelled = "T"
	_, err := r.db.Where("cn_code = ?", creditNote.CnCode).Cols("cancelled").Update(creditNote)
	if err == nil {
		r.log("DELETE", []*entities.CmsCreditnoteSales{creditNote})
	}
	return err
}

// DeleteMany sets the "Cancelled" attribute of each credit note in the input slice to "T"
// and updates them using a session. It returns an error if the update operation fails.
func (r *CmsCreditNoteSalesRepository) DeleteMany(creditNotes []*entities.CmsCreditnoteSales) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}

	var sessionErr error
	rollback := false
	for _, cn := range creditNotes {
		cn.Cancelled = "T"
		_, err = session.Where("cn_code = ?", cn.CnCode).Cols("cancelled").Update(cn)
		if err != nil {
			rollback = true
			sessionErr = err
			break
		}
	}
	if rollback {
		err = session.Rollback()
		if err != nil {
			return err
		}
		return sessionErr
	}

	err = session.Commit()
	if err != nil {
		return err
	}

	r.log("DELETE", creditNotes)

	return nil
}

// log logs the operation and payload to the audit log.
func (r *CmsCreditNoteSalesRepository) log(op string, payload []*entities.CmsCreditnoteSales) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsCreditnoteSales) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      item.CnCode,
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
