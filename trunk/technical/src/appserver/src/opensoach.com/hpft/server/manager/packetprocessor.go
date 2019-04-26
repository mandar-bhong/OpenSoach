package manager

import (
	"errors"
	"fmt"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	serconst "opensoach.com/hpft/server/constants"
	"opensoach.com/hpft/server/dbaccess"
	lhelper "opensoach.com/hpft/server/helper"
	"opensoach.com/hpft/server/processor/endpoint"
	epproc "opensoach.com/hpft/server/processor/endpoint"
	repo "opensoach.com/hpft/server/repository"
	gmodels "opensoach.com/models"
	pcconst "opensoach.com/prodcore/constants"
	pchelper "opensoach.com/prodcore/helper"
	pcmodels "opensoach.com/prodcore/models"
)

type CommandExecutor func(context *pcmodels.DevicePacketProccessExecution, result *gmodels.PacketProcessingTaskResult)

var PacketProcessExecutor map[string]CommandExecutor

func init() {
	PacketProcessExecutor = make(map[string]CommandExecutor)
}

func InitProcessor() {

	CAT_DR_DEV_REG := pchelper.GetDeviceCmdKey(serconst.DEVICE_CMD_CAT_DEVICE_REG,
		serconst.DEVICE_CMD_CAT_DATA)

	CAT_DR_DEV_REG_ACK := pchelper.GetDeviceCmdKey(serconst.DEVICE_CMD_CAT_DEVICE_REG,
		serconst.DEVICE_CMD_CONFIG_PART_DATA)

	CAT_CONFIG_DEV_SYNC_COMP := pchelper.GetDeviceCmdKey(serconst.DEVICE_CMD_CAT_CONFIG,
		serconst.DEVICE_CMD_CONFIG_DEVICE_SYNC_COMPLETED)

	CAT_DATA_SERVICE_INST_DATA := pchelper.GetDeviceCmdKey(serconst.DEVICE_CMD_CAT_DATA,
		serconst.DEVICE_CMD_SERVICE_INST_DATA)

	CAT_DATA_COMPLAINT_DATA := pchelper.GetDeviceCmdKey(serconst.DEVICE_CMD_CAT_DATA,
		serconst.DEVICE_CMD_COMPLAINT_DATA)

	CAT_DATA_FEEDBACK_DATA := pchelper.GetDeviceCmdKey(serconst.DEVICE_CMD_CAT_DATA,
		serconst.DEVICE_CMD_FEEDBACK_DATA)

	CAT_DATA_DEVICE_STATE_BATTERY_LEVEL_DATA := pchelper.GetDeviceCmdKey(serconst.DEVICE_CMD_CAT_DATA,
		serconst.DEVICE_CMD_DEVICE_STATE_BATTERY_LEVEL_DATA)

	CAT_DATA_STORE_GET_SYNC := pchelper.GetDeviceCmdKey(serconst.DEVICE_CMD_CAT_DATA,
		serconst.DEVICE_CMD_STORE_GET_SYNC)

	CAT_DATA_STORE_APPLY_SYNC := pchelper.GetDeviceCmdKey(serconst.DEVICE_CMD_CAT_DATA,
		serconst.DEVICE_CMD_STORE_APPLY_SYNC)

	PacketProcessExecutor[CAT_DR_DEV_REG] = endpoint.ProcessDevReg

	PacketProcessExecutor[CAT_CONFIG_DEV_SYNC_COMP] = epproc.ProcessDeviceSyncCompleted
	PacketProcessExecutor[CAT_DR_DEV_REG_ACK] = endpoint.ProcessDevReg

	PacketProcessExecutor[CAT_DATA_SERVICE_INST_DATA] = endpoint.ProcessServiceInstanceData

	PacketProcessExecutor[CAT_DATA_COMPLAINT_DATA] = endpoint.ProcessComplaintData
	PacketProcessExecutor[CAT_DATA_FEEDBACK_DATA] = endpoint.ProcessFeedbackData
	PacketProcessExecutor[CAT_DATA_DEVICE_STATE_BATTERY_LEVEL_DATA] = endpoint.ProcessDeviceStateBatteryLevelData

	PacketProcessExecutor[CAT_DATA_STORE_GET_SYNC] = endpoint.ProcessGetStoreSync
	PacketProcessExecutor[CAT_DATA_STORE_APPLY_SYNC] = endpoint.ProcessApplyStoreSync

}

