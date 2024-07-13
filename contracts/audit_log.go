package contracts

import "github.com/easytech-international-sdn-bhd/core/entities"

type IAuditLog interface {
	Log(audits []*entities.AuditLog)
}
