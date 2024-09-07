package hooks

import (
	"context"
	"github.com/easytech-international-sdn-bhd/esynx-server-core/contracts"
	hook "xorm.io/xorm/contexts"
)

type SQLHook struct {
	logger contracts.IDatabaseLogger
}

func NewSQLHook(logger contracts.IDatabaseLogger) *SQLHook {
	return &SQLHook{
		logger,
	}
}

func (h *SQLHook) BeforeProcess(c *hook.ContextHook) (context.Context, error) {
	return nil, nil
}

func (h *SQLHook) AfterProcess(c *hook.ContextHook) error {
	if c.Err != nil {
		h.logger.Errorf("Error: %s, [SQL] - %s", c.Err.Error(), c.SQL)
	}
	if c.ExecuteTime.Seconds() > 10 {
		h.logger.Errorf("Slow SQL: %s, [SQL] - %s", c.ExecuteTime.String(), c.SQL)
	}
	return nil
}
