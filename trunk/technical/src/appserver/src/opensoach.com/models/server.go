package models

type ExecutionContext struct {
	SessionInfo             UserSessionInfo
	SelectedCustomerProduct int64
	RequestJSON             string
	SessionToken            string
}

type ServerListingResultModel struct {
	RecordCount int
	RecordList  interface{}
}

type PacketProcessingTaskModel struct {
	Token         string `json:"token"`
	EPTaskListner string `json:"eptsklistner"`
	ChannelID     int    `json:"chnid"`
	Message       []byte `json:"msg"`
}

type PacketProcessingTaskResult struct {
	IsSuccess  bool            `json:"issuccess"`
	StatusCode int             `json:"status"`
	AckPayload []*DevicePacket `json:"ackpayload"`
	ChannelID  int             `json:"chnid"`
}

type ExcelData struct {
	SheetName  string
	StartCell  string
	IsVertical bool
	Headers    []string
	Data       [][]string
}
