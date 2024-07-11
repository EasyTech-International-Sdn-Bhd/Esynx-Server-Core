package stock

import (
	"github.com/easytech-international-sdn-bhd/core/entities"
	iterator "github.com/ledongthuc/goterators"
	"xorm.io/xorm"
)

type CmsProductAtchRepository struct {
	db *xorm.Engine
}

func NewCmsProductAtchRepository(db *xorm.Engine) *CmsProductAtchRepository {
	return &CmsProductAtchRepository{
		db: db,
	}
}

func (r *CmsProductAtchRepository) Get(productCode string) ([]*entities.CmsProductAtch, error) {
	var record []*entities.CmsProductAtch
	err := r.db.Where("product_code = ? AND active_status = ?", productCode, 1).Find(&record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (r *CmsProductAtchRepository) InsertMany(records []*entities.CmsProductAtch) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}

	_, err = session.Insert(iterator.Map(records, func(item *entities.CmsProductAtch) *entities.CmsProductAtch {
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

func (r *CmsProductAtchRepository) Update(record *entities.CmsProductAtch) error {
	_, err := r.db.Where("id = ?", record.Id).Update(record)
	if err != nil {
		return err
	}
	return nil
}

func (r *CmsProductAtchRepository) UpdateMany(records []*entities.CmsProductAtch) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}
	var sessionErr error
	rollback := false
	for _, product := range records {
		product.Validate()
		product.ToUpdate()
		_, err = session.Where("id = ?", product.Id).Update(product)
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

func (r *CmsProductAtchRepository) Delete(record *entities.CmsProductAtch) error {
	record.ActiveStatus = 0
	record.ToUpdate()
	return r.Update(record)
}

func (r *CmsProductAtchRepository) DeleteMany(records []*entities.CmsProductAtch) error {
	for _, record := range records {
		record.ActiveStatus = 0
		record.ToUpdate()
	}
	return r.UpdateMany(records)
}
