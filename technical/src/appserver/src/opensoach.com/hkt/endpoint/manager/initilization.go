package manager

import (
	"fmt"
	"strconv"

	"opensoach.com/core/logger"
	"opensoach.com/hkt/constants/dbquery"

	"opensoach.com/core"
	taskque "opensoach.com/core/manager/taskqueue"
	repo "opensoach.com/hkt/endpoint/repository"
	gmodels "opensoach.com/models"
	wm "opensoach.com/prodcore/endpoint/webSocketManager"
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

	initErr := pcmgr.InitModules(masterConfigSetting)

	if initErr != nil {
		fmt.Println("Error occured while fetching configuration data: ", initErr.Error())
		return initErr
	}

	initModules(masterConfigSetting)

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

	webServerStartErr := wm.Init(8080, OnEPConnection, OnEPDisConnection, OnEPMessage)

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
	prodTaskConfig.ResultsExpireIn = 60 * 1000 // 60 Sec

	prodTaskCtx := &taskque.TaskContext{}

	prodTaskQueErr := prodTaskCtx.CreateServer(prodTaskConfig)

	if prodTaskQueErr != nil {
		return prodTaskQueErr
	}

	repo.Instance().Config = configSetting
	repo.Instance().Context = &core.Context{}
	repo.Instance().MasterTaskContext = mstTaskCtx
	repo.Instance().ProdTaskContext = prodTaskCtx

	return nil
}

func initTaskQueHandler() {

	var hkthandler map[string]interface{}
	hkthandler = make(map[string]interface{})

	hkthandler["EPAck"] = ProcessEPPacket
	repo.Instance().ProdTaskContext.RegisterTaskHandlers(hkthandler)

	go repo.Instance().ProdTaskContext.StartWorker("EPAck")
}
