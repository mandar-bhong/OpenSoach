package manager

import (
	"errors"
	"fmt"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	gmodels "opensoach.com/models"
	pchelper "opensoach.com/prodcore/helper"
	serconst "opensoach.com/vst/server/constants"
	"opensoach.com/vst/server/dbaccess"
	lhelper "opensoach.com/vst/server/helper"
	lmodels "opensoach.com/vst/server/models"
	"opensoach.com/vst/server/processor/endpoint"
	epproc "opensoach.com/vst/server/processor/endpoint"
	repo "opensoach.com/vst/server/repository"
)

type CommandExecutor func(context *lmodels.PacketProccessExecution, result *gmodels.PacketProcessingTaskResult)

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

	CAT_CONFIG_DEV_TOKEN_LIST := pchelper.GetDeviceCmdKey(serconst.DEVICE_CMD_CAT_CONFIG,
		serconst.DEVICE_CMD_CONFIG_DEVICE_TOKEN_LIST)

	CAT_DATA_SERVICE_INST_DATA := pchelper.GetDeviceCmdKey(serconst.DEVICE_CMD_CAT_DATA,
		serconst.DEVICE_CMD_SERVICE_INST_DATA)

	CAT_DATA_COMPLAINT_DATA := pchelper.GetDeviceCmdKey(serconst.DEVICE_CMD_CAT_DATA,
		serconst.DEVICE_CMD_COMPLAINT_DATA)

	CAT_DATA_FEEDBACK_DATA := pchelper.GetDeviceCmdKey(serconst.DEVICE_CMD_CAT_DATA,
		serconst.DEVICE_CMD_FEEDBACK_DATA)

	CAT_DATA_DEVICE_STATE_BATTERY_LEVEL_DATA := pchelper.GetDeviceCmdKey(serconst.DEVICE_CMD_CAT_DATA,
		serconst.DEVICE_CMD_DEVICE_STATE_BATTERY_LEVEL_DATA)

	CAT_DATA_DEVICE_VEHICLE_TOKEN_DATA := pchelper.GetDeviceCmdKey(serconst.DEVICE_CMD_CAT_DATA,
		serconst.DEVICE_CMD_VEHICLE_TOKEN_DATA)

	CAT_DATA_DEVICE_VEHICLE_DETAILS_DATA := pchelper.GetDeviceCmdKey(serconst.DEVICE_CMD_CAT_DATA,
		serconst.DEVICE_CMD_VEHICLE_DETAILS_DATA)

	CAT_DATA_DEVICE_TOKEN_GENERATION_DATA := pchelper.GetDeviceCmdKey(serconst.DEVICE_CMD_CAT_DATA,
		serconst.DEVICE_CMD_TOKEN_GENERATION_DATA)

	CAT_DATA_DEVICE_JOB_CREATION_DATA := pchelper.GetDeviceCmdKey(serconst.DEVICE_CMD_CAT_DATA,
		serconst.DEVICE_CMD_JOB_CREATION_DATA)

	CAT_DATA_DEVICE_JOB_EXE_DATA := pchelper.GetDeviceCmdKey(serconst.DEVICE_CMD_CAT_DATA,
		serconst.DEVICE_CMD_JOB_EXEC_DATA)

	CAT_DATA_DEVICE_TOKEN_GENERATION_CLAIM_DATA := pchelper.GetDeviceCmdKey(serconst.DEVICE_CMD_CAT_DATA,
		serconst.DEVICE_CMD_TOKEN_GENERATION_CLAIM_DATA)

	CAT_DATA_DEVICE_JOB_EXE_CLAIM_DATA := pchelper.GetDeviceCmdKey(serconst.DEVICE_CMD_CAT_DATA,
		serconst.DEVICE_CMD_JOB_EXEC_CLAIM_DATA)

	PacketProcessExecutor[CAT_DR_DEV_REG] = endpoint.ProcessDevReg

	PacketProcessExecutor[CAT_CONFIG_DEV_SYNC_COMP] = epproc.ProcessDeviceSyncCompleted
	PacketProcessExecutor[CAT_CONFIG_DEV_TOKEN_LIST] = epproc.ProcessDeviceTokenList

	PacketProcessExecutor[CAT_DR_DEV_REG_ACK] = endpoint.ProcessDevReg

	PacketProcessExecutor[CAT_DATA_SERVICE_INST_DATA] = endpoint.ProcessServiceInstanceData

	PacketProcessExecutor[CAT_DATA_COMPLAINT_DATA] = endpoint.ProcessComplaintData
	PacketProcessExecutor[CAT_DATA_FEEDBACK_DATA] = endpoint.ProcessFeedbackData
	PacketProcessExecutor[CAT_DATA_DEVICE_STATE_BATTERY_LEVEL_DATA] = endpoint.ProcessDeviceStateBatteryLevelData

	PacketProcessExecutor[CAT_DATA_DEVICE_VEHICLE_TOKEN_DATA] = endpoint.ProcessVehicleTokenData
	PacketProcessExecutor[CAT_DATA_DEVICE_VEHICLE_DETAILS_DATA] = endpoint.ProcessVehicleDetailsData
	PacketProcessExecutor[CAT_DATA_DEVICE_TOKEN_GENERATION_DATA] = endpoint.ProcessTokenGenerationData
	PacketProcessExecutor[CAT_DATA_DEVICE_JOB_CREATION_DATA] = endpoint.ProcessJobCreationData
	PacketProcessExecutor[CAT_DATA_DEVICE_JOB_EXE_DATA] = endpoint.ProcessJobExeData
	PacketProcessExecutor[CAT_DATA_DEVICE_TOKEN_GENERATION_CLAIM_DATA] = endpoint.ProcessGenerateTokenClaimData
	PacketProcessExecutor[CAT_DATA_DEVICE_JOB_EXE_CLAIM_DATA] = endpoint.ProcessJobExeClaimData

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

	isTokenGetSuccess, tokenInfo := lhelper.GetEPTokenInfo(repo.Instance().Context.Master.Cache, endPointToServerTaskModel.Token)

	if isTokenGetSuccess == false {
		return "", fmt.Errorf("Unable to get token. Token: ", endPointToServerTaskModel.Token)
	}

	isDBInstGetSuccess, dbInstConn := dbaccess.EPGetInstanceDB(repo.Instance().Context.Master.DBConn, tokenInfo.CpmID, tokenInfo.DevID)

	if isDBInstGetSuccess != nil {
		return "", fmt.Errorf("Unable to get dbconn. CPMID: %d, DeviceID: %d ", tokenInfo.CpmID, tokenInfo.DevID)
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

	packetProccessExecution := &lmodels.PacketProccessExecution{}
	packetProccessExecution.Token = endPointToServerTaskModel.Token
	packetProccessExecution.DevicePacket = endPointToServerTaskModel.Message
	packetProccessExecution.InstanceDBConn = dbInstConn
	packetProccessExecution.TokenInfo = tokenInfo

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
