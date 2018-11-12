package repository

import (
	"sync"

	"opensoach.com/core"
	"opensoach.com/core/manager/taskqueue"
	gmodels "opensoach.com/models"
	pcmodels "opensoach.com/prodcore/models"
)

var (
	r    *pcmodels.Repo
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
		r = &pcmodels.Repo{}
	})
}

func Instance() *pcmodels.Repo {
	return r
}
