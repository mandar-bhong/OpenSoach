package manager

import (
	"fmt"

	"errors"

	"opensoach.com/core"
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	coremodels "opensoach.com/core/models"
	repo "opensoach.com/hkt/api/repository"
	"opensoach.com/hkt/api/webserver"
	"opensoach.com/hkt/constants/dbquery"
	gmodels "opensoach.com/models"
	pchelper "opensoach.com/prodcore/helper"
	pcmgr "opensoach.com/prodcore/manager"
	taskque "opensoach.com/core/manager/taskqueue"
	"strconv"

)

func InitilizeModues(dbconfig *gmodels.ConfigDB) error {

	dbConnnErr := pchelper.VerifyDBConnection(dbconfig)

	if dbConnnErr != nil {
		logger.Context().WithField("DbConn ", dbconfig.ConnectionString).
			LogError(SUB_MODULE_NAME, logger.Normal, "Unable to connect database.", dbConnnErr)
		return dbConnnErr
	}

	err, masterConfigData := pchelper.GetMasterConfiguration(dbconfig)

	if err != nil {
		fmt.Println("Error occured while fetching configuration data: ", err.Error())
		logger.Context().WithField("Error ", err.Error()).
			LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while fetching configuration data.", err)
		return err
	}

	errPrepareConfig, masterConfigSetting := pcmgr.PrepareMasterConfiguration(dbconfig, masterConfigData, gmodels.PRODUCT_TYPE_HKT)

	if errPrepareConfig != nil {
		fmt.Println("Error occured while fetching configuration data: ", err.Error())
		logger.Context().WithField("DbConn ", dbconfig.ConnectionString).
			LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while fetching configuration data.", errPrepareConfig)
		return err
	}

	dbConnnErr = pchelper.VerifyDBConnection(masterConfigSetting.ProdMstDBConfig)

	if dbConnnErr != nil {
		fmt.Println("Error occured while fetching configuration data: ", dbConnnErr.Error())
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while fetching configuration data.", dbConnnErr)
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
		case 2: //Redis server connection error
			return verifyErr
		case 3: //Redis server master que cache error
			return verifyErr
		case 4: //Redis server product cache error
			return verifyErr
		}
	}

	setGlobalErr := SetGlobal(dbconfig, masterConfigSetting)

	if setGlobalErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Server, "Error occured while setting global values.", setGlobalErr)
		return setGlobalErr
	}

	initErr := initModules(masterConfigSetting)

	if initErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Server, "Initiliazation module error occured", initErr)
		return initErr
	}

	return nil
}

func SetGlobal(dbconfig *gmodels.ConfigDB, configSetting *gmodels.ConfigSettings) error {

	ghelper.BaseDir = configSetting.ServerConfig.BaseDir

	isJsonConvMstCacheSuccess, jsonMstCacheRedisAddress := ghelper.ConvertToJSON(configSetting.MasterCache)

	if isJsonConvMstCacheSuccess == false {
		logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while converting ConfigCacheAddress structure to JSON")
		return errors.New("Error occured while converting ConfigCacheAddress to json")
	}

	isJsonConvProdCacheSuccess, jsonProdCacheRedisAddress := ghelper.ConvertToJSON(configSetting.ProductCache)

	if isJsonConvProdCacheSuccess == false {
		logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while converting ConfigCacheAddress structure to JSON")
		return errors.New("Error occured while converting ConfigCacheAddress to json")
	}

	ctx := &core.Context{}
	ctx.Master.DBConn = configSetting.DBConfig.ConnectionString
	ctx.Master.Cache.CacheAddress = jsonMstCacheRedisAddress

	ctx.ProdMst.Cache.CacheAddress = jsonProdCacheRedisAddress
	ctx.ProdMst.DBConn = configSetting.ProdMstDBConfig.ConnectionString

	repo.Init(configSetting, ctx)

	return nil

}

func initModules(configSetting *gmodels.ConfigSettings) error {

	logger.Init()
	logger.SetModule("HKT")
	logger.SetLogLevel(logger.Debug)
	logger.SetLoggingService(logger.LoggingServiceFmt)

	coreConfig := &coremodels.CoreConfig{}
	err := core.Init(coreConfig)

	if err != nil {
		return err
	}


	prodTaskConfig := taskque.TaskConfig{}
	prodTaskConfig.Broker = "redis://" + configSetting.ProductCache.Address + ":" + strconv.Itoa(configSetting.MasterQueCache.Port)
	prodTaskConfig.ResultBackend = "redis://" + configSetting.ProductCache.Address + ":" + strconv.Itoa(configSetting.MasterQueCache.Port)
	prodTaskConfig.DefaultQueue = gmodels.HKT_SERVER_DEFAULT_TASK_QUEUE
	prodTaskConfig.ResultsExpireIn = 1 // in min

	prodTaskCtx := &taskque.TaskContext{}

	prodTaskQueErr := prodTaskCtx.CreateServer(prodTaskConfig)

	if prodTaskQueErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME,logger.Normal,"Error occured while creating queue server",prodTaskQueErr)
		return prodTaskQueErr
	}

	repo.Instance().ProdTaskContext = prodTaskCtx

	webServerInitError := webserver.Init(configSetting)

	if webServerInitError != nil {
		return webServerInitError
	}

	return nil
}
