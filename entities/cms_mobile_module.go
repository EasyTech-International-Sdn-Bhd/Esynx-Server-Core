package entities

type CmsMobileModule struct {
	Id     uint64 `xorm:"not null pk autoincr unique UNSIGNED BIGINT"`
	Module string `xorm:"unique VARCHAR(20)"`
	Status []byte `xorm:"comment('0 = disable, 1 = enable') LONGBLOB"`
}

func (m *CmsMobileModule) TableName() string {
	return "cms_mobile_module"
}
