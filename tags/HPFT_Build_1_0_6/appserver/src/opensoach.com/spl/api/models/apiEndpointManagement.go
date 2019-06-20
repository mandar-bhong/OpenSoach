package models

type APIDeviceAuthRequest struct {
	SerialNo    string `db:"serialno" json:"serialno"`
	ProductCode string `db:"prod_code" json:"prodcode"`
}

type APIDeviceAuthResponse struct {
	Token       string `json:"token"`
	LocationUrl string `json:"locationurl"`
}

type APIDeviceUserLoginRequest struct {
	UserName string `db:"usr_name" json:"username"`
	Password string `db:"usr_password" json:"password"`
	ProdCode string `db:"prod_code" json:"prodcode"`
}

type APIDeviceUserListRequest struct {
	DeviceToken string `json:"devicetoken"`
}

type APIDeviceUserLoginResponse struct {
	Token       string `json:"token"`
	LocationUrl string `json:"locationurl"`
	UserID      int64  `json:"userid"`
	UserRoleID  int64  `json:"userroleid"`
	CpmID       int64  `json:"cpmid"`
}

type APIDeviceSharedUserAuthRequest struct {
	UserName    string `db:"usr_name" json:"username"`
	Password    string `db:"usr_password" json:"password"`
	DeviceToken string `json:"devicetoken"`
}

type APIDeviceUserCPMListRequest struct {
	UserName string `db:"usr_name" json:"username"`
}
