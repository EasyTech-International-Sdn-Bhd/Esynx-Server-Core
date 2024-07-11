package stock

import (
	"github.com/easytech-international-sdn-bhd/core/entities"
	iterator "github.com/ledongthuc/goterators"
	"xorm.io/xorm"
)

type CmsProductPriceTagRepository struct {
	db *xorm.Engine
}

func NewCmsProductPriceTagRepository(db *xorm.Engine) *CmsProductPriceTagRepository {
	return &CmsProductPriceTagRepository{
		db: db,
	}
}

func (r *CmsProductPriceTagRepository) Get(productCode string) ([]*entities.CmsProductPriceV2, error) {
	var records []*entities.CmsProductPriceV2
	err := r.db.Where("product_code = ? AND active_status = ?", productCode, 1).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsProductPriceTagRepository) GetByCustCode(productCode string, custCode string) ([]*entities.CmsProductPriceV2, error) {
	var records []*entities.CmsProductPriceV2
	err := r.db.Where("product_code = ? AND active_status = ? AND cust_code = ?", productCode, 1, custCode).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsProductPriceTagRepository) GetByPriceType(productCode string, priceType string) ([]*entities.CmsProductPriceV2, error) {
	var records []*entities.CmsProductPriceV2
	err := r.db.Where("product_code = ? AND active_status = ? AND price_cat = ?", productCode, 1, priceType).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsProductPriceTagRepository) InsertMany(records []*entities.CmsProductPriceV2) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}
	_, err = session.Insert(iterator.Map(records, func(item *entities.CmsProductPriceV2) *entities.CmsProductPriceV2 {
		item.Validate()
		item.ToUpdate()
		return item
	}))
	if err != nil {
		err := session.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	err = session.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (r *CmsProductPriceTagRepository) Update(record *entities.CmsProductPriceV2) error {
	_, err := r.db.Where("product_price_id = ?", record.ProductPriceId).Update(record)
	if err != nil {
		return err
	}
	return nil
}

func (r *CmsProductPriceTagRepository) UpdateMany(records []*entities.CmsProductPriceV2) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}
	var sessionErr error
	rollback := false
	for _, record := range records {
		record.Validate()
		record.ToUpdate()
		_, err = session.Where("product_price_id = ?", record.ProductPriceId).Update(record)
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

func (r *CmsProductPriceTagRepository) Delete(record *entities.CmsProductPriceV2) error {
	record.ActiveStatus = 0
	record.ToUpdate()
	return r.Update(record)
}

func (r *CmsProductPriceTagRepository) DeleteMany(records []*entities.CmsProductPriceV2) error {
	for _, record := range records {
		record.ActiveStatus = 0
		record.ToUpdate()
	}
	return r.UpdateMany(records)
}
