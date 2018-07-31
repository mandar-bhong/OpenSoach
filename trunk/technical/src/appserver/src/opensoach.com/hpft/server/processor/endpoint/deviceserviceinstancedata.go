package endpoint

import (
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	hktmodels "opensoach.com/hpft/models"
	lconst "opensoach.com/hpft/server/constants"
	"opensoach.com/hpft/server/dbaccess"
	lhelper "opensoach.com/hpft/server/helper"
	lmodels "opensoach.com/hpft/server/models"
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
		dbServiceInstanceDataRowModel.FOPCode = packetServiceInstanceDataItem.FOPCode
		dbServiceInstanceDataRowModel.Status = packetServiceInstanceDataItem.Status

		dbErr := dbaccess.EPInsertServiceInstanceData(ctx.InstanceDBConn, dbServiceInstanceDataRowModel)

		if dbErr != nil {
			logger.Context().WithField("Token", ctx.Token).
				WithField("InstanceData", dbServiceInstanceDataRowModel).LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while saving service instance data.", dbErr)
		}
	}

	commandAck := lhelper.GetEPAckPacket(lconst.DEVICE_CMD_CAT_ACK_DEFAULT,
		devicePacket.Header.SeqID,
		true, 0, nil)

	packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, commandAck)

	packetProcessingResult.IsSuccess = true
}
