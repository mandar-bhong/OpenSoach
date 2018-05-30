package endpoint

import (
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	hktmodels "opensoach.com/hkt/models"
	lconst "opensoach.com/hkt/server/constants"
	"opensoach.com/hkt/server/dbaccess"
	lhelper "opensoach.com/hkt/server/helper"
	lmodels "opensoach.com/hkt/server/models"
	gmodels "opensoach.com/models"
)

func ProcessFeedbackData(ctx *lmodels.PacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult) {

	devicePacket := &gmodels.DevicePacket{}
	devicePacket.Payload = &[]lmodels.PacketFeedbackData{}

	convErr := ghelper.ConvertFromJSONBytes(ctx.DevicePacket, devicePacket)

	if convErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while converting from json data.", convErr)
		packetProcessingResult.IsSuccess = false
		return
	}

	packetFeedbackDataList := *devicePacket.Payload.(*[]lmodels.PacketFeedbackData)

	for _, packetFeedbackDataItem := range packetFeedbackDataList {
		dbFeedbackInsertRowModel := hktmodels.DBFeedbackInsertRowModel{}
		dbFeedbackInsertRowModel.CpmIdFk = ctx.TokenInfo.CpmID
		dbFeedbackInsertRowModel.SpIdFk = devicePacket.Header.SPID
		dbFeedbackInsertRowModel.Feedback = packetFeedbackDataItem.Feedback
		dbFeedbackInsertRowModel.RaisedOn = packetFeedbackDataItem.RaisedOn

		dbErr := dbaccess.EPInsertFeedbackData(ctx.InstanceDBConn, dbFeedbackInsertRowModel)

		if dbErr != nil {
			logger.Context().WithField("Token", ctx.Token).
				WithField("Feedback Data", dbFeedbackInsertRowModel).LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while saving complaint data.", dbErr)
		}
	}

	commandAck := lhelper.GetEPAckPacket(lconst.DEVICE_CMD_CAT_ACK_DEFAULT,
		devicePacket.Header.SeqID,
		true, 0, nil)

	packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, commandAck)

	packetProcessingResult.IsSuccess = true
}
