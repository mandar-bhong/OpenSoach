package endpoint

import (
	"fmt"

	"opensoach.com/core/logger"
	repo "opensoach.com/hkt/server/repository"
	gmodels "opensoach.com/models"
	pchelper "opensoach.com/prodcore/helper"
)

func ProcessDeviceConnected(token string) error {

	fmt.Println("Device Connected task handled at hkt server")
	//repo.Instance().MasterTaskContext.SubmitTask(gmodels.TASK_SPL_EP_CONNECTED, "Client connected")

	fmt.Printf("Device connect Token: %s", token)

	return nil
}

func ProcessDeviceDisConnected(token string) error {

	logger.Context().WithField("Token", token).LogDebug(SUB_MODULE_NAME, logger.Normal, "Device disconnect task is handled by HKT server")

	isSuccess, _, jsonData := pchelper.CacheGetDeviceInfo(repo.Instance().Context.Master.Cache, token)

	if isSuccess == false {
		//log
		logger.Context().WithField("Token", token).Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Unable to get information for provided token")
		return fmt.Errorf("Unable to get information for provided token. Token: %s", token)
	}

	subErr := repo.Instance().MasterTaskContext.SubmitTask(gmodels.TASK_SPL_EP_DISCONNECTED, jsonData)

	if subErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while submitting device disconnection task to SPL server", subErr)
		return fmt.Errorf("Error occured while submittin task to SPL server. Token: %s ", token)
	}

	return nil
}
