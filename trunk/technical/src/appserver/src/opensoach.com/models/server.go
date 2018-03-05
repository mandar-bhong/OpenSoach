package models

type ExecutionContext struct {
	SessionInfo     UserSessionInfo
	SelectedProduct int64
	Request         interface{}
}
