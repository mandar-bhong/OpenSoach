package server

import (
	"fmt"

	"opensoach.com/hpft/server/manager"
	gmodels "opensoach.com/models"
)

func Init(config *gmodels.ConfigSettings) bool {

	fmt.Println("Starting HPFT Server")

	initErr := manager.InitilizeModues(config)

	if initErr != nil {
		fmt.Printf("Unable to start server. Error: %s \n", initErr)
		return false
	}

	return true
}

func DeInit() {

}
