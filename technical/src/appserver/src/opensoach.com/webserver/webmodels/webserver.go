package webmodels

type UserSessionInfo struct {
	UserID     int64 `json:"userid"`
	UserRoleID int   `json:"userroleid"`
	UserType   int   `json:"usertype"`
}

type ExecutionContext struct {
	SessionInfo     *UserSessionInfo
	SelectedProduct int
	Request         interface{}
}
