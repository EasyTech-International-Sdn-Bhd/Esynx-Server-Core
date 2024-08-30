package audit

import (
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	iterator "github.com/ledongthuc/goterators"
	"time"
)

// AuditLogRepository represents a repository for storing audit log information.
type AuditLogRepository struct {
	opt  contracts.IDatabaseUserSession
	dest contracts.IAuditLogger
}

// NewAuditLogRepository creates a new instance of AuditLogRepository with the given db and session.
// It returns the pointer to the AuditLogRepository struct.
func NewAuditLogRepository(session contracts.IDatabaseUserSession) *AuditLogRepository {
	return &AuditLogRepository{
		dest: session.GetAuditLogger(),
		opt:  session,
	}
}

// Log inserts the given audit logs into the database.
// Each audit log's OperationTime is set to the current time.
// The AppName and UserCode fields of each audit log are set based on the
// options provided in the AuditLogRepository.
func (r *AuditLogRepository) Log(audits []*entities.AuditLog) {
	r.dest.Insert(iterator.Map(audits, func(item *entities.AuditLog) *entities.AuditLog {
		item.OperationTime = time.Now()
		item.AppName = r.opt.GetApp()
		item.UserCode = r.opt.GetUser()
		return item
	}))
}
