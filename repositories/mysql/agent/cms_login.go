package agent

import (
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/entities"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goccy/go-json"
	iterator "github.com/ledongthuc/goterators"
	"xorm.io/xorm"
)

type CmsLoginRepository struct {
	db    *xorm.Engine
	audit contracts.IAuditLog
}

func NewCmsLoginRepository(option *contracts.IRepository) *CmsLoginRepository {
	return &CmsLoginRepository{
		db:    option.Db,
		audit: option.Audit,
	}
}

func (r *CmsLoginRepository) Get(agentId int64) (*entities.CmsLogin, error) {
	var record entities.CmsLogin
	has, err := r.db.Where("login_id = ?", agentId).Get(&record)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return &record, nil
}

func (r *CmsLoginRepository) GetByAgentCode(agentCode string) (*entities.CmsLogin, error) {
	var record entities.CmsLogin
	has, err := r.db.Where("staff_code = ?", agentCode).Get(&record)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return &record, nil
}

func (r *CmsLoginRepository) GetAll() ([]*entities.CmsLogin, error) {
	var records []*entities.CmsLogin
	err := r.db.Find(&records)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, nil
	}
	return records, nil
}

func (r *CmsLoginRepository) InsertMany(records []*entities.CmsLogin) error {
	_, err := r.db.Insert(iterator.Map(records, func(item *entities.CmsLogin) *entities.CmsLogin {
		return item
	}))
	if err != nil {
		return err
	}

	r.log("INSERT", records)

	return nil
}

func (r *CmsLoginRepository) Update(record *entities.CmsLogin) error {
	_, err := r.db.Where("login_id = ?", record.LoginId).Update(record)
	if err != nil {
		return err
	}

	r.log("UPDATE", []*entities.CmsLogin{record})

	return nil
}

func (r *CmsLoginRepository) UpdateMany(records []*entities.CmsLogin) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}
	var sessionErr error
	rollback := false
	for _, record := range records {
		_, err = session.Where("login_id = ?", record.LoginId).Update(record)
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

func (r *CmsLoginRepository) Delete(record *entities.CmsLogin) error {
	record.LoginStatus = 0
	return r.Update(record)
}

func (r *CmsLoginRepository) DeleteMany(records []*entities.CmsLogin) error {
	for _, record := range records {
		record.LoginStatus = 0
	}
	return r.UpdateMany(records)
}

func (r *CmsLoginRepository) log(op string, payload []*entities.CmsLogin) {
	record, _ := json.Marshal(payload)
	body := iterator.Map(payload, func(item *entities.CmsLogin) *entities.AuditLog {
		return &entities.AuditLog{
			OperationType: op,
			RecordTable:   item.TableName(),
			RecordId:      item.StaffCode,
			RecordBody:    string(record),
		}
	})
	r.audit.Log(body)
}
