package contracts

import (
	"github.com/easytech-international-sdn-bhd/core/entities"
	"github.com/easytech-international-sdn-bhd/core/models"
)

type ICmsInvoiceDetails interface {
	Get(invoiceCode string) ([]entities.CmsInvoiceDetails, error)
	GetMany(invoiceCodes []string) ([]entities.CmsInvoiceDetails, error)
	GetByProductCode(productCode string) ([]entities.CmsInvoiceDetails, error)
	GetWithInvoice(invoiceCode string) (models.InvoiceWithItems, error)
	GetWithProduct(invoiceCode string) ([]models.InvoiceDetailsWithProduct, error)
	InsertBatch(details []entities.CmsInvoiceDetails) error
	Update(details entities.CmsInvoiceDetails) error
	UpdateBatch(details []entities.CmsInvoiceDetails) error
	Delete(details entities.CmsInvoiceDetails) error
	DeleteBatch(details []entities.CmsInvoiceDetails) error
}
