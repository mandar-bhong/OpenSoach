package main

import (
	"os"
	"time"

	"fmt"

	ghelper "opensoach.com/core/helper"
	gmodels "opensoach.com/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	dbmgr "opensoach.com/core/manager/db"

	"opensoach.com/core"
	"opensoach.com/core/logger"
	coremodels "opensoach.com/core/models"
)

func main() {

	isSuccess, dbconfig := ReadConfiguration()

	if !isSuccess {
		ShutDown(100)
		return
	}

	ReadConfiguration()

	err, configData := GetConfiguration(dbconfig)

	if err != nil {
		fmt.Println("Error occured while fetching configuration data: ", err.Error())
	}

	InitModules(configData)

	logger.Context().LogDebug("Main", "Starting Application")

	time.Sleep(time.Second * 5)
}

func ReadConfiguration() (bool, *gmodels.ConfigDB) {

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

	return true, &settings.DBConfig
}

func GetConfiguration(config *gmodels.ConfigDB) (error, *[]gmodels.DBMasterConfigRowModel) {

	dbEngine, err := sqlx.Connect(config.DBDriver, config.ConnectionString)

	configRows := &[]gmodels.DBMasterConfigRowModel{}

	if err != nil {
		fmt.Printf("DB Error %#+v \n", err.Error())
		return err, nil
	}

	selCtx := dbmgr.SelectProcContext{}
	selCtx.Engine = dbEngine
	selCtx.SPName = "sp_mst_get_configuration"
	selCtx.Dest = configRows

	selErr := selCtx.Select()

	if selErr != nil {
		fmt.Printf("DB Error %#+v \n", selErr.Error())
		return selErr, nil
	}

	return nil, configRows
}

func InitModules(configData *[]gmodels.DBMasterConfigRowModel) {
	InitCoreModule(configData)
}

func InitCoreModule(configData *[]gmodels.DBMasterConfigRowModel) {

	logger.Init()
	logger.SetModule("SPL")
	logger.SetLogLevel(logger.Debug)
	logger.SetLoggingService(logger.LoggingServiceFmt)

	coreConfig := &coremodels.CoreConfig{}
	core.Init(coreConfig)
}

func ShutDown(errorCode int) {
	os.Exit(errorCode)
}
