package models

type ConfigSettings struct {
	DBConfig ConfigDB `json:"databaseconfiguration"`
}

type ConfigDB struct {
	ConnectionString string `json:"connectionstring"`
	DBDriver         string `json:"driver"`
}
