package core

import (
	"github.com/easytech-international-sdn-bhd/core/contracts"
	"github.com/easytech-international-sdn-bhd/core/repositories/mysql"
)

type DatabaseProvider int

const (
	MySQL DatabaseProvider = iota
	Firestore
)

type ESynx struct {
	engine contracts.IDatabase
}

func NewEsynxProvider(provider DatabaseProvider, conn string) (*ESynx, error) {
	if provider == MySQL {
		db := mysql.NewMySqlDb()
		err := db.Open(conn)
		if err != nil {
			return nil, err
		}
		return &ESynx{engine: db}, nil
	}
	if provider == Firestore {

	}
	return nil, nil
}

func (e *ESynx) Destroy() error {
	return e.engine.Close()
}
