package logger

import (
	"bytes"
	"encoding/json"
	"strings"

	"fmt"
	"net/http"
)

var fluentHost string
var fluentDBPoint string

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

	resp, err := http.Post(host+"/"+dbPoint, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Printf("Post Error: %#v, Response: %#v \n", err.Error(), resp)
		return
	}
}
