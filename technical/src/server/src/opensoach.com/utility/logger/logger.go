package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
)

var Instance *logrus.Logger

func Init() {

	//logrus.SetFormatter(&logrus.JSONFormatter{})

	Instance := logrus.New()

	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   "logfile.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     7,
		Level:      logrus.DebugLevel,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		},
	})

	if err != nil {
		fmt.Println("Error occured")
		fmt.Println(err.Error())
		return
	}

	Instance.AddHook(rotateFileHook)

	//	Instance.Formatter = &logrus.JSONFormatter{
	//		//TimestampFormat: "2006-01-02 15:04:05",

	//	}

	Instance.Out = os.Stdout

	//	for i := 0; ; i++ {

	//		fmt.Println(time.Now())
	//		Instance.WithField("Count", i).WithField("Count1", i).Error("this is logger test", i)
	//		Instance.WithField("Mount", i).Warn("this is logger test", i)
	//		Instance.WithField("Xount", i).Debug("this is logger test", i)
	//		time.Sleep(time.Microsecond * 100)
	//	}

}
