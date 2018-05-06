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

	pcmgr.UpdateProductConfiguration(masterConfigSetting, prodConfigData)

	SetGlobal(dbconfig, masterConfigSetting)

	initErr := pcmgr.InitModules(masterConfigSetting)

	if initErr != nil {
		fmt.Println("Error occured while fetching configuration data: ", initErr.Error())
		return initErr
	}

	initModules(masterConfigSetting)

	fmt.Printf("RepoContx 111 : %#v \n", repo.Instance().Context)
	fmt.Printf("RepoContx 111222 : %#v \n", repo.Instance().MasterTaskContext)
	fmt.Printf("RepoContx 111333 : %#v \n", repo.Instance().ProdTaskContext)

	return nil
}

func initModules(configSetting *gmodels.ConfigSettings) error {
	logger.SetModule("HKT.Endpoint")

	//	broker: 'redis://localhost:6379'
	//default_queue: machinery_tasks
	//result_backend: 'redis://127.0.0.1:6379'
	//results_expire_in: 3600000

	repo.Init()

	initTaskQue(configSetting)

	initTaskQueHandler()

	webServerStartErr := pcmgr.HandleEndPoint(8080, EPHandler{})

	return webServerStartErr
}

func initTaskQue(configSetting *gmodels.ConfigSettings) error {

	mstTaskConfig := taskque.TaskConfig{}
	mstTaskConfig.Broker = "redis://" + configSetting.MasterQueCache.Address + ":" + strconv.Itoa(configSetting.MasterQueCache.Port)
	mstTaskConfig.ResultBackend = "redis://" + configSetting.MasterQueCache.Address + ":" + strconv.Itoa(configSetting.MasterQueCache.Port)
	mstTaskConfig.DefaultQueue = "SPL"

	mstTaskCtx := &taskque.TaskContext{}

	mstTaskQueErr := mstTaskCtx.CreateServer(mstTaskConfig)

	if mstTaskQueErr != nil {
		return mstTaskQueErr
	}

	prodTaskConfig := taskque.TaskConfig{}
	prodTaskConfig.Broker = "redis://" + configSetting.ProductCache.Address + ":" + strconv.Itoa(configSetting.MasterQueCache.Port)
	prodTaskConfig.ResultBackend = "redis://" + configSetting.ProductCache.Address + ":" + strconv.Itoa(configSetting.MasterQueCache.Port)
	prodTaskConfig.DefaultQueue = "HKT"
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
