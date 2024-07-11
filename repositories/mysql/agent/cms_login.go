package agent

import (
	"github.com/easytech-international-sdn-bhd/core/entities"
	_ "github.com/go-sql-driver/mysql"
	iterator "github.com/ledongthuc/goterators"
	"xorm.io/xorm"
)

type CmsLoginRepository struct {
	db *xorm.Engine
}

func NewCmsLoginRepository(db *xorm.Engine) *CmsLoginRepository {
	return &CmsLoginRepository{
		db: db,
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
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}

	_, err = session.Insert(iterator.Map(records, func(item *entities.CmsLogin) *entities.CmsLogin {
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

func (r *CmsLoginRepository) Update(record *entities.CmsLogin) error {
	_, err := r.db.Where("login_id = ?", record.LoginId).Update(record)
	if err != nil {
		return err
	}
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
		record.Validate()
		record.ToUpdate()
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
	return nil
}

func (r *CmsLoginRepository) Delete(record *entities.CmsLogin) error {
	record.LoginStatus = 0
	record.ToUpdate()
	return r.Update(record)
}

func (r *CmsLoginRepository) DeleteMany(records []*entities.CmsLogin) error {
	for _, record := range records {
		record.LoginStatus = 0
		record.ToUpdate()
	}
	return r.UpdateMany(records)
}
