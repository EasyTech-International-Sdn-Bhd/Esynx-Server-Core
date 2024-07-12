package invoice

import (
	"github.com/easytech-international-sdn-bhd/core/entities"
	"github.com/easytech-international-sdn-bhd/core/models"
	"github.com/easytech-international-sdn-bhd/core/repositories/mysql/customer"
	iterator "github.com/ledongthuc/goterators"
	"time"
	"xorm.io/builder"
	"xorm.io/xorm"
)

type CmsInvoiceSalesRepository struct {
	db *xorm.Engine
	c  *customer.CmsCustomerRepository
	d  *CmsInvoiceDetailsRepository
}

func NewCmsInvoiceSalesRepository(db *xorm.Engine) *CmsInvoiceSalesRepository {
	return &CmsInvoiceSalesRepository{
		db: db,
		c:  customer.NewCmsCustomerRepository(db),
		d:  NewCmsInvoiceDetailsRepository(db),
	}
}

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

func (r *CmsInvoiceSalesRepository) GetByCustomer(custCode string) ([]*entities.CmsInvoiceSales, error) {
	var records []*entities.CmsInvoiceSales
	err := r.db.Where("cust_code = ? AND cancelled = ?", custCode, "F").OrderBy("invoice_date DESC").Limit(100).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsInvoiceSalesRepository) GetByDate(from time.Time, to time.Time) ([]*entities.CmsInvoiceSales, error) {
	var records []*entities.CmsInvoiceSales
	err := r.db.Where(builder.Between{Col: "DATE(invoice_date)", LessVal: from.Format("2006-01-02"), MoreVal: to.Format("2006-01-02")}.And(builder.Eq{"cancelled": "F"})).OrderBy("invoice_date DESC").Limit(100).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsInvoiceSalesRepository) InsertMany(invoices []*entities.CmsInvoiceSales) error {
	_, err := r.db.Insert(iterator.Map(invoices, func(item *entities.CmsInvoiceSales) *entities.CmsInvoiceSales {
		item.Validate()
		item.ToUpdate()
		return item
	}))
	if err != nil {
		return err
	}
	return nil
}

func (r *CmsInvoiceSalesRepository) Update(invoice *entities.CmsInvoiceSales) error {
	_, err := r.db.Where("invoice_code = ?", invoice.InvoiceCode).Update(invoice)
	if err != nil {
		return err
	}
	return nil
}

func (r *CmsInvoiceSalesRepository) UpdateMany(invoices []*entities.CmsInvoiceSales) error {
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
	return nil
}

func (r *CmsInvoiceSalesRepository) Delete(invoice *entities.CmsInvoiceSales) error {
	invoice.Cancelled = "T"
	invoice.ToUpdate()
	return r.Update(invoice)
}

func (r *CmsInvoiceSalesRepository) DeleteMany(invoices []*entities.CmsInvoiceSales) error {
	for _, invoice := range invoices {
		invoice.Cancelled = "T"
		invoice.ToUpdate()
	}
	return r.UpdateMany(invoices)
}
