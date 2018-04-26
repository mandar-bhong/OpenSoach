package main

import (
	"fmt"
	"os"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	"opensoach.com/hkt/api"
	gmodels "opensoach.com/models"
)

func main() {

	isSuccess, dbconfig := readConfiguration()

	if !isSuccess {
		shutDown(100)
		return
	}

	api.Init(dbconfig)

}

func readConfiguration() (bool, *gmodels.ConfigDB) {

	currentPath := ghelper.GetExeFolder()

	err, readContent := ghelper.ReadFileContent(currentPath, "settings", "win.config.json")

	if err != nil {
		logger.Context().LogError("", logger.Normal, "Error occured while reading configuration file", err)
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
