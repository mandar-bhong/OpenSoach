package main

import (
	"fmt"

	"log"
	"sync"
	"time"

	"github.com/bitly/go-nsq"
	"opensoach.com/utility/logger"
	webServer "opensoach.com/webserver"
)

func main() {

	logger.Init()

	webServer.Init()
	//CreateProducer()

	//CreatCunsumer()

	//doEvery(4*time.Second, CreateProducer)

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
