package api

import (
	"fmt"
	apimgr "opensoach.com/hkt/api/manager"
	gmodels "opensoach.com/models"
)

func Init(config *gmodels.ConfigDB) bool {

	initilizationErr := apimgr.InitilizeModues(config)

	if initilizationErr != nil{
		fmt.Println(initilizationErr.Error())
	}

	return true
}
