package invoice

import (
	"github.com/easytech-international-sdn-bhd/core/entities"
	"github.com/easytech-international-sdn-bhd/core/models"
	"github.com/easytech-international-sdn-bhd/core/repositories/mysql/customer"
	"xorm.io/xorm"
)

type CmsInvoiceRepository struct {
	db *xorm.Engine
	c  *customer.CmsCustomerRepository
}

func NewCmsInvoiceRepository(db *xorm.Engine) *CmsInvoiceRepository {
	return &CmsInvoiceRepository{
		db: db,
		c:  customer.NewCmsCustomerRepository(db),
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
