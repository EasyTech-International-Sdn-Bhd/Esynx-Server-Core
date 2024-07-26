package entities

import (
	"time"
)

type CmsDoJob struct {
	RunningId        uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"runningId,omitempty" xml:"runningId"`
	JobId            int       `xorm:"not null unique INT" json:"jobId,omitempty" xml:"jobId"`
	DoCode           string    `xorm:"not null default '' VARCHAR(500)" json:"doCode,omitempty" xml:"doCode"`
	StartTime        time.Time `xorm:"DATETIME" json:"startTime,omitempty" xml:"startTime"`
	EndTime          time.Time `xorm:"DATETIME" json:"endTime,omitempty" xml:"endTime"`
	RiderName        string    `xorm:"not null default '' VARCHAR(50)" json:"riderName,omitempty" xml:"riderName"`
	JobRemark        string    `xorm:"VARCHAR(200)" json:"jobRemark,omitempty" xml:"jobRemark"`
	JobStatus        string    `xorm:"not null default '' VARCHAR(20)" json:"jobStatus,omitempty" xml:"jobStatus"`
	ActiveStatus     int       `xorm:"not null default 1 INT" json:"activeStatus,omitempty" xml:"activeStatus"`
	Misc             int       `xorm:"default 0 INT" json:"misc,omitempty" xml:"misc"`
	CancelReason     string    `xorm:"VARCHAR(200)" json:"cancelReason,omitempty" xml:"cancelReason"`
	UpdatedAt        time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updatedAt,omitempty" xml:"updatedAt"`
	UpdatedCount     int       `xorm:"not null default 0 INT" json:"updatedCount,omitempty" xml:"updatedCount"`
	CreatedAt        time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP" json:"createdAt,omitempty" xml:"createdAt"`
	TransportNo      int       `xorm:"default 10001 INT" json:"transportNo,omitempty" xml:"transportNo"`
	TransportNote    string    `xorm:"VARCHAR(200)" json:"transportNote,omitempty" xml:"transportNote"`
	TransportCompany string    `xorm:"VARCHAR(200)" json:"transportCompany,omitempty" xml:"transportCompany"`
	TransportStatus  int       `xorm:"INT" json:"transportStatus,omitempty" xml:"transportStatus"`
	LastPrintAt      time.Time `xorm:"DATETIME" json:"lastPrintAt,omitempty" xml:"lastPrintAt"`
}

func (m *CmsDoJob) TableName() string {
	return "cms_do_job"
}
