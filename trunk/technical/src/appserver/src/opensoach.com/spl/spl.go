package main

import (
	"os"

	"fmt"

	ghelper "opensoach.com/core/helper"
	gmodels "opensoach.com/models"
)

func main() {

	isSuccess, dbconfig := ReadDBConfiguration()

	if !isSuccess {
		ShutDown(100)
		return
	}

	ReadConfiguration(dbconfig)

}

func ReadDBConfiguration() (bool, *gmodels.ConfigDB) {

	currentPath := ghelper.GetExeFolder()

	isReadSuccess, readContent, errorMsg := ghelper.ReadFileContent(currentPath, "settings", "win.config.json")

	if !isReadSuccess {
		fmt.Printf("Unable to configuration file. Error: %s \n", errorMsg)
		return false, nil
	}

	fmt.Println(string(readContent))

	settings := gmodels.ConfigSettings{}
	isJSONConvertSuccess := ghelper.ConvertFromJSONBytes(readContent, &settings)

	if !isJSONConvertSuccess {
		return false, nil
	}

	return isReadSuccess, &settings.DBConfig
}

func ReadConfiguration(config *gmodels.ConfigDB) {

}

func ShutDown(errorCode int) {
	os.Exit(errorCode)
}
