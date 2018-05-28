package endpoint

import (
	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	"opensoach.com/hkt/constants"
	hktmodels "opensoach.com/hkt/models"
	lconst "opensoach.com/hkt/server/constants"
	"opensoach.com/hkt/server/dbaccess"
	lhelper "opensoach.com/hkt/server/helper"
	lmodels "opensoach.com/hkt/server/models"
	gmodels "opensoach.com/models"
)

func ProcessComplaintData(ctx *lmodels.PacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult) {

	devicePacket := &gmodels.DevicePacket{}
	devicePacket.Payload = &[]lmodels.PacketComplaintData{}

	convErr := ghelper.ConvertFromJSONBytes(ctx.DevicePacket, devicePacket)

	if convErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while converting from json data.", convErr)
		packetProcessingResult.IsSuccess = false
		return
	}

	packetComplaintDataList := *devicePacket.Payload.(*[]lmodels.PacketComplaintData)

	for _, packetComplaintDataItem := range packetComplaintDataList {
		dbComplaintInsertRowModel := hktmodels.DBComplaintInsertRowModel{}
		dbComplaintInsertRowModel.CpmId = ctx.TokenInfo.CpmID
		dbComplaintInsertRowModel.SpId = packetComplaintDataItem.SpId
		dbComplaintInsertRowModel.Description = packetComplaintDataItem.Description
		dbComplaintInsertRowModel.ComplaintBy = packetComplaintDataItem.ComplaintBy
		dbComplaintInsertRowModel.RaisedOn = ghelper.GetCurrentTime()
		dbComplaintInsertRowModel.ComplaintState = constants.DB_COMPLAINT_STATE_OPEN

		dbErr := dbaccess.EPInsertComplaintData(ctx.InstanceDBConn, dbComplaintInsertRowModel)

		if dbErr != nil {
			logger.Context().WithField("Token", ctx.Token).
				WithField("ComplaintData", dbComplaintInsertRowModel).LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while saving complaint data.", dbErr)
		}
	}

	commandAck := lhelper.GetEPAckPacket(lconst.DEVICE_CMD_CAT_ACK_DEFAULT,
		devicePacket.Header.SeqID,
		true, 0, nil)

	packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, commandAck)

	packetProcessingResult.IsSuccess = true
}
