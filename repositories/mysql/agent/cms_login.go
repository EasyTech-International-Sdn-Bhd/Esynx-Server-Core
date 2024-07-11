package agent

import (
	"github.com/easytech-international-sdn-bhd/core/entities"
	_ "github.com/go-sql-driver/mysql"
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

func (r *CmsLoginRepository) Resolve(agentId int64) (*entities.CmsLogin, error) {
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
