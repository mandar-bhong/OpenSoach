package manager

import (
	"errors"
	"fmt"
	"strconv"

	"opensoach.com/core/logger"
	"opensoach.com/hpft/constants/dbquery"

	"opensoach.com/core"
	taskque "opensoach.com/core/manager/taskqueue"
	repo "opensoach.com/hpft/server/repository"
	gmodels "opensoach.com/models"

	ghelper "opensoach.com/core/helper"
	pchelper "opensoach.com/prodcore/helper"
	pcmgr "opensoach.com/prodcore/manager"
)

var SUB_MODULE_NAME = "HPFT.Endpoint.Manager"

func InitilizeModues(dbconfig *gmodels.ConfigDB) error {

	dbConnnErr := pchelper.VerifyDBConnection(dbconfig)

	if dbConnnErr != nil {
		fmt.Println("Unable to connect database: ", dbConnnErr.Error())
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

	errPrepareConfig, masterConfigSetting := pcmgr.PrepareMasterConfiguration(dbconfig, masterConfigData, gmodels.PRODUCT_TYPE_HPFT)

	if errPrepareConfig != nil {
		fmt.Println("Error occured while fetching configuration data: ", errPrepareConfig.Error())
		logger.Context().WithField("DbConn ", dbconfig.ConnectionString).
			LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while fetching configuration data.", errPrepareConfig)
		return errPrepareConfig
	}

	dbConnnErr = pchelper.VerifyDBConnection(masterConfigSetting.ProdMstDBConfig)

	if dbConnnErr != nil {
		fmt.Println("Error occured while fetching configuration data: ", dbConnnErr.Error())
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while fetching configuration data.", dbConnnErr)
		return dbConnnErr
	}

	prodConfigErr, prodConfigData := pchelper.GetConfiguration(masterConfigSetting.ProdMstDBConfig, dbquery.QUERY_GET_PRODUCT_MASTER_CONFIGURATION)

	if prodConfigErr != nil {
		fmt.Println("Error occured while fetching configuration data: ", prodConfigErr.Error())
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while fetching configuration data.", prodConfigErr)
		return err
	}

	pcmgr.UpdateProductConfiguration(masterConfigSetting, prodConfigData)

	SetGlobal(dbconfig, masterConfigSetting)

	initErr := pcmgr.InitModules(masterConfigSetting)

	if initErr != nil {
		fmt.Println("Error occured while fetching configuration data: ", initErr.Error())
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while fetching configuration data.", initErr)
		return initErr
	}

	pcmgr.SetLogger(masterConfigSetting)
	logger.SetModule("HPFT.Server")

	InitProcessor()

	initModules(masterConfigSetting)

	return nil
}

func initModules(configSetting *gmodels.ConfigSettings) error {

	//	broker: 'redis://localhost:6379'
	//default_queue: machinery_tasks
	//result_backend: 'redis://127.0.0.1:6379'
	//results_expire_in: 3600000

	// configSetting.MasterQueCache.Address = "localhost"
	// configSetting.MasterQueCache.Port = 6379
	// configSetting.ProductCache.Address = "localhost"
	// configSetting.ProductCache.Port = 6379

	mstTaskConfig := taskque.TaskConfig{}
	mstTaskConfig.Broker = "redis://" + configSetting.MasterQueCache.Address + ":" + strconv.Itoa(configSetting.MasterQueCache.Port)
	mstTaskConfig.ResultBackend = "redis://" + configSetting.MasterQueCache.Address + ":" + strconv.Itoa(configSetting.MasterQueCache.Port)
	mstTaskConfig.DefaultQueue = gmodels.SPL_SERVER_DEFAULT_TASK_QUEUE
	mstTaskConfig.ResultsExpireIn = 1 // in min

	mstTaskCtx := &taskque.TaskContext{}

	mstTaskQueErr := mstTaskCtx.CreateServer(mstTaskConfig)

	if mstTaskQueErr != nil {
		return mstTaskQueErr
	}

	prodTaskConfig := taskque.TaskConfig{}
	prodTaskConfig.Broker = "redis://" + configSetting.ProductCache.Address + ":" + strconv.Itoa(configSetting.MasterQueCache.Port)
	prodTaskConfig.ResultBackend = "redis://" + configSetting.ProductCache.Address + ":" + strconv.Itoa(configSetting.MasterQueCache.Port)
	prodTaskConfig.DefaultQueue = gmodels.HPFT_SERVER_DEFAULT_TASK_QUEUE
	prodTaskConfig.ResultsExpireIn = 1 // in min

	prodTaskCtx := &taskque.TaskContext{}

	prodTaskQueErr := prodTaskCtx.CreateServer(prodTaskConfig)

	if prodTaskQueErr != nil {
		return prodTaskQueErr
	}

	var hkthandler map[string]interface{}
	hkthandler = make(map[string]interface{})

	RegisterHandler(hkthandler)

	prodTaskCtx.RegisterTaskHandlers(hkthandler)

	repo.Instance().MasterTaskContext = mstTaskCtx
	repo.Instance().ProdTaskContext = prodTaskCtx

	go func() {
		err := prodTaskCtx.StartWorker("HKT_SERVER_WORKER")

		if err != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while starting worker", err)
		}
	}()

	return nil
}

func SetGlobal(dbconfig *gmodels.ConfigDB, configSetting *gmodels.ConfigSettings) error {

	//ghelper.BaseDir = configSetting.ServerConfig.BaseDir

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

	repo.Init()
	repo.Instance().Context = ctx
	repo.Instance().Config = configSetting

	return nil

}
