package core

import (
	"github.com/easytech-international-sdn-bhd/core/contracts"
	"github.com/easytech-international-sdn-bhd/core/options"
	"github.com/easytech-international-sdn-bhd/core/repositories/mysql"
	"github.com/easytech-international-sdn-bhd/core/repositories/mysql/agent"
	"github.com/easytech-international-sdn-bhd/core/repositories/mysql/audit"
	"github.com/easytech-international-sdn-bhd/core/repositories/mysql/creditnote"
	"github.com/easytech-international-sdn-bhd/core/repositories/mysql/customer"
	"github.com/easytech-international-sdn-bhd/core/repositories/mysql/debitnote"
	"github.com/easytech-international-sdn-bhd/core/repositories/mysql/invoice"
	"github.com/easytech-international-sdn-bhd/core/repositories/mysql/module"
	"github.com/easytech-international-sdn-bhd/core/repositories/mysql/stock"
)

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

func NewEsynxProvider(session contracts.IDatabaseUserSession) (*ESynx, error) {
	if session.GetStore() == options.MySQL {
		db := mysql.NewMySqlDb()
		err := db.Open(session.GetConnection())
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

func (e *ESynx) Destroy() error {
	return e.engine.Close()
}
