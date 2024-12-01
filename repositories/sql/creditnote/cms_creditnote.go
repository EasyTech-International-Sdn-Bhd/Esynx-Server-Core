package creditnote

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

// CmsCreditNoteRepository represents a repository for managing credit notes in a CMS system.
// It is responsible for interacting with the database (db) to perform CRUD operations on credit notes,
// using the provided db engine (xorm.Engine).
// It also utilizes the IAuditLog interface for logging audit data and the CmsCustomerRepository and
// CmsCreditNoteDetailsRepository for accessing related customer and credit note details data respectively.
type CmsCreditNoteRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
	c     *customer.CmsCustomerRepository
	d     *CmsCreditNoteDetailsRepository
}

// NewCmsCreditNoteRepository creates a new instance of CmsCreditNoteRepository.
// It takes an option parameter of type *contracts.IRepository and returns a pointer to CmsCreditNoteRepository.
func NewCmsCreditNoteRepository(option *contracts.IRepository) *CmsCreditNoteRepository {
	return &CmsCreditNoteRepository{
		db:    option.Db,
		audit: option.Audit,
		c:     customer.NewCmsCustomerRepository(option),
		d:     NewCmsCreditNoteDetailsRepository(option),
	}
}

// Get retrieves a CmsCreditnote entity from the database based on the creditNoteCode provided.
// If the entity is found, it is returned with a nil error. If the entity is not found, nil is returned.
// If an error occurs during the retrieval process, nil is returned with the error.
func (r *CmsCreditNoteRepository) Get(creditNoteCode string) (*entities.CmsCreditnote, error) {
	var cmsCreditNote entities.CmsCreditnote
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
func (r *CmsCreditNoteRepository) GetWithCustomer(creditNoteCode string) (*models.CreditNoteWithCustomer, error) {
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
	return &models.CreditNoteWithCustomer{
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
func (r *CmsCreditNoteRepository) GetWithItems(creditNoteCode string) (*models.CreditNoteWithItems, error) {
	cn, err := r.Get(creditNoteCode)
	if err != nil {
		return nil, err
	}
	details, err := r.d.Get(creditNoteCode)
	if err != nil {
		return nil, err
	}
	return &models.CreditNoteWithItems{
		M: cn,
		D: details,
	}, nil
}

// GetByCustomer retrieves credit notes for a specific customer identified by custCode.
// It returns a slice of CmsCreditnote entities and an error if any occurred.
func (r *CmsCreditNoteRepository) GetByCustomer(custCode string) ([]*entities.CmsCreditnote, error) {
	var records []*entities.CmsCreditnote
	err := r.db.Where("cust_code = ? AND cancelled = ?", custCode, "F").OrderBy("cn_date DESC").Limit(100).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

// GetByDate retrieves credit notes within a specified date range.
func (r *CmsCreditNoteRepository) GetByDate(from time.Time, to time.Time) ([]*entities.CmsCreditnote, error) {
	var records []*entities.CmsCreditnote
	err := r.db.Where(builder.Between{Col: "DATE(cn_date)", LessVal: from.Format("2006-01-02"), MoreVal: to.Format("2006-01-02")}.And(builder.Eq{"cancelled": "F"})).OrderBy("cn_date DESC").Limit(100).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsCreditNoteRepository) Find(predicate *builder.Builder) ([]*entities.CmsCreditnote, error) {
	var records []*entities.CmsCreditnote
	var t entities.CmsCreditnote
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
func (r *CmsCreditNoteRepository) InsertMany(creditNotes []*entities.CmsCreditnote) error {
	_, err := r.db.Insert(iterator.Map(creditNotes, func(item *entities.CmsCreditnote) *entities.CmsCreditnote {
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
func (r *CmsCreditNoteRepository) Update(creditNote *entities.CmsCreditnote) error {
	_, err := r.db.Where("cn_code = ?", creditNote.CnCode).Omit("cn_code").Update(creditNote)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsCreditnote{creditNote})

	return nil
}

// Delete sets the "Cancelled" attribute of the given creditNote to "T"
// and updates it directly using r.db. It returns an error if the update operation fails.
func (r *CmsCreditNoteRepository) Delete(creditNote *entities.CmsCreditnote) error {
	creditNote.Cancelled = "T"
	_, err := r.db.Where("cn_code = ?", creditNote.CnCode).Cols("cancelled", "ref_no").Update(&entities.CmsCreditnote{
		Cancelled: "T",
		RefNo:     fmt.Sprintf("DELETED-%s", time.Now().Format("20060102")),
	})
	if err == nil {
		r.log("DELETE", []*entities.CmsCreditnote{creditNote})
	}
	return err
}

// UpdateMany updates multiple credit notes in the database.
// It takes a slice of credit notes as the input and returns an error if any.
func (r *CmsCreditNoteRepository) UpdateMany(creditNotes []*entities.CmsCreditnote) error {
	for _, cn := range creditNotes {
		_, err := r.db.Where("cn_code = ?", cn.CnCode).Omit("cn_code").Update(cn)
		if err != nil {
			return err
		}
	}

	r.log("UPDATE", creditNotes)

	return nil
}

// DeleteMany sets the "Cancelled" attribute of multiple credit notes to "T"
// and updates them directly using a session. It takes a slice of credit notes
// as input and returns an error if any operation fails.
func (r *CmsCreditNoteRepository) DeleteMany(creditNotes []*entities.CmsCreditnote) error {
	ids := iterator.Map(creditNotes, func(item *entities.CmsCreditnote) string {
		return item.CnCode
	})

	_, err := r.db.In("cn_code", ids).Cols("cancelled", "ref_no").Update(&entities.CmsCreditnote{
		Cancelled: "T",
		RefNo:     fmt.Sprintf("DELETED-%s", time.Now().Format("20060102")),
	})
	if err != nil {
		return err
	}

	r.log("DELETE", creditNotes)

	return nil
}

func (r *CmsCreditNoteRepository) DeleteByAny(predicate *builder.Builder) ([]*entities.CmsCreditnote, error) {
	var t entities.CmsCreditnote

	var records []*entities.CmsCreditnote
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

// log logs the operation and payload to the audit log.
func (r *CmsCreditNoteRepository) log(op string, payload []*entities.CmsCreditnote) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsCreditnote) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      item.CnCode,
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
