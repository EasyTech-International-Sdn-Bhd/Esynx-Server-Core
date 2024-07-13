package audit

import (
	"github.com/easytech-international-sdn-bhd/core/contracts"
	"github.com/easytech-international-sdn-bhd/core/entities"
	iterator "github.com/ledongthuc/goterators"
	"time"
	"xorm.io/xorm"
)

type AuditLogRepository struct {
	opt contracts.IDatabaseUserSession
	db  *xorm.Engine
}

func NewAuditLogRepository(db *xorm.Engine, session contracts.IDatabaseUserSession) *AuditLogRepository {
	return &AuditLogRepository{
		db:  db,
		opt: session,
	}
}

func (r *AuditLogRepository) Log(audits []*entities.AuditLog) {
	_, _ = r.db.Insert(iterator.Map(audits, func(item *entities.AuditLog) *entities.AuditLog {
		item.OperationTime = time.Now()
		item.AppName = r.opt.GetApp()
		item.UserCode = r.opt.GetUser()
		return item
	}))
}
