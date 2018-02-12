package main

import (
	"fmt"
	"os"

	//"opensoach.com/manager/dbmanager"
	//msgbkr "opensoach.com/manager/messagebroker"
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

	loggerInitErr := logger.Init(config.LoggerConfig.Filename, config.LoggerConfig.MaxSize, config.LoggerConfig.MaxBackups, config.LoggerConfig.MaxAge, config.LoggerConfig.Level)

	if loggerInitErr != nil {
		fmt.Println(loggerInitErr.Error())
		ShutDown(51)
		return
	}

	isModSuccess, code := InitModules(config)

	if !isModSuccess {
		fmt.Printf("Error occured while initializing modules %#v \n", code)
	}

	//filename string, maxsize int, maxbackups int, maxage int, loglevel string

	//logger.Instance.Error("Starting Server")

	webServer.Init()

	//go	msgbkr.CreatCunsumer()

	//	msgbkr.CreateProducer()

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

func InitModules(config *gmodels.ConfigSettings) (bool, int) {

	return true, 0
}

func ShutDown(errorCode int) {
	os.Exit(errorCode)
}
