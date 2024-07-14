package stock

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/entities"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"xorm.io/xorm"
)

type CmsProductBatchRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
}

func NewCmsProductBatchRepository(option *contracts.IRepository) *CmsProductBatchRepository {
	return &CmsProductBatchRepository{
		db:    option.Db,
		audit: option.Audit,
	}
}

func (r *CmsProductBatchRepository) Get(productCode string) ([]*entities.CmsProductBatch, error) {
	var records []*entities.CmsProductBatch
	err := r.db.Where("product_code = ? AND active_status = ?", productCode, 1).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsProductBatchRepository) GetByWarehouse(productCode string, warehouse string) ([]*entities.CmsProductBatch, error) {
	var records []*entities.CmsProductBatch
	err := r.db.Where("product_code = ? AND wh_code = ? AND active_status = ?", productCode, warehouse, 1).Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *CmsProductBatchRepository) InsertMany(records []*entities.CmsProductBatch) error {
	_, err := r.db.Insert(iterator.Map(records, func(item *entities.CmsProductBatch) *entities.CmsProductBatch {
		item.Validate()
		item.ToUpdate()
		return item
	}))
	if err != nil {
		return err
	}

	go r.log("INSERT", records)

	return nil
}

func (r *CmsProductBatchRepository) Update(record *entities.CmsProductBatch) error {
	_, err := r.db.Where("id = ?", record.Id).Update(record)
	if err != nil {
		return err
	}

	go r.log("UPDATE", []*entities.CmsProductBatch{record})

	return nil
}

func (r *CmsProductBatchRepository) UpdateMany(records []*entities.CmsProductBatch) error {
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

	go r.log("UPDATE", records)

	return nil
}

func (r *CmsProductBatchRepository) Delete(record *entities.CmsProductBatch) error {
	record.ActiveStatus = 0
	record.ToUpdate()
	return r.Update(record)
}

func (r *CmsProductBatchRepository) DeleteMany(records []*entities.CmsProductBatch) error {
	for _, record := range records {
		record.ActiveStatus = 0
		record.ToUpdate()
	}
	return r.UpdateMany(records)
}

func (r *CmsProductBatchRepository) log(op string, payload []*entities.CmsProductBatch) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsProductBatch) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordID:      fmt.Sprintf("%s.%s", item.BatchCode, item.ProductCode),
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
