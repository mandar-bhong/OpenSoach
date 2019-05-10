package processor

import (
	"fmt"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	wsm "opensoach.com/prodcore/endpoint/websocketmanager"
	pcmodels "opensoach.com/prodcore/models"
)

func SendEPPacketHandler(msg string) error {

	taskEPPacketDataModels := []pcmodels.TaskEPPacketSendDataModel{}
	isJsonSuccess := ghelper.ConvertFromJSONString(msg, &taskEPPacketDataModels)

	if isJsonSuccess == false {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Unable to convert json packet for ep send packet", nil)
		return fmt.Errorf("Unable to convert json packet", nil)
	}

	for _, epItem := range taskEPPacketDataModels {

		chnID, hasChn := tokenvsChnID[epItem.Token]

		if hasChn == false {
			logger.Context().WithField("Token", epItem.Token).LogDebug(SUB_MODULE_NAME, logger.Normal, "Unable to find channel id for token")
			continue
		}

		issucess := wsm.SendMessage(chnID, []byte(epItem.Packet))
		if issucess == false {
			logger.Context().WithField("Token", epItem.Token).LogDebug(SUB_MODULE_NAME, logger.Normal, "Unable to send packet")

		}

		logger.Context().WithField("Token", epItem.Token).LogDebug(SUB_MODULE_NAME, logger.Normal, "Send packet succesfully.")

	}

	return nil
}
