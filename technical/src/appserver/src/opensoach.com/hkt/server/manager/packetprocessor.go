package manager

import (
	"errors"
	"fmt"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	serconst "opensoach.com/hkt/server/constants"
	"opensoach.com/hkt/server/dbaccess"
	lhelper "opensoach.com/hkt/server/helper"
	lmodels "opensoach.com/hkt/server/models"
	"opensoach.com/hkt/server/processor/endpoint"
	epproc "opensoach.com/hkt/server/processor/endpoint"
	repo "opensoach.com/hkt/server/repository"
	gmodels "opensoach.com/models"
	pchelper "opensoach.com/prodcore/helper"
)

type CommandExecutor func(context *lmodels.PacketProccessExecution, result *gmodels.PacketProcessingTaskResult)

var PacketProcessExecutor map[string]CommandExecutor

func init() {
	PacketProcessExecutor = make(map[string]CommandExecutor)
}

func InitProcessor() {

	CAT_DR_DEV_REG := pchelper.GetDeviceCmdKey(serconst.DEVICE_CMD_CAT_DEVICE_REG,
		serconst.DEVICE_CMD_CAT_DATA,
		serconst.DEVICE_CMD_KEY)

	CAT_DR_DEV_REG_ACK := pchelper.GetDeviceCmdKey(serconst.DEVICE_CMD_CAT_DEVICE_REG,
		serconst.DEVICE_CMD_CONFIG_PART_DATA,
		serconst.DEVICE_CMD_ACK_KEY)

	CAT_CONFIG_DEV_SYNC_COMP := pchelper.GetDeviceCmdKey(serconst.DEVICE_CMD_CAT_CONFIG,
		serconst.DEVICE_CMD_CONFIG_DEVICE_SYNC_COMPLETED,
		serconst.DEVICE_CMD_KEY)

	PacketProcessExecutor[CAT_CONFIG_DEV_SYNC_COMP] = epproc.ProcessDeviceSyncCompleted

	PacketProcessExecutor[CAT_DR_DEV_REG] = endpoint.ProcessDevReg
	PacketProcessExecutor[CAT_DR_DEV_REG_ACK] = endpoint.ProcessDevReg

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
