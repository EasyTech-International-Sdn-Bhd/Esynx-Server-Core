package entities

type CmsProject struct {
	Id             uint64  `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	ProjectCode    string  `xorm:"unique VARCHAR(200)" json:"projectCode,omitempty" xml:"projectCode"`
	ProjectDesc    string  `xorm:"VARCHAR(500)" json:"projectDesc,omitempty" xml:"projectDesc"`
	ProjectValue   float64 `xorm:"default 0 DOUBLE" json:"projectValue,omitempty" xml:"projectValue"`
	ProjectCost    float64 `xorm:"default 0 DOUBLE" json:"projectCost,omitempty" xml:"projectCost"`
	ProjectStatus  string  `xorm:"ENUM('ACTIVE','CANCELLED')" json:"projectStatus,omitempty" xml:"projectStatus"`
	ProjectDetails string  `xorm:"JSON" json:"projectDetails,omitempty" xml:"projectDetails"`
}

func (m *CmsProject) TableName() string {
	return "cms_project"
}
