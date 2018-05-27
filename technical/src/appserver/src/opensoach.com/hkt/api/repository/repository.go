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
	Config          *gmodels.ConfigSettings
	Context         *core.Context
	ProdTaskContext *taskqueue.TaskContext
}

func Init(config *gmodels.ConfigSettings, ctx *core.Context) {
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
