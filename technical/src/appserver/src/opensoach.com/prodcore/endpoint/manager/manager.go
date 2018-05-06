package manager

import (
	gmodels "opensoach.com/models"
	wm "opensoach.com/prodcore/endpoint/webSocketManager"
)

var epHandler EPHandler

type WSHandler struct {
}

type EPHandler interface {
	OnEPConnection(int)
	OnEPDisConnection(int)
	OnEPMessage(*gmodels.EndPointToServerTaskModel) *gmodels.PacketProcessingResult
	RegisterHandler(map[string]PacketProcessHandlerFunc)
}

func Init(port int, handler EPHandler) error {

	epHandler = handler

	handler.RegisterHandler(packetHandlers)

	webServerStartErr := wm.Init(port, WSHandler{})

	return webServerStartErr
}
