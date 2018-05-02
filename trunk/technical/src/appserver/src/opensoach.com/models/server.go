package models

type ExecutionContext struct {
	SessionInfo             UserSessionInfo
	SelectedCustomerProduct int64
	RequestJSON             string
}

type ServerListingResultModel struct {
	RecordCount int
	RecordList  interface{}
}
