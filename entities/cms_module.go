package entities

type CmsModule struct {
	ModuleId uint64 `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	Name     string `xorm:"not null VARCHAR(100)"`
	Value    []byte `xorm:"BLOB"`
}

func (m *CmsModule) TableName() string {
	return "cms_module"
}
