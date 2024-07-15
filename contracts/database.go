package contracts

import (
	"github.com/easytech-international-sdn-bhd/esynx-server-core/options"
	"xorm.io/xorm/log"
)

type IDatabase interface {
	Open(conn string, logger *log.Logger) error
	DefineSchema() error
	Close() error
}

type IDatabaseUserSession interface {
	GetUser() string
	GetApp() string
	GetStore() options.DatabaseStore
	GetConnection() string
	GetLogger() *log.Logger
}
