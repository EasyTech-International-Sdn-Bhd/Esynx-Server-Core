package entities

import (
	"time"
)

type CmsSalespersonUploads struct {
	UploadId            int       `xorm:"not null pk autoincr INT"`
	UploadImage         string    `xorm:"not null unique VARCHAR(200)"`
	UploadTypeName      string    `xorm:"not null comment('module name from cms_module_camera') VARCHAR(200)"`
	UploadRemark        string    `xorm:"VARCHAR(200)"`
	UploadStatus        int       `xorm:"not null default 1 comment('1 - not deleted 0 - deleted') INT"`
	UploadSalespersonId int       `xorm:"not null INT"`
	UploadLocation      string    `xorm:"VARCHAR(200)"`
	UploadBindId        string    `xorm:"not null comment('id - from mobile') VARCHAR(200)"`
	UploadBindType      string    `xorm:"not null comment('SO/CO') VARCHAR(200)"`
	UploadDate          time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
	TakenDate           time.Time `xorm:"DATETIME"`
}

func (m *CmsSalespersonUploads) TableName() string {
	return "cms_salesperson_uploads"
}
