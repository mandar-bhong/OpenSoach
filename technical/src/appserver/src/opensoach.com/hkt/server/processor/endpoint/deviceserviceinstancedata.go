package endpoint

import (
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	hktmodels "opensoach.com/hkt/models"
	"opensoach.com/hkt/server/dbaccess"
	lmodels "opensoach.com/hkt/server/models"
	gmodels "opensoach.com/models"
)

func ProcessServiceInstanceData(ctx *lmodels.PacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult) {

	devicePacket := &gmodels.DevicePacket{}
	devicePacket.Payload = &[]lmodels.PacketServiceInstanceData{}

	convErr := ghelper.ConvertFromJSONBytes(ctx.DevicePacket, devicePacket)

	if convErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while converting from json data.", convErr)
		packetProcessingResult.IsSuccess = false
		return
	}

	packetServiceInstanceDataList := *devicePacket.Payload.(*[]lmodels.PacketServiceInstanceData)

	for _, packetServiceInstanceDataItem := range packetServiceInstanceDataList {
		dbServiceInstanceDataRowModel := hktmodels.DBServiceInstanceTxDataRowModel{}
		dbServiceInstanceDataRowModel.CpmId = ctx.TokenInfo.CpmID
		dbServiceInstanceDataRowModel.ServiceInstanceID = packetServiceInstanceDataItem.ServiceInstanceID
		dbServiceInstanceDataRowModel.TransactionData = packetServiceInstanceDataItem.TxnData
		dbServiceInstanceDataRowModel.TransactionDate = packetServiceInstanceDataItem.TxnDate

		dbErr := dbaccess.EPInsertServiceInstanceData(ctx.InstanceDBConn, dbServiceInstanceDataRowModel)

		if dbErr != nil {
			logger.Context().WithField("Token", ctx.Token).
				WithField("InstanceData", dbServiceInstanceDataRowModel).LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while saving service instance data.", dbErr)
		}
	}

	dataReceivedAckPacket := &gmodels.DevicePacket{}
	dataReceivedAckPacket.Header = devicePacket.Header
	

	packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, dataReceivedAckPacket)

	packetProcessingResult.IsSuccess = true
}
