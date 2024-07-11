package invoice

import (
	"github.com/easytech-international-sdn-bhd/core/entities"
	"github.com/easytech-international-sdn-bhd/core/models"
	"github.com/easytech-international-sdn-bhd/core/repositories/mysql/stock"
	"xorm.io/xorm"
)

type CmsInvoiceDetailsRepository struct {
	db *xorm.Engine
	p  *stock.CmsProductRepository
}

func NewCmsInvoiceDetailsRepository(db *xorm.Engine) *CmsInvoiceDetailsRepository {
	return &CmsInvoiceDetailsRepository{
		db: db,
		p:  stock.NewCmsProductRepository(db),
	}
}

func (r *CmsInvoiceDetailsRepository) Get(invoiceCode string) ([]*entities.CmsInvoiceDetails, error) {
	var details []*entities.CmsInvoiceDetails
	err := r.db.Where("invoice_code = ?", invoiceCode).Find(&details)
	if err != nil {
		return nil, err
	}
	return details, nil
}

func (r *CmsInvoiceDetailsRepository) GetMany(invoiceCodes []string) ([]*entities.CmsInvoiceDetails, error) {
	var details []*entities.CmsInvoiceDetails
	err := r.db.In("invoice_code", invoiceCodes).Find(&details)
	if err != nil {
		return nil, err
	}
	return details, nil
}

func (r *CmsInvoiceDetailsRepository) GetWithProduct(invoiceCode string) ([]*models.InvoiceDetailsWithProduct, error) {
	details, err := r.Get(invoiceCode)
	if err != nil {
		return nil, err
	}
	var productCodes []string
	for _, detail := range details {
		if detail.ItemCode != "" {
			productCodes = append(productCodes, detail.ItemCode)
		}
	}
	products, err := r.p.GetMany(productCodes)
	if err != nil {
		return nil, err
	}
	var results []*models.InvoiceDetailsWithProduct
	for _, detail := range details {
		for _, product := range products {
			if detail.ItemCode == product.ProductCode {
				results = append(results, &models.InvoiceDetailsWithProduct{
					D: detail,
					P: product,
				})
			}
		}
	}
	return results, nil
}

func (r *CmsInvoiceDetailsRepository) InsertMany(details []*entities.CmsInvoiceDetails) error {

}

func (r *CmsInvoiceDetailsRepository) Update(details *entities.CmsInvoiceDetails) error {

}

func (r *CmsInvoiceDetailsRepository) UpdateMany(details []*entities.CmsInvoiceDetails) error {

}

func (r *CmsInvoiceDetailsRepository) Delete(details *entities.CmsInvoiceDetails) error {

}

func (r *CmsInvoiceDetailsRepository) DeleteMany(details []*entities.CmsInvoiceDetails) error {

}
