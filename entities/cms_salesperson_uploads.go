package entities

import (
	"time"
)

type CmsSalespersonUploads struct {
	UploadId            uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"uploadId,omitempty" xml:"uploadId"`
	UploadImage         string    `xorm:"not null unique VARCHAR(200)" json:"uploadImage,omitempty" xml:"uploadImage"`
	UploadTypeName      string    `xorm:"not null comment('module name from cms_module_camera') VARCHAR(200)" json:"uploadTypeName,omitempty" xml:"uploadTypeName"`
	UploadRemark        string    `xorm:"VARCHAR(200)" json:"uploadRemark,omitempty" xml:"uploadRemark"`
	UploadStatus        int       `xorm:"not null default 1 comment('1 - not deleted 0 - deleted') INT" json:"uploadStatus,omitempty" xml:"uploadStatus"`
	UploadSalespersonId int       `xorm:"not null INT" json:"uploadSalespersonId,omitempty" xml:"uploadSalespersonId"`
	UploadLocation      string    `xorm:"VARCHAR(200)" json:"uploadLocation,omitempty" xml:"uploadLocation"`
	UploadBindId        string    `xorm:"not null comment('id - from mobile') VARCHAR(200)" json:"uploadBindId,omitempty" xml:"uploadBindId"`
	UploadBindType      string    `xorm:"not null comment('SO/CO') VARCHAR(200)" json:"uploadBindType,omitempty" xml:"uploadBindType"`
	UploadDate          time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"uploadDate,omitempty" xml:"uploadDate"`
	TakenDate           time.Time `xorm:"DATETIME" json:"takenDate,omitempty" xml:"takenDate"`
}

func (m *CmsSalespersonUploads) TableName() string {
	return "cms_salesperson_uploads"
}
