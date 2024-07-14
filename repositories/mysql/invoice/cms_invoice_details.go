package invoice

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/models"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/mysql/stock"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"xorm.io/xorm"
)

type CmsInvoiceDetailsRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
	p     *stock.CmsProductRepository
}

func NewCmsInvoiceDetailsRepository(option *contracts.IRepository) *CmsInvoiceDetailsRepository {
	return &CmsInvoiceDetailsRepository{
		db:    option.Db,
		audit: option.Audit,
		p:     stock.NewCmsProductRepository(option),
	}
}

func (r *CmsInvoiceDetailsRepository) Get(invoiceCode string) ([]*entities.CmsInvoiceDetails, error) {
	var details []*entities.CmsInvoiceDetails
	err := r.db.Where("invoice_code = ? AND active_status = ?", invoiceCode, 1).Find(&details)
	if err != nil {
		return nil, err
	}
	return details, nil
}

func (r *CmsInvoiceDetailsRepository) GetMany(invoiceCodes []string) ([]*entities.CmsInvoiceDetails, error) {
	var details []*entities.CmsInvoiceDetails
	err := r.db.In("invoice_code", invoiceCodes).Where("active_status = ?", 1).Find(&details)
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
	_, err := r.db.Insert(iterator.Map(details, func(item *entities.CmsInvoiceDetails) *entities.CmsInvoiceDetails {
		item.Validate()
		item.ToUpdate()
		return item
	}))
	if err != nil {
		return err
	}

	go r.log("INSERT", details)

	return nil
}

func (r *CmsInvoiceDetailsRepository) Update(details *entities.CmsInvoiceDetails) error {
	_, err := r.db.Where("id = ?", details.Id).Update(details)
	if err != nil {
		return err
	}

	go r.log("UPDATE", []*entities.CmsInvoiceDetails{details})

	return nil
}

func (r *CmsInvoiceDetailsRepository) UpdateMany(details []*entities.CmsInvoiceDetails) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}
	var sessionErr error
	rollback := false
	for _, detail := range details {
		detail.Validate()
		detail.ToUpdate()
		_, err = session.Where("id = ?", detail.Id).Update(detail)
		if err != nil {
			rollback = true
			sessionErr = err
			break
		}
	}
	if rollback {
		err := session.Rollback()
		if err != nil {
			return err
		}
		return sessionErr
	}
	err = session.Commit()
	if err != nil {
		return err
	}

	go r.log("UPDATE", details)

	return nil
}

func (r *CmsInvoiceDetailsRepository) Delete(details *entities.CmsInvoiceDetails) error {
	details.ActiveStatus = 0
	details.ToUpdate()
	return r.Update(details)
}

func (r *CmsInvoiceDetailsRepository) DeleteMany(details []*entities.CmsInvoiceDetails) error {
	for _, detail := range details {
		detail.ActiveStatus = 0
		detail.ToUpdate()
	}
	return r.UpdateMany(details)
}

func (r *CmsInvoiceDetailsRepository) log(op string, payload []*entities.CmsInvoiceDetails) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsInvoiceDetails) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordID:      fmt.Sprintf("%s.%s", item.InvoiceCode, item.ItemCode),
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
