package entities

type CmsMobileModule struct {
	Id     uint64 `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	Module string `xorm:"unique VARCHAR(20)" json:"module,omitempty" xml:"module"`
	Status []byte `xorm:"comment('0 = disable, 1 = enable') LONGBLOB" json:"status,omitempty" xml:"status"`
}

func (m *CmsMobileModule) TableName() string {
	return "cms_mobile_module"
}
