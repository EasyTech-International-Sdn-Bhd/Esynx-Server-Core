package invoice

import (
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/models"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/mysql/customer"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"time"
	"xorm.io/builder"
	"xorm.io/xorm"
)

type CmsInvoiceRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
	c     *customer.CmsCustomerRepository
	d     *CmsInvoiceDetailsRepository
}

func NewCmsInvoiceRepository(option *contracts.IRepository) *CmsInvoiceRepository {
	return &CmsInvoiceRepository{
		db:    option.Db,
		audit: option.Audit,
		c:     customer.NewCmsCustomerRepository(option),
		d:     NewCmsInvoiceDetailsRepository(option),
	}
}

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

func (r *CmsInvoiceRepository) GetByCustomer(custCode string) ([]*entities.CmsInvoice, error) {
	var records []*entities.CmsInvoice
	err := r.db.Where("cust_code = ? AND cancelled = ?", custCode, "F").OrderBy("invoice_date DESC").Limit(100).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsInvoiceRepository) GetByDate(from time.Time, to time.Time) ([]*entities.CmsInvoice, error) {
	var records []*entities.CmsInvoice
	err := r.db.Where(builder.Between{Col: "DATE(invoice_date)", LessVal: from.Format("2006-01-02"), MoreVal: to.Format("2006-01-02")}.And(builder.Eq{"cancelled": "F"})).OrderBy("invoice_date DESC").Limit(100).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsInvoiceRepository) InsertMany(invoices []*entities.CmsInvoice) error {
	_, err := r.db.Insert(iterator.Map(invoices, func(item *entities.CmsInvoice) *entities.CmsInvoice {
		item.Validate()
		item.ToUpdate()
		return item
	}))
	if err != nil {
		return err
	}

	r.log("INSERT", invoices)

	return nil
}

func (r *CmsInvoiceRepository) Update(invoice *entities.CmsInvoice) error {
	_, err := r.db.Where("invoice_code = ?", invoice.InvoiceCode).Update(invoice)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsInvoice{invoice})

	return nil
}

func (r *CmsInvoiceRepository) UpdateMany(invoices []*entities.CmsInvoice) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}
	var sessionErr error
	rollback := false
	for _, inv := range invoices {
		inv.Validate()
		inv.ToUpdate()
		_, err = session.Where("invoice_code = ?", inv.InvoiceCode).Update(inv)
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

	r.log("UPDATE", invoices)

	return nil
}

func (r *CmsInvoiceRepository) Delete(invoice *entities.CmsInvoice) error {
	invoice.Cancelled = "T"
	invoice.ToUpdate()
	return r.Update(invoice)
}

func (r *CmsInvoiceRepository) DeleteMany(invoices []*entities.CmsInvoice) error {
	for _, invoice := range invoices {
		invoice.Cancelled = "T"
		invoice.ToUpdate()
	}
	return r.UpdateMany(invoices)
}

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
