package main

import (
	"schemagenerator.com/ReadConfig"
	"schemagenerator.com/process"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	configStruct := ReadConfig.ReadCongiguration()

	process.CreateSchema(configStruct)

}
