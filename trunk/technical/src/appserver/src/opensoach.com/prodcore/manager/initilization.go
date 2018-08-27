package manager

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"

	"github.com/go-redis/redis"

	gmodels "opensoach.com/models"

	"opensoach.com/core"
	"opensoach.com/core/logger"
	coremodels "opensoach.com/core/models"
	pcconst "opensoach.com/prodcore/constants"
	pcepmgr "opensoach.com/prodcore/endpoint/manager"
)

func PrepareMasterConfiguration(dbconfig *gmodels.ConfigDB, configData *[]gmodels.DBMasterConfigRowModel, productType string) (error, *gmodels.ConfigSettings) {

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

	prodMstDBConfig := &gmodels.ConfigDB{}
	globalConfiguration.ProdMstDBConfig = prodMstDBConfig

	loggerConfig := &gmodels.ConfigLogger{}
	globalConfiguration.LoggerConfig = loggerConfig

	emailConfig := &gmodels.ConfigEmail{}
	globalConfiguration.EmailConfig = emailConfig

	for _, dbRow := range *configData {

		switch dbRow.ConfigKey {
		case pcconst.DB_CONFIG_LOGGER_LOG_LEVEL:
			loggerConfig.LogLevel = dbRow.ConfigValue
			break
		case pcconst.DB_CONFIG_LOGGER_LOGGING_TYPE:
			loggerConfig.LoggingType = dbRow.ConfigValue
			break
		case pcconst.DB_CONFIG_LOGGER_FLUENT_HOST:
			loggerConfig.LoggingFluentHost = dbRow.ConfigValue
			break
		case pcconst.DB_CONFIG_LOGGER_INFLUXDB_HOST:
			loggerConfig.LoggingInfluxDBHost = dbRow.ConfigValue
			break
		case pcconst.DB_CONFIG_WEB_SERVICE_ADDRESS:
			webConfig.ServiceAddress = dbRow.ConfigValue
			break
		case pcconst.DB_CONFIG_CACHE_ADDRESS_HOST:
			mstCacheConfig.Address = dbRow.ConfigValue
			break
		case pcconst.DB_CONFIG_CACHE_ADDRESS_PORT:
			mstAddPort, err := strconv.Atoi(dbRow.ConfigValue)
			if err != nil {
				return errors.New(fmt.Sprintf("Unable to convert Master Cache Port value to interger. Received Value : %s", dbRow.ConfigValue)), nil
			}
			mstCacheConfig.Port = mstAddPort
			break
		case pcconst.DB_CONFIG_CACHE_ADDRESS_PASSWORD:
			mstCacheConfig.Password = dbRow.ConfigValue
			break
		case pcconst.DB_CONFIG_CACHE_ADDRESS_DB:
			mstDBPort, err := strconv.Atoi(dbRow.ConfigValue)
			if err != nil {
				return errors.New(fmt.Sprintf("Unable to convert Master Cache DB value to interger. Received Value : %s", dbRow.ConfigValue)), nil
			}
			mstCacheConfig.DB = mstDBPort
			break

		case pcconst.DB_CONFIG_QUE_ADDRESS_HOST:
			mstQueCacheConfig.Address = dbRow.ConfigValue
			break
		case pcconst.DB_CONFIG_QUE_ADDRESS_PORT:
			mstQueAddPort, err := strconv.Atoi(dbRow.ConfigValue)
			if err != nil {
				return errors.New(fmt.Sprintf("Unable to convert Master Que Cache Port value to interger. Received Value : %s", dbRow.ConfigValue)), nil
			}
			mstQueCacheConfig.Port = mstQueAddPort

			break

		case pcconst.DB_CONFIG_QUE_ADDRESS_PASSWORD:
			mstQueCacheConfig.Password = dbRow.ConfigValue
			break
		case pcconst.DB_CONFIG_QUE_ADDRESS_DB:
			mstDBPort, err := strconv.Atoi(dbRow.ConfigValue)
			if err != nil {
				return errors.New(fmt.Sprintf("Unable to convert Master Que Cache DB value to interger. Received Value : %s", dbRow.ConfigValue)), nil
			}
			mstQueCacheConfig.DB = mstDBPort
			break

		case pcconst.DB_CONFIG_SERVER_WIN_BASE_DIRECTORY:
			if runtime.GOOS == "windows" {
				serverConfig.BaseDir = dbRow.ConfigValue
			}
			break

		case pcconst.DB_CONFIG_SERVER_LIN_BASE_DIRECTORY:
			if runtime.GOOS == "linux" {
				serverConfig.BaseDir = dbRow.ConfigValue
			}
			break

		case pcconst.DB_CONFIG_HKT_MASTER_DB_CONNECTION: //Addition product connection will be listed here
			switch productType {
			case gmodels.PRODUCT_TYPE_HKT:
				prodMstDBConfig.ConnectionString = dbRow.ConfigValue
				prodMstDBConfig.DBDriver = "mysql"
				break
			}
			break

		case pcconst.DB_CONFIG_HPFT_MASTER_DB_CONNECTION: //Addition product connection will be listed here
			switch productType {
			case gmodels.PRODUCT_TYPE_HPFT:
				prodMstDBConfig.ConnectionString = dbRow.ConfigValue
				prodMstDBConfig.DBDriver = "mysql"
				break
			}
			break

		case pcconst.DB_CONFIG_VST_MASTER_DB_CONNECTION: //Addition product connection will be listed here
			switch productType {
			case gmodels.PRODUCT_TYPE_VST:
				prodMstDBConfig.ConnectionString = dbRow.ConfigValue
				prodMstDBConfig.DBDriver = "mysql"
				break
			}
			break

		case pcconst.DB_CONFIG_SMTP_FROM:
			emailConfig.From = dbRow.ConfigValue
		case pcconst.DB_CONFIG_SMTP_ADDRESS:
			emailConfig.SMTPAddress = dbRow.ConfigValue
		case pcconst.DB_CONFIG_SMTP_USERNAME:
			emailConfig.SMTPUsername = dbRow.ConfigValue
		case pcconst.DB_CONFIG_SMTP_PASSWORD:
			emailConfig.SMTPPassword = dbRow.ConfigValue
		case pcconst.DB_CONFIG_SMTP_PORT:
			mstSMTPPort, err := strconv.Atoi(dbRow.ConfigValue)
			if err != nil {
				return errors.New(fmt.Sprintf("Unable to convert Master Cache Port value to interger. Received Value : %s", dbRow.ConfigValue)), nil
			}

			emailConfig.SMTPPort = mstSMTPPort

		}
	}

	return nil, globalConfiguration
}

