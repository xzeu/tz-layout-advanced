package task

import (
	"github.com/xzeu/tz-layout-advanced/internal/repository"
	"github.com/xzeu/tz-layout-advanced/pkg/jwt"
	"github.com/xzeu/tz-layout-advanced/pkg/log"
	"github.com/xzeu/tz-layout-advanced/pkg/sid"
)

type Task struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewTask(
	tm repository.Transaction,
	logger *log.Logger,
	sid *sid.Sid,
) *Task {
	return &Task{
		logger: logger,
		sid:    sid,
		tm:     tm,
	}
}
