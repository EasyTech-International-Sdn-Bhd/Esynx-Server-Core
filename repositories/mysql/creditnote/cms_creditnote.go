package creditnote

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/models"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/mysql/customer"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"time"
	"xorm.io/builder"
	"xorm.io/xorm"
)

type CmsCreditNoteRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
	c     *customer.CmsCustomerRepository
	d     *CmsCreditNoteDetailsRepository
}

func NewCmsCreditNoteRepository(option *contracts.IRepository) *CmsCreditNoteRepository {
	return &CmsCreditNoteRepository{
		db:    option.Db,
		audit: option.Audit,
		c:     customer.NewCmsCustomerRepository(option),
		d:     NewCmsCreditNoteDetailsRepository(option),
	}
}

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

func (r *CmsCreditNoteRepository) GetByCustomer(custCode string) ([]*entities.CmsCreditnote, error) {
	var records []*entities.CmsCreditnote
	err := r.db.Where("cust_code = ? AND cancelled = ?", custCode, "F").OrderBy("cn_date DESC").Limit(100).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsCreditNoteRepository) GetByDate(from time.Time, to time.Time) ([]*entities.CmsCreditnote, error) {
	var records []*entities.CmsCreditnote
	err := r.db.Where(builder.Between{Col: "DATE(cn_date)", LessVal: from.Format("2006-01-02"), MoreVal: to.Format("2006-01-02")}.And(builder.Eq{"cancelled": "F"})).OrderBy("cn_date DESC").Limit(100).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

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

func (r *CmsCreditNoteRepository) Update(creditNote *entities.CmsCreditnote) error {
	_, err := r.db.Where("cn_code = ?", creditNote.CnCode).Update(creditNote)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsCreditnote{creditNote})

	return nil
}

func (r *CmsCreditNoteRepository) UpdateMany(creditNotes []*entities.CmsCreditnote) error {
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

func (r *CmsCreditNoteRepository) Delete(creditNote *entities.CmsCreditnote) error {
	creditNote.Cancelled = "T"
	return r.Update(creditNote)
}

func (r *CmsCreditNoteRepository) DeleteMany(creditNotes []*entities.CmsCreditnote) error {
	for _, cn := range creditNotes {
		cn.Cancelled = "T"
	}
	return r.UpdateMany(creditNotes)
}

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
