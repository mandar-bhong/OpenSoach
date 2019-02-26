package endpoint

import (
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	lmodels "opensoach.com/hpft/server/models"
	repo "opensoach.com/hpft/server/repository"
	gmodels "opensoach.com/models"
	pchelper "opensoach.com/prodcore/helper"
	pcmodels "opensoach.com/prodcore/models"
	pcstoresync "opensoach.com/prodcore/server/storesync"
	pcservices "opensoach.com/prodcore/services"
)

func ProcessGetStoreSync(ctx *lmodels.PacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult) {

	packetProcessingResult.IsSuccess = true
	deviceCommandAck := gmodels.DeviceCommandAck{}
	deviceCommandAck.Ack = true

	reqModel := pcmodels.StoreSyncGetRequestModel{}

	devPacket := &gmodels.DevicePacket{}
	devPacket.Payload = &reqModel

	convErr := ghelper.ConvertFromJSONBytes(ctx.DevicePacket, devPacket)
	if convErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while converting from json", convErr)
		deviceCommandAck.Ack = false
		packetProcessingResult.IsSuccess = false
	} else {
		err, data := pcstoresync.GetChanges(ctx.InstanceDBConn, reqModel)
		if err != nil {
			logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while getting db changes", err)
			deviceCommandAck.Ack = false
			packetProcessingResult.IsSuccess = false
		}

		deviceCommandAck.Data = data
	}

	serviceCtx := &pcservices.ServiceContext{}
	serviceCtx.Repo = *repo.Instance()
	serviceCtx.ServiceConfig.SourcePacket = devPacket
	serviceCtx.ServiceConfig.AckData = deviceCommandAck
	serviceCtx.ServiceConfig.SourceToken = ctx.Token

	packetbldAckSourceService := &pcservices.PacketbldAckSourceService{}
	packetbldAckSourceService.ServiceContext = serviceCtx

	senderService := &pcservices.SenderService{}
	senderService.ServiceContext = serviceCtx

	packetbldAckSourceService.NextHandler = senderService

	serviceErr := packetbldAckSourceService.Handle(serviceCtx)
	if serviceErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while implementing services", serviceErr)
		packetProcessingResult.IsSuccess = false
		return
	}

}

func ProcessApplyStoreSync(ctx *lmodels.PacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult) {

	packetProcessingResult.IsSuccess = true
	deviceCommandAck := gmodels.DeviceCommandAck{}
	deviceCommandAck.Ack = true

	convErr, reqModel, devPacket := pchelper.GetStoreTableStruct(ctx.DevicePacket, pcmodels.StoreConfigModel{})
	if convErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while getting store struct", convErr)
		packetProcessingResult.IsSuccess = false
		deviceCommandAck.Ack = false

		serviceCtx := &pcservices.ServiceContext{}
		serviceCtx.Repo = *repo.Instance()
		serviceCtx.ServiceConfig.SourcePacket = devPacket
		serviceCtx.ServiceConfig.SourceToken = ctx.Token
		serviceCtx.ServiceConfig.AckData = deviceCommandAck
		notifyErr := pcstoresync.NotifyAck(serviceCtx)
		if notifyErr != nil {
			logger.Context().WithField("Service Context", serviceCtx).LogError(SUB_MODULE_NAME, logger.Normal, "Failed to notify apply sync changes.", notifyErr)
		}

		return
	}

	// add cpmid in request data
	for i := 0; i < len(reqModel.Data.([]map[string]interface{})); i++ {
		reqModel.Data.([]map[string]interface{})[i]["cpm_id_fk"] = ctx.TokenInfo.CpmID
	}

	err, _ := pcstoresync.ApplyChangesNotify(ctx.InstanceDBConn, reqModel, devPacket, ctx.Token, *repo.Instance())

	if err != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while apply sync changes", err)
		packetProcessingResult.IsSuccess = false
		return
	}

}
