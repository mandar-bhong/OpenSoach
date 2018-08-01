package endpoint

import (
	"fmt"

	"opensoach.com/hpft/endpoint/manager"
	gmodels "opensoach.com/models"
)

func Init(dbconfig *gmodels.ConfigDB) bool {

	fmt.Println("Starting HPFT EndPoint")

	initErr := manager.InitilizeModues(dbconfig)

	if initErr != nil {
		fmt.Printf("Error occured while starting endpoint server. Error %s \n", initErr)
		return false
	}

	return true
}

func DeInit() {

}
