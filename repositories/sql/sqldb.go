package sql

import (
	"database/sql"
	"fmt"
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
	engine, err := connectWithRetry(conn)
	if err != nil {
		return err
	}
	m.Engine = engine
	if logger != nil {
		m.Engine.SetLogger(logger)
	}
	m.Engine.ShowSQL(true)
	m.Engine.SetLogLevel(0)
	return nil
}

func connectWithRetry(conn string) (*xorm.Engine, error) {
	const maxRetries = 3
	const retryDelay = 2 * time.Second

	var engine *xorm.Engine
	var err error

	for i := 1; i <= maxRetries; i++ {
		engine, err = xorm.NewEngine("mysql", conn, func(db *sql.DB) error {
			db.SetMaxOpenConns(1)
			db.SetMaxIdleConns(1)
			db.SetConnMaxLifetime(time.Minute)
			return nil
		})
		if err == nil {
			if pingErr := engine.Ping(); pingErr == nil {
				return engine, nil
			} else {
				err = pingErr
			}
		}
		if i < maxRetries {
			time.Sleep(retryDelay)
		}
	}

	return nil, fmt.Errorf("could not connect to the database after %d attempts: %v", maxRetries, err)
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
