package entities

import (
	"time"
)

type CmsDoJob struct {
	RunningId        int       `xorm:"not null pk autoincr INT"`
	JobId            int       `xorm:"not null unique INT"`
	DoCode           string    `xorm:"not null default '' VARCHAR(500)"`
	StartTime        time.Time `xorm:"DATETIME"`
	EndTime          time.Time `xorm:"DATETIME"`
	RiderName        string    `xorm:"not null default '' VARCHAR(50)"`
	JobRemark        string    `xorm:"VARCHAR(200)"`
	JobStatus        string    `xorm:"not null default '' VARCHAR(20)"`
	ActiveStatus     int       `xorm:"not null default 1 INT"`
	Misc             int       `xorm:"default 0 INT"`
	CancelReason     string    `xorm:"VARCHAR(200)"`
	UpdatedAt        time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdatedCount     int       `xorm:"not null default 0 INT"`
	CreatedAt        time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
	TransportNo      int       `xorm:"default 10001 INT"`
	TransportNote    string    `xorm:"VARCHAR(200)"`
	TransportCompany string    `xorm:"VARCHAR(200)"`
	TransportStatus  int       `xorm:"INT"`
	LastPrintAt      time.Time `xorm:"DATETIME"`
}

func (m *CmsDoJob) TableName() string {
	return "cms_do_job"
}
