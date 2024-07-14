package contracts

import (
	"github.com/easytech-international-sdn-bhd/esynx-server-core/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/models"
)

type ICmsInvoiceDetails interface {
	Get(invoiceCode string) ([]*entities.CmsInvoiceDetails, error)
	GetMany(invoiceCodes []string) ([]*entities.CmsInvoiceDetails, error)
	GetWithProduct(invoiceCode string) ([]*models.InvoiceDetailsWithProduct, error)
	InsertMany(details []*entities.CmsInvoiceDetails) error
	Update(details *entities.CmsInvoiceDetails) error
	UpdateMany(details []*entities.CmsInvoiceDetails) error
	Delete(details *entities.CmsInvoiceDetails) error
	DeleteMany(details []*entities.CmsInvoiceDetails) error
}
