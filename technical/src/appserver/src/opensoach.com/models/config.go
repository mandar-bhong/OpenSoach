package models

type ConfigSettings struct {
	DBConfig ConfigDB `json:"databaseconfiguration"`
}

type ConfigDB struct {
	ConnectionString int `json:"connectionstring"`
	DBDriver         int `json:"driver"`
}
