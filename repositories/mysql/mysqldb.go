package mysql

import (
	"database/sql"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	migrate "github.com/easytech-international-sdn-bhd/esynx-server-core/migrate/sql"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

// MySqlDb represents a wrapper for the xorm.Engine that provides functions for interacting with a MySQL database.
// Open initializes a connection to the MySQL database using the provided connection string and logger.
// Returns an error if the connection fails.
type MySqlDb struct {
	Engine *xorm.Engine
}

// NewMySqlDb returns a new instance of MySqlDb.
func NewMySqlDb() *MySqlDb {
	return &MySqlDb{}
}

// Open initializes a connection to the MySQL database using the provided connection string and logger.
// Returns an error if the connection fails.
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

// DefineSchema creates all the tables in the database.
// !! Use this only when you create a new database
func (m *MySqlDb) DefineSchema() error {
	return migrate.DefineSchema(m.Engine)
}

// Close closes the connection to the MySQL database.
// Returns an error if closing the connection fails.
func (m *MySqlDb) Close() error {
	err := m.Engine.Close()
	if err != nil {
		return err
	}
	return nil
}

func (m *MySqlDb) GetEngine() *xorm.Engine {
	return m.Engine
}
