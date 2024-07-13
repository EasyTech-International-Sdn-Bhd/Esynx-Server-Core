package debitnote

import (
	"github.com/easytech-international-sdn-bhd/core/contracts"
	"github.com/easytech-international-sdn-bhd/core/entities"
	"github.com/easytech-international-sdn-bhd/core/models"
	"github.com/easytech-international-sdn-bhd/core/repositories/mysql/customer"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"time"
	"xorm.io/builder"
	"xorm.io/xorm"
)

type CmsDebitNoteRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
	c     *customer.CmsCustomerRepository
	d     *CmsDebitNoteDetailsRepository
}

func NewCmsDebitNoteRepository(option *contracts.IRepository) *CmsDebitNoteRepository {
	return &CmsDebitNoteRepository{
		db:    option.Db,
		audit: option.Audit,
		c:     customer.NewCmsCustomerRepository(option),
		d:     NewCmsDebitNoteDetailsRepository(option),
	}
}

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

func (r *CmsDebitNoteRepository) GetByCustomer(custCode string) ([]*entities.CmsDebitnote, error) {
	var records []*entities.CmsDebitnote
	err := r.db.Where("cust_code = ? AND cancelled = ?", custCode, "F").OrderBy("dn_date DESC").Limit(100).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsDebitNoteRepository) GetByDate(from time.Time, to time.Time) ([]*entities.CmsDebitnote, error) {
	var records []*entities.CmsDebitnote
	err := r.db.Where(builder.Between{Col: "DATE(dn_date)", LessVal: from.Format("2006-01-02"), MoreVal: to.Format("2006-01-02")}.And(builder.Eq{"cancelled": "F"})).OrderBy("dn_date DESC").Limit(100).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsDebitNoteRepository) InsertMany(debitNotes []*entities.CmsDebitnote) error {
	_, err := r.db.Insert(iterator.Map(debitNotes, func(item *entities.CmsDebitnote) *entities.CmsDebitnote {
		item.Validate()
		item.ToUpdate()
		return item
	}))
	if err != nil {
		return err
	}

	go r.log("INSERT", debitNotes)

	return nil
}

func (r *CmsDebitNoteRepository) Update(debitNote *entities.CmsDebitnote) error {
	_, err := r.db.Where("dn_code = ?", debitNote.DnCode).Update(debitNote)
	if err != nil {
		return err
	}

	go r.log("UPDATE", []*entities.CmsDebitnote{debitNote})

	return nil
}

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
		dn.Validate()
		dn.ToUpdate()
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

	go r.log("UPDATE", debitNotes)

	return nil
}

func (r *CmsDebitNoteRepository) Delete(debitNote *entities.CmsDebitnote) error {
	debitNote.Cancelled = "T"
	debitNote.ToUpdate()
	return r.Update(debitNote)
}

func (r *CmsDebitNoteRepository) DeleteMany(debitNotes []*entities.CmsDebitnote) error {
	for _, debitNote := range debitNotes {
		debitNote.Cancelled = "T"
		debitNote.ToUpdate()
	}
	return r.UpdateMany(debitNotes)
}

func (r *CmsDebitNoteRepository) log(op string, payload []*entities.CmsDebitnote) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsDebitnote) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordID:      item.DnCode,
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
