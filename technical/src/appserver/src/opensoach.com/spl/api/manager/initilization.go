package manager

import (
	"errors"
	"fmt"
	"runtime"

	ghelper "opensoach.com/core/helper"
	gmodels "opensoach.com/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	dbmgr "opensoach.com/core/manager/db"

	"strconv"

	"github.com/go-redis/redis"
	"opensoach.com/core"
	"opensoach.com/core/logger"

	coremodels "opensoach.com/core/models"
	pcconstants "opensoach.com/prodcore/constants"
	pchelper "opensoach.com/prodcore/helper"
	pcmgr "opensoach.com/prodcore/manager"
	"opensoach.com/spl/api/models"
	repo "opensoach.com/spl/api/repository"
	"opensoach.com/spl/api/webserver"

	taskque "opensoach.com/core/manager/taskqueue"
)

func InitilizeModues(dbconfig *gmodels.ConfigDB) error {

	dbConnnErr := pchelper.VerifyDBConnection(dbconfig)

	if dbConnnErr != nil {
		logger.Context().WithField("DBConfig", dbconfig).LogError(SUB_MODULE_NAME, logger.Normal, "Unable to verify master db connection.", dbConnnErr)
		return dbConnnErr
	}

	err, masterConfigData := pchelper.GetMasterConfiguration(dbconfig)

	if err != nil {
		logger.Context().WithField("DBConfig", dbconfig).LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while fetching configuration data.", err)
		return err
	}

	errPrepareConfigErr, masterConfigSetting := pcmgr.PrepareMasterConfiguration(dbconfig, masterConfigData, gmodels.PRODUCT_TYPE_HKT)

	if errPrepareConfigErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while preparing master configuration data.", errPrepareConfigErr)
		return errPrepareConfigErr
	}

	verifyErr, moduleType := pcmgr.VerifyConnection(dbconfig, masterConfigSetting)

	if verifyErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while verifing connections.", verifyErr)

		switch moduleType {
		case 2: //Redis server connection error
			return verifyErr
		case 3: //Redis server master que cache error
			return verifyErr
		case 4: //Redis server product cache error
			return verifyErr
		}
	}

	pcmgr.SetLogger(masterConfigSetting)

	setGlobalErr := SetGlobal(dbconfig, masterConfigSetting)

	if setGlobalErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Server, "Error occured while setting global values.", setGlobalErr)
		return setGlobalErr
	}

	if taskErr := createTaskQue(masterConfigSetting); taskErr != nil {
		return taskErr
	}

	initErr := initModules(masterConfigSetting)

	if initErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Server, "Initiliazation module error occured", initErr)
		return initErr
	}

	return nil
}

func SetGlobal(dbconfig *gmodels.ConfigDB, masterConfigSetting *gmodels.ConfigSettings) error {

	ghelper.BaseDir = masterConfigSetting.ServerConfig.BaseDir

	isJsonConvMstCacheSuccess, jsonMstCacheRedisAddress := ghelper.ConvertToJSON(masterConfigSetting.MasterCache)

	if isJsonConvMstCacheSuccess == false {
		logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Error occured while converting ConfigCacheAddress structure to JSON")
		return errors.New("Error occured while converting ConfigCacheAddress to json")

	}
	ctx := &core.Context{}
	ctx.Master.Cache.CacheAddress = jsonMstCacheRedisAddress
	ctx.Master.DBConn = dbconfig.ConnectionString

	repo.Init(masterConfigSetting, ctx)

	return nil
}

func getConfiguration(config *gmodels.ConfigDB) (error, *[]gmodels.DBMasterConfigRowModel) {

	configRows := &[]gmodels.DBMasterConfigRowModel{}

	selCtx := dbmgr.SelectContext{}
	selCtx.DBConnection = config.ConnectionString
	selCtx.Query = models.QUERY_GET_CONFIGURATION
	selCtx.QueryType = dbmgr.Query
	selCtx.Dest = configRows

	selErr := selCtx.SelectAll()

	if selErr != nil {
		fmt.Printf("DB Error %#+v \n", selErr.Error())
		return selErr, nil
	}

	return nil, configRows
}

