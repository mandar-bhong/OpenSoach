package models

type ExecutionContext struct {
	SessionInfo             UserSessionInfo
	SelectedCustomerProduct int64
	RequestJSON             string
	Product                 ProductInfoModel
}

type ProductInfoModel struct {
	NodeDbConn string
}
