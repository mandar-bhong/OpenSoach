package manager

import (
	"fmt"

	ghelper "opensoach.com/core/helper"
	"opensoach.com/core/logger"
	"opensoach.com/hkt/constants"
	repo "opensoach.com/hkt/endpoint/repository"
	gmodels "opensoach.com/models"
	wh "opensoach.com/prodcore/endpoint/webSocketHelper"
	wm "opensoach.com/prodcore/endpoint/webSocketManager"
)

var connID int

var connectionDeviceInfo map[int]string

func SendPacket(connID int, message string) {

	wm.SendMessage(connID, []byte(message))

}

func OnEPConnection(wsconn int) {
	connID = wsconn
	fmt.Printf("Client connected %v\n", wsconn)

}
func OnEPDisConnection(wsconn int) {
	fmt.Printf("Client disconnected %v\n", wsconn)
}
func OnEPMessage(message wh.WebsocketDataReceivedMessageStruct) {
	fmt.Printf("Client message %v\n", string(message.Message))

	endPointToServerTaskModel := gmodels.EndPointToServerTaskModel{}
	endPointToServerTaskModel.ChannelID = message.ChannelID
	endPointToServerTaskModel.Token = "Token1"
	endPointToServerTaskModel.EPTaskListner = "TaskListner"
	endPointToServerTaskModel.Message = message.Message

	if isSuccess, jsonData := ghelper.ConvertToJSON(endPointToServerTaskModel); isSuccess == true {
		repo.Instance().ProdTaskContext.
			SubmitTask(constants.TASK_HANDLER_END_POINT_TO_SERVER_KEY, jsonData)
	} else {
		logger.Context().Log("", logger.Normal, logger.Error, "Unable to convert to JSON packet")
	}

	repo.Instance().ProdTaskContext.
		SubmitTask(constants.TASK_HANDLER_END_POINT_TO_SERVER_KEY, string(message.Message))
}

func ProcessEPPacket(msg string) (string, error) {

	fmt.Printf("In ProcessEPPacket Message : %s \n", msg)

	SendPacket(connID, msg)
	return msg, nil

}
