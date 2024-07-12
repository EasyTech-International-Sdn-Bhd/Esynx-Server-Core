package stock

import (
	"github.com/easytech-international-sdn-bhd/core/entities"
	iterator "github.com/ledongthuc/goterators"
	"xorm.io/xorm"
)

type CmsProductImageRepository struct {
	db *xorm.Engine
}

func NewCmsProductImageRepository(db *xorm.Engine) *CmsProductImageRepository {
	return &CmsProductImageRepository{
		db: db,
	}
}

func (r *CmsProductImageRepository) Get(productId string) ([]*entities.CmsProductImage, error) {
	var records []*entities.CmsProductImage
	err := r.db.Where("product_id = ? AND active_status = ?", productId, 1).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsProductImageRepository) GetByProductCode(productCode string) ([]*entities.CmsProductImage, error) {
	var records []*entities.CmsProductImage
	err := r.db.Where("product_code = ? AND active_status = ?", productCode, 1).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsProductImageRepository) InsertMany(records []*entities.CmsProductImage) error {
	_, err := r.db.Insert(iterator.Map(records, func(item *entities.CmsProductImage) *entities.CmsProductImage {
		item.Validate()
		item.ToUpdate()
		return item
	}))
	if err != nil {
		return err
	}
	return nil
}

func (r *CmsProductImageRepository) Update(record *entities.CmsProductImage) error {
	_, err := r.db.Where("product_image_id = ?", record.ProductImageId).Update(record)
	if err != nil {
		return err
	}
	return nil
}

func (r *CmsProductImageRepository) UpdateMany(records []*entities.CmsProductImage) error {
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
		_, err = session.Where("product_image_id = ?", record.ProductImageId).Update(record)
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

func (r *CmsProductImageRepository) Delete(record *entities.CmsProductImage) error {
	record.ActiveStatus = 0
	record.ToUpdate()
	return r.Update(record)
}

func (r *CmsProductImageRepository) DeleteMany(records []*entities.CmsProductImage) error {
	for _, record := range records {
		record.ActiveStatus = 0
		record.ToUpdate()
	}
	return r.UpdateMany(records)
}
