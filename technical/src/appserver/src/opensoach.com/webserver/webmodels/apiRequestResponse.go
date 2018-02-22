package webmodels

type PayloadResponse struct {
	Success bool        `json:"issuccess"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Code int         `json:"code"`
	Data interface{} `json:"errordata"`
}
