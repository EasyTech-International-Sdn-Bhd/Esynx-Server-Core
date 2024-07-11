package contracts

import (
	"github.com/easytech-international-sdn-bhd/core/entities"
	"github.com/easytech-international-sdn-bhd/core/models"
	"time"
)

type ICmsInvoice interface {
	Get(invoiceCode string) (*entities.CmsInvoice, error)
	GetWithCustomer(invoiceCode string) (*models.InvoiceWithCustomer, error)
	GetWithItems(invoiceCode string) (*models.InvoiceWithItems, error)
	GetByCustomer(custCode string) ([]*entities.CmsInvoice, error)
	GetByDate(from time.Time, to time.Time) ([]*entities.CmsInvoice, error)
	InsertBatch(invoices []*entities.CmsInvoice) error
	Update(invoice *entities.CmsInvoice) error
	UpdateBatch(invoices []*entities.CmsInvoice) error
	Delete(invoice *entities.CmsInvoice) error
	DeleteBatch(invoices []*entities.CmsInvoice) error
}
