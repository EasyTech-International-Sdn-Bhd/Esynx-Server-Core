package stock

import (
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/entities"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"strconv"
	"xorm.io/xorm"
)

type CmsProductImageRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
}

func NewCmsProductImageRepository(option *contracts.IRepository) *CmsProductImageRepository {
	return &CmsProductImageRepository{
		db:    option.Db,
		audit: option.Audit,
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

	r.log("INSERT", records)

	return nil
}

func (r *CmsProductImageRepository) Update(record *entities.CmsProductImage) error {
	_, err := r.db.Where("product_image_id = ?", record.ProductImageId).Update(record)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsProductImage{record})

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

	r.log("UPDATE", records)

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

func (r *CmsProductImageRepository) log(op string, payload []*entities.CmsProductImage) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsProductImage) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      strconv.Itoa(item.ProductId),
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