func ProcessEndPointReceivedPacket(msg string) (string, error) {

	packetProcessingResult := &gmodels.PacketProcessingTaskResult{}
	var packetResult string

	fmt.Println("Server: Task received for processing")

	endPointToServerTaskModel := &gmodels.PacketProcessingTaskModel{}

	if isSuccess := ghelper.ConvertFromJSONString(msg, endPointToServerTaskModel); isSuccess == false {

		logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Unable to convert from endPointToServerTaskModel JSON data")

		isJsonSuccess, jsonData := ghelper.ConvertToJSON(packetProcessingResult)

		if isJsonSuccess == false {

			return "", errors.New("Unable to convert result to json")
		}

		return jsonData, nil
	}

	isTokenGetSuccess, contextType, tokenInfo, userTokenInfo := lhelper.GetEPTokenInfo(repo.Instance().Context.Master.Cache, endPointToServerTaskModel.Token)

	if isTokenGetSuccess == false {
		return "", fmt.Errorf("Unable to get token. Token: ", endPointToServerTaskModel.Token)
	}

	var cpmid int64

	if contextType == pcconst.DEVICE_TYPE_SHARED_DEVICE {
		cpmid = tokenInfo.CpmID
	} else if contextType == pcconst.DEVICE_TYPE_USER_DEVICE {
		cpmid = userTokenInfo.Product.CustProdID
	}

	isDBInstGetSuccess, dbInstConn := dbaccess.EPGetInstanceDB(repo.Instance().Context.Master.DBConn, cpmid)

	if isDBInstGetSuccess != nil {
		return "", fmt.Errorf("Unable to get dbconn. CPMID: %d ", cpmid)
	}

	devicePacket := &gmodels.DevicePacket{}

	if err := ghelper.ConvertFromJSONBytes(endPointToServerTaskModel.Message, devicePacket); err != nil {
		logger.Context().Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Unable to convert from devicePacket JSON data")

		isJsonSuccess, jsonData := ghelper.ConvertToJSON(packetProcessingResult)

		if isJsonSuccess == false {
			//Handle error

			return "", errors.New("Unable to convert result to json")
		}

		return jsonData, nil
	}

	receivedCommand := pchelper.GetDeviceCmdKeyFromHeader(devicePacket.Header)

	executor, hasExecutor := PacketProcessExecutor[receivedCommand]

	if hasExecutor == false {

		isJsonSuccess, jsonData := ghelper.ConvertToJSON(packetProcessingResult)

		if isJsonSuccess == false {

			return "", fmt.Errorf("Unable to convert result to json", endPointToServerTaskModel.Token)
		}

		return jsonData, nil
	}

	packetProcessingResult.AckPayload = []*gmodels.DevicePacket{}

	packetProccessExecution := &pcmodels.DevicePacketProccessExecution{}
	packetProccessExecution.Token = endPointToServerTaskModel.Token
	packetProccessExecution.DevicePacket = endPointToServerTaskModel.Message
	packetProccessExecution.InstanceDBConn = dbInstConn

	if contextType == pcconst.DEVICE_TYPE_SHARED_DEVICE {
		packetProccessExecution.DeviceContext = tokenInfo

	} else if contextType == pcconst.DEVICE_TYPE_USER_DEVICE {
		packetProccessExecution.DeviceContext = userTokenInfo
	}

	executor(packetProccessExecution, packetProcessingResult)

	if packetProcessingResult.IsSuccess == false {

		return packetResult, nil
	}

	packetProcessingResult.IsSuccess = true

	return ConvertResult(packetProcessingResult)

}

func ConvertResult(packetProcessingResult *gmodels.PacketProcessingTaskResult) (string, error) {
	isJsonSuccess, jsonData := ghelper.ConvertToJSON(packetProcessingResult)

	if isJsonSuccess == false {
		//Handle error
		return "", errors.New("Unable to convert result to json")
	}

	fmt.Println("Result calculation completed returning")

	return jsonData, nil
}

func GetInstanceDB(token string) (bool, string) {

	isSuccess, jsondata := repo.Instance().Context.Master.Cache.Get(token)

	if isSuccess == false {
		return false, ""
	}

	deviceTokenModel := &gmodels.DeviceTokenModel{}
	isSuccess = ghelper.ConvertFromJSONString(jsondata, deviceTokenModel)

	if isSuccess == false {
		return false, ""
	}
	return true, ""
}
