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
	EPTaskListner string `json:"eptasklistner"`
	ChannelID     int    `json:"channelid"`
	Message       []byte `json:"message"`
}
