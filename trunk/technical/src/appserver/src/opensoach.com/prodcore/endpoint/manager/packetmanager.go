package manager

import (
	gmodels "opensoach.com/models"
	wh "opensoach.com/prodcore/endpoint/websockethelper"

	"opensoach.com/prodcore/endpoint/processor"
	ws "opensoach.com/prodcore/endpoint/websocketmanager"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	pcconst "opensoach.com/prodcore/constants"
	pchelper "opensoach.com/prodcore/helper"
)

var SUB_MODULE_NAME = "ProdCore.Endpoint.Manager"

var chIDDeviceInfo map[int]string
var packetHandlers map[string]PacketProcessHandlerFunc

type PacketProcessHandlerFunc func(*gmodels.PacketProcessingTaskModel) *gmodels.PacketProcessingTaskResult

func init() {
	chIDDeviceInfo = make(map[int]string)
	packetHandlers = make(map[string]PacketProcessHandlerFunc)
}

func (WSHandler) OnConnection(wsconn int) {

	if epHandler == nil {
		logger.Context().WithField("ConnectionID", wsconn).Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Unable to raise event for device connected. End Point Handler is nil")
		return
	}

	epHandler.OnEPConnection(wsconn)
}

func (WSHandler) OnDisConnection(wsconn int) {

	if epHandler == nil {
		logger.Context().WithField("ConnectionID", wsconn).Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Unable to raise event for device disconnected. End Point Handler is nil")
		return
	}

	epHandler.OnEPDisConnection(wsconn)
}

func (WSHandler) OnMessage(packet wh.WebsocketDataReceivedMessageStruct) {

	packetProcessingTaskModel := &gmodels.PacketProcessingTaskModel{}
	packetProcessingTaskModel.ChannelID = packet.ChannelID
	packetProcessingTaskModel.Token = ghelper.GenerateTaskToken()
	packetProcessingTaskModel.EPTaskListner = "TaskListner"
	packetProcessingTaskModel.Message = packet.Message

	err, packetHeader := processor.DecodeHeader(packet.Message)

	if err != nil {
		logger.Context().WithField("Packet Ddata", string(packet.Message)).LogError(SUB_MODULE_NAME, logger.Normal, "Header decoding failed.", err)
		return
	}

	cmd := pchelper.GetDeviceCmdKeyFromHeader(packetHeader)

	var packetProcessingResult *gmodels.PacketProcessingTaskResult
	preExecutor, hasPreExecutor := packetHandlers[pcconst.DEVICE_CMD_PRE_EXECUTOR]

	if hasPreExecutor {
		packetProcessingResult = preExecutor(packetProcessingTaskModel)

		//If preexecutor available the result should be success

		if packetProcessingResult.IsSuccess == false {

			logger.Context().WithField("Status Code", packetProcessingResult.StatusCode).Log(SUB_MODULE_NAME, logger.Normal, logger.Error, "Pre-executor failed")

			switch packetProcessingResult.StatusCode {
			case gmodels.DEVICE_PROCESSING_AUTH_TOKEN_NOT_FOUND:

				isSuccess, jsonPacket := processor.GetUnauthorizedDevicePacket()

				if isSuccess == false {
					//TODO: Need to stop communication if token not valid
					//ws.DisconnectClient(packet.ChannelID)
					return
				}

				ws.SendMessage(packet.ChannelID, []byte(jsonPacket))
			}

			return
		}
	}

	executor, hasHandler := packetHandlers[cmd]

	if hasHandler == true {
		packetProcessingResult = executor(packetProcessingTaskModel)
	} else {
		packetProcessingResult = epHandler.OnEPMessage(packetProcessingTaskModel)
	}

	if len(packetProcessingResult.AckPayload) > 0 {

		for _, packet := range packetProcessingResult.AckPayload {

			isJsonConvSucc, jsonData := ghelper.ConvertToJSON(packet)

			if isJsonConvSucc == false {
				logger.Context().WithField("Packet", packet).Log("ProdCore.EP.Manager", logger.Normal, logger.Error, "Unable to convert to json data")
				continue
			}

			ws.SendMessage(packetProcessingTaskModel.ChannelID, []byte(jsonData))
		}
	}
}
