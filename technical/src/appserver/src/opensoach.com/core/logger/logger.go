package logger

import (
	"time"
)

type serviceType int
type severity int
type logmsgType int

type keyValueFields struct {
	Key   string
	Value interface{}
}

type loggerContext struct {
	fields       []keyValueFields
	level        severity
	logTime      time.Time
	msg          string
	msgType      logmsgType
	err          error
	appComponent string
	subComponent string
}

const (
	Error severity = 1
	Debug          = 2
	Info           = 3
)

const (
	Normal          logmsgType = 0
	Instrumentation            = 1
	Performace                 = 2
	Server                     = 3
)

const (
	LoggingServiceFmt      serviceType = 1
	LoggingServiceFile                 = 2
	LoggingServiceFluent               = 3
	LoggingServiceInfluxDB             = 4
)

var logLevel severity = Error
var appComponent string
var loggingServiceType serviceType

func init() {
	loggingServiceType = LoggingServiceFmt
}

func Init() {
	initDispatcher()
}

func Context() *loggerContext {
	return &loggerContext{}

}

func SetModule(module string) {
	appComponent = module
}

func SetLogLevel(level severity) {
	logLevel = level
}

func SetLoggingService(serType serviceType) {
	loggingServiceType = serType
}

func GetLogLevel() severity {
	return logLevel
}

func (ctx *loggerContext) Log(subcomp string, logtype logmsgType, logginglevel severity, message string) {
	ctx.logTime = time.Now()
	ctx.appComponent = appComponent
	ctx.subComponent = subcomp
	ctx.level = logginglevel
	ctx.msgType = logtype
	ctx.msg = message
	verifyAndPush(ctx)
}

func (ctx *loggerContext) LogDebug(subcomp string, logtype logmsgType, message string) {
	ctx.logTime = time.Now()
	ctx.appComponent = appComponent
	ctx.subComponent = subcomp
	ctx.msg = message
	ctx.level = Debug
	ctx.msgType = logtype
	verifyAndPush(ctx)
}

func (ctx *loggerContext) LogError(subcomp string, logtype logmsgType, message string, err error) {
	ctx.logTime = time.Now()
	ctx.appComponent = appComponent
	ctx.subComponent = subcomp
	ctx.msg = message
	ctx.err = err
	ctx.level = Error
	ctx.msgType = logtype
	verifyAndPush(ctx)
}

func (ctx *loggerContext) WithField(field string, value interface{}) *loggerContext {
	ctx.fields = append(ctx.fields, keyValueFields{Key: field, Value: value})
	return ctx
}

func verifyAndPush(ctx *loggerContext) {
	switch logLevel {
	case Info:
		chanbuffLog <- ctx
	case Debug:
		switch ctx.level {
		case Info:
		default:
			chanbuffLog <- ctx
		}
	case Error:
		switch ctx.level {
		case Info, Debug:
		default:
			chanbuffLog <- ctx
		}
	}
}
