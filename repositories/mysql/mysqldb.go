package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

type MySqlDb struct {
	Engine *xorm.Engine
}

func NewMySqlDb() *MySqlDb {
	return &MySqlDb{}
}

func (m *MySqlDb) Open(conn string) (err error) {
	m.Engine, err = xorm.NewEngine("mysql", conn, func(db *sql.DB) error {
		db.SetMaxOpenConns(4)
		db.SetConnMaxLifetime(-1)
		db.SetMaxIdleConns(1)
		db.SetConnMaxIdleTime(time.Minute * 1)
		err := db.Ping()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *MySqlDb) Close() error {
	err := m.Engine.Close()
	if err != nil {
		return err
	}
	return nil
}
