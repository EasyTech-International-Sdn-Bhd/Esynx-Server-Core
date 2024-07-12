package entities

type CmsProject struct {
	Id             uint64  `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	ProjectCode    string  `xorm:"unique VARCHAR(200)"`
	ProjectDesc    string  `xorm:"VARCHAR(500)"`
	ProjectValue   float64 `xorm:"default 0 DOUBLE"`
	ProjectCost    float64 `xorm:"default 0 DOUBLE"`
	ProjectStatus  string  `xorm:"ENUM('ACTIVE','CANCELLED')"`
	ProjectDetails string  `xorm:"JSON"`
}

func (m *CmsProject) TableName() string {
	return "cms_project"
}
