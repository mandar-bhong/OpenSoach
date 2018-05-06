package manager

import (
	"fmt"

	gmodels "opensoach.com/models"
	wh "opensoach.com/prodcore/endpoint/websockethelper"

	"opensoach.com/prodcore/endpoint/processor"
	ws "opensoach.com/prodcore/endpoint/websocketmanager"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	pcconst "opensoach.com/prodcore/constants"
	pchelper "opensoach.com/prodcore/helper"
)

var chIDDeviceInfo map[int]string
var packetHandlers map[string]PacketProcessHandlerFunc

type PacketProcessHandlerFunc func(*gmodels.PacketProcessingTaskModel) *gmodels.PacketProcessingTaskResult

func init() {
	chIDDeviceInfo = make(map[int]string)
	packetHandlers = make(map[string]PacketProcessHandlerFunc)
}

func (WSHandler) OnConnection(wsconn int) {
	//	connID = wsconn
	fmt.Printf("Client connected %v\n", wsconn)

}
func (WSHandler) OnDisConnection(wsconn int) {
	fmt.Printf("Client disconnected %v\n", wsconn)
}

func (WSHandler) OnMessage(message wh.WebsocketDataReceivedMessageStruct) {
	fmt.Printf("Client message %v\n", string(message.Message))

	packetProcessingTaskModel := &gmodels.PacketProcessingTaskModel{}
	packetProcessingTaskModel.ChannelID = message.ChannelID
	packetProcessingTaskModel.Token = "Token1"
	packetProcessingTaskModel.EPTaskListner = "TaskListner"
	packetProcessingTaskModel.Message = message.Message

	err, packetHeader := processor.DecodeHeader(message.Message)

	if err != nil {
		//Packet err
		return
	}

	cmd := pchelper.GetDeviceCmdKeyFromHeader(packetHeader)

	var packetProcessingResult *gmodels.PacketProcessingTaskResult
	preExecutor, hasPreExecutor := packetHandlers[pcconst.DEVICE_CMD_PRE_EXECUTOR]

	if hasPreExecutor {
		packetProcessingResult = preExecutor(packetProcessingTaskModel)
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
