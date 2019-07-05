package services

import (
	"errors"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	pcconst "opensoach.com/prodcore/constants"
)

//SenderService this service sends packets
type SenderService struct {
	*ServiceContext
	NextHandler IHandler
}

func (r *SenderService) Handle(serctx *ServiceContext) error {

	if r.ServiceResult.AckPacket != nil {

		isJsonSucc, jsonData := ghelper.ConvertToJSON(r.ServiceResult.AckPacket)

		if isJsonSucc == false {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Unable to convert packet to json data", nil)
			return errors.New("Unable to convert packet to json")
		}

		sendErr := r.Repo.ProdTaskContext.SubmitTask(pcconst.TASK_EP_SEND_PACKET, jsonData)
		if sendErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while send packet task.", sendErr)
			return errors.New("Error occured while send packet task")
		}

	}

	if r.ServiceResult.DestinationPackets != nil {

		isJsonSucc, jsonData := ghelper.ConvertToJSON(r.ServiceResult.DestinationPackets)

		if isJsonSucc == false {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Unable to convert packet to json data", nil)
			return errors.New("Unable to convert packet to json")
		}

		sendErr := r.Repo.ProdTaskContext.SubmitTask(pcconst.TASK_EP_SEND_PACKET, jsonData)

		if sendErr != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while send packet task.", sendErr)
			return errors.New("Error occured while send packet task")
		}
	}

	if r.NextHandler != nil {
		err := r.NextHandler.Handle(serctx)
		return err
	}

	return nil
}