func prepareConfiguration(dbconfig *gmodels.ConfigDB, configData *[]gmodels.DBMasterConfigRowModel) (error, *gmodels.ConfigSettings) {

	globalConfiguration := &gmodels.ConfigSettings{}
	globalConfiguration.DBConfig = dbconfig

	webConfig := &gmodels.ConfigWebSettings{}
	globalConfiguration.WebConfig = webConfig

	mstCacheConfig := &gmodels.ConfigCacheAddress{}
	globalConfiguration.MasterCache = mstCacheConfig

	mstQueCacheConfig := &gmodels.ConfigCacheAddress{}
	globalConfiguration.MasterQueCache = mstQueCacheConfig

	serverConfig := &gmodels.ConfigServer{}
	globalConfiguration.ServerConfig = serverConfig

	loggerConfig := &gmodels.ConfigLogger{}
	globalConfiguration.LoggerConfig = loggerConfig

	for _, dbRow := range *configData {

		switch dbRow.ConfigKey {
		case pcconstants.DB_CONFIG_WEB_SERVICE_ADDRESS:
			webConfig.ServiceAddress = dbRow.ConfigValue
			break
		case pcconstants.DB_CONFIG_CACHE_ADDRESS_HOST:
			mstCacheConfig.Address = dbRow.ConfigValue
			break
		case pcconstants.DB_CONFIG_CACHE_ADDRESS_PORT:
			mstAddPort, err := strconv.Atoi(dbRow.ConfigValue)
			if err != nil {
				return errors.New(fmt.Sprintf("Unable to convert Master Cache Port value to interger. Received Value : %s", dbRow.ConfigValue)), nil
			}
			mstCacheConfig.Port = mstAddPort
			break
		case pcconstants.DB_CONFIG_CACHE_ADDRESS_PASSWORD:
			mstCacheConfig.Password = dbRow.ConfigValue
			break
		case pcconstants.DB_CONFIG_CACHE_ADDRESS_DB:
			mstDBPort, err := strconv.Atoi(dbRow.ConfigValue)
			if err != nil {
				return errors.New(fmt.Sprintf("Unable to convert Master Cache DB value to interger. Received Value : %s", dbRow.ConfigValue)), nil
			}
			mstCacheConfig.DB = mstDBPort
			break

		case pcconstants.DB_CONFIG_QUE_ADDRESS_HOST:
			mstQueCacheConfig.Address = dbRow.ConfigValue
			break
		case pcconstants.DB_CONFIG_QUE_ADDRESS_PORT:
			mstQueAddPort, err := strconv.Atoi(dbRow.ConfigValue)
			if err != nil {
				return errors.New(fmt.Sprintf("Unable to convert Master Que Cache Port value to interger. Received Value : %s", dbRow.ConfigValue)), nil
			}
			mstQueCacheConfig.Port = mstQueAddPort

			break

		case pcconstants.DB_CONFIG_QUE_ADDRESS_PASSWORD:
			mstQueCacheConfig.Password = dbRow.ConfigValue
			break
		case pcconstants.DB_CONFIG_QUE_ADDRESS_DB:
			mstDBPort, err := strconv.Atoi(dbRow.ConfigValue)
			if err != nil {
				return errors.New(fmt.Sprintf("Unable to convert Master Que Cache DB value to interger. Received Value : %s", dbRow.ConfigValue)), nil
			}
			mstQueCacheConfig.DB = mstDBPort
			break

		case pcconstants.DB_CONFIG_SERVER_WIN_BASE_DIRECTORY:
			if runtime.GOOS == "windows" {
				serverConfig.BaseDir = dbRow.ConfigValue
			}
			break

		case pcconstants.DB_CONFIG_SERVER_LIN_BASE_DIRECTORY:
			if runtime.GOOS == "linux" {
				serverConfig.BaseDir = dbRow.ConfigValue
			}
			break
		}
	}

	return nil, globalConfiguration
}

func initModules(configSetting *gmodels.ConfigSettings) error {

	logger.SetModule("SPL")

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

func verifyConnectionSetGlobal(dbconfig *gmodels.ConfigDB, configSetting *gmodels.ConfigSettings) error {

	ghelper.BaseDir = configSetting.ServerConfig.BaseDir

	_, dbErr := sqlx.Connect(configSetting.DBConfig.DBDriver, configSetting.DBConfig.ConnectionString)

	if dbErr != nil {
		return dbErr
	}

	client := redis.NewClient(&redis.Options{
		Addr:     configSetting.MasterCache.Address + ":" + strconv.Itoa(configSetting.MasterCache.Port),
		Password: configSetting.MasterCache.Password,
		DB:       configSetting.MasterCache.DB,
	})

	_, redisMstErr := client.Ping().Result()

	if redisMstErr != nil {
		fmt.Printf("Unable to connect redis server Address : '%s', Password : '%s', DB: '%d' \n", configSetting.MasterCache.Address, configSetting.MasterCache.Password, configSetting.MasterCache.DB)
		return redisMstErr
	}

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

func createTaskQue(configSetting *gmodels.ConfigSettings) error {

	mstTaskConfig := taskque.TaskConfig{}
	mstTaskConfig.Broker = "redis://" + configSetting.MasterQueCache.Address + ":" + strconv.Itoa(configSetting.MasterQueCache.Port)
	mstTaskConfig.ResultBackend = "redis://" + configSetting.MasterQueCache.Address + ":" + strconv.Itoa(configSetting.MasterQueCache.Port)
	mstTaskConfig.DefaultQueue = gmodels.SPL_SERVER_DEFAULT_TASK_QUEUE
	mstTaskConfig.ResultsExpireIn = 1 // in min

	mstTaskCtx := &taskque.TaskContext{}

	if mstTaskQueErr := mstTaskCtx.CreateServer(mstTaskConfig); mstTaskQueErr != nil {
		return mstTaskQueErr
	}

	repo.Instance().TaskQue = mstTaskCtx

	return nil
}
