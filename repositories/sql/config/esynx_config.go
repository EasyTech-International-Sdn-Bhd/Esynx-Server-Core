package config

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/goccy/go-json"
	"xorm.io/builder"
	"xorm.io/xorm"
)

type EsynxConfigRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
}

func NewEsynxConfigRepository(option *contracts.IRepository) *EsynxConfigRepository {
	return &EsynxConfigRepository{
		db:    option.Db,
		audit: option.Audit,
	}
}

// GetAll retrieves all EsynxConfig entities from the database.
func (r *EsynxConfigRepository) GetAll() ([]*entities.EsynxConfig, error) {
	var records []*entities.EsynxConfig
	err := r.db.Find(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

// Insert inserts a new EsynxConfig entity into the database.
func (r *EsynxConfigRepository) Insert(esynxConfig *entities.EsynxConfig) error {
	_, err := r.db.Insert(esynxConfig)
	if err != nil {
		return err
	}

	r.log("INSERT", []*entities.EsynxConfig{esynxConfig})
	return nil
}

// Update updates an EsynxConfig entity in the database.
func (r *EsynxConfigRepository) Update(esynxConfig *entities.EsynxConfig) error {
	_, err := r.db.Where("service_id = ?", esynxConfig.ServiceId).Update(esynxConfig)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.EsynxConfig{esynxConfig})
	return nil
}

// Delete deletes an EsynxConfig entity from the database based on the configCode.
func (r *EsynxConfigRepository) Delete(serviceId string) error {
	_, err := r.db.Where("service_id = ?", serviceId).Delete(&entities.EsynxConfig{})
	if err != nil {
		return err
	}

	r.log("DELETE", []*entities.EsynxConfig{{ServiceId: serviceId}})
	return nil
}

// Find retrieves EsynxConfig entities based on a custom query defined by the given predicate.
func (r *EsynxConfigRepository) Find(predicate *builder.Builder) ([]*entities.EsynxConfig, error) {
	var records []*entities.EsynxConfig
	var t entities.EsynxConfig
	err := r.db.SQL(predicate.From(t.TableName())).Find(&records)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, nil
	}
	return records, nil
}

// log logs the operation and payload to the audit log.
func (r *EsynxConfigRepository) log(op string, payload []*entities.EsynxConfig) {
	record, _ := json.Marshal(payload)
	body := make([]*entities.AuditLog, len(payload))
	for i, item := range payload {
		body[i] = &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      item.ServiceId,
			RecordBody:    string(record),
		}
	}
	r.audit.Log(body)
}
