package logger

import (
	"bytes"
	"encoding/json"
	"strings"

	"fmt"
	"net/http"
)

var fluentHost string

type FluentLoggingService struct {
}

type logMsg struct {
	AppComponent string `json:"AppComponent"`
	SubComponent string `json:"SubComponent"`
	Fields       string `json:"Fields"`
	Level        string `json:"LogLevel"`
	LogTime      string `json:"Time"`
	Msg          string `json:"Message"`
	MsgType      string `json:"MessageType"`
	Err          string `json:"Error"`
}

func SetFluentHost(host string) {
	fluentHost = strings.TrimRight(host, "/")
}

func (FluentLoggingService) prepareLogMessage(l *loggerContext) string {
	lmsg := convertToLogMsg(*l)
	dataBytes, _ := json.Marshal(lmsg)
	return string(dataBytes)
}

func (r FluentLoggingService) logMessage(l *loggerContext) {
	msg := r.prepareLogMessage(l)
	postMessage(fluentHost, []byte(msg))
}

func convertToJSON(dataStruct interface{}) (bool, []byte) {
	dataBytes, err := json.Marshal(dataStruct)

	if err != nil {
		fmt.Printf("Error occured while converting json data. Error: %s", err.Error())
		return false, []byte{}
	}

	return true, dataBytes
}

func postMessage(host string, jsonData []byte) {

	resp, err := http.Post(host+"/spl.spl", "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Printf("Post Error: %#v, Response: %#v \n", err.Error(), resp)
		return
	}
}

func convertToLogMsg(l loggerContext) logMsg {
	lmsg := logMsg{}

	lmsg.AppComponent = l.appComponent
	lmsg.LogTime = l.logTime.Format("Jan 2, 2006 at 3:04:05pm (MST)")
	lmsg.Msg = l.msg
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