func UpdateProductConfiguration(globalConfiguration *gmodels.ConfigSettings, configData *[]gmodels.DBMasterConfigRowModel) error {

	productCache := &gmodels.ConfigCacheAddress{}
	globalConfiguration.ProductCache = productCache

	productQueCache := &gmodels.ConfigCacheAddress{}
	globalConfiguration.ProductQueCache = productQueCache

	webconfig := &gmodels.ConfigWebSettings{}
	globalConfiguration.WebConfig = webconfig

	for _, dbRow := range *configData {
		switch dbRow.ConfigKey {

		case pcconst.DB_CONFIG_PRODUCT_WEB_SERVICE_ADDRESS:
			webconfig.ServiceAddress = dbRow.ConfigValue
			break
		case pcconst.DB_CONFIG_PRODUCT_CACHE_ADDRESS_HOST:
			productCache.Address = dbRow.ConfigValue
			break
		case pcconst.DB_CONFIG_PRODUCT_CACHE_ADDRESS_PORT:
			prodHostPort, err := strconv.Atoi(dbRow.ConfigValue)
			if err != nil {
				return errors.New(fmt.Sprintf("Unable to convert Master Cache DB value to interger. Received Value : %s", dbRow.ConfigValue))
			}
			productCache.Port = prodHostPort
			break
		case pcconst.DB_CONFIG_PRODUCT_CACHE_ADDRESS_PASSWORD:
			productCache.Password = dbRow.ConfigValue
			break
		case pcconst.DB_CONFIG_PRODUCT_CACHE_ADDRESS_DB:
			prodDB, err := strconv.Atoi(dbRow.ConfigValue)
			if err != nil {
				return errors.New(fmt.Sprintf("Unable to convert Master Cache DB value to interger. Received Value : %s", dbRow.ConfigValue))
			}
			productCache.DB = prodDB
			break

		case pcconst.DB_CONFIG_PRODUCT_QUE_ADDRESS_HOST:
			productQueCache.Address = dbRow.ConfigValue
			break
		case pcconst.DB_CONFIG_PRODUCT_QUE_ADDRESS_PORT:
			prodQueHostPort, err := strconv.Atoi(dbRow.ConfigValue)
			if err != nil {
				return errors.New(fmt.Sprintf("Unable to convert Master Cache DB value to interger. Received Value : %s", dbRow.ConfigValue))
			}
			productQueCache.Port = prodQueHostPort
			break
		case pcconst.DB_CONFIG_PRODUCT_QUE_ADDRESS_PASSWORD:
			productQueCache.Password = dbRow.ConfigValue
			break
		case pcconst.DB_CONFIG_PRODUCT_QUE_ADDRESS_DB:
			prodQueDB, err := strconv.Atoi(dbRow.ConfigValue)
			if err != nil {
				return errors.New(fmt.Sprintf("Unable to convert Master Cache DB value to interger. Received Value : %s", dbRow.ConfigValue))
			}
			productQueCache.DB = prodQueDB
			break
		}
	}

	return nil
}

