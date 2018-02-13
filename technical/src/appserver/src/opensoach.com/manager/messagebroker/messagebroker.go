package messagebroker

import (
	"sync"

	"fmt"
	"time"

	"github.com/bitly/go-nsq"
)

var nsqAddress string
var producer *nsq.Producer

func Init(nsqInstanceAdd string) {
	nsqAddress = nsqInstanceAdd

	CreatePublisher(nsqAddress)
}

func CreatePublisher(nsqAddress string) {
	config := nsq.NewConfig()
	prd, producerErr := nsq.NewProducer("127.0.0.1:4150", config)

	if producerErr != nil {
		fmt.Println(producerErr.Error())
	}

	producer = prd
}

func Publish(topic string, msg []byte) error {
	return producer.Publish(topic, msg)
}

func CreatCunsumerChannel() {
}

func CreatCunsumer() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("write_test", "ch", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		fmt.Printf("Got a message: %v \n", string(message.Body))

		fmt.Printf("message.NSQDAddress %s \n", message.NSQDAddress)
		//wg.Done()
		return nil
	}))
	err := q.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		fmt.Println("Could not connect")
	}
	wg.Wait()
}

func CreateProducer() {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	err := w.Publish("write_test", []byte("test"))
	if err != nil {
		//log.Panic("Could not connect")
		fmt.Printf("Could not connect")
	}

	//doEvery(4*time.Second, fun(w){})
	//SendMsgdoEvery(4*time.Second, CreateProducer)

	//w.Stop()

	//	for i := 0; ; i++ {
	//		perr := w.Publish("write_test", []byte("test"))

	//		if perr != nil {
	//			fmt.Println(perr.Error())
	//		}

	//		time.Sleep(time.Second)
	//	}
}

func doEvery(d time.Duration, f func()) {

	for t := range time.Tick(d) {
		fmt.Println("Tick at", t)

		f()
	}
}
