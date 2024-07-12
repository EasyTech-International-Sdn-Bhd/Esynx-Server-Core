package sql

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/core/entities"
	_ "github.com/go-sql-driver/mysql"
	"src.techknowlogick.com/xormigrate"
	"xorm.io/xorm"
)

func DefineSchema(db *xorm.Engine) error {
	m := xormigrate.New(db, migrations())
	return m.Migrate()
}

func migrations() []*xormigrate.Migration {
	var schemas []*xormigrate.Migration
	for i, schema := range defaults() {
		schemas = append(schemas, &xormigrate.Migration{
			ID: fmt.Sprintf("define_schema_%d", i),
			Migrate: func(tx *xorm.Engine) error {
				return tx.Sync(schema)
			},
			Rollback: func(db *xorm.Engine) error {
				return nil
			},
		})
	}
	return schemas
}

func defaults() []interface{} {
	return []interface{}{
		&entities.CmsAccExistingOrder{},
		&entities.CmsAccExistingOrderItem{},
		&entities.CmsAppAdvertisement{},
		&entities.CmsAppAnnouncement{},
		&entities.CmsArCreditnoteDetails{},
		&entities.CmsCreditnote{},
		&entities.CmsCreditnoteDetails{},
		&entities.CmsCreditnoteSales{},
		&entities.CmsCustomer{},
		&entities.CmsCustomerAgeingKo{},
		&entities.CmsCustomerBranch{},
		&entities.CmsCustomerContraLocal{},
		&entities.CmsCustomerMerchandSched{},
		&entities.CmsCustomerMerchandSchedDtl{},
		&entities.CmsCustomerMerchandSchedSeq{},
		&entities.CmsCustomerProducts{},
		&entities.CmsCustomerRefund{},
		&entities.CmsCustomerSalesperson{},
		&entities.CmsCustomerVisitSched{},
		&entities.CmsCustomerVisitSchedLog{},
		&entities.CmsCustomerZone{},
		&entities.CmsDebitnote{},
		&entities.CmsDebitnoteDetails{},
		&entities.CmsDebitnoteSales{},
		&entities.CmsDeliveryInfo{},
		&entities.CmsDo{},
		&entities.CmsDoDetails{},
		&entities.CmsDoJob{},
		&entities.CmsInvoice{},
		&entities.CmsInvoiceDetails{},
		&entities.CmsInvoiceSales{},
		&entities.CmsLogin{},
		&entities.CmsMobileModule{},
		&entities.CmsModule{},
		&entities.CmsOrder{},
		&entities.CmsOrderItem{},
		&entities.CmsOutstandingSo{},
		&entities.CmsPackage{},
		&entities.CmsPackageDtl{},
		&entities.CmsPayment{},
		&entities.CmsPaymentDetail{},
		&entities.CmsPaymentGatewayBills{},
		&entities.CmsPaymentGatewayCollections{},
		&entities.CmsPaymentGatewayLog{},
		&entities.CmsPdfImage{},
		&entities.CmsProduct{},
		&entities.CmsProductAtch{},
		&entities.CmsProductAttribute{},
		&entities.CmsProductBatch{},
		&entities.CmsProductBrand{},
		&entities.CmsProductCategory{},
		&entities.CmsProductGroup{},
		&entities.CmsProductImage{},
		&entities.CmsProductPriceV2{},
		&entities.CmsProductUomPriceV2{},
		&entities.CmsProject{},
		&entities.CmsPurchaseReturn{},
		&entities.CmsPurchaseReturnDtl{},
		&entities.CmsReceipt{},
		&entities.CmsReportCollection{},
		&entities.CmsReportSales{},
		&entities.CmsSalespersonDevice{},
		&entities.CmsSalespersonUploads{},
		&entities.CmsSalespersonUploadsType{},
		&entities.CmsSerialNo{},
		&entities.CmsSetting{},
		&entities.CmsStockAdjustment{},
		&entities.CmsStockAdjustmentDtl{},
		&entities.CmsStockCard{},
		&entities.CmsStockTake{},
		&entities.CmsStockTakeDtl{},
		&entities.CmsStockTmplt{},
		&entities.CmsStockTmpltBind{},
		&entities.CmsStockTmpltDtl{},
		&entities.CmsStockTransfer{},
		&entities.CmsStockTransferDtl{},
		&entities.CmsVisitReport{},
		&entities.CmsVisitReportDocuments{},
		&entities.CmsWarehouse{},
		&entities.CmsWarehouseStock{},
	}
}
