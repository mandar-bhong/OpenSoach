package models

type APITaskHandlerFunc func(msg string, sessionkey string, token string, taskData interface{}) (error, APITaskResultModel)

type APITaskHandlerModel struct {
	Handler     APITaskHandlerFunc
	PayloadType interface{}
}

type APITaskResultModel struct {
}
