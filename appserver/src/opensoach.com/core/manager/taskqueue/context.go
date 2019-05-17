package taskqueue

import (
	"github.com/RichardKnop/machinery/v1"
)

type TaskContext struct {
	Server *machinery.Server
}

type TaskConfig struct {
	//	broker: 'redis://localhost:6379'
	//default_queue: machinery_tasks
	//result_backend: 'redis://127.0.0.1:6379'
	//results_expire_in: 3600000

	Broker          string
	DefaultQueue    string
	ResultBackend   string
	ResultsExpireIn int
}
