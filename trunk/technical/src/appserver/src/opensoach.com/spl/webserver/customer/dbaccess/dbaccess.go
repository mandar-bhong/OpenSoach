package dbaccess

import (
	dbmgr "opensoach.com/core/manager/db"
	"opensoach.com/spl/constants/dbquery"
	lmodels "opensoach.com/spl/models"
)

func GetCustomerById(dbConn string, customerId int64) (error, *lmodels.DBSplMasterCustomerTableRowModel) {

	selDBCtx := dbmgr.SelectContext{}
	data := &lmodels.DBSplMasterCustomerTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_CUSTOMER_TABLE_INFO_BY_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Get(customerId)
	if selErr != nil {
		return selErr, &lmodels.DBSplMasterCustomerTableRowModel{}
	}
	return nil, data
}

func GetCustomerDetailsById(dbConn string, customerId int64) (error, *lmodels.DBSplMasterCustDetailsTableRowModel) {
	selDBCtx := dbmgr.SelectContext{}
	data := &lmodels.DBSplMasterCustDetailsTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_CUSTOMER_DETAILS_TABLE_INFO_BY_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Get(customerId)
	if selErr != nil {
		return selErr, &lmodels.DBSplMasterCustDetailsTableRowModel{}
	}
	return nil, data
}

func GetCorpDetailsById(dbConn string, customerId int64) (error, *lmodels.DBSplMasterCorpTableRowModel) {
	selDBCtx := dbmgr.SelectContext{}
	data := &lmodels.DBSplMasterCorpTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_CORP_TABLE_INFO_BY_CUSTOMER_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Get(customerId)
	if selErr != nil {
		return selErr, &lmodels.DBSplMasterCorpTableRowModel{}
	}
	return nil, data
}
