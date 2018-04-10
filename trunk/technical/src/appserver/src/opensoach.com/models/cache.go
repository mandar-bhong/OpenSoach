package models

type UserSessionInfo struct {
	UserID     int64    `json:"userid"`
	CpmID      int64    `json:"cpmid"`
	UserRoleID int64    `json:"userroleid"`
	CustomerID int64    `json:customerid`
	UserType   int      `json:"usertype"`
	ModDB      ConfigDB `json:"moddb"`
}
