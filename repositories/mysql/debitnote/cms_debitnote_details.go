package debitnote

import (
	"github.com/easytech-international-sdn-bhd/core/entities"
	"github.com/easytech-international-sdn-bhd/core/models"
	"github.com/easytech-international-sdn-bhd/core/repositories/mysql/stock"
	iterator "github.com/ledongthuc/goterators"
	"xorm.io/xorm"
)

type CmsDebitNoteDetailsRepository struct {
	db *xorm.Engine
	p  *stock.CmsProductRepository
}

func NewCmsDebitNoteDetailsRepository(db *xorm.Engine) *CmsDebitNoteDetailsRepository {
	return &CmsDebitNoteDetailsRepository{
		db: db,
		p:  stock.NewCmsProductRepository(db),
	}
}

func (r *CmsDebitNoteDetailsRepository) Get(debitNoteCode string) ([]*entities.CmsDebitnoteDetails, error) {
	var details []*entities.CmsDebitnoteDetails
	err := r.db.Where("dn_code = ? AND active_status = ?", debitNoteCode, 1).Find(&details)
	if err != nil {
		return nil, err
	}
	return details, nil
}

func (r *CmsDebitNoteDetailsRepository) GetMany(debitNoteCodes []string) ([]*entities.CmsDebitnoteDetails, error) {
	var details []*entities.CmsDebitnoteDetails
	err := r.db.In("dn_code", debitNoteCodes).Where("active_status = ?", 1).Find(&details)
	if err != nil {
		return nil, err
	}
	return details, nil
}

func (r *CmsDebitNoteDetailsRepository) GetWithProduct(debitNoteCode string) ([]*models.DebitNoteDetailsWithProduct, error) {
	details, err := r.Get(debitNoteCode)
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
	var results []*models.DebitNoteDetailsWithProduct
	for _, detail := range details {
		for _, product := range products {
			if detail.ItemCode == product.ProductCode {
				results = append(results, &models.DebitNoteDetailsWithProduct{
					D: detail,
					P: product,
				})
			}
		}
	}
	return results, nil
}

func (r *CmsDebitNoteDetailsRepository) InsertMany(details []*entities.CmsDebitnoteDetails) error {
	_, err := r.db.Insert(iterator.Map(details, func(item *entities.CmsDebitnoteDetails) *entities.CmsDebitnoteDetails {
		item.Validate()
		item.ToUpdate()
		return item
	}))
	if err != nil {
		return err
	}
	return nil
}

func (r *CmsDebitNoteDetailsRepository) Update(details *entities.CmsDebitnoteDetails) error {
	_, err := r.db.Where("id = ?", details.Id).Update(details)
	if err != nil {
		return err
	}
	return nil
}

func (r *CmsDebitNoteDetailsRepository) UpdateMany(details []*entities.CmsDebitnoteDetails) error {
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
	return nil
}

func (r *CmsDebitNoteDetailsRepository) Delete(details *entities.CmsDebitnoteDetails) error {
	details.ActiveStatus = 0
	details.ToUpdate()
	return r.Update(details)
}

func (r *CmsDebitNoteDetailsRepository) DeleteMany(details []*entities.CmsDebitnoteDetails) error {
	for _, detail := range details {
		detail.ActiveStatus = 0
		detail.ToUpdate()
	}
	return r.UpdateMany(details)
}
