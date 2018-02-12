package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/bitly/go-nsq"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"opensoach.com/manager/dbmanager"
	msgbkr "opensoach.com/manager/messagebroker"
	gmodels "opensoach.com/models"
	ghelper "opensoach.com/utility/helper"
	"opensoach.com/utility/logger"
	webServer "opensoach.com/webserver"
)

func main() {

	//DBTest()

	isSuccess, config := ReadConfiguration()

	if !isSuccess {
		fmt.Println("Unable to read configuration")
		ShutDown(50)
		return
	}

	//filename string, maxsize int, maxbackups int, maxage int, loglevel string

	logger.Init(config.LoggerConfig.Filename, config.LoggerConfig.MaxSize, config.LoggerConfig.MaxBackups, config.LoggerConfig.MaxAge, config.LoggerConfig.Level)

	//logger.Instance.Error("Starting Server")

	webServer.Init()

	//time.Sleep(time.Second * 5)
	//	fmt.Println(time.Local)
	//	fmt.Println(time.Now())
	//	location, err := time.LoadLocation("Asia/Kolkata")

	//	if err != nil {
	//		fmt.Println(err.Error())
	//	}
	//	time.Local = location
	//	fmt.Println("Changed")
	//	fmt.Println(time.Now())
	//	fmt.Println(time.Local)

	//go	msgbkr.CreatCunsumer()

	//	msgbkr.CreateProducer()

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

var engine *sqlx.DB

func DBTest() {

	//hlist := []HCodeStruct{}

	engine = sqlx.MustConnect("mysql", "root:welcome@tcp(localhost:3306)/hktdb?parseTime=true")

	InsertTest()

	//UpdateTest()

	return

	tx := engine.MustBegin()

	//	selerr := engine.Select(&hlist, "call proc_get_hcode(3333)")

	//	if selerr != nil {
	//		fmt.Printf("Select error : %#v \n", selerr)
	//		return
	//	}

	hcodeUpdate := HCodeStruct{}

	hcodeUpdate.CustomerId = 1
	hcodeUpdate.Hcode = "HCode"
	hcodeUpdate.HcodeDesc = "Desc"
	hcodeUpdate.HcodeName = "Name"

	insertresult, inserterr := tx.NamedExec("call proc_insert_hcode(:customer_id_fk,:hcode,:hcode_name,:hcode_desc); ", hcodeUpdate)

	//insertresult, inserterr := engine.Exec("insert into hcode_tbl1 (customer_id_fk,hcode,hcode_name,hcode_desc) values (:customer_id_fk,:hcode,:hcode_name,:hcode_desc)", hcodeUpdate)

	if inserterr != nil {
		fmt.Printf("Insert error : %#v\n", inserterr)
		return
	}

	go InsertSecond()
	time.Sleep(time.Second * 3)

	var lastinsertid int

	tx.Get(&lastinsertid, "SELECT Last_Insert_ID()")

	fmt.Printf("Lst Insert id error : %#v\n", lastinsertid)

	insertid, errinsid := insertresult.LastInsertId()

	if errinsid != nil {
		fmt.Printf("Insert id error : %#v\n", errinsid)
	}

	fmt.Printf("Inserted id %#v", insertid)

	tx.Commit()

}

type HCodeStruct struct {
	ID         int    `db:"id" json:"customerid"`
	CustomerId int    `db:"customer_id_fk" json:"customerid"`
	Hcode      string `db:"hcode" json:"hcode"`
	HcodeName  string `db:"hcode_name" json:"hcodename"`
	HcodeDesc  string `db:"hcode_desc" json:"hcodedesc"`
}

func InsertTest() {
	hcodeUpdate := HCodeStruct{}

	hcodeUpdate.CustomerId = 1
	hcodeUpdate.Hcode = "abc2"
	hcodeUpdate.HcodeDesc = "Desc2"
	hcodeUpdate.HcodeName = "Name2"

	insertSPCtx := dbmanager.InsertProcContext{}
	insertSPCtx.Engine = engine
	insertSPCtx.SPName = "proc_insert_hcode"
	insertSPCtx.SPArgs = hcodeUpdate

	insertMgrErr := insertSPCtx.Insert()

	if insertMgrErr != nil {
		fmt.Printf("Error occured while insert row: %#v\n", insertMgrErr.Error())
	}

	fmt.Printf("InsertID: ", insertSPCtx.InsertID)

}

func UpdateTest() {

	hcodeUpdate := HCodeStruct{}
	//hcodeUpdate.CustomerId = 2
	hcodeUpdate.Hcode = "abc2"
	hcodeUpdate.HcodeName = "This is test description"

	spctx := StoredProcContext{}
	spctx.Engine = engine
	spctx.SPName = "proc_update_hcode"
	spctx.SPArgs = hcodeUpdate

	//spctx.Update()

	fmt.Printf("Affected Rows are : %#v\n", spctx.AffectedRows)

	selspctx := SelectProcContext{}
	selspctx.Dest = &[]HCodeStruct{}
	selspctx.SPName = "procc_get_hcode_by_id"
	selspctx.SPArgs = hcodeUpdate
	selspctx.Engine = engine

	selerr := selspctx.Select("90")

	if selerr != nil {
		fmt.Println(selerr.Error())
		return
	}

	fmt.Printf("Result: %#v \n", selspctx.Dest)

}

func InsertSecond() {

	hcodeUpdate := HCodeStruct{}

	hcodeUpdate.CustomerId = 1
	hcodeUpdate.Hcode = "abc2"
	hcodeUpdate.HcodeDesc = "Desc2"
	hcodeUpdate.HcodeName = "Name2"

	tx := engine.MustBegin()

	tx.NamedExec("call proc_insert_hcode(:customer_id_fk,:hcode,:hcode_name,:hcode_desc); ", hcodeUpdate)

	var lastinsertid int

	tx.Get(&lastinsertid, "SELECT Last_Insert_ID()")

	fmt.Printf("2Lst Insert id error : %#v\n", lastinsertid)

	tx.Commit()

}
