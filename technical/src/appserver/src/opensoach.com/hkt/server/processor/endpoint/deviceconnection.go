package endpoint

import (
	"fmt"

	"opensoach.com/core/logger"
	//ghelper "opensoach.com/core/helper"
	//	lmodels "opensoach.com/hkt/server/models"
	repo "opensoach.com/hkt/server/repository"
	//gmodels "opensoach.com/models"
	pchelper "opensoach.com/prodcore/helper"
)

func ProcessDeviceConnected(token string) error {

	fmt.Println("Device Connected task handled at hkt server")
	//repo.Instance().MasterTaskContext.SubmitTask(gmodels.TASK_SPL_EP_CONNECTED, "Client connected")

	fmt.Printf("Device connect Token: %s", token)

	return nil
}

func ProcessDeviceDisConnected(token string) error {
	fmt.Println("Device DisConnected task handled at hkt server")
	fmt.Printf("HKT SERVER: Device connect Token: %s", token)

	//isSuccess, deviceTokenModel, jsonData := pchelper.CacheGetDeviceInfo(repo.Instance().Context.Master.Cache, token)
	isSuccess, _, _ := pchelper.CacheGetDeviceInfo(repo.Instance().Context.Master.Cache, token)

	if isSuccess == false {
		//log
		logger.Context().WithField("Token", token).Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Unable to get information for provided token")
		return fmt.Errorf("Unable to get information for provided token. Token: %s", token)
	}

	//repo.Instance().MasterTaskContext.SubmitTask(gmodels.TASK_SPL_EP_DISCONNECTED, jsonData)

	return nil
}
