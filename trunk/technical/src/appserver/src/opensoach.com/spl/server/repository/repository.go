package repository

import (
	"sync"

	"opensoach.com/core"
	"opensoach.com/core/manager/taskqueue"
	gmodels "opensoach.com/models"
)

var (
	r    *repo
	once sync.Once
)

type repo struct {
	Config  *gmodels.SPLConfigSettings
	Context *core.Context
	TaskQue *taskqueue.TaskContext // Producer Task context
}

func Init(config *gmodels.SPLConfigSettings, ctx *core.Context) {
	once.Do(func() {
		r = &repo{
			Config:  config,
			Context: ctx,
		}
	})
}

func Instance() *repo {
	return r
}
