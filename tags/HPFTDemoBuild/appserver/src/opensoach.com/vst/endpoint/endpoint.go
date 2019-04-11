package endpoint

import (
	"fmt"

	gmodels "opensoach.com/models"
	"opensoach.com/vst/endpoint/manager"
)

func Init(dbconfig *gmodels.ConfigDB) bool {

	fmt.Println("Starting VST EndPoint")

	initErr := manager.InitilizeModues(dbconfig)

	if initErr != nil {
		fmt.Printf("Error occured while starting endpoint server. Error %s \n", initErr)
		return false
	}

	return true
}

func DeInit() {

}
