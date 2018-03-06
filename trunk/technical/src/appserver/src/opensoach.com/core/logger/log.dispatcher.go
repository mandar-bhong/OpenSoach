package logger

import (
	"fmt"
	"sync"
)

type LoggingService interface {
	prepareLogMessage(*loggerContext) string
	logMessage(*loggerContext)
}

var chanbuffLog chan *loggerContext
var wg sync.WaitGroup

func initDispatcher() {
	chanbuffLog = make(chan *loggerContext)
	wg.Add(1)
	go dispatch()
}

func dispatch() {
	defer wg.Done()
	for {
		select {
		case msg := <-chanbuffLog:
			logMessage(msg)
		}
	}
}

func logMessage(logMsg *loggerContext) {
	//TODO: As per configuration send msg to respective component, currently printing to screen
	// need to modify preparing and sending messages with unified interface

	switch loggingServiceType {
	case LoggingServiceFmt:
		stdService := StandardOutLoggingService{}
		stdService.logMessage(logMsg)
		break
	case LoggingServiceFile:
		break
	case LoggingServiceFluent:
		break
	default:
		stdService := StandardOutLoggingService{}
		stdService.logMessage(logMsg)
		break
	}
}

func prepareLogMessage(logMsg *loggerContext) string {

	msg := ""

	switch logMsg.level {
	case Debug:
		msg = "Debug:"
		break
	case Error:
		msg = "Error: "
		break
	case Info:
		msg = "Info: "
		break
	default:
		msg = "Verbose: "
		break
	}

	msg = msg + "Time: " + logMsg.logTime.Format("2-Jan-06 03:04:05.000 PM MST") + "\n"

	if logMsg.msg != "" {
		msg = msg + logMsg.msg + "\n"
	}

	if logMsg.err != nil {
		msg = msg + fmt.Sprintf("%+v", logMsg.err) + "\n"
	}

	if len(logMsg.fields) > 0 {
		fieldMsg := "Fields: "
		for _, kvp := range logMsg.fields {
			fieldMsg = fieldMsg + fmt.Sprintf("%s: %+v\n", kvp.Key, kvp.Value)
		}

		msg = msg + fieldMsg
	}

	return msg
}
