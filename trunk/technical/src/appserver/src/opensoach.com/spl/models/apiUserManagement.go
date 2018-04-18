package models

type AuthRequest struct {
	UserName string `db:"usr_name" json:"username"`
	Password string `db:"usr_password" json:"password"`
	ProdCode string `db:"prod_code" json:"prodcode"`
}

type AuthResponse struct {
	Token     string `json:"token"`
	UroleCode string `json:"urolecode"`
}

type ValidateAuthTokenRequest struct {
	Token string `josn:"token"`
}
