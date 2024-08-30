package test

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/options"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/sql"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/repositories/sql/audit"
)

type fakeAuditLogger struct{}

func newFakeAuditLogger() *fakeAuditLogger {
	return &fakeAuditLogger{}
}
func (l *fakeAuditLogger) Insert(data []*entities.AuditLog) {

}

type TestSession struct {
}

func (s *TestSession) GetAuditLogger() contracts.IAuditLogger {
	return newFakeAuditLogger()
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
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=2s",
		"root",
		"mysql",
		"localhost",
		3306,
		"esynx_vit",
	)
}

func (s *TestSession) GetLogger() contracts.IDatabaseLogger {
	return nil
}

func TestOption() (*contracts.IRepository, error) {
	session := NewTestSession()
	db := sql.NewSqlDb()
	err := db.Open(session.GetConnection(), session.GetLogger())
	if err != nil {
		return nil, err
	}
	return &contracts.IRepository{
		Db:      db.Engine,
		User:    session.GetUser(),
		AppName: session.GetApp(),
		Audit:   audit.NewAuditLogRepository(session),
	}, nil
}
