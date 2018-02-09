package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/bitly/go-nsq"
	gmodels "opensoach.com/models"
	ghelper "opensoach.com/utility/helper"
	"opensoach.com/utility/logger"
	webServer "opensoach.com/webserver"
)

func main() {

	isSuccess, config := ReadConfiguration()

	if !isSuccess {
		fmt.Println("Unable to read configuration")
		ShutDown(50)
		return
	}

	//filename string, maxsize int, maxbackups int, maxage int, loglevel string

	logger.Init(config.LoggerConfig.Filename, config.LoggerConfig.MaxSize, config.LoggerConfig.MaxBackups, config.LoggerConfig.MaxAge, config.LoggerConfig.Level)

	logger.Instance.Error("Starting Server")

	logger.Instance.Error("Starting Server")

	webServer.Init()

	time.Sleep(time.Second * 5)

	//CreateProducer()

	//CreatCunsumer()

	//doEvery(4*time.Second, CreateProducer)

}

func ReadConfiguration() (bool, *gmodels.ConfigSettings) {

	currentPath := ghelper.GetExeFolder()

	isReadSuccess, readContent, errorMsg := ghelper.ReadFileContent(currentPath, "settings", "win.config.json")
	//ioutil.ReadFile()

	if !isReadSuccess {
		fmt.Printf("Unable to configuration file. Error: %s", errorMsg)
		return false, nil
	}

	settings := gmodels.ConfigSettings{}
	isJSONConvertSuccess := ghelper.ConvertFromJSONBytes(readContent, &settings)

	if !isJSONConvertSuccess {
		return false, nil
	}

	fmt.Println(string(readContent))

	return isReadSuccess, &settings
}

func ShutDown(errorCode int) {
	os.Exit(errorCode)
}

func CreatCunsumer() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("write_test", "ch", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message: %v \n", string(message.Body))

		fmt.Printf("message.NSQDAddress %s \n", message.NSQDAddress)
		//wg.Done()
		return nil
	}))
	err := q.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect")
	}
	wg.Wait()
}

func CreateProducer() {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	err := w.Publish("write_test", []byte("test"))
	if err != nil {
		log.Panic("Could not connect")
	}

	//SendMsgdoEvery(4*time.Second, CreateProducer)

	//w.Stop()
}

func doEvery(d time.Duration, f func()) {

	for t := range time.Tick(d) {
		fmt.Println("Tick at", t)

		f()
	}
}

//func SendMsgdoEvery(d time.Duration, nsqprod *nsq.Producer, msg []byte) {
//	for t := range time.Tick(d) {

//		nsqprod.Publish("write_test", msg)
//	}
//}
