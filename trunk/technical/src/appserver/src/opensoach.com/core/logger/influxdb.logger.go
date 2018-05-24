package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	client "github.com/influxdata/influxdb/client/v2"
)

const (
	MyDB = "influxd"
)

var influxDBHost string

type InfluxDBLoggingService struct {
}

func SetInfluxDBHost(host string) {
	influxDBHost = strings.TrimRight(host, "/")
}

func (i InfluxDBLoggingService) logMessage(l *loggerContext) {
	msg := i.influxDBprepareLogMessage(l)

	influxDBpostMessage(influxDBHost, msg)
}

func (InfluxDBLoggingService) influxDBprepareLogMessage(l *loggerContext) *client.Point {

	tags := map[string]string{"": ""}

	fields := map[string]interface{}{}

	fields["AppComponent"] = l.appComponent
	fields["SubComponent"] = l.subComponent
	fields["Message"] = l.msg
	fields["Time"] = l.logTime.Format("Jan 2, 2006 at 3:04:05pm (MST)")

	switch l.level {
	case Error:
		fields["LogLevel"] = "Error"
	case Debug:
		fields["LogLevel"] = "Debug"
	case Info:
		fields["LogLevel"] = "Info"
	}

	switch l.msgType {
	case Normal:
		fields["MessageType"] = "Normal"
	case Instrumentation:
		fields["MessageType"] = "Instrumentation"
	case Performace:
		fields["MessageType"] = "Performace"
	case Server:
		fields["MessageType"] = "Server"
	default:
		fields["MessageType"] = "Unknown"
	}

	if l.err != nil {
		fields["Error"] = l.err.Error()
	} else {
		fields["Error"] = ""
	}

	if l.fields == nil {
		fields["Fields"] = ""
	} else {
		dataBytes, _ := json.Marshal(l.fields)
		fields["Fields"] = string(dataBytes)
	}

	pt, err := client.NewPoint("spl.spl", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}

	return pt

}

func influxDBpostMessage(host string, pt *client.Point) {

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: host,
	})
	if err != nil {
		fmt.Printf("Error occured while creating new endpoint client: %#v \n", err)
	}
	defer c.Close()

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})
	if err != nil {
		fmt.Printf("Error occured while creating new batch points %#v \n", err)
	}

	bp.AddPoint(pt)

	if err := c.Write(bp); err != nil {
		fmt.Printf("Error occured while writing to db: %#v \n", err)
	}

	if err := c.Close(); err != nil {
		fmt.Printf("Error occured while closing connection:  %#v \n", err)
	}

}
