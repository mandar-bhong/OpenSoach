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
	Config            *gmodels.ConfigSettings
	Context           *core.Context
	MasterTaskContext *taskqueue.TaskContext
	ProdTaskContext   *taskqueue.TaskContext
}

func Init() {
	once.Do(func() {
		r = &repo{}
	})
}

func Instance() *repo {
	return r
}
