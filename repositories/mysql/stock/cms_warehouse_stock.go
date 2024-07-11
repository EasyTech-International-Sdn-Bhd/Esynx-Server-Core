package stock

import (
	"github.com/easytech-international-sdn-bhd/core/entities"
	iterator "github.com/ledongthuc/goterators"
	"xorm.io/xorm"
)

type CmsWarehouseStockRepository struct {
	db *xorm.Engine
}

func NewCmsWarehouseStockRepository(db *xorm.Engine) *CmsWarehouseStockRepository {
	return &CmsWarehouseStockRepository{
		db: db,
	}
}

func (r *CmsWarehouseStockRepository) Get(productCode string) ([]*entities.CmsWarehouseStock, error) {
	var records []*entities.CmsWarehouseStock
	err := r.db.Where("product_code = ? AND active_status = ?", productCode, 1).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsWarehouseStockRepository) InsertMany(records []*entities.CmsWarehouseStock) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}

	_, err = session.Insert(iterator.Map(records, func(item *entities.CmsWarehouseStock) *entities.CmsWarehouseStock {
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

func (r *CmsWarehouseStockRepository) Update(record *entities.CmsWarehouseStock) error {
	_, err := r.db.Where("id = ?", record.Id).Update(record)
	if err != nil {
		return err
	}
	return nil
}

func (r *CmsWarehouseStockRepository) UpdateMany(records []*entities.CmsWarehouseStock) error {
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
		_, err = session.Where("id = ?", record.Id).Update(record)
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

func (r *CmsWarehouseStockRepository) Delete(record *entities.CmsWarehouseStock) error {
	record.ActiveStatus = 0
	record.ToUpdate()
	return r.Update(record)
}

func (r *CmsWarehouseStockRepository) DeleteMany(records []*entities.CmsWarehouseStock) error {
	for _, record := range records {
		record.ActiveStatus = 0
		record.ToUpdate()
	}
	return r.UpdateMany(records)
}
