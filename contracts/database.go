package contracts

import (
	"github.com/easytech-international-sdn-bhd/esynx-server-core/options"
	"xorm.io/xorm/log"
)

type IDatabase interface {
	Open(conn string, logger IDatabaseLogger) error
	DefineSchema() error
	Close() error
}

type IDatabaseUserSession interface {
	GetUser() string
	GetApp() string
	GetStore() options.DatabaseStore
	GetConnection() string
	GetLogger() IDatabaseLogger
}

type IDatabaseLogger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Level() log.LogLevel
	SetLevel(level log.LogLevel)
	ShowSQL(show ...bool)
	IsShowSQL() bool
	BeforeSQL(context *log.LogContext)
	AfterSQL(context *log.LogContext)
}
