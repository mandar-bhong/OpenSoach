package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
)

//var Instance *logrus.Logger

type Instance struct {
}

func Init(filename string, maxsize int, maxbackups int, maxage int, loglevel string) {

	var level logrus.Level
	switch loglevel {
	case "DEBUG":
		level = logrus.DebugLevel
		break
	case "ERROR":
		level = logrus.ErrorLevel
		break
	case "WARN":
		level = logrus.WarnLevel
		break
	}

	instance := logrus.New()

	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   filename,
		MaxSize:    maxsize,
		MaxBackups: maxbackups,
		MaxAge:     maxage,
		Level:      level,
		//		Formatter: &logrus.JSONFormatter{
		//			TimestampFormat: "2006-01-02 15:04:05",
		//		},
	})

	if err != nil {
		fmt.Println("Error occured")
		fmt.Println(err.Error())
		return
	}

	instance.AddHook(rotateFileHook)

	//	Instance.Formatter = &logrus.JSONFormatter{
	//		//TimestampFormat: "2006-01-02 15:04:05",

	//	}

	instance.Out = os.Stdout

	//	Instance = instance

	//writeLog(0)
	//writeLog(500)
}

func Debug(msg string, args ...interface{}) {

}

func Error(msg string, args ...interface{}) {

}

//func writeLog(startpt int) {
//	for i := startpt; i < 100; i++ {
//		fmt.Println(time.Now())
//		Instance.Error("this is logger test", i)
//		Instance.WithField("Count", i).WithField("Count1", i).Error("this is logger test", i)
//		Instance.WithField("Mount", i).Warn("this is logger test", i)
//		Instance.WithField("Xount", i).Debug("this is logger test", i)
//		time.Sleep(time.Millisecond * 100)
//	}
//}
