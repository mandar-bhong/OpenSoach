package endpoint

import (
	"opensoach.com/core/logger"
	lconst "opensoach.com/hkt/server/constants"
	"opensoach.com/hkt/server/dbaccess"
	lmodels "opensoach.com/hkt/server/models"
	gmodels "opensoach.com/models"
)

func ProcessDeviceSyncCompleted(ctx *lmodels.PacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult) {

	dberr, splst := dbaccess.EPGetDeviceServicePoints(ctx.InstanceDBConn, ctx.TokenInfo.CpmID, ctx.TokenInfo.DevID)

	if dberr != nil {

		packetProcessingResult.IsSuccess = false
		return
	}

	packetProcessingResult.IsSuccess = true

	spinfo := &gmodels.DevicePacket{}
	spinfo.Header = gmodels.DeviceHeaderData{}
	spinfo.Header.Category = lconst.DEVICE_CMD_CAT_CONFIG
	spinfo.Header.CommandID = lconst.DEVICE_CMD_CONFIG_DEVICE_SERVICE_POINTS

	spinfo.Payload = splst

	packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, spinfo)

	for _, spitem := range *splst {

		dberr, authlist := dbaccess.EPGetSPAuthCodes(ctx.InstanceDBConn, ctx.TokenInfo.CpmID, spitem.ID)

		if dberr != nil {
			logger.Context().WithField("CPMID", ctx.TokenInfo.CpmID).WithField("SPID", spitem.ID).LogError(SUB_MODULE_NAME, logger.Normal, "Unable to get auth code.", dberr)
			continue
		}

		authcode := &gmodels.DevicePacket{}
		authcode.Header = gmodels.DeviceHeaderData{}
		authcode.Header.Category = lconst.DEVICE_CMD_CAT_CONFIG
		authcode.Header.CommandID = lconst.DEVICE_CMD_CONFIG_SERVICE_POINTS_AUTH_CODE
		authcode.Header.SPID = spitem.ID

		authcode.Payload = authlist

		packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, authcode)

	}

}
