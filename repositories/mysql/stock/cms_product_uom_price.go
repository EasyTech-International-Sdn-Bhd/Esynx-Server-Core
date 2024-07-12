package stock

import (
	"github.com/easytech-international-sdn-bhd/core/entities"
	iterator "github.com/ledongthuc/goterators"
	"xorm.io/xorm"
)

type CmsProductUomPriceRepository struct {
	db *xorm.Engine
}

func NewCmsProductUomPriceRepository(db *xorm.Engine) *CmsProductUomPriceRepository {
	return &CmsProductUomPriceRepository{
		db: db,
	}
}

func (r *CmsProductUomPriceRepository) Get(productCode string) ([]*entities.CmsProductUomPriceV2, error) {
	var records []*entities.CmsProductUomPriceV2
	err := r.db.Where("product_code = ? AND active_status = ?", productCode, 1).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsProductUomPriceRepository) InsertMany(records []*entities.CmsProductUomPriceV2) error {
	_, err := r.db.Insert(iterator.Map(records, func(item *entities.CmsProductUomPriceV2) *entities.CmsProductUomPriceV2 {
		item.Validate()
		item.ToUpdate()
		return item
	}))
	if err != nil {
		return err
	}
	return nil
}

func (r *CmsProductUomPriceRepository) Update(record *entities.CmsProductUomPriceV2) error {
	_, err := r.db.Where("product_uom_price_id = ?", record.ProductUomPriceId).Update(record)
	if err != nil {
		return err
	}
	return nil
}

func (r *CmsProductUomPriceRepository) UpdateMany(records []*entities.CmsProductUomPriceV2) error {
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
		_, err = session.Where("product_uom_price_id = ?", record.ProductUomPriceId).Update(record)
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

func (r *CmsProductUomPriceRepository) Delete(record *entities.CmsProductUomPriceV2) error {
	record.ActiveStatus = 0
	record.ToUpdate()
	return r.Update(record)
}

func (r *CmsProductUomPriceRepository) DeleteMany(records []*entities.CmsProductUomPriceV2) error {
	for _, record := range records {
		record.ActiveStatus = 0
		record.ToUpdate()
	}
	return r.UpdateMany(records)
}
