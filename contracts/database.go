package contracts

import "github.com/easytech-international-sdn-bhd/esynx-server-core/options"

type IDatabase interface {
	Open(conn string) error
	DefineSchema() error
	Close() error
}

type IDatabaseUserSession interface {
	GetUser() string
	GetApp() string
	GetStore() options.DatabaseStore
	GetConnection() string
}
