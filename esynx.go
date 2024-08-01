package core

import (
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/options"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/mysql"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/mysql/agent"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/mysql/audit"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/mysql/creditnote"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/mysql/customer"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/mysql/debitnote"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/mysql/invoice"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/mysql/module"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/mysql/stock"
)

// ESynx represents a struct with various contracts/interfaces for database operations.
// The contracts are related to credit notes, customers, branches, salespersons, debit notes,
// invoices, products, warehouse stocks, etc.
type ESynx struct {
	engine                contracts.IDatabase
	CreditNote            contracts.ICmsCreditNote
	CreditNoteDetails     contracts.ICmsCreditNoteDetails
	Customer              contracts.ICmsCustomer
	CustomerBranch        contracts.ICmsCustomerBranch
	CustomerSalesperson   contracts.ICmsCustomerSalesperson
	DebitNote             contracts.ICmsDebitNote
	DebitNoteDetails      contracts.ICmsDebitNoteDetails
	Invoice               contracts.ICmsInvoice
	InvoiceDetails        contracts.ICmsInvoiceDetails
	InvoiceSales          contracts.ICmsInvoiceSales
	Agent                 contracts.ICmsLogin
	MobileAppModule       contracts.ICmsMobileModule
	Product               contracts.ICmsProduct
	ProductAttachment     contracts.ICmsProductAtch
	ProductBatch          contracts.ICmsProductBatch
	ProductImage          contracts.ICmsProductImage
	ProductPriceTag       contracts.ICmsProductPriceTag
	ProductStandardPrice  contracts.ICmsProductUomPrice
	ProductWarehouseStock contracts.ICmsWarehouseStock
}

// NewEsynxProvider creates a new instance of ESynx with the given database user session.
// It checks the database store type from the session and initializes the appropriate database engine.
// If the store type is MySQL, it creates a new instance of MySqlDb and opens a connection to the MySQL database.
// Then it creates an IRepository instance with the database engine, user, app name, and audit log repository.
// Finally, it returns the initialized ESynx instance with all the repository instances.
// If the store type is not MySQL, it returns nil and no error.
// If there's an error opening the MySQL connection, it returns nil and the error.
func NewEsynxProvider(session contracts.IDatabaseUserSession) (*ESynx, error) {
	if session.GetStore() == options.MySQL {
		db := mysql.NewMySqlDb()
		err := db.Open(session.GetConnection(), session.GetLogger())
		if err != nil {
			return nil, err
		}
		userOptions := contracts.IRepository{
			Db:      db.Engine,
			User:    session.GetUser(),
			AppName: session.GetApp(),
			Audit:   audit.NewAuditLogRepository(db.Engine, session),
		}
		return &ESynx{
			engine:                db,
			CreditNote:            creditnote.NewCmsCreditNoteRepository(&userOptions),
			CreditNoteDetails:     creditnote.NewCmsCreditNoteDetailsRepository(&userOptions),
			Customer:              customer.NewCmsCustomerRepository(&userOptions),
			CustomerBranch:        customer.NewCmsCustomerBranchRepository(&userOptions),
			CustomerSalesperson:   agent.NewCmsCustomerSalespersonRepository(&userOptions),
			DebitNote:             debitnote.NewCmsDebitNoteRepository(&userOptions),
			DebitNoteDetails:      debitnote.NewCmsDebitNoteDetailsRepository(&userOptions),
			Invoice:               invoice.NewCmsInvoiceRepository(&userOptions),
			InvoiceDetails:        invoice.NewCmsInvoiceDetailsRepository(&userOptions),
			InvoiceSales:          invoice.NewCmsInvoiceSalesRepository(&userOptions),
			Agent:                 agent.NewCmsLoginRepository(&userOptions),
			MobileAppModule:       module.NewCmsMobileModuleRepository(&userOptions),
			Product:               stock.NewCmsProductRepository(&userOptions),
			ProductAttachment:     stock.NewCmsProductAtchRepository(&userOptions),
			ProductBatch:          stock.NewCmsProductBatchRepository(&userOptions),
			ProductImage:          stock.NewCmsProductImageRepository(&userOptions),
			ProductPriceTag:       stock.NewCmsProductPriceTagRepository(&userOptions),
			ProductStandardPrice:  stock.NewCmsProductUomPriceRepository(&userOptions),
			ProductWarehouseStock: stock.NewCmsWarehouseStockRepository(&userOptions),
		}, nil
	}
	if session.GetStore() == options.Firestore {
		// TODO: implement firestore or others for realtime data. **Not Now**
	}
	return nil, nil
}

// Destroy closes the database connection used by the ESynx instance.
// It returns an error if there is an issue closing the connection.
func (e *ESynx) Destroy() error {
	return e.engine.Close()
}
