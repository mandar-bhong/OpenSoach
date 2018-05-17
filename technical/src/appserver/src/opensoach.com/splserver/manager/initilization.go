package manager

import (
	"errors"
	"fmt"
	"strconv"

	"opensoach.com/core"
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	taskque "opensoach.com/core/manager/taskqueue"
	gmodels "opensoach.com/models"
	pchelper "opensoach.com/prodcore/helper"
	pcmgr "opensoach.com/prodcore/manager"
	"opensoach.com/splserver/processor"
	repo "opensoach.com/splserver/repository"
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

	initModules(masterConfigSetting)

	return nil
}

func SetGlobal(dbconfig *gmodels.ConfigDB, configSetting *gmodels.ConfigSettings) error {

	ghelper.BaseDir = configSetting.ServerConfig.BaseDir

	isJsonConvMstCacheSuccess, jsonMstCacheRedisAddress := ghelper.ConvertToJSON(configSetting.MasterCache)

	if isJsonConvMstCacheSuccess == false {
		logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while converting ConfigCacheAddress structure to JSON")
		return errors.New("Error occured while converting ConfigCacheAddress to json")
	}

	ctx := &core.Context{}
	ctx.Master.DBConn = configSetting.DBConfig.ConnectionString
	ctx.Master.Cache.CacheAddress = jsonMstCacheRedisAddress
	ctx.ProdMst.DBConn = configSetting.ProdMstDBConfig.ConnectionString

	repo.Init(configSetting, ctx)

	return nil
}

func initModules(configSetting *gmodels.ConfigSettings) error {
	logger.SetModule("SPL.Server")

	mstTaskConfig := taskque.TaskConfig{}
	mstTaskConfig.Broker = "redis://" + configSetting.MasterQueCache.Address + ":" + strconv.Itoa(configSetting.MasterQueCache.Port)
	mstTaskConfig.ResultBackend = "redis://" + configSetting.MasterQueCache.Address + ":" + strconv.Itoa(configSetting.MasterQueCache.Port)
	mstTaskConfig.DefaultQueue = "SPL"
	mstTaskConfig.ResultsExpireIn = 1 // in min

	mstTaskCtx := &taskque.TaskContext{}

	mstTaskQueErr := mstTaskCtx.CreateServer(mstTaskConfig)

	if mstTaskQueErr != nil {
		return mstTaskQueErr
	}

	//handlerTaskCtx := &taskque.TaskContext{}

	var handler map[string]interface{}
	handler = make(map[string]interface{})

	processor.RegisterHandler(handler)

	mstTaskCtx.RegisterTaskHandlers(handler)

	repo.Instance().TaskQue = mstTaskCtx

	mstTaskCtx.StartWorker(gmodels.SPL_TASK_QUEUE)

	return nil
}
