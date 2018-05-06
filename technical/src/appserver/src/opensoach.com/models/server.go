package models

type ExecutionContext struct {
	SessionInfo             UserSessionInfo
	SelectedCustomerProduct int64
	RequestJSON             string
}

type ServerListingResultModel struct {
	RecordCount int
	RecordList  interface{}
}

type EndPointToServerTaskModel struct {
	Token         string `json:"token"`
	EPTaskListner string `json:"eptsklistner"`
	ChannelID     int    `json:"chnid"`
	Message       []byte `json:"msg"`
}

type PacketProcessingResult struct {
	IsSuccess  bool            `json:"issuccess"`
	StatusCode int             `json:"status"`
	AckPayload []*DevicePacket `json:"ackpayload"`
	ChannelID  int             `json:"chnid"`
}
