package entities

type CmsModule struct {
	ModuleId int    `xorm:"not null pk autoincr INT"`
	Name     string `xorm:"not null VARCHAR(100)"`
	Value    []byte `xorm:"BLOB"`
}

func (m *CmsModule) TableName() string {
	return "cms_module"
}
