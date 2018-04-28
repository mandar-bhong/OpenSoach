package models

type UserSessionInfo struct {
	UserID     int64            `json:"userid"`
	UserRoleID int64            `json:"userroleid"`
	CustomerID int64            `json:customerid`
	Product    ProductInfoModel `json:productinfo`
}

type ProductInfoModel struct {
	CustProdID int64  `json:cpmid`
	NodeDbConn string `json:dbconn`
}
