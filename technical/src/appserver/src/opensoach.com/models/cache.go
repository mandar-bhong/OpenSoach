package models

type UserSessionInfo struct {
	UserID     int64    `json:"userid"`
	CpmID      int      `json:"cpmid"`
	UserRoleID int      `json:"userroleid"`
	CustomerID int      `json:customerid`
	UserType   int      `json:"usertype"`
	ModDB      ConfigDB `json:"moddb"`
}
