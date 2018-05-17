package manager

import (
	"fmt"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	"opensoach.com/hkt/constants"
	repo "opensoach.com/hkt/endpoint/repository"
	gmodels "opensoach.com/models"
	//wh "opensoach.com/prodcore/endpoint/webSocketHelper"
	"opensoach.com/hkt/endpoint/processor"
	pcconst "opensoach.com/prodcore/constants"
	pcepmgr "opensoach.com/prodcore/endpoint/manager"

	wm "opensoach.com/prodcore/endpoint/websocketmanager"
)

type EPHandler struct {
}

var connectionDeviceInfo map[int]string

func SendPacket(connID int, message string) {

	wm.SendMessage(connID, []byte(message))

}

func (EPHandler) RegisterHandler(handler map[string]pcepmgr.PacketProcessHandlerFunc) {
	handler[processor.GetAuthKey()] = processor.AuthProcessor
	handler[pcconst.DEVICE_CMD_PRE_EXECUTOR] = processor.PreProcessExecutor
}

func (EPHandler) OnEPConnection(chnid int) {
	logger.Context().WithField("ChannelID", chnid).LogDebug(SUB_MODULE_NAME, logger.Normal, "Device Connected")
	processor.EPOnConnectProcessExecutor(chnid)
}

func (EPHandler) OnEPDisConnection(chnid int) {
	logger.Context().WithField("ChannelID", chnid).LogDebug(SUB_MODULE_NAME, logger.Normal, "Device Disconnected")
	processor.EPOnDisConnectProcessExecutor(chnid)
}

func (EPHandler) OnEPMessage(endPointToServerTaskModel *gmodels.PacketProcessingTaskModel) *gmodels.PacketProcessingTaskResult {

	fmt.Println("Packet received at OnEPMessage")
	logger.Context().WithField("ChannelID: ", endPointToServerTaskModel.ChannelID).LogDebug(SUB_MODULE_NAME, logger.Normal, "Packet received from enpoint")

	packetProcessingResult := &gmodels.PacketProcessingTaskResult{}

	isJSONSuccess, jsonData := ghelper.ConvertToJSON(endPointToServerTaskModel)

	if isJSONSuccess == false {
		logger.Context().WithField("endPointToServerTaskModel", endPointToServerTaskModel).Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Unable to convert to JSON packet")
		packetProcessingResult.StatusCode = 0 //Assign status code
		return packetProcessingResult
	}

	executionErr, exeResult := repo.Instance().ProdTaskContext.
		ProcessTask(constants.TASK_HANDLER_END_POINT_TO_SERVER_KEY, jsonData)

	if executionErr != nil {
		//TODO: Task processing is failed no data sending to device
		logger.Context().LogError(SUB_MODULE_NAME, logger.Server, "Sever returned error ", executionErr)
		packetProcessingResult.StatusCode = 0 //Assign status code
		return packetProcessingResult
	}

	isConvertionSuccess := ghelper.ConvertFromJSONString(exeResult, packetProcessingResult)

	if isConvertionSuccess == false {
		logger.Context().WithField("JSONData", exeResult).Log(SUB_MODULE_NAME, logger.Server, logger.Error, "Unable to convert json data to structure")
		return packetProcessingResult
	}

	packetProcessingResult.IsSuccess = true

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Message successfully processed from server.")

	return packetProcessingResult
}
