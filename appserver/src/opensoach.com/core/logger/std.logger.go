package logger

import (
	"fmt"
)

type StandardOutLoggingService struct {
}

func (StandardOutLoggingService) prepareLogMessage(l *loggerContext) string {
	return prepareLogMessage(l)
}

func (StandardOutLoggingService) logMessage(l *loggerContext) {
	msg := prepareLogMessage(l)
	fmt.Println(msg)
}
