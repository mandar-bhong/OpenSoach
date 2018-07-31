package server

import (
	"fmt"

	"opensoach.com/hpft/server/manager"
	gmodels "opensoach.com/models"
)

func Init(dbconfig *gmodels.ConfigDB) bool {

	fmt.Println("Starting HPFT Server")

	initErr := manager.InitilizeModues(dbconfig)

	if initErr != nil {
		fmt.Printf("Unable to start server. Error: %s \n", initErr)
		return false
	}

	return true
}

func DeInit() {

}
