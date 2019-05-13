package websockethelper

import (
	"sync"
)

type WebSocketConnectionReceivedFunc func(int)
type WebSocketClientDisconnectedFunc func(int)
type WebSocketDataReceivedFunc func(WebsocketDataReceivedMessageStruct)

type WebSocketConnectionReceivedHelperFunc func(*WebSocketConnection)
type WebSocketClientDisconnectedHelperFunc func(*WebSocketConnection)
type WebSocketDataReceivedHelperFunc func(*WebSocketConnection, []byte)

type WebsocketDataReceivedMessageStruct struct {
	ChannelID int
	Message   []byte
}

type WebsocketInitStruct struct {
	WebSocketPort                 int
	OnWebSocketDataReceiver       WebSocketDataReceivedFunc
	OnWebSocketConnection         WebSocketConnectionReceivedFunc
	OnWebSocketClientDisconnected WebSocketClientDisconnectedFunc
}

type WebsocketInitHelperStruct struct {
	WebSocketPort                 int
	WebSocketDataReceiver         WebSocketDataReceivedHelperFunc
	OnWebSocketConnection         WebSocketConnectionReceivedHelperFunc
	OnWebSocketClientDisconnected WebSocketClientDisconnectedHelperFunc
}

type WebsockData struct {
	RecLock                 *sync.Mutex
	DeviceConnectionMap     map[int]*WebSocketConnection
	ConnectionDeviceMap     map[*WebSocketConnection]int
	DeviceConnectionCounter int
}
