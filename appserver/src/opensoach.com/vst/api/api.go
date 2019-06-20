package api

import (
	"fmt"

	gmodels "opensoach.com/models"
	apimgr "opensoach.com/vst/api/manager"
)

func Init(config *gmodels.ConfigDB) bool {

	initilizationErr := apimgr.InitilizeModues(config)

	if initilizationErr != nil {
		fmt.Println(initilizationErr.Error())
		return false
	}

	return true
}

func DeInit() {

}
