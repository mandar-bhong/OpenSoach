package endpoint

import (
	"fmt"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	hktmodels "opensoach.com/vst/models"
	lconst "opensoach.com/vst/server/constants"
	"opensoach.com/vst/server/dbaccess"
	lhelper "opensoach.com/vst/server/helper"
	lmodels "opensoach.com/vst/server/models"
)

func ProcessDeviceStateBatteryLevelData(ctx *lmodels.PacketProccessExecution, packetProcessingResult *gmodels.PacketProcessingTaskResult) {

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
	dbDevStatusBatteryLevelUpdateDataModel.DevId = ctx.TokenInfo.DevID
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
