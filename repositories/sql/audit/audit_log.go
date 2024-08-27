package audit

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	iterator "github.com/ledongthuc/goterators"
	"log"
	"time"
	"xorm.io/xorm"
)

// AuditLogRepository represents a repository for storing audit log information.
type AuditLogRepository struct {
	opt contracts.IDatabaseUserSession
	db  *xorm.Engine
}

// NewAuditLogRepository creates a new instance of AuditLogRepository with the given db and session.
// It returns the pointer to the AuditLogRepository struct.
func NewAuditLogRepository(db *xorm.Engine, session contracts.IDatabaseUserSession) *AuditLogRepository {
	return &AuditLogRepository{
		db:  db,
		opt: session,
	}
}

// Log inserts the given audit logs into the database.
// Each audit log's OperationTime is set to the current time.
// The AppName and UserCode fields of each audit log are set based on the
// options provided in the AuditLogRepository.
func (r *AuditLogRepository) Log(audits []*entities.AuditLog) {
	_, err := r.db.Insert(iterator.Map(audits, func(item *entities.AuditLog) *entities.AuditLog {
		item.OperationTime = time.Now()
		item.AppName = r.opt.GetApp()
		item.UserCode = r.opt.GetUser()
		return item
	}))
	if err != nil {
		log.Println("AuditLogRepository.Log:", err)
	}
}
