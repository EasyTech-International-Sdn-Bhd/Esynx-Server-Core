package sql

import (
	"database/sql"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	migrate "github.com/easytech-international-sdn-bhd/esynx-server-core/migrate/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

// SqlDb represents a wrapper for the xorm.Engine that provides functions for interacting with a MySQL database.
// Open initializes a connection to the MySQL database using the provided connection string and logger.
// Returns an error if the connection fails.
type SqlDb struct {
	Engine *xorm.Engine
}

// NewSqlDb returns a new instance of SqlDb.
func NewSqlDb() *SqlDb {
	return &SqlDb{}
}

// Open initializes a connection to the MySQL database using the provided connection string and logger.
// Returns an error if the connection fails.
func (m *SqlDb) Open(conn string, logger contracts.IDatabaseLogger) (err error) {
	m.Engine, err = xorm.NewEngine("mysql", conn, func(db *sql.DB) error {
		db.SetMaxOpenConns(1)
		db.SetMaxIdleConns(0)
		db.SetConnMaxLifetime(time.Second * 5)
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

// DefineSchema creates all the tables in the database.
// !! Use this only when you create a new database
func (m *SqlDb) DefineSchema() error {
	return migrate.DefineSchema(m.Engine)
}

// Close closes the connection to the MySQL database.
// Returns an error if closing the connection fails.
func (m *SqlDb) Close() error {
	err := m.Engine.Close()
	if err != nil {
		return err
	}
	return nil
}

func (m *SqlDb) GetEngine() *xorm.Engine {
	return m.Engine
}
