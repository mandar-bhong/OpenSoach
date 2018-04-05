package manager

import (
	"errors"
	"fmt"

	gmodels "opensoach.com/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	dbmgr "opensoach.com/core/manager/db"

	"strconv"

	"github.com/go-redis/redis"
	"opensoach.com/core"
	"opensoach.com/core/logger"
	coremodels "opensoach.com/core/models"
	"opensoach.com/spl/models"
	repo "opensoach.com/spl/repository"
	"opensoach.com/spl/webserver"
)

func InitilizeModues(dbconfig *gmodels.ConfigDB) error {

	err, configData := getConfiguration(dbconfig)

	if err != nil {
		fmt.Println("Error occured while fetching configuration data: ", err.Error())
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

	dbEngine, err := sqlx.Connect(config.DBDriver, config.ConnectionString)

	configRows := &[]gmodels.DBMasterConfigRowModel{}

	if err != nil {
		fmt.Printf("DB Error %#+v \n", err.Error())
		return err, nil
	}

	filter := gmodels.DBMasterConfigRowModel{}
	filter.Category = models.DB_CONFIG_CATEGORY_SPL

	selCtx := dbmgr.SelectContext{}
	selCtx.Engine = dbEngine
	selCtx.Query = models.QUERY_GET_CONFIGURATION
	selCtx.Type = dbmgr.Query
	selCtx.Dest = configRows
	selCtx.TableName = ""

	selErr := selCtx.SelectByFilter(filter, "Category")

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
		if dbRow.Category == "SPL" {

			switch dbRow.Key {
			case "Web.Service.Address":
				webConfig.ServiceAddress = dbRow.Value
				break
			case "Cache.Address":
				mstCacheConfig.Address = dbRow.Value
				break
			case "Cache.Address.Password":
				mstCacheConfig.Password = dbRow.Value
				break
			case "Cache.Address.DB":
				mstDBPort, err := strconv.Atoi(dbRow.Value)
				if err != nil {
					return errors.New("Unable to convert Master Cache DB value to interger"), nil
				}
				mstCacheConfig.DB = mstDBPort
				break
			}
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

	dbEngine, dbErr := sqlx.Connect(configSetting.DBConfig.DBDriver, configSetting.DBConfig.ConnectionString)

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
		return redisMstErr
	}

	ctx := &core.Context{}
	ctx.Dynamic.Cache.RedisClient = client
	ctx.Dynamic.DB = dbEngine

	//ctx.Dynamic.Cache = configSetting.MasterCache

	repo.Init(configSetting, ctx)

	return nil

}
