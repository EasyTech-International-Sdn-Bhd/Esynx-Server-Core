package contracts

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/models"
	"time"
)

type ICmsInvoiceSales interface {
	Get(invoiceCode string) (*entities.CmsInvoiceSales, error)
	GetWithCustomer(invoiceCode string) (*models.InvoiceSalesWithCustomer, error)
	GetWithItems(invoiceCode string) (*models.InvoiceSalesWithItems, error)
	GetByCustomer(custCode string) ([]*entities.CmsInvoiceSales, error)
	GetByDate(from time.Time, to time.Time) ([]*entities.CmsInvoiceSales, error)
	InsertMany(invoices []*entities.CmsInvoiceSales) error
	Update(invoice *entities.CmsInvoiceSales) error
	UpdateMany(invoices []*entities.CmsInvoiceSales) error
	Delete(invoice *entities.CmsInvoiceSales) error
	DeleteMany(invoices []*entities.CmsInvoiceSales) error
}
