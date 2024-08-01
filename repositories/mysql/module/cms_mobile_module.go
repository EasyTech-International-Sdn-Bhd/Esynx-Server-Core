package module

import (
	"encoding/json"
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/mysql/agent"
	"xorm.io/xorm"
)

type CmsMobileModuleRepository struct {
	db *xorm.Engine
	l  *agent.CmsLoginRepository
}

func NewCmsMobileModuleRepository(option *contracts.IRepository) *CmsMobileModuleRepository {
	return &CmsMobileModuleRepository{
		db: option.Db,
		l:  agent.NewCmsLoginRepository(option),
	}
}

func (r *CmsMobileModuleRepository) Get(module string) (string, error) {
	var result entities.CmsMobileModule
	has, err := r.db.Where("module=?", module).Get(&result)
	if err != nil {
		return "", err
	}
	if !has {
		return "", nil
	}
	return string(result.Status), nil
}

func (r *CmsMobileModuleRepository) SalesmanGroup() (map[string][]string, error) {
	s, err := r.Get("app_sp_group")
	if err != nil {
		return nil, err
	}
	if s != "" && s != "0" {
		return r.parseSalesmanGroup(s)
	}
	return nil, nil
}

func (r *CmsMobileModuleRepository) parseSalesmanGroup(s string) (map[string][]string, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		return nil, err
	}
	group := make(map[string][]string)
	for k, v := range result {
		var each []string
		if v.(string) == "1" {
			records, err := r.l.GetAll()
			if err != nil {
				return nil, err
			}
			for _, record := range records {
				each = append(each, record.StaffCode)
			}
			group[k] = each
		}
		if v.([]map[string]interface{}) != nil {
			agentGroup := v.([]map[string]interface{})
			for _, agentGroup := range agentGroup {
				each = append(each, agentGroup["id"].(string))
			}
			group[k] = each
		}
	}
	return group, nil
}
