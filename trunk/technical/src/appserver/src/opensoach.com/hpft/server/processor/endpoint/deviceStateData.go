package endpoint

import (
	"fmt"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	hktmodels "opensoach.com/hpft/models"
	lconst "opensoach.com/hpft/server/constants"
	"opensoach.com/hpft/server/dbaccess"
	lhelper "opensoach.com/hpft/server/helper"
	lmodels "opensoach.com/hpft/server/models"
	gmodels "opensoach.com/models"
	pcmodels "opensoach.com/prodcore/models"
)

func ProcessDeviceStateBatteryLevelData(ctx *pcmodels.DevicePacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult) {

	devicePacket := &gmodels.DevicePacket{}
	devicePacket.Payload = &lmodels.PacketDeviceBatteryLevelUpdateData{}

	convErr := ghelper.ConvertFromJSONBytes(ctx.DevicePacket, devicePacket)

	if convErr != nil {
		logger.Context().LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while converting from json data.", convErr)
		packetProcessingResult.IsSuccess = false
		return
	}

	packetDeviceBatteryLevelUpdateData := *devicePacket.Payload.(*lmodels.PacketDeviceBatteryLevelUpdateData)

	dbDevStatusBatteryLevelUpdateDataModel := hktmodels.DBDevStatusBatteryLevelUpdateDataModel{}
	_, dbDevStatusBatteryLevelUpdateDataModel.DevId = ctx.GetDeviceID()
	dbDevStatusBatteryLevelUpdateDataModel.BatteryLevel = packetDeviceBatteryLevelUpdateData.BatteryLevel
	dbDevStatusBatteryLevelUpdateDataModel.BatteryLevelSince = ghelper.GetCurrentTime()
	fmt.Println(dbDevStatusBatteryLevelUpdateDataModel.BatteryLevelSince)

	dbErr := dbaccess.EPUpdateDeviceBatteryLevelData(ctx.InstanceDBConn, dbDevStatusBatteryLevelUpdateDataModel)

	if dbErr != nil {
		logger.Context().WithField("Token", ctx.Token).
			WithField("DeviceStatusBatterylevelData", dbDevStatusBatteryLevelUpdateDataModel).LogError(SUB_MODULE_NAME, logger.Normal, "Error occured while saving battery status data.", dbErr)
	}

	commandAck := lhelper.GetEPAckPacket(lconst.DEVICE_CMD_CAT_ACK_DEFAULT,
		devicePacket.Header.SeqID,
		true, 0, nil)

	packetProcessingResult.AckPayload = append(packetProcessingResult.AckPayload, commandAck)

	packetProcessingResult.IsSuccess = true
}
