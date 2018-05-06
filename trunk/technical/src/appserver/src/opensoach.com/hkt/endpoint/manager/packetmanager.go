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
	wm "opensoach.com/prodcore/endpoint/webSocketManager"
)

type EPHandler struct {
}

var connID int

var connectionDeviceInfo map[int]string

func SendPacket(connID int, message string) {

	wm.SendMessage(connID, []byte(message))

}

//func (EPHandler) GetMasterCache() coremodels.CacheContext {
//	return repo.Instance().Context.Master.Cache
//}

func (EPHandler) RegisterHandler(handler map[string]pcepmgr.PacketProcessHandlerFunc) {
	handler[processor.GetAuthKey()] = processor.AuthProcessor
	handler[pcconst.DEVICE_CMD_PRE_EXECUTOR] = processor.PreProcessExecutor
}

func (EPHandler) OnEPConnection(wsconn int) {
	connID = wsconn
	fmt.Printf("Client connected %v\n", wsconn)
}

func (EPHandler) OnEPDisConnection(wsconn int) {
	fmt.Printf("Client disconnected %v\n", wsconn)
}

func (EPHandler) OnEPMessage(endPointToServerTaskModel *gmodels.EndPointToServerTaskModel) *gmodels.PacketProcessingResult {

	fmt.Println("Packet received at OnEPMessage")

	packetProcessingResult := &gmodels.PacketProcessingResult{}

	if isSuccess, jsonData := ghelper.ConvertToJSON(endPointToServerTaskModel); isSuccess == true {
		executionErr, exeResult := repo.Instance().ProdTaskContext.
			ProcessTask(constants.TASK_HANDLER_END_POINT_TO_SERVER_KEY, jsonData)

		if executionErr != nil {

		}

		isConvertionSuccess := ghelper.ConvertFromJSONString(exeResult, packetProcessingResult)

		if isConvertionSuccess == false {
			fmt.Println("Packet received at if isConvertionSuccess == false {")
			fmt.Println(exeResult)
			return packetProcessingResult
		}

	} else {
		logger.Context().Log("", logger.Normal, logger.Error, "Unable to convert to JSON packet")
	}

	packetProcessingResult.IsSuccess = true

	fmt.Println("Returning from OnEPMessage")
	return packetProcessingResult

}

func ProcessEPPacket(msg string) (string, error) {

	fmt.Printf("In ProcessEPPacket Message : %s \n", msg)

	SendPacket(connID, msg)
	return msg, nil

}
