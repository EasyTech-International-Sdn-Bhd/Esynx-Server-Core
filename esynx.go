package core

import (
	"github.com/easytech-international-sdn-bhd/core/contracts"
	"github.com/easytech-international-sdn-bhd/core/repositories/mysql"
	"github.com/easytech-international-sdn-bhd/core/repositories/mysql/agent"
	"github.com/easytech-international-sdn-bhd/core/repositories/mysql/creditnote"
	"github.com/easytech-international-sdn-bhd/core/repositories/mysql/customer"
	"github.com/easytech-international-sdn-bhd/core/repositories/mysql/debitnote"
	"github.com/easytech-international-sdn-bhd/core/repositories/mysql/invoice"
	"github.com/easytech-international-sdn-bhd/core/repositories/mysql/module"
	"github.com/easytech-international-sdn-bhd/core/repositories/mysql/stock"
)

type DatabaseProvider int

const (
	MySQL DatabaseProvider = iota
	Firestore
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

func NewEsynxProvider(provider DatabaseProvider, conn string) (*ESynx, error) {
	if provider == MySQL {
		db := mysql.NewMySqlDb()
		err := db.Open(conn)
		if err != nil {
			return nil, err
		}
		return &ESynx{
			engine:                db,
			CreditNote:            creditnote.NewCmsCreditNoteRepository(db.Engine),
			CreditNoteDetails:     creditnote.NewCmsCreditNoteDetailsRepository(db.Engine),
			Customer:              customer.NewCmsCustomerRepository(db.Engine),
			CustomerBranch:        customer.NewCmsCustomerBranchRepository(db.Engine),
			CustomerSalesperson:   agent.NewCmsCustomerSalespersonRepository(db.Engine),
			DebitNote:             debitnote.NewCmsDebitNoteRepository(db.Engine),
			DebitNoteDetails:      debitnote.NewCmsDebitNoteDetailsRepository(db.Engine),
			Invoice:               invoice.NewCmsInvoiceRepository(db.Engine),
			InvoiceDetails:        invoice.NewCmsInvoiceDetailsRepository(db.Engine),
			InvoiceSales:          invoice.NewCmsInvoiceSalesRepository(db.Engine),
			Agent:                 agent.NewCmsLoginRepository(db.Engine),
			MobileAppModule:       module.NewCmsMobileModuleRepository(db.Engine),
			Product:               stock.NewCmsProductRepository(db.Engine),
			ProductAttachment:     stock.NewCmsProductAtchRepository(db.Engine),
			ProductBatch:          stock.NewCmsProductBatchRepository(db.Engine),
			ProductImage:          stock.NewCmsProductImageRepository(db.Engine),
			ProductPriceTag:       stock.NewCmsProductPriceTagRepository(db.Engine),
			ProductStandardPrice:  stock.NewCmsProductUomPriceRepository(db.Engine),
			ProductWarehouseStock: stock.NewCmsWarehouseStockRepository(db.Engine),
		}, nil
	}
	if provider == Firestore {
		// TODO: implement firestore or others for realtime data. **Not Now**
	}
	return nil, nil
}

func (e *ESynx) Destroy() error {
	return e.engine.Close()
}
