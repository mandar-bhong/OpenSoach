package logger

import (
	"fmt"
)

type InfluxDBLoggingService struct {
}

func (InfluxDBLoggingService) prepareLogMessage(l *loggerContext) string {
	return prepareLogMessage(l)
}

func (InfluxDBLoggingService) logMessage(l *loggerContext) {
	msg := prepareLogMessage(l)
	fmt.Println(msg)
}
