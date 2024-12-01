package contracts

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/models"
	"time"
	"xorm.io/builder"
)

type ICmsInvoice interface {
	Get(invoiceCode string) (*entities.CmsInvoice, error)
	GetWithCustomer(invoiceCode string) (*models.InvoiceWithCustomer, error)
	GetWithItems(invoiceCode string) (*models.InvoiceWithItems, error)
	GetByCustomer(custCode string) ([]*entities.CmsInvoice, error)
	GetByDate(from time.Time, to time.Time) ([]*entities.CmsInvoice, error)
	InsertMany(invoices []*entities.CmsInvoice) error
	Update(invoice *entities.CmsInvoice) error
	UpdateMany(invoices []*entities.CmsInvoice) error
	Delete(invoice *entities.CmsInvoice) error
	DeleteMany(invoices []*entities.CmsInvoice) error
	DeleteByAny(predicate *builder.Builder) error
	Find(predicate *builder.Builder) ([]*entities.CmsInvoice, error)
}
