package models

type ConfigSettings struct {
	DBConfig    *ConfigDB           `json:"databaseconfiguration"`
	WebConfig   *ConfigWebSettings  `json:"webconfiguration"`
	MasterCache *ConfigCacheAddress `json:"mstcacheconfiguration"`
	ModuleCache *ConfigCacheAddress `json:"modcacheconfiguration"`
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
	Password string `json:"password"`
	DB       int    `json:"db"`
}
