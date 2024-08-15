package module

import (
	"encoding/json"
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/sql/agent"
	"xorm.io/xorm"
)

// CmsMobileModuleRepository is a repository for accessing CMS mobile modules.
type CmsMobileModuleRepository struct {
	db *xorm.Engine
	l  *agent.CmsLoginRepository
}

// NewCmsMobileModuleRepository creates a new instance of CmsMobileModuleRepository with the given IRepository option. The option should have a valid db connection and audit log implementation.
func NewCmsMobileModuleRepository(option *contracts.IRepository) *CmsMobileModuleRepository {
	return &CmsMobileModuleRepository{
		db: option.Db,
		l:  agent.NewCmsLoginRepository(option),
	}
}

// Get retrieves the status of a given mobile module from the CmsMobileModuleRepository.
// It takes a string parameter `module` representing the name of the module to be retrieved.
// It returns a string representing the status of the module and an error object if any error occurs.
// If the module does not exist, it returns an empty string and a nil error.
// Example usage:
//
//	r := &CmsMobileModuleRepository{...}
//	status, err := r.Get("module_name")
//	if err != nil {
//	    // handle error
//	}
//	fmt.Println(status)
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

// SalesmanGroup retrieves the salesman group from the "app_sp_group" module in the CmsMobileModuleRepository.
// It first retrieves the module value using the Get method, then, if the value is not empty or "0", it invokes
// the parseSalesmanGroup method to parse the value into a map[string][]string format, representing the salesman
// group. The result map contains the group name as the key and an array of salesman IDs as the value. If any error
// occurs during the process, it returns nil and the error.
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

// parseSalesmanGroup parses a JSON string `s` and returns a map of sales group
// where each group is represented by a key-value pair, where the key is the
// group name and the value is a slice of staff codes or agent IDs.
// The function uses the `json.Unmarshal` function to parse the input JSON string.
// It then iterates over each key-value pair in the parsed result.
// If the value is a string with value "1", it retrieves all records using
// `r.l.GetAll` method and appends each record's staff code to the slice.
// If the value is a slice of maps, it retrieves each agent ID from the maps
// and appends it to the slice.
// Finally, it returns the map of sales group and an error, if any.
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
