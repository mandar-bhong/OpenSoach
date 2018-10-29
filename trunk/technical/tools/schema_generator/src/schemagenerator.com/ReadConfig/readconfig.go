package ReadConfig

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Driver           string `json:"driver"`
	ConnectionString string `json:"connectionstring"`
	DBName           string `json:"dbname"`
}

func ReadCongiguration() Config {
	config := Config{}

	configFile, err := os.Open("settings/DBConfig.json")
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	configData, err := ioutil.ReadAll(configFile)
	json.Unmarshal(configData, &config)

	return config
}
