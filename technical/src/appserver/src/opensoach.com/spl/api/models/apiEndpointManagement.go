package models

type APIDeviceAuthRequest struct {
	SerialNo    string `db:"serialno" json:"serialno"`
	ProductCode string `db:"prod_code" json:"prodcode"`
}

type APIDeviceAuthResponse struct {
	Token       string `json:"token"`
	LocationUrl string `json:"locationurl"`
}

type APIDeviceUserAuthRequest struct {
	UserName    string `db:"usr_name" json:"username"`
	Password    string `db:"usr_password" json:"password"`
	DeviceToken string `json:"devicetoken"`
}
