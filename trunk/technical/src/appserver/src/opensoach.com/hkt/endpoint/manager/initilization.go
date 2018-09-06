package manager

import (
	"errors"
	"fmt"
	"strconv"

	"opensoach.com/core/logger"
	"opensoach.com/hkt/constants/dbquery"

	"opensoach.com/core"
	taskque "opensoach.com/core/manager/taskqueue"
	repo "opensoach.com/hkt/endpoint/repository"
	gmodels "opensoach.com/models"
	//wm "opensoach.com/prodcore/endpoint/webSocketManager"
	ghelper "opensoach.com/core/helper"
	pchelper "opensoach.com/prodcore/helper"
	pcmgr "opensoach.com/prodcore/manager"
)

var SUB_MODULE_NAME = "HKT.Endpoint.Manager"

func InitilizeModues(dbconfig *gmodels.ConfigDB) error {

	dbConnnErr := pchelper.VerifyDBConnection(dbconfig)

	if dbConnnErr != nil {
		//TODO: LOG
		fmt.Printf("Error occured while verifing db conn.DB Conn: %s. Error: %s", dbconfig.ConnectionString, dbConnnErr.Error())
		logger.Context().WithField("", dbconfig.ConnectionString).
			LogError(SUB_MODULE_NAME, logger.Server, "Error occured while verifing db conn", dbConnnErr)
		return dbConnnErr
	}

	err, masterConfigData := pchelper.GetMasterConfiguration(dbconfig)

	if err != nil {
		fmt.Println("Error occured while fetching configuration data: ", err.Error())
		logger.Context().LogError(SUB_MODULE_NAME, logger.Server, "Error occured while fetching configuration data ", err)
		return err
	}

	errPrepareConfig, masterConfigSetting := pcmgr.PrepareMasterConfiguration(dbconfig, masterConfigData, gmodels.PRODUCT_TYPE_HKT)

	if errPrepareConfig != nil {
		fmt.Println("Error occured while fetching configuration data: ", err.Error())
		logger.Context().LogError(SUB_MODULE_NAME, logger.Server, "Error occured while fetching configuration data", errPrepareConfig)
		return err
	}

	dbConnnErr = pchelper.VerifyDBConnection(masterConfigSetting.ProdMstDBConfig)

	if dbConnnErr != nil {
		fmt.Printf("Error occured while verifing prod mst db connection. Error: %s", dbConnnErr.Error())
		logger.Context().LogError(SUB_MODULE_NAME, logger.Server, "Error occured while verifing prod mst db connection.", dbConnnErr)
		return dbConnnErr
	}

	prodConfigErr, prodConfigData := pchelper.GetConfiguration(masterConfigSetting.ProdMstDBConfig, dbquery.QUERY_GET_PRODUCT_MASTER_CONFIGURATION)

	if prodConfigErr != nil {
		fmt.Println("Error occured while fetching configuration data: ", err.Error())
		logger.Context().LogError(SUB_MODULE_NAME, logger.Server, "Error occured while fetching configuration data", prodConfigErr)
		return err
	}

	prodconfUpdateErr := pcmgr.UpdateProductConfiguration(masterConfigSetting, prodConfigData)

	if prodconfUpdateErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Server, "Error occured while update product configuration", prodconfUpdateErr)
		return prodconfUpdateErr
	}

	setglobalErr := SetGlobal(dbconfig, masterConfigSetting)

	if setglobalErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Server, "Error occured while setting global values", setglobalErr)
		return setglobalErr
	}

	initErr := pcmgr.InitModules(masterConfigSetting)

	if initErr != nil {
		fmt.Println("Error occured while fetching configuration data: ", initErr.Error())
		return initErr
	}

	pcmgr.SetLogger(masterConfigSetting, gmodels.PRODUCT_TYPE_HKT)
	logger.SetModule("HKT.Endpoint")

	initModErr := initModules(masterConfigSetting)

	if initModErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Server, "Error occured while initilizing modules.", initModErr)
		return initModErr
	}

	return nil
}

func initModules(configSetting *gmodels.ConfigSettings) error {

	//	broker: 'redis://localhost:6379'
	//default_queue: machinery_tasks
	//result_backend: 'redis://127.0.0.1:6379'
	//results_expire_in: 3600000

	repo.Init()

	initTaskQueErr := initTaskQue(configSetting)

	if initTaskQueErr != nil {
		fmt.Printf("Error occured while iniTaskQue. Error: %s", initTaskQueErr.Error())
		logger.Context().LogError(SUB_MODULE_NAME, logger.Server, "Error occured while initilizing task queue.", initTaskQueErr)
		return initTaskQueErr
	}

	initTaskQueHandler()

	webServerStartErr := pcmgr.HandleEndPoint(8080, EPHandler{})

	return webServerStartErr
}

func initTaskQue(configSetting *gmodels.ConfigSettings) error {

	mstTaskConfig := taskque.TaskConfig{}
	mstTaskConfig.Broker = "redis://" + configSetting.MasterQueCache.Address + ":" + strconv.Itoa(configSetting.MasterQueCache.Port)
	mstTaskConfig.ResultBackend = "redis://" + configSetting.MasterQueCache.Address + ":" + strconv.Itoa(configSetting.MasterQueCache.Port)
	mstTaskConfig.DefaultQueue = gmodels.SPL_SERVER_DEFAULT_TASK_QUEUE

	mstTaskCtx := &taskque.TaskContext{}

	mstTaskQueErr := mstTaskCtx.CreateServer(mstTaskConfig)

	if mstTaskQueErr != nil {
		return mstTaskQueErr
	}

	prodTaskConfig := taskque.TaskConfig{}
	prodTaskConfig.Broker = "redis://" + configSetting.ProductCache.Address + ":" + strconv.Itoa(configSetting.MasterQueCache.Port)
	prodTaskConfig.ResultBackend = "redis://" + configSetting.ProductCache.Address + ":" + strconv.Itoa(configSetting.MasterQueCache.Port)
	prodTaskConfig.DefaultQueue = gmodels.HKT_SERVER_DEFAULT_TASK_QUEUE
	prodTaskConfig.ResultsExpireIn = 1 // In minute

	prodTaskCtx := &taskque.TaskContext{}

	prodTaskQueErr := prodTaskCtx.CreateServer(prodTaskConfig)

	if prodTaskQueErr != nil {
		return prodTaskQueErr
	}

	repo.Instance().MasterTaskContext = mstTaskCtx
	repo.Instance().ProdTaskContext = prodTaskCtx

	return nil
}

func initTaskQueHandler() {

	var hkthandler map[string]interface{}
	hkthandler = make(map[string]interface{})

	RegisterTaskHandler(hkthandler)

	repo.Instance().ProdTaskContext.RegisterTaskHandlers(hkthandler)

	go repo.Instance().ProdTaskContext.StartWorker("EPAck")
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

	repo.Init()
	repo.Instance().Context = ctx
	repo.Instance().Config = configSetting

	return nil

}
