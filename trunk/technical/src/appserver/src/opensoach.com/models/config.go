package models

type ConfigSettings struct {
	ServerConfig    *ConfigServer       `json:"serverconfiguration"`
	DBConfig        *ConfigDB           `json:"databaseconfiguration"`
	ProdMstDBConfig *ConfigDB           `json:"productmasterdbconfig"`
	WebConfig       *ConfigWebSettings  `json:"webconfiguration"`
	MasterCache     *ConfigCacheAddress `json:"mstcacheconfiguration"`
	MasterQueCache  *ConfigCacheAddress `json:"mstquecacheconfiguration"`
	ProductCache    *ConfigCacheAddress `json:"prodcacheconfiguration"`
	ProductQueCache *ConfigCacheAddress `json:"prodquecacheconfiguration"`
}

type ConfigDB struct {
	ConnectionString string `json:"connectionstring"`
	DBDriver         string `json:"driver"`
}

type ConfigWebSettings struct {
	ServiceAddress string `json:"webaddress"`
	SessionTimeOut int    `json:"sessiontimeout"`
}

type ConfigCacheAddress struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type ConfigServer struct {
	BaseDir string `json:"basefolder"`
}
