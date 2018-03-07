package main

import (
	"os"
	"time"

	"fmt"

	ghelper "opensoach.com/core/helper"
	gmodels "opensoach.com/models"

	"opensoach.com/core/logger"
	"opensoach.com/spl/manager"
)

func main() {

	isSuccess, dbconfig := readConfiguration()

	if !isSuccess {
		shutDown(100)
		return
	}

	manager.InitilizeModues(dbconfig)

	logger.Context().LogDebug("Main", "Starting Application")

	time.Sleep(time.Second * 5)
}

func readConfiguration() (bool, *gmodels.ConfigDB) {

	currentPath := ghelper.GetExeFolder()

	err, readContent := ghelper.ReadFileContent(currentPath, "settings", "win.config.json")

	if err != nil {
		logger.Context().LogError("", "Error occured while reading configuration file", err)
		return false, nil
	}

	//fmt.Println(string(readContent))

	settings := gmodels.ConfigSettings{}
	jsonConvertErr := ghelper.ConvertFromJSONBytes(readContent, &settings)

	if jsonConvertErr != nil {
		fmt.Printf("Error occured while converting from JSON %+v\n", err)
		return false, nil
	}

	return true, settings.DBConfig
}

func shutDown(errorCode int) {
	os.Exit(errorCode)
}
