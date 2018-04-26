package manager

import (
	"fmt"

	"errors"

	"opensoach.com/core"
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	coremodels "opensoach.com/core/models"
	"opensoach.com/hkt/api/constants/dbquery"
	repo "opensoach.com/hkt/api/repository"
	"opensoach.com/hkt/api/webserver"
	gmodels "opensoach.com/models"
	pchelper "opensoach.com/prodcore/helper"
	pcmgr "opensoach.com/prodcore/manager"
)

func InitilizeModues(dbconfig *gmodels.ConfigDB) error {

	dbConnnErr := pchelper.VerifyDBConnection(dbconfig)

	if dbConnnErr != nil {
		//TODO: LOG
		return dbConnnErr
	}

	err, masterConfigData := pchelper.GetMasterConfiguration(dbconfig)

	if err != nil {
		fmt.Println("Error occured while fetching configuration data: ", err.Error())
		return err
	}

	errPrepareConfig, masterConfigSetting := pcmgr.PrepareMasterConfiguration(dbconfig, masterConfigData, gmodels.PRODUCT_TYPE_HKT)

	if errPrepareConfig != nil {
		fmt.Println("Error occured while fetching configuration data: ", err.Error())
		return err
	}

	dbConnnErr = pchelper.VerifyDBConnection(masterConfigSetting.ProdMstDBConfig)

	if dbConnnErr != nil {
		//TODO: LOG
		return dbConnnErr
	}

	prodConfigErr, prodConfigData := pchelper.GetConfiguration(masterConfigSetting.ProdMstDBConfig, dbquery.QUERY_GET_PRODUCT_MASTER_CONFIGURATION)

	if prodConfigErr != nil {
		fmt.Println("Error occured while fetching configuration data: ", err.Error())
		return err
	}

	prodUpdateConfigErr := pcmgr.UpdateProductConfiguration(masterConfigSetting, prodConfigData)

	if prodUpdateConfigErr != nil {
		//TODO: log message, need to identify fmt or file base or both
		return errPrepareConfig
	}

	//init logger for fmt or file or both for temp then switch mode as per configuration after component connection verification

	verifyErr, moduleType := pcmgr.VerifyConnection(dbconfig, masterConfigSetting)

	if verifyErr != nil {

		switch moduleType {
		case 1:
		case 2:
		}
	}

	connErr := SetGlobal(dbconfig, masterConfigSetting)

	if connErr != nil {
		//TODO: log message, need to identify fmt or file base or both
		return connErr
	}

	initModules(masterConfigSetting)

	return nil
}

func SetGlobal(dbconfig *gmodels.ConfigDB, configSetting *gmodels.ConfigSettings) error {

	ghelper.BaseDir = configSetting.ServerConfig.BaseDir

	isJsonConvertionSuccess, jsonRedisAddress := ghelper.ConvertToJSON(configSetting.MasterCache)

	if isJsonConvertionSuccess == false {
		logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while converting ConfigCacheAddress structure to JSON")
		return errors.New("Error occured while converting ConfigCacheAddress to json")
	}

	ctx := &core.Context{}
	ctx.Master.Cache.CacheAddress = jsonRedisAddress
	ctx.Master.DBConn = configSetting.DBConfig.ConnectionString

	repo.Init(configSetting, ctx)

	return nil

}

func initModules(configSetting *gmodels.ConfigSettings) error {

	logger.Init()
	logger.SetModule("SPL")
	logger.SetLogLevel(logger.Debug)
	logger.SetLoggingService(logger.LoggingServiceFmt)

	coreConfig := &coremodels.CoreConfig{}
	err := core.Init(coreConfig)

	if err != nil {
		return err
	}

	webServerInitError := webserver.Init(configSetting)

	if webServerInitError != nil {
		return webServerInitError
	}

	return nil
}
