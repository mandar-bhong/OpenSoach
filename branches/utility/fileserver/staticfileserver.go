package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type Config struct {
	Port       string `json:"port"`
	IsSSL      bool   `json:"isssl"`
	ServerCert string `json:"cert"`
	ServerKey  string `json:"certkey"`
}

func main() {

	currDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	configFile := filepath.Join(currDir, filepath.Join("config.json"))
	fmt.Printf("configuration: %s\n", configFile)

	byteData, readError := ioutil.ReadFile("config.json")
	if readError != nil {
		fmt.Println("Error:", readError)
		return
	}

	config := &Config{}
	if err := json.Unmarshal(byteData, config); err != nil {
		fmt.Printf("ERROR: Env configuration error: %s\n", err.Error())
		return
	}

	http.Handle("/", http.FileServer(http.Dir("./web/")))
	fmt.Printf("Running on Port %s\n", config.Port)

	if config.IsSSL {
		err := http.ListenAndServeTLS(config.Port, config.ServerCert, config.ServerKey, nil)
		if err != nil {
			fmt.Println("Error:", err)
		}
	} else {
		err := http.ListenAndServe(config.Port, nil)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

}
