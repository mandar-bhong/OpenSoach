package models

type APITaskHandlerFunc func(exeCtx *APITaskExecutionCtx) (error, *APITaskResultModel)

type APITaskHandlerModel struct {
	Handler     APITaskHandlerFunc
	PayloadType interface{}
}

type APITaskExecutionCtx struct {
	Message    string
	SessionKey string
	Token      string
	TaskData   interface{}
}

type APITaskResultModel struct {
	IsSuccess bool                        `json:"isSuceess"`
	Data      interface{}                 `json:"data"`
	ErrorData APITaskResultErrorDataModel `json:"errordata"`
}

type APITaskResultErrorDataModel struct {
	ErrorCode int         `json:"code"`
	Data      interface{} `json:"data"`
}
