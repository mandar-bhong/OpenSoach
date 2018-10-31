package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	"opensoach.com/hpft/api"
	"opensoach.com/hpft/endpoint"
	"opensoach.com/hpft/server"
	gmodels "opensoach.com/models"
)

func main() {

	isSuccess, config := readConfiguration()

	if !isSuccess {
		shutDown(100)
		return
	}

	if startOk := api.Init(config); startOk == false {
		return
	}

	if startOk := endpoint.Init(config); startOk == false {
		return
	}

	if startOk := server.Init(config); startOk == false {
		return
	}

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	go func() {
		sig := <-gracefulStop
		fmt.Printf("caught sig: %+v", sig)
		DeInit()
		os.Exit(0)
	}()

	select {}
}

func DeInit() {
	api.DeInit()
	endpoint.DeInit()
	server.DeInit()
}

func readConfiguration() (bool, *gmodels.ConfigSettings) {

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

	return true, &settings
}

func shutDown(errorCode int) {
	os.Exit(errorCode)
}
