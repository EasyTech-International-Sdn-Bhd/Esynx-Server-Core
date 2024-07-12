package creditnote

import (
	"github.com/easytech-international-sdn-bhd/core/entities"
	"github.com/easytech-international-sdn-bhd/core/models"
	"github.com/easytech-international-sdn-bhd/core/repositories/mysql/stock"
	iterator "github.com/ledongthuc/goterators"
	"xorm.io/xorm"
)

type CmsCreditNoteDetailsRepository struct {
	db *xorm.Engine
	p  *stock.CmsProductRepository
}

func NewCmsCreditNoteDetailsRepository(db *xorm.Engine) *CmsCreditNoteDetailsRepository {
	return &CmsCreditNoteDetailsRepository{
		db: db,
		p:  stock.NewCmsProductRepository(db),
	}
}

func (r *CmsCreditNoteDetailsRepository) Get(creditNoteCode string) ([]*entities.CmsCreditnoteDetails, error) {
	var details []*entities.CmsCreditnoteDetails
	err := r.db.Where("cn_code = ? AND active_status = ?", creditNoteCode, 1).Find(&details)
	if err != nil {
		return nil, err
	}
	return details, nil
}

func (r *CmsCreditNoteDetailsRepository) GetMany(creditNoteCodes []string) ([]*entities.CmsCreditnoteDetails, error) {
	var details []*entities.CmsCreditnoteDetails
	err := r.db.In("cn_code", creditNoteCodes).Where("active_status = ?", 1).Find(&details)
	if err != nil {
		return nil, err
	}
	return details, nil
}

func (r *CmsCreditNoteDetailsRepository) GetWithProduct(creditNoteCode string) ([]*models.CreditNoteDetailsWithProduct, error) {
	details, err := r.Get(creditNoteCode)
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
	var results []*models.CreditNoteDetailsWithProduct
	for _, detail := range details {
		for _, product := range products {
			if detail.ItemCode == product.ProductCode {
				results = append(results, &models.CreditNoteDetailsWithProduct{
					D: detail,
					P: product,
				})
			}
		}
	}
	return results, nil
}

func (r *CmsCreditNoteDetailsRepository) InsertMany(details []*entities.CmsCreditnoteDetails) error {
	_, err := r.db.Insert(iterator.Map(details, func(item *entities.CmsCreditnoteDetails) *entities.CmsCreditnoteDetails {
		item.Validate()
		item.ToUpdate()
		return item
	}))
	if err != nil {
		return err
	}
	return nil
}

func (r *CmsCreditNoteDetailsRepository) Update(details *entities.CmsCreditnoteDetails) error {
	_, err := r.db.Where("id = ?", details.Id).Update(details)
	if err != nil {
		return err
	}
	return nil
}

func (r *CmsCreditNoteDetailsRepository) UpdateMany(details []*entities.CmsCreditnoteDetails) error {
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

func (r *CmsCreditNoteDetailsRepository) Delete(details *entities.CmsCreditnoteDetails) error {
	details.ActiveStatus = 0
	details.ToUpdate()
	return r.Update(details)
}

func (r *CmsCreditNoteDetailsRepository) DeleteMany(details []*entities.CmsCreditnoteDetails) error {
	for _, detail := range details {
		detail.ActiveStatus = 0
		detail.ToUpdate()
	}
	return r.UpdateMany(details)
}
