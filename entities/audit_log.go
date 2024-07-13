package entities

import "time"

type AuditLog struct {
	AuditID       int64     `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	OperationType string    `xorm:"ENUM('INSERT', 'UPDATE') not null"`
	RecordTable   string    `xorm:"VARCHAR(100)"`
	RecordID      string    `xorm:"VARCHAR(80)"`
	RecordBody    string    `xorm:"JSON"`
	UserCode      string    `xorm:"VARCHAR(50) not null"`
	AppName       string    `xorm:"VARCHAR(20) not null"`
	OperationTime time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (m *AuditLog) TableName() string {
	return "audit_log"
}
