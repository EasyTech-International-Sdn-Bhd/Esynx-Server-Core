package mysql

import (
	"database/sql"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	migrate "github.com/easytech-international-sdn-bhd/esynx-server-core/migrate/sql"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type MySqlDb struct {
	Engine *xorm.Engine
}

func NewMySqlDb() *MySqlDb {
	return &MySqlDb{}
}

func (m *MySqlDb) Open(conn string, logger contracts.IDatabaseLogger) (err error) {
	m.Engine, err = xorm.NewEngine("mysql", conn, func(db *sql.DB) error {
		db.SetMaxOpenConns(1)
		db.SetMaxIdleConns(1)
		db.SetConnMaxLifetime(-1)
		err := db.Ping()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	if logger != nil {
		m.Engine.SetLogger(logger)
	}
	m.Engine.ShowSQL(true)
	m.Engine.SetLogLevel(0)
	return nil
}

func (m *MySqlDb) DefineSchema() error {
	return migrate.DefineSchema(m.Engine)
}

func (m *MySqlDb) Close() error {
	err := m.Engine.Close()
	if err != nil {
		return err
	}
	return nil
}
