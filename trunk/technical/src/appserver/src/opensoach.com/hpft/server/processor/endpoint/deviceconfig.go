package endpoint

import (
	"opensoach.com/core/logger"
	lconst "opensoach.com/hpft/server/constants"
	"opensoach.com/hpft/server/dbaccess"
	gmodels "opensoach.com/models"
	pcmodels "opensoach.com/prodcore/models"
)

func ProcessDeviceSyncCompleted(ctx *pcmodels.DevicePacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult) {

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

	for _, spitem := range *splst {

		dbErr, servconflist := dbaccess.EPGetSPServConf(ctx.InstanceDBConn, ctx.TokenInfo.CpmID, spitem.ID)

		if dbErr != nil {
			logger.Context().WithField("CPMID", ctx.TokenInfo.CpmID).WithField("SPID", spitem.ID).LogError(SUB_MODULE_NAME, logger.Normal, "Unable to get serv conf.", dberr)
			continue
		}

		servconfinfo := &gmodels.DevicePacket{}
		servconfinfo.Header = gmodels.DeviceHeaderData{}
		servconfinfo.Header.Category = lconst.DEVICE_CMD_CAT_CONFIG
		servconfinfo.Header.CommandID = lconst.DEVCIE_CMD_CONFIG_SERVICE_POINTS_SERV_CONF
		servconfinfo.Header.SPID = spitem.ID

		servconfinfo.Payload = servconflist

		packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, servconfinfo)

	}

	for _, spitem := range *splst {

		dbErr, servconflist := dbaccess.EPGetSPPatientConf(ctx.InstanceDBConn, ctx.TokenInfo.CpmID, spitem.ID)

		if dbErr != nil {
			logger.Context().WithField("CPMID", ctx.TokenInfo.CpmID).WithField("SPID", spitem.ID).LogError(SUB_MODULE_NAME, logger.Normal, "Unable to get serv conf.", dberr)
			continue
		}

		servconfinfo := &gmodels.DevicePacket{}
		servconfinfo.Header = gmodels.DeviceHeaderData{}
		servconfinfo.Header.Category = lconst.DEVICE_CMD_CAT_CONFIG
		servconfinfo.Header.CommandID = lconst.DEVCIE_CMD_CONFIG_SERVICE_POINTS_PATIENT_CONF
		servconfinfo.Header.SPID = spitem.ID

		servconfinfo.Payload = servconflist

		packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, servconfinfo)

	}

	serverSyncCompleted := &gmodels.DevicePacket{}
	serverSyncCompleted.Header = gmodels.DeviceHeaderData{}
	serverSyncCompleted.Header.Category = lconst.DEVICE_CMD_CAT_CONFIG
	serverSyncCompleted.Header.CommandID = lconst.DEVCIE_CMD_CONFIG_SERVER_SYNC_COMPLETED

	packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, serverSyncCompleted)
}
