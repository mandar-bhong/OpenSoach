package endpoint

import (
	"fmt"

	"opensoach.com/hkt/endpoint/manager"
	gmodels "opensoach.com/models"
)

func Init(dbconfig *gmodels.ConfigDB) bool {

	fmt.Println("Starting HKT EndPoint")

	initErr := manager.InitilizeModues(dbconfig)

	if initErr != nil {
		fmt.Printf("Error occured while starting endpoint server. Error %s \n", initErr)
		return false
	}

	return true
}

func DeInit() {

}
