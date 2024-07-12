package contracts

type IDatabase interface {
	Open(conn string) error
	DefineSchema() error
	Close() error
}
