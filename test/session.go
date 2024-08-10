package test

import (
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/mock"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/options"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/mysql/audit"
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
	return options.Mock
}

func (s *TestSession) GetConnection() string {
	return "./mock/mock.db"
}

func (s *TestSession) GetLogger() contracts.IDatabaseLogger {
	return nil
}

func TestOption() (*contracts.IRepository, error) {
	session := NewTestSession()
	db := mock.NewMockDb()
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
