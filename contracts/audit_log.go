package contracts

import "github.com/easytech-international-sdn-bhd/esynx-common/entities"

type IAuditLog interface {
	Log(audits []*entities.AuditLog)
}
