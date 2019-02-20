package models

type APIDocumentDownloadRequest struct {
	Uuid string `json:"uuid"`
}

type APIDocumentUploadResponse struct {
	Uuid string `json:"uuid"`
}
