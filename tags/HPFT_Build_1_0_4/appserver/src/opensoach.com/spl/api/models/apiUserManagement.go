package models

type APIAuthRequest struct {
	UserName string `db:"usr_name" json:"username"`
	Password string `db:"usr_password" json:"password"`
	ProdCode string `db:"prod_code" json:"prodcode"`
}

type APIAuthResponse struct {
	Token        string `json:"token"`
	UroleCode    string `json:"urolecode"`
	UserCategory int    `json:"usrcategory"`
}

type APIValidateAuthTokenRequest struct {
	Token string `json:"token"`
}

type APICustomerAssociateUserRequest struct {
	UserName string `json:"usrname"`
	DBUsrCpmRowModel
}

type APIUroleRequest struct {
	Prodcode string `json:"prodcode"`
}

type APIUpdatePasswordRequest struct {
	OldPassword string `json:"oldpassword"`
	NewPassword string `json:"newpassword"`
}

type APICUUserUpdateRequestModel struct {
	DBCUUserUpateRowModel
	UroleId int64 `json:"uroleid"`
}

type APIUserActivateRequestModel struct {
	Code string `json:"code"`
}

type APICreatePasswordRequest struct {
	UserID      int64  `json:"usrid"`
	NewPassword string `json:"newpassword"`
}

type APIForgotPasswordRequest struct {
	UserName string `json:"usrname"`
}

type APIResetPasswordRequest struct {
	UserName    string `json:"usrname"`
	Otp         string `json:"otp"`
	NewPassword string `json:"newpassword"`
}
