package stock

import (
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/entities"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"xorm.io/xorm"
)

type CmsProductAtchRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
}

func NewCmsProductAtchRepository(option *contracts.IRepository) *CmsProductAtchRepository {
	return &CmsProductAtchRepository{
		db:    option.Db,
		audit: option.Audit,
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
	_, err := r.db.Insert(iterator.Map(records, func(item *entities.CmsProductAtch) *entities.CmsProductAtch {
		return item
	}))
	if err != nil {
		return err
	}

	r.log("INSERT", records)

	return nil
}

func (r *CmsProductAtchRepository) Update(record *entities.CmsProductAtch) error {
	_, err := r.db.Where("id = ?", record.Id).Update(record)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsProductAtch{record})

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

	r.log("UPDATE", records)

	return nil
}

func (r *CmsProductAtchRepository) Delete(record *entities.CmsProductAtch) error {
	record.ActiveStatus = 0
	return r.Update(record)
}

func (r *CmsProductAtchRepository) DeleteMany(records []*entities.CmsProductAtch) error {
	for _, record := range records {
		record.ActiveStatus = 0
	}
	return r.UpdateMany(records)
}

func (r *CmsProductAtchRepository) log(op string, payload []*entities.CmsProductAtch) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsProductAtch) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      item.ProductCode,
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
