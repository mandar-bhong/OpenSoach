package manager

import (
	"errors"
	"fmt"

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
	"opensoach.com/spl/constants"
	"opensoach.com/spl/models"
	repo "opensoach.com/spl/repository"
	"opensoach.com/spl/webserver"
)

func InitilizeModues(dbconfig *gmodels.ConfigDB) error {

	err, configData := getConfiguration(dbconfig)

	if err != nil {
		fmt.Println("Error occured while fetching configuration data: ", err.Error())
		return err
	}

	errPrepareConfig, configSetting := prepareConfiguration(dbconfig, configData)

	if errPrepareConfig != nil {
		//TODO: log message, need to identify fmt or file base or both
		return errPrepareConfig
	}

	//init logger for fmt or file or both for temp then switch mode as per configuration after component connection verification

	connErr := verifyConnectionSetGlobal(dbconfig, configSetting)

	if connErr != nil {
		//TODO: log message, need to identify fmt or file base or both
		return connErr
	}

	initModules(configSetting)

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

	for _, dbRow := range *configData {

		switch dbRow.ConfigKey {
		case constants.DB_CONFIG_WEB_SERVICE_ADDRESS:
			webConfig.ServiceAddress = dbRow.ConfigValue
			break
		case constants.DB_CONFIG_CACHE_ADDRESS:
			mstCacheConfig.Address = dbRow.ConfigValue
			break
		case constants.DB_CONFIG_CACHE_ADDRESS_PASSWORD:
			mstCacheConfig.Password = dbRow.ConfigValue
			break
		case constants.DB_CONFIG_ADDRESS_DB:
			mstDBPort, err := strconv.Atoi(dbRow.ConfigValue)
			if err != nil {
				return errors.New("Unable to convert Master Cache DB value to interger"), nil
			}
			mstCacheConfig.DB = mstDBPort
			break
		}
	}

	return nil, globalConfiguration
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

func verifyConnectionSetGlobal(dbconfig *gmodels.ConfigDB, configSetting *gmodels.ConfigSettings) error {

	_, dbErr := sqlx.Connect(configSetting.DBConfig.DBDriver, configSetting.DBConfig.ConnectionString)

	if dbErr != nil {
		return dbErr
	}

	client := redis.NewClient(&redis.Options{
		Addr:     configSetting.MasterCache.Address,
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
