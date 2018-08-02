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
	LoggerConfig    *ConfigLogger       `json:"loggerconfig"`
	EmailConfig     *ConfigEmail        `json:"emailconfig"`
}

type SPLConfigSettings struct {
	ConfigSettings
	SPLProdMstDBConfig map[string]*ConfigDB `json:"splproductmasterdbconfig"`
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

type ConfigLogger struct {
	LogLevel            string `json:"loglevel"`
	LoggingType         string `json:"logtype"`
	LoggingFluentHost   string `json:"fluenthost"`
	LoggingInfluxDBHost string `json:"fluenthost"`
}

type ConfigEmail struct {
	From         string `json:"from"`
	SMTPAddress  string `json:"smtpaddr"`
	SMTPPort     int    `json:"smtpport"`
	SMTPUsername string `json:"smtpusername"`
	SMTPPassword string `json:"smtpassword"`
}
