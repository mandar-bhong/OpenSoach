package logger

import (
	"fmt"
	"log"
	"strings"
	"time"

	client "github.com/influxdata/influxdb/client/v2"
)

const (
	loginApplicationInfluxDB = "spl"
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

	lmsg := convertToLogMsg(*l)

	tags := map[string]string{"": ""}
	fields := map[string]interface{}{}

	fields["AppComponent"] = lmsg.AppComponent
	fields["SubComponent"] = lmsg.SubComponent
	fields["Message"] = lmsg.Msg
	fields["Time"] = lmsg.LogTime
	fields["LogLevel"] = lmsg.Level
	fields["MessageType"] = lmsg.MsgType
	fields["Error"] = lmsg.Err
	fields["Fields"] = lmsg.Fields

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
		Database:  loginApplicationInfluxDB,
		Precision: "ns", //precision set nano seconds because other than this packets are not logged appropriately few packets are discarded
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
