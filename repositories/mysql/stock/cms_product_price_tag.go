package stock

import (
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/entities"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"xorm.io/xorm"
)

type CmsProductPriceTagRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
}

func NewCmsProductPriceTagRepository(option *contracts.IRepository) *CmsProductPriceTagRepository {
	return &CmsProductPriceTagRepository{
		db:    option.Db,
		audit: option.Audit,
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
	_, err := r.db.Insert(iterator.Map(records, func(item *entities.CmsProductPriceV2) *entities.CmsProductPriceV2 {
		item.Validate()
		item.ToUpdate()
		return item
	}))
	if err != nil {
		return err
	}

	r.log("INSERT", records)

	return nil
}

func (r *CmsProductPriceTagRepository) Update(record *entities.CmsProductPriceV2) error {
	_, err := r.db.Where("product_price_id = ?", record.ProductPriceId).Update(record)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsProductPriceV2{record})

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

	r.log("UPDATE", records)

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

func (r *CmsProductPriceTagRepository) log(op string, payload []*entities.CmsProductPriceV2) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsProductPriceV2) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      item.ProductCode,
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
