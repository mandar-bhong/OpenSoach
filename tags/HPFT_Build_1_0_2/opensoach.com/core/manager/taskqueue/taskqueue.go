package taskqueue

import (
	"context"
	"fmt"
	"time"

	"github.com/RichardKnop/machinery/v1/backends/result"

	"github.com/RichardKnop/machinery/v1"
	machineryconfig "github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"

	opentracing "github.com/opentracing/opentracing-go"
	uuid "github.com/satori/go.uuid"
)

var serverTasks map[string]interface{}

func init() {
	serverTasks = make(map[string]interface{})
}

func (ctx *TaskContext) CreateServer(config TaskConfig) error {

	cnf := &machineryconfig.Config{}

	cnf.Broker = config.Broker
	cnf.DefaultQueue = config.DefaultQueue
	cnf.ResultBackend = config.ResultBackend
	cnf.ResultsExpireIn = config.ResultsExpireIn

	cnf.AMQP = &machineryconfig.AMQPConfig{}

	cnf.AMQP.BindingKey = "machinery_task"
	cnf.AMQP.Exchange = "machinery_exchange"
	cnf.AMQP.ExchangeType = "direct"
	cnf.AMQP.PrefetchCount = 3

	//	cnf, err := loadConfig()
	//	if err != nil {
	//		return nil, err
	//	}

	// Create server instance
	server, err := machinery.NewServer(cnf)
	if err != nil {
		return err
	}

	ctx.Server = server

	return nil
}

func (ctx *TaskContext) RegisterTaskHandlers(tasks map[string]interface{}) {

	ctx.Server.RegisterTasks(tasks)
}

func (ctx *TaskContext) StartWorker(consumerTag string) error {
	worker := ctx.Server.NewWorker(consumerTag, 0)
	return worker.Launch()
}

func (ctx *TaskContext) processAsync(taskname string, jsonData string) (error, *result.AsyncResult) {

	command := tasks.Signature{
		Name: taskname,
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: jsonData,
			},
		},
	}

	span, machineryctx := opentracing.StartSpanFromContext(context.Background(), "send")
	defer span.Finish()

	batchUUID, err := uuid.NewV4()

	if err != nil {
		return fmt.Errorf("Error generating batch id: %s", err.Error()), nil
	}

	batchID := batchUUID.String()
	span.SetBaggageItem("batch.id", batchID)

	asyncResult, err := ctx.Server.SendTaskWithContext(machineryctx, &command)

	return err, asyncResult
}

func (ctx *TaskContext) SubmitTask(taskname string, jsonData string) error {
	err, _ := ctx.processAsync(taskname, jsonData)
	return err
}

func (ctx *TaskContext) ProcessTask(taskname string, jsonData string) (error, string) {

	asyncResultErr, asyncResult := ctx.processAsync(taskname, jsonData)

	if asyncResultErr != nil {
		return asyncResultErr, ""
	}

	results, err := asyncResult.Get(time.Duration(time.Millisecond * 5))

	resultData := tasks.HumanReadableResults(results)

	return err, resultData
}
