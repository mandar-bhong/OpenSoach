package models

type APIDeviceAuthRequest struct {
	SerialNo string `db:"serialno" json:"serialno"`
}

type APIDeviceAuthResponse struct {
	Token       string `json:"token"`
	LocationUrl string `json:"locationurl"`
}
