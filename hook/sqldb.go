package hook

import (
	"context"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	"xorm.io/xorm/contexts"
)

type SqlDBHook struct {
	logger contracts.IDatabaseLogger
}

func NewSqlDBHook(logger contracts.IDatabaseLogger) *SqlDBHook {
	return &SqlDBHook{
		logger: logger,
	}
}

func (h *SqlDBHook) BeforeProcess(c *contexts.ContextHook) (context.Context, error) {
	return c.Ctx, nil
}

func (h *SqlDBHook) AfterProcess(c *contexts.ContextHook) error {
	if c.Err != nil {
		h.logger.Errorf("Error: %s, [SQL] - %s", c.Err.Error(), c.SQL)
	}
	if c.ExecuteTime.Seconds() > 10 {
		h.logger.Warnf("Slow SQL: %s, [SQL] - %s", c.ExecuteTime.String(), c.SQL)
	}
	return nil
}
