package mock

import (
	_ "database/sql"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	migrate "github.com/easytech-international-sdn-bhd/esynx-server-core/migrate/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"xorm.io/xorm"
)

// MockDb represents a wrapper for the xorm.Engine that provides functions for interacting with a SQLite database.
// Open initializes a connection to the SQLite database using the provided connection string and logger.
// Returns an error if the connection fails.
type MockDb struct {
	Engine *xorm.Engine
}

// NewMockDb returns a new instance of MockDb.
func NewMockDb() *MockDb {
	return &MockDb{}
}

// Open initializes a connection to the SQLite database using the provided connection string and logger.
// If the database file does not exist, it will create one.
// Returns an error if the connection fails.
func (m *MockDb) Open(conn string, logger contracts.IDatabaseLogger) (err error) {
	// Check if the database file exists, if not, create an empty file
	if _, err := os.Stat(conn); os.IsNotExist(err) {
		file, err := os.Create(conn)
		if err != nil {
			return err
		}
		file.Close()
	}

	m.Engine, err = xorm.NewEngine("sqlite3", conn)
	if err != nil {
		return err
	}
	m.Engine.DB().SetMaxOpenConns(1)
	m.Engine.DB().SetMaxIdleConns(1)
	m.Engine.DB().SetConnMaxLifetime(-1)
	if err := m.Engine.DB().Ping(); err != nil {
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
func (m *MockDb) DefineSchema() error {
	return migrate.DefineSchema(m.Engine)
}

// Close closes the connection to the SQLite database.
// Returns an error if closing the connection fails.
func (m *MockDb) Close() error {
	err := m.Engine.Close()
	if err != nil {
		return err
	}
	return nil
}

func (m *MockDb) GetEngine() *xorm.Engine {
	return m.Engine
}
