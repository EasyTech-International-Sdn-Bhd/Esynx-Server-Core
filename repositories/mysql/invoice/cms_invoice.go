package invoice

import (
	"github.com/easytech-international-sdn-bhd/core/entities"
	"github.com/easytech-international-sdn-bhd/core/models"
	"xorm.io/xorm"
)

type CmsInvoiceRepository struct {
	db *xorm.Engine
}

func NewCmsInvoiceRepository(db *xorm.Engine) *CmsInvoiceRepository {
	return &CmsInvoiceRepository{
		db: db,
	}
}

func (r *CmsInvoiceRepository) Resolve(invoiceCode string) (*entities.CmsInvoice, error) {
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
	invoice, err := r.Resolve(invoiceCode)
	if err != nil {
		return nil, err
	}
	if invoice == nil {
		return nil, nil
	}

}
