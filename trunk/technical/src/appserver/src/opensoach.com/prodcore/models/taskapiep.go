package models

type APITaskProcessorHandlerFunc func(exeCtx *APITaskExecutionCtx) (error, *APITaskProcessorResultModel)

type APITaskProcessorHandlerModel struct {
	Handler     APITaskProcessorHandlerFunc
	PayloadType interface{}
}

type APITaskProcessorResultModel struct {
	IsSuccess  bool                        `json:"isSuceess"`
	Data       interface{}                 `json:"data"`
	IsEPSync   bool                        `json:"isepsync"`
	EPSyncData []EPTaskSendPacketDataModel `json:"epsyncdata"`
}

type EPTaskSendPacketDataModel struct {
	Token    string      `json:"token"`
	TaskType string      `json:"tasktype"`
	Data     interface{} `json:"data"`
}


type TaskEPPacketSendDataModel struct {
	Token  string `json:"token"`
	Packet string `json:"data"`
}