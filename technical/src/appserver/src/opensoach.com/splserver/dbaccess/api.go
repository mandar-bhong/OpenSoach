package dbaccess

func UpdateCPMIDToInstDB(dbConn string, customerId int64) (error, *[]lmodels.DBSplMasterCustomerTableRowModel) {

	logger.Context().LogDebug(SUB_MODULE_NAME, logger.Normal, "Executing GetCustomerById")

	selDBCtx := dbmgr.SelectContext{}
	data := &[]lmodels.DBSplMasterCustomerTableRowModel{}
	selDBCtx.DBConnection = dbConn
	selDBCtx.Query = dbquery.QUERY_GET_CUSTOMER_TABLE_INFO_BY_ID
	selDBCtx.QueryType = dbmgr.Query
	selDBCtx.Dest = data
	selErr := selDBCtx.Select(customerId)
	if selErr != nil {
		return selErr, &[]lmodels.DBSplMasterCustomerTableRowModel{}
	}
	return nil, data
}
