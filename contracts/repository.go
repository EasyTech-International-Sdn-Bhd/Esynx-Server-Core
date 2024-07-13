package contracts

import (
	"xorm.io/xorm"
)

type IRepository struct {
	Db      *xorm.Engine
	User    string
	AppName string
	Audit   IAuditLog
}
