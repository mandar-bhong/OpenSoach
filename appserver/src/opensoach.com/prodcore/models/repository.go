package models

import (
	gmodels "opensoach.com/models"
	"opensoach.com/core"
	"opensoach.com/core/manager/taskqueue"
)

type Repo struct {
	Config            *gmodels.ConfigSettings
	Context           *core.Context
	MasterTaskContext *taskqueue.TaskContext
	ProdTaskContext   *taskqueue.TaskContext
}