func VerifyConnection(dbconfig *gmodels.ConfigDB, configSetting *gmodels.ConfigSettings) (error, int) {

	var moduleErrType int

	client := redis.NewClient(&redis.Options{
		Addr:     configSetting.MasterCache.Address + ":" + strconv.Itoa(configSetting.MasterCache.Port),
		Password: configSetting.MasterCache.Password,
		DB:       configSetting.MasterCache.DB,
	})

	_, redisMstErr := client.Ping().Result()

	if redisMstErr != nil {
		fmt.Printf("Unable to connect redis server Address : '%s', Password : '%s', DB: '%d' \n", configSetting.MasterCache.Address, configSetting.MasterCache.Password, configSetting.MasterCache.DB)
		moduleErrType = 2
		return redisMstErr, moduleErrType
	}

	queClient := redis.NewClient(&redis.Options{
		Addr:     configSetting.MasterQueCache.Address + ":" + strconv.Itoa(configSetting.MasterQueCache.Port),
		Password: configSetting.MasterQueCache.Password,
		DB:       configSetting.MasterQueCache.DB,
	})

	_, redisQueMstErr := queClient.Ping().Result()

	if redisQueMstErr != nil {
		fmt.Printf("Unable to connect redis server Address : '%s', Password : '%s', DB: '%d' \n", configSetting.MasterQueCache.Address, configSetting.MasterQueCache.Password, configSetting.MasterQueCache.DB)
		moduleErrType = 3
		return redisQueMstErr, moduleErrType
	}

	if configSetting.ProductCache != nil {

		client := redis.NewClient(&redis.Options{
			Addr:     configSetting.ProductCache.Address + ":" + strconv.Itoa(configSetting.ProductCache.Port),
			Password: configSetting.ProductCache.Password,
			DB:       configSetting.ProductCache.DB,
		})

		_, redisMstErr := client.Ping().Result()

		if redisMstErr != nil {
			fmt.Printf("Unable to connect redis server Address : '%s', Password : '%s', DB: '%d' \n", configSetting.ProductCache.Address, configSetting.ProductCache.Password, configSetting.ProductCache.DB)
			moduleErrType = 4
			return redisMstErr, moduleErrType
		}

	}

	if configSetting.ProductQueCache != nil {
		client := redis.NewClient(&redis.Options{
			Addr:     configSetting.ProductQueCache.Address + ":" + strconv.Itoa(configSetting.ProductQueCache.Port),
			Password: configSetting.ProductQueCache.Password,
			DB:       configSetting.ProductQueCache.DB,
		})

		_, redisMstErr := client.Ping().Result()

		if redisMstErr != nil {
			fmt.Printf("Unable to connect redis server Address : '%s', Password : '%s', DB: '%d' \n", configSetting.ProductQueCache.Address, configSetting.ProductQueCache.Password, configSetting.ProductQueCache.DB)
			moduleErrType = 5
			return redisMstErr, moduleErrType
		}
	}

	return nil, moduleErrType

}

func InitModules(configSetting *gmodels.ConfigSettings) error {
	logger.Init()
	logger.SetLogLevel(logger.Debug)
	logger.SetLoggingService(logger.LoggingServiceFmt)

	coreConfig := &coremodels.CoreConfig{}
	err := core.Init(coreConfig)

	if err != nil {
		return err
	}

	return nil
}

func HandleEndPoint(wsport int, handler pcepmgr.EPHandler) error {
	return pcepmgr.Init(wsport, handler)
}

func SetLogger(configSetting *gmodels.ConfigSettings) {
	logger.Init()

	if configSetting.LoggerConfig == nil {
		logger.Context().LogError("Prod.Core", logger.Server, "Logger configuration is nil", nil)
		return
	}

	switch configSetting.LoggerConfig.LogLevel {
	case pcconst.LOGGER_LOGGING_LEVEL_ERROR:
		logger.SetLogLevel(logger.Error)
	case pcconst.LOGGER_LOGGING_LEVEL_DEBUG:
		logger.SetLogLevel(logger.Debug)
	case pcconst.LOGGER_LOGGING_LEVEL_INFO:
		logger.SetLogLevel(logger.Info)
	default:
		logger.SetLogLevel(logger.Error)
	}

	switch configSetting.LoggerConfig.LoggingType {
	case pcconst.LOGGER_LOGGING_TYPE_STDIO:
		logger.SetLoggingService(logger.LoggingServiceFmt)
	case pcconst.LOGGER_LOGGING_TYPE_FLUENT:
		logger.SetLoggingService(logger.LoggingServiceFluent)
		logger.SetFluentHost(configSetting.LoggerConfig.LoggingFluentHost)
	case pcconst.LOGGER_LOGGING_TYPE_INFLUXDB:
		logger.SetLoggingService(logger.LoggingServiceInfluxDB)
		logger.SetInfluxDBHost(configSetting.LoggerConfig.LoggingInfluxDBHost)
	default:
		logger.SetLoggingService(logger.LoggingServiceFmt)
	}

}
