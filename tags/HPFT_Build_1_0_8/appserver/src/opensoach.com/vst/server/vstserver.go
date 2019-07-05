package server

import (
	"fmt"

	gmodels "opensoach.com/models"
	"opensoach.com/vst/server/manager"
)

func Init(dbconfig *gmodels.ConfigDB) bool {

	fmt.Println("Starting VST Server")

	initErr := manager.InitilizeModues(dbconfig)

	if initErr != nil {
		fmt.Printf("Unable to start server. Error: %s \n", initErr)
		return false
	}

	return true
}

func DeInit() {

}
