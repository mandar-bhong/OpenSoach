package models

type UserSessionInfo struct {
	UserID     int64 `json:"userid"`
	UserRoleID int   `json:"userroleid"`
	UserType   int   `json:"usertype"`
}
