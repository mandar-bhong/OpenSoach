package websocketmanager

import (
	"sync"

	//"opensoach.com/core/logger"
	wh "opensoach.com/prodcore/endpoint/webSocketHelper"
)

var websocketInitData *wh.WebsocketInitStruct
var websockDataRec wh.WebsockData

func Init(port int,
	onconnection wh.WebSocketConnectionReceivedFunc,
	ondisconntion wh.WebSocketClientDisconnectedFunc,
	onmessage wh.WebSocketDataReceivedFunc) error {

	websocketInitData.WebSocketPort = port
	websocketInitData.OnWebSocketConnection = onconnection
	websocketInitData.OnWebSocketClientDisconnected = ondisconntion
	websocketInitData.OnWebSocketDataReceiver = onmessage

	var websocketInitStructModel wh.WebsocketInitHelperStruct

	websocketInitStructModel.WebSocketPort = port
	websocketInitStructModel.WebSocketDataReceiver = websocketDataReceived
	websocketInitStructModel.OnWebSocketConnection = websocketClientConnected
	websocketInitStructModel.OnWebSocketClientDisconnected = websocketClientDisConnected
	return wh.Init(&websocketInitStructModel)
}

func DeInit() {

}

func websocketDataReceived(c *wh.WebSocketConnection, message []byte) {
	//logger.Log(logger.CORESERVER, logger.DEBUG, "Websocket data received")
	go PrepareAndDispatch(c, message)
}

func websocketClientConnected(c *wh.WebSocketConnection) {
	//logger.Log(logger.CORESERVER, logger.DEBUG, "Websocket client received")
	//TODO: This code should be thread safe

	websockDataRec.RecLock.Lock()
	websockDataRec.DeviceConnectionCounter++
	tmpCounter := websockDataRec.DeviceConnectionCounter
	websockDataRec.DeviceConnectionMap[websockDataRec.DeviceConnectionCounter] = c
	websockDataRec.ConnectionDeviceMap[c] = websockDataRec.DeviceConnectionCounter
	//logger.Log(logger.CORESERVER, logger.DEBUG, "WebSocketManager: WS request channelID: %d", websockDataRec.DeviceConnectionCounter)
	websockDataRec.RecLock.Unlock()

	go websocketInitData.OnWebSocketConnection(tmpCounter)
}

func websocketClientDisConnected(c *wh.WebSocketConnection) {
	//logger.Log(logger.CORESERVER, logger.WARNING, "Websocket client disconnected")
	//TODO: This code should be thread safe

	websockDataRec.RecLock.Lock()
	counter := websockDataRec.ConnectionDeviceMap[c]
	delete(websockDataRec.DeviceConnectionMap, counter)
	delete(websockDataRec.ConnectionDeviceMap, c)
	websockDataRec.RecLock.Unlock()

	go websocketInitData.OnWebSocketClientDisconnected(counter)
}

func PrepareAndDispatch(c *wh.WebSocketConnection, message []byte) {
	websockDataRec.RecLock.Lock() /* fetches channel-ID in record-lock. */
	/* channelID := websockDataRec.ConnectionDeviceMap[c] */

	var dataMessage wh.WebsocketDataReceivedMessageStruct
	dataMessage.ChannelID = websockDataRec.ConnectionDeviceMap[c]
	dataMessage.Message = message
	websockDataRec.RecLock.Unlock()

	go websocketInitData.OnWebSocketDataReceiver(dataMessage)
}

func SendMessage(channelID int, message []byte) bool {
	websockDataRec.RecLock.Lock() /* fetches channel-ID in record-lock. */
	defer websockDataRec.RecLock.Unlock()

	ws := websockDataRec.DeviceConnectionMap[channelID]
	if ws == nil {
		//logger.Log(logger.CORESERVER, logger.WARNING, "Websocket not found for channel ID: %d", channelID)
		return false
	}

	wh.SendMessage(ws, message)

	return true
}

func DisconnectClient(channelID int) {
	//logger.Log(logger.CORESERVER, logger.DEBUG, "Disconnection ws connection for channelid: %d", channelID)

	websockDataRec.RecLock.Lock() /* fetches channel-ID in record-lock. */
	defer websockDataRec.RecLock.Unlock()

	ws := websockDataRec.DeviceConnectionMap[channelID]
	wh.CloseConnection(ws)
}

func init() {
	websocketInitData = &wh.WebsocketInitStruct{}
	websockDataRec = wh.WebsockData{}
	websockDataRec.RecLock = &sync.Mutex{}
	websockDataRec.DeviceConnectionMap = make(map[int]*wh.WebSocketConnection)
	websockDataRec.ConnectionDeviceMap = make(map[*wh.WebSocketConnection]int)
	websockDataRec.DeviceConnectionCounter = 0
}
