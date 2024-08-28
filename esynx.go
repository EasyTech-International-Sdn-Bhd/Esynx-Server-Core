package core

import (
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/mock"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/options"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/sql"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/sql/agent"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/sql/audit"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/sql/creditnote"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/sql/customer"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/sql/debitnote"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/sql/invoice"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/sql/module"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/sql/stock"
)

// ESynx represents a struct with various contracts/interfaces for database operations.
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
func NewEsynxProvider(session contracts.IDatabaseUserSession) (*ESynx, error) {
	var db contracts.IDatabase
	var err error

	switch session.GetStore() {
	case options.SqlDb:
		db = sql.NewSqlDb()
		err = db.Open(session.GetConnection(), session.GetLogger())
	case options.Mock:
		db = mock.NewMockDb()
		err = db.Open(session.GetConnection(), session.GetLogger())
	case options.Firestore:
		// TODO: implement Firestore or other databases for real-time data. **Not Now**
		return nil, nil
	default:
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	userOptions := contracts.IRepository{
		Db:      db.GetEngine(),
		User:    session.GetUser(),
		AppName: session.GetApp(),
		Audit:   audit.NewAuditLogRepository(db.GetEngine(), session),
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

func (e *ESynx) DefineSchema() error {
	return e.engine.DefineSchema()
}

// Destroy closes the database connection used by the ESynx instance.
func (e *ESynx) Destroy() error {
	return e.engine.Close()
}
