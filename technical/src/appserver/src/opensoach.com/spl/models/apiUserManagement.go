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

type RecordChangePassRequest struct {
	OldPassword string `json:"oldpassword"`
	NewPassword string `json:"newpassword"`
}

type DataListRequest struct {
	Limit          int         `json:"limit"`
	OrderBy        string      `json:"orderby"`
	OrderDirection string      `json:"orderdirection"`
	CurrentPage    int         `json:"page"`
	Filter         interface{} `json:"filter"`
}

type DataListResponse struct {
	TotalRecords    int           `json:"totalrecords"`
	FilteredRecords int           `json:"filteredrecords"`
	Records         []interface{} `json:"records"`
}
