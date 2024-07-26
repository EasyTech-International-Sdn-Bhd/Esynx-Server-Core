package entities

type CmsSalespersonUploadsType struct {
	TypeId     uint64 `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"typeId,omitempty" xml:"typeId"`
	TypeName   string `xorm:"not null VARCHAR(100)" json:"typeName,omitempty" xml:"typeName"`
	TypeStatus int    `xorm:"not null default 0 comment('0 means inactive and 1 means active') INT" json:"typeStatus,omitempty" xml:"typeStatus"`
}

func (m *CmsSalespersonUploadsType) TableName() string {
	return "cms_salesperson_uploads_type"
}
