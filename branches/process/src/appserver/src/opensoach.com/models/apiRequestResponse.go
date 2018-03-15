package models

type APIPayloadResponse struct {
	Success bool        `json:"issuccess"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type APIResponseError struct {
	Code int         `json:"code"`
	Data interface{} `json:"errordata"`
}
