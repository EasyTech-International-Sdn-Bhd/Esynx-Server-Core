package contracts

import "github.com/easytech-international-sdn-bhd/esynx-server-core/entities"

type IAuditLog interface {
	Log(audits []*entities.AuditLog)
}
