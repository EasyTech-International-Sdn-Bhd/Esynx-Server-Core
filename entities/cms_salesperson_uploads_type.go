package entities

type CmsSalespersonUploadsType struct {
	TypeId     int    `xorm:"not null pk autoincr INT"`
	TypeName   string `xorm:"not null VARCHAR(100)"`
	TypeStatus int    `xorm:"not null default 0 comment('0 means inactive and 1 means active') INT"`
}

func (m *CmsSalespersonUploadsType) TableName() string {
	return "cms_salesperson_uploads_type"
}
