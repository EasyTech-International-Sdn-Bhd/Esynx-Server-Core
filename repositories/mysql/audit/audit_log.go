package audit

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
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
	_, err := r.db.Insert(iterator.Map(audits, func(item *entities.AuditLog) *entities.AuditLog {
		item.OperationTime = time.Now()
		item.AppName = r.opt.GetApp()
		item.UserCode = r.opt.GetUser()
		return item
	}))
	fmt.Printf("logging +++++++++++ %v\n", err)
}
