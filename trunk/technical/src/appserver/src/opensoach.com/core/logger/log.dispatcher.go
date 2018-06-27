package logger

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

type LoggingService interface {
	logMessage(*loggerContext)
}

var chanbuffLog chan *loggerContext
var wg sync.WaitGroup

func init() {
	chanbuffLog = make(chan *loggerContext, 1000)
	wg.Add(1)
	go dispatch()
}

func initDispatcher() {

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

	serviceLogger := GetServiceLogger()

	if serviceLogger == nil {
		serviceLogger = StandardOutLoggingService{}
	}

	serviceLogger.logMessage(logMsg)

}

func GetServiceLogger() LoggingService {
	switch loggingServiceType {
	case LoggingServiceFmt:
		return StandardOutLoggingService{}

	case LoggingServiceFluent:
		return FluentLoggingService{}
	case LoggingServiceInfluxDB:
		return InfluxDBLoggingService{}
	}
	return nil
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
		msg = msg + fmt.Sprintf("%+v", logMsg.err.Error()) + "\n"
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

func convertToLogMsg(l loggerContext) logMsg {
	lmsg := logMsg{}

	lmsg.AppComponent = l.appComponent
	lmsg.LogTime = l.logTime.Format("Jan 2, 2006 at 3:04:05pm (MST)")
	lmsg.Msg = l.msg
	lmsg.Msg = strings.Replace(lmsg.Msg, "\n", " ", -1)
	lmsg.SubComponent = l.subComponent

	switch l.level {
	case Error:
		lmsg.Level = "Error"
	case Debug:
		lmsg.Level = "Debug"
	case Info:
		lmsg.Level = "Info"
	}

	switch l.msgType {
	case Normal:
		lmsg.MsgType = "Normal"
	case Instrumentation:
		lmsg.MsgType = "Instrumentation"
	case Performace:
		lmsg.MsgType = "Performace"
	case Server:
		lmsg.MsgType = "Server"
	default:
		lmsg.MsgType = "Unknown"
	}

	if l.err != nil {
		lmsg.Err = l.err.Error()
	} else {
		lmsg.Err = ""
	}

	if l.fields == nil {
		lmsg.Fields = ""
	} else {
		dataBytes, _ := json.Marshal(l.fields)
		lmsg.Fields = string(dataBytes)
	}

	return lmsg

}
