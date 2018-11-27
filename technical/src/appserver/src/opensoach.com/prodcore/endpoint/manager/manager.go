package manager

import (
	gmodels "opensoach.com/models"
	wm "opensoach.com/prodcore/endpoint/websocketmanager"
	"opensoach.com/core/logger"
)

var epHandler EPHandler

type WSHandler struct {
}

type EPHandler interface {
	OnEPConnection(int)
	OnEPDisConnection(int)
	OnEPMessage(*gmodels.PacketProcessingTaskModel) *gmodels.PacketProcessingTaskResult
	RegisterHandler(map[string]PacketProcessHandlerFunc)
}

func Init(port int, handler EPHandler) error {

	epHandler = handler

	handler.RegisterHandler(packetHandlers)

	webServerStartErr := wm.Init(port, WSHandler{})

	logger.Context().WithField("Port",port) .LogDebug(SUB_MODULE_NAME,logger.Debug,"Starting WebSocket server" )
	return webServerStartErr
}
