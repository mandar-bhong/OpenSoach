package webmodels

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserID    int64  `json:"userid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Category  int    `json:"category"`
	State     int    `json:"status"`
	Token     string `json:"token"`
}
