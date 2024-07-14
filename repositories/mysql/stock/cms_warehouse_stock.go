package stock

import (
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/entities"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"xorm.io/xorm"
)

type CmsWarehouseStockRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
}

func NewCmsWarehouseStockRepository(option *contracts.IRepository) *CmsWarehouseStockRepository {
	return &CmsWarehouseStockRepository{
		db:    option.Db,
		audit: option.Audit,
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
	_, err := r.db.Insert(iterator.Map(records, func(item *entities.CmsWarehouseStock) *entities.CmsWarehouseStock {
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

func (r *CmsWarehouseStockRepository) Update(record *entities.CmsWarehouseStock) error {
	_, err := r.db.Where("id = ?", record.Id).Update(record)
	if err != nil {
		return err
	}

	go r.log("UPDATE", []*entities.CmsWarehouseStock{record})

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

	go r.log("UPDATE", records)

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

func (r *CmsWarehouseStockRepository) log(op string, payload []*entities.CmsWarehouseStock) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsWarehouseStock) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordID:      item.ProductCode,
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
