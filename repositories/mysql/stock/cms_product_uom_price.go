package stock

import (
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/entities"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"xorm.io/xorm"
)

type CmsProductUomPriceRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
}

func NewCmsProductUomPriceRepository(option *contracts.IRepository) *CmsProductUomPriceRepository {
	return &CmsProductUomPriceRepository{
		db:    option.Db,
		audit: option.Audit,
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

	go r.log("INSERT", records)

	return nil
}

func (r *CmsProductUomPriceRepository) Update(record *entities.CmsProductUomPriceV2) error {
	_, err := r.db.Where("product_uom_price_id = ?", record.ProductUomPriceId).Update(record)
	if err != nil {
		return err
	}

	go r.log("UPDATE", []*entities.CmsProductUomPriceV2{record})

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

	go r.log("UPDATE", records)

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

func (r *CmsProductUomPriceRepository) log(op string, payload []*entities.CmsProductUomPriceV2) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsProductUomPriceV2) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordID:      item.ProductCode,
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
