package endpoint

import (
	"errors"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	repo "opensoach.com/hpft/server/repository"
	gmodels "opensoach.com/models"
	pchelper "opensoach.com/prodcore/helper"
	pcmodels "opensoach.com/prodcore/models"
	pcstoresync "opensoach.com/prodcore/server/storesync"
	pcservices "opensoach.com/prodcore/services"
)

func ProcessGetStoreSync(ctx *pcmodels.DevicePacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult) {

	packetProcessingResult.IsSuccess = true
	deviceCommandAck := gmodels.DeviceCommandAck{}
	deviceCommandAck.Ack = true

	reqModel := pcmodels.StoreSyncGetRequestModel{}
	reqModel.FilterHandler = AttachServerFilter

	devPacket := &gmodels.DevicePacket{}
	devPacket.Payload = &reqModel

	convErr := ghelper.ConvertFromJSONBytes(ctx.DevicePacket, devPacket)
	if convErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while converting from json", convErr)
		deviceCommandAck.Ack = false
		packetProcessingResult.IsSuccess = false
	} else {

		dbConnections := make(map[int]string)

		dbConnections[gmodels.DB_CONNECTION_MASTER] = repo.Instance().Context.Master.DBConn
		dbConnections[gmodels.DB_CONNECTION_NODE] = ctx.InstanceDBConn

		err, data := pcstoresync.GetChanges(ctx, dbConnections, reqModel)
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

func ProcessApplyStoreSync(ctx *pcmodels.DevicePacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult) {

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

func AttachServerFilter(ctx *pcmodels.DevicePacketProccessExecution, filterModel *pcmodels.SyncConfigModel, request *pcmodels.StoreSyncGetRequestModel) error {

	queryDataModel := pcmodels.QueryDataModel{}
	isSuccess := ghelper.ConvertFromJSONString(filterModel.QueryData, &queryDataModel)
	if isSuccess == false {
		logger.Context().WithField("DB Server Filter", filterModel.QueryData).LogError(SUB_MODULE_NAME, logger.Normal, "Failed to convert query data json.", nil)
		return errors.New("Unable to server parse query parameter form json data")
	}

	if request.QueryParams == nil {
		request.QueryParams = make(map[string]interface{})
	}

	if len(queryDataModel.Select.Filters) > 0 {
		for _, each := range queryDataModel.Select.Filters {
			switch each {
			case "cpm":
				request.QueryParams["cpmid"] = ctx.TokenInfo.CpmID
				break
			case "updatedon":
				request.QueryParams["updatedon"] = request.UpdatedOn
				break
			}
		}
	}

	return nil
}
