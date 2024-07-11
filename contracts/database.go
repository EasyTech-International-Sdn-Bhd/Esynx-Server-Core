package contracts

type IDatabase interface {
	Open(conn string) error
	Close() error
}