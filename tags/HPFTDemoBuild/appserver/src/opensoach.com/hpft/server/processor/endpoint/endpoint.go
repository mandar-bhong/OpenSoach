package endpoint

import (
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	repo "opensoach.com/hpft/server/repository"
	pcconst "opensoach.com/prodcore/constants"
	pcmodels "opensoach.com/prodcore/models"
)

var SUB_MODULE_NAME = "HPFT.Server.Processor.EP"

func SendPacketToEP(epPackets []pcmodels.TaskEPPacketSendDataModel) {

	isJsonSucc, jsonData := ghelper.ConvertToJSON(epPackets)

	if isJsonSucc == false {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Unable to convert packet to json data", nil)
		return
	}

	sendErr := repo.Instance().ProdTaskContext.SubmitTask(pcconst.TASK_EP_SEND_PACKET, jsonData)

	if sendErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while submitting task.", sendErr)
	}

}
