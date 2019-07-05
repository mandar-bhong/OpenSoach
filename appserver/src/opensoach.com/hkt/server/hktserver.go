package server

import (
	"fmt"

	"opensoach.com/hkt/server/manager"
	gmodels "opensoach.com/models"
)

func Init(dbconfig *gmodels.ConfigDB) bool {

	fmt.Println("Starting HKT Server")

	initErr := manager.InitilizeModues(dbconfig)

	if initErr != nil {
		fmt.Printf("Unable to start server. Error: %s \n", initErr)
		return false
	}

	return true
}

func DeInit() {

}
