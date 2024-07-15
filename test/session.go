package test

import (
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/options"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/mysql"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/mysql/audit"
	"xorm.io/xorm/log"
)

type TestSession struct {
}

func NewTestSession() *TestSession {
	return &TestSession{}
}

func (s *TestSession) GetUser() string {
	return "_test_"
}

func (s *TestSession) GetApp() string {
	return "_test_app_"
}

func (s *TestSession) GetStore() options.DatabaseStore {
	return options.MySQL
}

func (s *TestSession) GetConnection() string {
	return "root:mysql@tcp(127.0.0.1:3306)/easysale_elk?charset=utf8mb4&parseTime=True&loc=Local&timeout=2s"
}

func (s *TestSession) GetLogger() *log.Logger {
	return nil
}

func TestOption() (*contracts.IRepository, error) {
	session := NewTestSession()
	db := mysql.NewMySqlDb()
	err := db.Open(session.GetConnection(), session.GetLogger())
	if err != nil {
		return nil, err
	}
	return &contracts.IRepository{
		Db:      db.Engine,
		User:    session.GetUser(),
		AppName: session.GetApp(),
		Audit:   audit.NewAuditLogRepository(db.Engine, session),
	}, nil
}
