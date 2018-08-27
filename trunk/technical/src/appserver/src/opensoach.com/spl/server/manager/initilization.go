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
	pcconst "opensoach.com/prodcore/constants"
	pchelper "opensoach.com/prodcore/helper"
	pcmgr "opensoach.com/prodcore/manager"
	repo "opensoach.com/spl/server/repository"
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

	splMasterConfigSetting := PrepareSPLProdMasterConfiguration(masterConfigData)
	splMasterConfigSetting.ConfigSettings = *masterConfigSetting

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

	setGlobalErr := SetGlobal(dbconfig, splMasterConfigSetting)

	if setGlobalErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Server, "Error occured while setting global values.", setGlobalErr)
		return setGlobalErr
	}

	initModules(masterConfigSetting)

	return nil
}

func SetGlobal(dbconfig *gmodels.ConfigDB, splconfigSetting *gmodels.SPLConfigSettings) error {

	ghelper.BaseDir = splconfigSetting.ConfigSettings.ServerConfig.BaseDir

	isJsonConvMstCacheSuccess, jsonMstCacheRedisAddress := ghelper.ConvertToJSON(splconfigSetting.ConfigSettings.MasterCache)

	if isJsonConvMstCacheSuccess == false {
		logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while converting ConfigCacheAddress structure to JSON")
		return errors.New("Error occured while converting ConfigCacheAddress to json")
	}

	ctx := &core.Context{}
	ctx.Master.DBConn = splconfigSetting.ConfigSettings.DBConfig.ConnectionString
	ctx.Master.Cache.CacheAddress = jsonMstCacheRedisAddress
	// ctx.ProdMst.DBConn = splconfigSetting.SPLProdMstDBConfig

	repo.Init(splconfigSetting, ctx)

	return nil
}

func initModules(configSetting *gmodels.ConfigSettings) error {
	logger.SetModule("SPL.Server")

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

	//handlerTaskCtx := &taskque.TaskContext{}

	var handler map[string]interface{}
	handler = make(map[string]interface{})

	RegisterHandler(handler)

	mstTaskCtx.RegisterTaskHandlers(handler)

	repo.Instance().TaskQue = mstTaskCtx

	mstTaskCtx.StartWorker(gmodels.SPL_TASK_QUEUE)

	return nil
}

func PrepareSPLProdMasterConfiguration(configData *[]gmodels.DBMasterConfigRowModel) *gmodels.SPLConfigSettings {

	splglobalConfiguration := &gmodels.SPLConfigSettings{}
	splglobalConfiguration.SPLProdMstDBConfig = make(map[string]*gmodels.ConfigDB)

	for _, dbRow := range *configData {

		switch dbRow.ConfigKey {
		case pcconst.DB_CONFIG_HKT_MASTER_DB_CONNECTION:
			prodMstDBConfig := &gmodels.ConfigDB{}
			prodMstDBConfig.ConnectionString = dbRow.ConfigValue
			prodMstDBConfig.DBDriver = "mysql"
			splglobalConfiguration.SPLProdMstDBConfig["SPL_HKT"] = prodMstDBConfig
			break

		case pcconst.DB_CONFIG_HPFT_MASTER_DB_CONNECTION:
			prodMstDBConfig := &gmodels.ConfigDB{}
			prodMstDBConfig.ConnectionString = dbRow.ConfigValue
			prodMstDBConfig.DBDriver = "mysql"
			splglobalConfiguration.SPLProdMstDBConfig["SPL_HPFT"] = prodMstDBConfig

			break

		}

	}

	return splglobalConfiguration
}
