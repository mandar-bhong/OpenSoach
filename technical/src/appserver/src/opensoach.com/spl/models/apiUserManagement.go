package models

type AuthRequest struct {
	UserName string `db:"usr_name" json:"username"`
	Password string `db:"usr_password" json:"password"`
	ProdCode string `db:"prod_code" json:"prodcode"`
}

type AuthResponse struct {
	Token        string `json:"token"`
	UroleCode    string `json:"urolecode"`
	UserCategory int    `json:"usrcategory"`
}

type ValidateAuthTokenRequest struct {
	Token string `json:"token"`
}

type CustomerAssociateUserRequest struct {
	UserName string `json:"usrname"`
	DBUsrCpmRowModel
}
