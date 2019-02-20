package models

type APIDocumentDownloadRequest struct {
	Uuid            string `json:"uuid"`
	DeviceAuthToken string `json:"token"`
}

type APIDocumentUploadResponse struct {
	Uuid string `json:"uuid"`
}
